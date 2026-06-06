package ssi

import (
	"context"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

type MarketDataRealtimeService struct {
	broker *Broker
}

func (s *MarketDataRealtimeService) SubscribeTicks(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.Tick], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeTicks); err != nil {
		return nil, err
	}
	return startSSIMarketDataSubscription(ctx, s.broker, buildSSIChannel("T", request.SymbolList()), mapSSITick)
}

func (s *MarketDataRealtimeService) SubscribeTopPrice(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.TopPrice], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeTop); err != nil {
		return nil, err
	}
	return startSSIMarketDataSubscription(ctx, s.broker, buildSSIChannel("X", request.SymbolList()), mapSSITopPrice)
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
	return strings.ToUpper(dataType) + ":" + strings.Join(normalized, ",")
}

func mapSSITick(message ssiRealtimeMessage) domain.Tick {
	return domain.Tick{
		Symbol:     ssiString(message.Data, "Symbol", "symbol"),
		Price:      decimalFrom(ssiValue(message.Data, "LastPrice", "MatchPrice", "Price", "price")),
		Quantity:   optionalDecimal(ssiValue(message.Data, "LastVol", "MatchQtty", "Quantity", "quantity")),
		ReceivedAt: ssiString(message.Data, "Time", "time", "TradingDate"),
		Raw:        ssiRawPayload(message),
	}
}

func mapSSITopPrice(message ssiRealtimeMessage) domain.TopPrice {
	return domain.TopPrice{
		Symbol:      ssiString(message.Data, "Symbol", "symbol"),
		BidPrice:    optionalDecimal(ssiValue(message.Data, "BidPrice1", "bidPrice1")),
		BidQuantity: optionalDecimal(ssiValue(message.Data, "BidVol1", "bidVol1")),
		AskPrice:    optionalDecimal(ssiValue(message.Data, "AskPrice1", "askPrice1")),
		AskQuantity: optionalDecimal(ssiValue(message.Data, "AskVol1", "askVol1")),
		ReceivedAt:  ssiString(message.Data, "Time", "time", "TradingDate"),
		Raw:         ssiRawPayload(message),
	}
}
