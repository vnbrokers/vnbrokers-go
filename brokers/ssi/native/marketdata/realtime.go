package marketdata

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
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
	DataToken           func() string
	MarketDataStreamURL string
	RequireCapability   func(core.Capability) error
	NewSignalRClient    func(baseURL string, hubs []string) SignalRClient
}

type realtimeService struct {
	deps RealtimeDependencies
}

func NewRealtimeService(deps RealtimeDependencies) RealtimeService {
	return &realtimeService{deps: deps}
}

const fcMarketDataHub = "FcMarketDataV2Hub"

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
	Raw      []byte
}

func (s *realtimeService) SubscribeTradingStatus(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.TradingStatusEvent], error) {
	return subscribeMarketData[dto.TradingStatusEvent](ctx, s.deps, "F", request.SymbolList())
}

func (s *realtimeService) SubscribeQuotes(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.QuoteEvent], error) {
	return subscribeMarketData[dto.QuoteEvent](ctx, s.deps, "X-QUOTE", request.SymbolList())
}

func (s *realtimeService) SubscribeTrades(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.TradeEvent], error) {
	return subscribeMarketData[dto.TradeEvent](ctx, s.deps, "X-TRADE", request.SymbolList())
}

func (s *realtimeService) SubscribeSnapshots(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.SnapshotEvent], error) {
	return subscribeMarketData[dto.SnapshotEvent](ctx, s.deps, "X", request.SymbolList())
}

func (s *realtimeService) SubscribeForeignRooms(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.ForeignRoomEvent], error) {
	return subscribeMarketData[dto.ForeignRoomEvent](ctx, s.deps, "R", request.SymbolList())
}

func (s *realtimeService) SubscribeMarketIndexes(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.MarketIndexEvent], error) {
	return subscribeMarketData[dto.MarketIndexEvent](ctx, s.deps, "MI", request.SymbolList())
}

func (s *realtimeService) SubscribeOHLCV(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.OHLCVEvent], error) {
	return subscribeMarketData[dto.OHLCVEvent](ctx, s.deps, "B", request.SymbolList())
}

func (s *realtimeService) SubscribeOddLots(ctx context.Context, request marketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.OddLotEvent], error) {
	return subscribeMarketData[dto.OddLotEvent](ctx, s.deps, "OL", request.SymbolList())
}

func (s *realtimeService) SubscribeRawChannel(ctx context.Context, channel string) (realtime.Subscription[domain.RawPayload], error) {
	if err := s.deps.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	return startMarketDataSubscription(ctx, s.deps, channel, func(message realtimeMessage) domain.RawPayload {
		return domain.RawPayload{Source: "ssi", Data: message.Data, Bytes: message.Raw}
	})
}

func startMarketDataSubscription[T any](
	ctx context.Context,
	deps RealtimeDependencies,
	channel string,
	mapEvent func(realtimeMessage) T,
) (realtime.Subscription[T], error) {
	token := deps.DataToken()
	if token == "" {
		return nil, sdkerrors.Auth("ssi", "realtime.subscribe", "SSI realtime requires an access token")
	}
	childCtx, cancel := context.WithCancel(ctx)
	client := deps.NewSignalRClient(deps.MarketDataStreamURL, []string{fcMarketDataHub})
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
	registerTypedHandler(client, subscription, fcMarketDataHub, "broadcast", mapEvent)
	registerTypedHandler(client, subscription, fcMarketDataHub, "update", mapEvent)

	if err := client.Connect(childCtx); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusConnected)
	if err := client.Invoke(fcMarketDataHub, "SwitchChannels", channel); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusSubscribed)

	go func() {
		<-childCtx.Done()
		_ = subscription.Close()
	}()
	return subscription, nil
}

func subscribeMarketData[T any](
	ctx context.Context,
	deps RealtimeDependencies,
	dataType string,
	symbols []string,
) (realtime.Subscription[T], error) {
	if err := deps.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	token := deps.DataToken()
	if token == "" {
		return nil, sdkerrors.Auth("ssi", "realtime.subscribe", "SSI realtime requires an access token")
	}
	childCtx, cancel := context.WithCancel(ctx)
	client := deps.NewSignalRClient(deps.MarketDataStreamURL, []string{fcMarketDataHub})
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
	registerJSONHandler[T](client, subscription, fcMarketDataHub, "broadcast")
	registerJSONHandler[T](client, subscription, fcMarketDataHub, "update")

	if err := client.Connect(childCtx); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusConnected)
	if err := client.Invoke(fcMarketDataHub, "SwitchChannels", buildChannel(dataType, symbols)); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusSubscribed)

	go func() {
		<-childCtx.Done()
		_ = subscription.Close()
	}()
	return subscription, nil
}

func registerTypedHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	hub string,
	method string,
	mapEvent func(realtimeMessage) T,
) {
	client.On(hub, method, func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeMessage(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w", hub, method, err))
				continue
			}
			subscription.PublishEvent(mapEvent(message))
		}
	})
}

func registerJSONHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	hub string,
	method string,
) {
	client.On(hub, method, func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeMessage(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w", hub, method, err))
				continue
			}
			var event T
			if err := json.Unmarshal(message.Raw, &event); err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w %s", hub, method, err, message.Raw))
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
		return realtimeMessage{DataType: broadcast.Type, Data: data, Raw: broadcast.Data}, nil
	}

	var envelope realtimeEnvelope
	if json.Unmarshal(payload, &envelope) == nil && envelope.Content != "" {
		content := []byte(envelope.Content)
		data := map[string]any{}
		if err := json.Unmarshal(content, &data); err != nil {
			return realtimeMessage{}, err
		}
		return realtimeMessage{DataType: envelope.DataType, Data: data, Raw: content}, nil
	}

	data := map[string]any{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return realtimeMessage{}, err
	}
	return realtimeMessage{Data: data, Raw: payload}, nil
}

func buildChannel(dataType string, symbols []string) string {
	if len(symbols) == 0 {
		return strings.ToUpper(dataType) + ":ALL"
	}
	normalized := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		if symbol = strings.TrimSpace(symbol); symbol != "" {
			normalized = append(normalized, strings.ToUpper(symbol))
		}
	}
	if len(normalized) == 0 {
		return strings.ToUpper(dataType) + ":ALL"
	}
	return strings.ToUpper(dataType) + ":" + strings.Join(normalized, "-")
}
