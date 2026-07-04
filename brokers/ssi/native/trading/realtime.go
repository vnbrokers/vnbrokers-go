package trading

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

type SignalRClient interface {
	Connect(context.Context) error
	Invoke(string, string, ...any) error
	On(string, string, func([]json.RawMessage))
	OnError(func(error))
	SetHeader(string, string)
	SetQuery(string, string)
	Close() error
}

type RealtimeDependencies struct {
	TradingToken      func() string
	TradingStreamURL  string
	RequireCapability func(core.Capability) error
	NewSignalRClient  func(baseURL string, hubs []string) SignalRClient
}

type realtimeService struct {
	deps RealtimeDependencies
}

func NewRealtimeService(deps RealtimeDependencies) RealtimeService {
	return &realtimeService{deps: deps}
}

const ssiTradingHub = "BroadcastHubV2"

type realtimeEnvelope struct {
	DataType string `json:"DataType"`
	Content  string `json:"Content"`
}

type broadcastEnvelope struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type realtimeMessage struct {
	DataType string
	Data     map[string]any
}

func (s *realtimeService) SubscribeOrderEvents(ctx context.Context) (realtime.Subscription[dto.OrderEvent], error) {
	return subscribeTradingEvent(ctx, s.deps, core.CapabilityTradingRealtimeOrders, "orderEvent", decodeTypedEvent[dto.OrderEvent])
}

func (s *realtimeService) SubscribeOrderErrors(ctx context.Context) (realtime.Subscription[dto.OrderError], error) {
	return subscribeTradingEvent(ctx, s.deps, core.CapabilityTradingRealtimeOrders, "orderError", decodeTypedEvent[dto.OrderError])
}

func (s *realtimeService) SubscribeOrderMatchEvents(ctx context.Context) (realtime.Subscription[dto.OrderMatchEvent], error) {
	return subscribeTradingEvent(ctx, s.deps, core.CapabilityTradingRealtimeOrders, "orderMatchEvent", decodeTypedEvent[dto.OrderMatchEvent])
}

func (s *realtimeService) SubscribeClientPortfolioEvents(ctx context.Context) (realtime.Subscription[dto.ClientPortfolioEvent], error) {
	return subscribeTradingEvent(ctx, s.deps, core.CapabilityTradingRealtimePosition, "clientPortfolioEvent", decodeTypedEvent[dto.ClientPortfolioEvent])
}

func (s *realtimeService) SubscribeFCOEvents(ctx context.Context) (realtime.Subscription[dto.FCOEvent], error) {
	return subscribeTradingEvent(ctx, s.deps, core.CapabilityTradingConditionalOrders, "fcoEvent", decodeTypedEvent[dto.FCOEvent])
}

func subscribeTradingEvent[T any](
	ctx context.Context,
	deps RealtimeDependencies,
	capability core.Capability,
	eventType string,
	decode func(realtimeMessage) (T, error),
) (realtime.Subscription[T], error) {
	if err := deps.RequireCapability(capability); err != nil {
		return nil, err
	}
	return startTradingSubscription(ctx, deps,
		func(client SignalRClient, subscription *realtime.QueueSubscription[T]) {
			client.SetQuery("notify_id", "-1")
			registerBroadcastHandler(client, subscription, []string{eventType}, decode)
		},
	)
}

func startTradingSubscription[T any](
	ctx context.Context,
	deps RealtimeDependencies,
	configure func(SignalRClient, *realtime.QueueSubscription[T]),
) (realtime.Subscription[T], error) {
	token := deps.TradingToken()
	if token == "" {
		return nil, sdkerrors.Auth("ssi", "realtime.subscribe", "SSI realtime requires an access token")
	}
	childCtx, cancel := context.WithCancel(ctx)
	cancelOnReturn := true
	defer func() {
		if cancelOnReturn {
			cancel()
		}
	}()
	client := deps.NewSignalRClient(deps.TradingStreamURL, []string{ssiTradingHub})
	client.SetHeader("Authorization", "Bearer "+token)

	subscription := realtime.NewQueueSubscription[T](128, func() error {
		cancel()
		return client.Close()
	})
	subscription.PublishStatus(realtime.StatusConnecting)
	client.OnError(func(err error) {
		if childCtx.Err() == nil {
			subscription.PublishError(err)
		}
	})
	configure(client, subscription)

	if err := client.Connect(childCtx); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusConnected)
	subscription.PublishStatus(realtime.StatusSubscribed)

	go func() {
		<-childCtx.Done()
		_ = subscription.Close()
	}()
	cancelOnReturn = false
	return subscription, nil
}

func registerBroadcastHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	eventTypes []string,
	mapEvent func(realtimeMessage) (T, error),
) {
	client.On(ssiTradingHub, "Broadcast", func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeMessage(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.Broadcast: %w", ssiTradingHub, err))
				continue
			}
			if !containsFold(eventTypes, message.DataType) {
				continue
			}
			event, err := mapEvent(message)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.Broadcast %s: %w", ssiTradingHub, message.DataType, err))
				continue
			}
			subscription.PublishEvent(event)
		}
	})
}

func decodeMessage(arg json.RawMessage) (realtimeMessage, error) {
	payload := []byte(arg)
	var text string
	if json.Unmarshal(arg, &text) == nil {
		payload = []byte(text)
	}

	var broadcast broadcastEnvelope
	if json.Unmarshal(payload, &broadcast) == nil && broadcast.Type != "" && len(broadcast.Data) > 0 && string(broadcast.Data) != "null" {
		data := map[string]any{}
		if err := json.Unmarshal(broadcast.Data, &data); err != nil {
			return realtimeMessage{}, err
		}
		return realtimeMessage{DataType: broadcast.Type, Data: data}, nil
	}

	var envelope realtimeEnvelope
	if json.Unmarshal(payload, &envelope) == nil && envelope.Content != "" {
		content := []byte(envelope.Content)
		data := map[string]any{}
		if err := json.Unmarshal(content, &data); err != nil {
			return realtimeMessage{}, err
		}
		return realtimeMessage{DataType: envelope.DataType, Data: data}, nil
	}

	data := map[string]any{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return realtimeMessage{}, err
	}
	return realtimeMessage{Data: data}, nil
}

func containsFold(values []string, target string) bool {
	for _, value := range values {
		if strings.EqualFold(value, target) {
			return true
		}
	}
	return false
}

func decodeEventData(message realtimeMessage, target any) error {
	payload, err := json.Marshal(message.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, target)
}

func decodeTypedEvent[T any](message realtimeMessage) (T, error) {
	var event T
	if err := decodeEventData(message, &event); err != nil {
		return event, err
	}
	return event, nil
}
