package tcbs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	stockMatchEndpoint       = "/ws/aither"
	stockOrderTopic          = "STOCK_ORDER"
	tcbsRealtimePingInterval = 2 * time.Second
)

type TradingRealtimeService struct {
	broker *Broker
}

func (s *TradingRealtimeService) SubscribeStockMatches(
	ctx context.Context,
) (realtime.Subscription[domain.OrderEvent], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimeOrders); err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	subscription := realtime.NewQueueSubscription[domain.OrderEvent](128, func() error {
		cancel()
		return nil
	})
	subscription.PublishStatus(realtime.StatusConnecting)
	socket, err := s.broker.config.WebSocketFactory(childCtx, s.broker.wsURL(stockMatchEndpoint), s.broker.headers(true, false))
	if err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription = realtime.NewQueueSubscription[domain.OrderEvent](128, func() error {
		cancel()
		return socket.Close()
	})
	subscription.PublishStatus(realtime.StatusConnected)
	subscription.PublishStatus(realtime.StatusAuthenticating)
	var sendMu sync.Mutex
	sendText := func(ctx context.Context, message string) error {
		sendMu.Lock()
		defer sendMu.Unlock()
		return socket.Send(ctx, transport.WebSocketMessage(message))
	}
	if err := sendText(childCtx, buildTCBSRealtimeAuthMessage(s.broker.config.AccessToken)); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	go func() {
		defer subscription.Close()
		authenticated := false
		subscribed := false
		pingStarted := false
		for {
			payload, err := socket.Receive(childCtx)
			if err != nil {
				if childCtx.Err() == nil {
					subscription.PublishError(err)
				}
				return
			}
			frame := parseTCBSRealtimeFrame(payload)
			switch frame.kind {
			case "authenticate":
				if !tcbsRealtimeAuthSucceeded(frame.payload) {
					subscription.PublishStatus(realtime.StatusFailed)
					subscription.PublishError(sdkerrors.Auth("tcbs", "trading.realtime.stock_matches", "TCBS websocket authentication failed"))
					return
				}
				authenticated = true
			case "pingTimeout":
				if !pingStarted {
					pingStarted = true
					startTCBSRealtimePingLoop(childCtx, tcbsRealtimePingInterval, sendText, subscription, cancel)
				}
				if authenticated && !subscribed {
					if err := sendText(childCtx, buildTCBSRealtimeSubscribeMessage(stockOrderTopic)); err != nil {
						subscription.PublishStatus(realtime.StatusFailed)
						subscription.PublishError(err)
						return
					}
					subscribed = true
					subscription.PublishStatus(realtime.StatusSubscribed)
				}
			case "session":
				continue
			case "ping":
				if err := sendText(childCtx, "ping|1"); err != nil {
					subscription.PublishStatus(realtime.StatusFailed)
					subscription.PublishError(err)
					return
				}
			case "message_proto":
				if frame.topic != stockOrderTopic {
					continue
				}
				message, err := decodeTCBSStockOrderRealtimeMessage(frame.payload)
				if err != nil {
					subscription.PublishError(sdkerrors.Decode("tcbs", "trading.realtime.stock_matches", "decode TCBS websocket message", payload, err))
					continue
				}
				if !isStockOrderMessage(message) {
					continue
				}
				subscription.PublishEvent(MapStockOrderEvent(message))
			default:
				message := map[string]any{}
				if err := json.Unmarshal(payload, &message); err != nil {
					subscription.PublishError(sdkerrors.Decode("tcbs", "trading.realtime.stock_matches", "decode TCBS websocket message", payload, err))
					continue
				}
				if !isStockMatchMessage(message) {
					continue
				}
				subscription.PublishEvent(MapStockMatchEvent(message))
			}
		}
	}()
	return subscription, nil
}

func decodeTCBSStockOrderRealtimeMessage(payload []byte) (StockOrderRealtimeMessage, error) {
	var message StockOrderRealtimeMessage
	err := json.Unmarshal(payload, &message)
	return message, err
}

func buildTCBSRealtimeAuthMessage(jwt string) string {
	payload, _ := json.Marshal(struct {
		JWT string `json:"jwt"`
	}{JWT: jwt})
	return "authenticate|" + base64.StdEncoding.EncodeToString(payload)
}

func buildTCBSRealtimeSubscribeMessage(topic string) string {
	payload, _ := json.Marshal(struct {
		Topic string `json:"topic"`
	}{Topic: topic})
	return "subscribe|" + base64.StdEncoding.EncodeToString(payload)
}

type tcbsRealtimeFrame struct {
	kind    string
	topic   string
	payload []byte
}

func parseTCBSRealtimeFrame(payload []byte) tcbsRealtimeFrame {
	parts := strings.SplitN(string(payload), "|", 3)
	if len(parts) == 0 {
		return tcbsRealtimeFrame{}
	}
	frame := tcbsRealtimeFrame{kind: parts[0]}
	switch {
	case parts[0] == "message_proto" && len(parts) == 3:
		frame.topic = parts[1]
		frame.payload = []byte(parts[2])
	case len(parts) >= 2:
		frame.payload = []byte(parts[1])
	}
	return frame
}

func isTCBSRealtimeControlFrame(frame tcbsRealtimeFrame) bool {
	switch frame.kind {
	case "session", "ping", "pingTimeout", "authenticate":
		return true
	default:
		return false
	}
}

func tcbsRealtimeAuthSucceeded(payload []byte) bool {
	message := struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}{}
	return json.Unmarshal(payload, &message) == nil && message.Success && message.Error == ""
}

func startTCBSRealtimePingLoop(
	ctx context.Context,
	interval time.Duration,
	send func(context.Context, string) error,
	subscription *realtime.QueueSubscription[domain.OrderEvent],
	cancel context.CancelFunc,
) {
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
				if err := send(ctx, "ping|1"); err != nil {
					if ctx.Err() == nil {
						subscription.PublishStatus(realtime.StatusFailed)
						subscription.PublishError(err)
						cancel()
					}
					return
				}
			}
		}
	}()
}

func (b *Broker) wsURL(path string) string {
	out := b.url(path)
	out = strings.Replace(out, "https://", "wss://", 1)
	out = strings.Replace(out, "http://", "ws://", 1)
	return out
}

func isStockMatchMessage(message map[string]any) bool {
	return firstString(message, "orderID", "orderId", "orderNo") != "" || stringify(message["symbol"]) != ""
}

func isStockOrderMessage(message StockOrderRealtimeMessage) bool {
	return message.OrderID != "" || message.Symbol != ""
}
