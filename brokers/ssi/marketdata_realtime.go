package ssi

import (
	"context"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

type MarketDataRealtimeService struct {
	broker *Broker
}

func (s *MarketDataRealtimeService) SubscribeTradingStatus(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.TradingStatusEvent], error) {
	return subscribeSSIMarketData[dto.TradingStatusEvent](ctx, s.broker, "F", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeQuotes(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.QuoteEvent], error) {
	return subscribeSSIMarketData[dto.QuoteEvent](ctx, s.broker, "X-QUOTE", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeTrades(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.TradeEvent], error) {
	return subscribeSSIMarketData[dto.TradeEvent](ctx, s.broker, "X-TRADE", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeSnapshots(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.SnapshotEvent], error) {
	return subscribeSSIMarketData[dto.SnapshotEvent](ctx, s.broker, "X", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeForeignRooms(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.ForeignRoomEvent], error) {
	return subscribeSSIMarketData[dto.ForeignRoomEvent](ctx, s.broker, "R", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeMarketIndexes(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.MarketIndexEvent], error) {
	return subscribeSSIMarketData[dto.MarketIndexEvent](ctx, s.broker, "MI", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeOHLCV(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.OHLCVEvent], error) {
	return subscribeSSIMarketData[dto.OHLCVEvent](ctx, s.broker, "B", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeOddLots(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[dto.OddLotEvent], error) {
	return subscribeSSIMarketData[dto.OddLotEvent](ctx, s.broker, "OL", request.SymbolList())
}

func (s *MarketDataRealtimeService) SubscribeRawChannel(
	ctx context.Context,
	channel string,
) (realtime.Subscription[domain.RawPayload], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	return startSSIMarketDataSubscription(ctx, s.broker, channel, func(message ssiRealtimeMessage) domain.RawPayload {
		return ssiRawPayload(message)
	})
}

func startSSIMarketDataSubscription[T any](
	ctx context.Context,
	broker *Broker,
	channel string,
	mapEvent func(ssiRealtimeMessage) T,
) (realtime.Subscription[T], error) {
	return startSSISignalRSubscription(
		ctx,
		broker,
		broker.dataToken(),
		broker.config.MarketDataStreamURL,
		ssiMarketDataHub,
		func(client SignalRClient, subscription *realtime.QueueSubscription[T]) {
			registerSSIHandler(client, subscription, ssiMarketDataHub, "broadcast", mapEvent)
			registerSSIHandler(client, subscription, ssiMarketDataHub, "update", mapEvent)
		},
		func(client SignalRClient) error {
			return client.Invoke(ssiMarketDataHub, "SwitchChannels", channel)
		},
	)
}

func subscribeSSIMarketData[T any](
	ctx context.Context,
	broker *Broker,
	dataType string,
	symbols []string,
) (realtime.Subscription[T], error) {
	if err := broker.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	return startSSISignalRSubscription(
		ctx,
		broker,
		broker.dataToken(),
		broker.config.MarketDataStreamURL,
		ssiMarketDataHub,
		func(client SignalRClient, subscription *realtime.QueueSubscription[T]) {
			registerSSIJSONHandler(client, subscription, ssiMarketDataHub, "broadcast")
			registerSSIJSONHandler(client, subscription, ssiMarketDataHub, "update")
		},
		func(client SignalRClient) error {
			return client.Invoke(ssiMarketDataHub, "SwitchChannels", buildSSIChannel(dataType, symbols))
		},
	)
}

func buildSSIChannel(dataType string, symbols []string) string {
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
