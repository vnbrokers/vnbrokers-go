package trading

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityRealtimeStockOrders                 core.Capability = "native.trading.realtime.stock_orders"
	CapabilityRealtimeDerivativeOrders            core.Capability = "native.trading.realtime.derivative_orders"
	CapabilityRealtimeDerivativeOpenPositions     core.Capability = "native.trading.realtime.derivative_open_positions"
	CapabilityRealtimeDerivativeConditionalOrders core.Capability = "native.trading.realtime.derivative_conditional_orders"
)

type RealtimeService interface {
	SubscribeStockOrders(context.Context, dto.SubscribeStockOrdersRequest) (realtime.Subscription[dto.StockOrderEvent], error)
	SubscribeDerivativeOrders(context.Context, dto.SubscribeDerivativeOrdersRequest) (realtime.Subscription[dto.DerivativeOrderEvent], error)
	SubscribeDerivativeOpenPositions(context.Context, dto.SubscribeDerivativeOpenPositionsRequest) (realtime.Subscription[dto.DerivativeOpenPositionEvent], error)
	SubscribeDerivativeConditionalOrders(context.Context, dto.SubscribeDerivativeConditionalOrdersRequest) (realtime.Subscription[dto.DerivativeConditionalOrderEvent], error)
}

type RealtimeDependencies struct {
	BaseURL           string
	AccessToken       func() string
	Headers           func(bool, bool) map[string]string
	RequireCapability func(core.Capability) error
	WebSocketFactory  func(context.Context, string, map[string]string) (transport.WebSocketTransport, error)
	PingInterval      time.Duration
}

type realtimeService struct {
	dependencies RealtimeDependencies
}

func NewRealtimeService(dependencies RealtimeDependencies) RealtimeService {
	if dependencies.PingInterval == 0 {
		dependencies.PingInterval = 2 * time.Second
	}
	return &realtimeService{dependencies: dependencies}
}

func (s *realtimeService) SubscribeStockOrders(ctx context.Context, _ dto.SubscribeStockOrdersRequest) (realtime.Subscription[dto.StockOrderEvent], error) {
	return subscribeProto(ctx, s.dependencies, CapabilityRealtimeStockOrders, "/ws/aither", "STOCK_ORDER", func(payload []byte) (dto.StockOrderEvent, error) {
		var event dto.StockOrderEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeOrders(ctx context.Context, _ dto.SubscribeDerivativeOrdersRequest) (realtime.Subscription[dto.DerivativeOrderEvent], error) {
	return subscribeProto(ctx, s.dependencies, CapabilityRealtimeDerivativeOrders, "/ws/nesoi", "DE_ORDER", func(payload []byte) (dto.DerivativeOrderEvent, error) {
		var event dto.DerivativeOrderEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeOpenPositions(ctx context.Context, _ dto.SubscribeDerivativeOpenPositionsRequest) (realtime.Subscription[dto.DerivativeOpenPositionEvent], error) {
	return subscribeProto(ctx, s.dependencies, CapabilityRealtimeDerivativeOpenPositions, "/ws/nesoi", "DE_OPEN_POSITION", func(payload []byte) (dto.DerivativeOpenPositionEvent, error) {
		var event dto.DerivativeOpenPositionEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeConditionalOrders(ctx context.Context, _ dto.SubscribeDerivativeConditionalOrdersRequest) (realtime.Subscription[dto.DerivativeConditionalOrderEvent], error) {
	return subscribeProto(ctx, s.dependencies, CapabilityRealtimeDerivativeConditionalOrders, "/ws/nesoi", "DE_CONDITIONAL_ORDER", func(payload []byte) (dto.DerivativeConditionalOrderEvent, error) {
		var event dto.DerivativeConditionalOrderEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func subscribeProto[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, path, topic string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	if err := dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	headers := map[string]string{}
	if dependencies.Headers != nil {
		headers = dependencies.Headers(true, false)
	}
	socket, err := dependencies.WebSocketFactory(childCtx, websocketURL(dependencies.BaseURL, path), headers)
	if err != nil {
		cancel()
		return nil, err
	}
	var sendMu sync.Mutex
	send := func(message string) error {
		sendMu.Lock()
		defer sendMu.Unlock()
		return socket.Send(childCtx, transport.WebSocketMessage(message))
	}
	subscription := realtime.NewQueueSubscription[T](128, func() error {
		cancel()
		return socket.Close()
	})
	subscription.PublishStatus(realtime.StatusConnected)
	subscription.PublishStatus(realtime.StatusAuthenticating)
	if err := send(protoAuthMessage(dependencies.AccessToken())); err != nil {
		_ = subscription.Close()
		return nil, err
	}
	go runProtoSubscription(childCtx, cancel, dependencies.PingInterval, capability, topic, socket, send, decode, subscription)
	return subscription, nil
}

func runProtoSubscription[T any](ctx context.Context, cancel context.CancelFunc, pingInterval time.Duration, capability core.Capability, topic string, socket transport.WebSocketTransport, send func(string) error, decode func([]byte) (T, error), subscription *realtime.QueueSubscription[T]) {
	defer subscription.Close()
	authenticated := false
	subscribed := false
	for {
		payload, err := socket.Receive(ctx)
		if err != nil {
			if ctx.Err() == nil {
				subscription.PublishError(err)
			}
			return
		}
		frame := parseProtoFrame(payload)
		switch frame.kind {
		case "authenticate":
			if !authSucceeded(frame.payload) {
				subscription.PublishStatus(realtime.StatusFailed)
				subscription.PublishError(sdkerrors.Auth("tcbs", string(capability), "TCBS websocket authentication failed"))
				return
			}
			authenticated = true
		case "pingTimeout":
			if authenticated && !subscribed {
				if err := send(protoSubscribeMessage(topic)); err != nil {
					subscription.PublishError(err)
					return
				}
				subscribed = true
				subscription.PublishStatus(realtime.StatusSubscribed)
				startPingLoop(ctx, cancel, pingInterval, "ping|1", send, subscription)
			}
		case "ping":
			if err := send("ping|1"); err != nil {
				subscription.PublishError(err)
				return
			}
		case "session":
			continue
		case "message_proto":
			if frame.topic != topic {
				continue
			}
			event, err := decode(frame.payload)
			if err != nil {
				subscription.PublishError(sdkerrors.Decode("tcbs", string(capability), "decode TCBS websocket event", payload, err))
				continue
			}
			subscription.PublishEvent(event)
		}
	}
}

type protoFrame struct {
	kind    string
	topic   string
	payload []byte
}

func parseProtoFrame(payload []byte) protoFrame {
	parts := strings.SplitN(string(payload), "|", 3)
	if len(parts) == 0 {
		return protoFrame{}
	}
	frame := protoFrame{kind: parts[0]}
	if parts[0] == "message_proto" && len(parts) == 3 {
		frame.topic = parts[1]
		frame.payload = []byte(parts[2])
	} else if len(parts) >= 2 {
		frame.payload = []byte(parts[1])
	}
	return frame
}

func protoAuthMessage(token string) string {
	payload, _ := json.Marshal(struct {
		JWT string `json:"jwt"`
	}{JWT: token})
	return "authenticate|" + base64.StdEncoding.EncodeToString(payload)
}

func protoSubscribeMessage(topic string) string {
	payload, _ := json.Marshal(struct {
		Topic string `json:"topic"`
	}{Topic: topic})
	return "subscribe|" + base64.StdEncoding.EncodeToString(payload)
}

func authSucceeded(payload []byte) bool {
	message := struct {
		Success bool            `json:"success"`
		Error   json.RawMessage `json:"error"`
	}{}
	return json.Unmarshal(payload, &message) == nil && message.Success
}

func websocketURL(baseURL, path string) string {
	endpoint := strings.TrimRight(baseURL, "/") + path
	endpoint = strings.Replace(endpoint, "https://", "wss://", 1)
	return strings.Replace(endpoint, "http://", "ws://", 1)
}

func startPingLoop[T any](ctx context.Context, cancel context.CancelFunc, interval time.Duration, message string, send func(string) error, subscription *realtime.QueueSubscription[T]) {
	if interval <= 0 {
		return
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := send(message); err != nil {
					if ctx.Err() == nil {
						subscription.PublishError(err)
						cancel()
					}
					return
				}
			}
		}
	}()
}
