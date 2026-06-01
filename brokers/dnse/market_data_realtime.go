package dnse

import (
	"context"
	"fmt"

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
	return startMarketDataSubscription(
		ctx,
		s.broker,
		BuildMarketDataChannel("tick", s.broker.config.MarketDataBoardID, "", "", s.broker.config.StreamEncoding),
		request.SymbolList(),
		func(subscription *realtime.QueueSubscription[domain.Tick], message map[string]any) {
			subscription.PublishEvent(decodeTick(message))
		},
	)
}

func (s *MarketDataRealtimeService) SubscribeTopPrice(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.TopPrice], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeTop); err != nil {
		return nil, err
	}
	return startMarketDataSubscription(
		ctx,
		s.broker,
		BuildMarketDataChannel("top_price", s.broker.config.MarketDataBoardID, "", "", s.broker.config.StreamEncoding),
		request.SymbolList(),
		func(subscription *realtime.QueueSubscription[domain.TopPrice], message map[string]any) {
			subscription.PublishEvent(decodeTopPrice(message))
		},
	)
}

func (s *MarketDataRealtimeService) SubscribeCandles(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
	interval string,
) (realtime.Subscription[domain.Candle], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeCandle); err != nil {
		return nil, err
	}
	return startMarketDataSubscription(
		ctx,
		s.broker,
		BuildMarketDataChannel("ohlc", "", interval, "", s.broker.config.StreamEncoding),
		request.SymbolList(),
		func(subscription *realtime.QueueSubscription[domain.Candle], message map[string]any) {
			subscription.PublishEvent(decodeCandle(message))
		},
	)
}

func (s *MarketDataRealtimeService) SubscribeClosedCandles(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
	interval string,
) (realtime.Subscription[domain.Candle], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeCandle); err != nil {
		return nil, err
	}
	return startMarketDataSubscription(
		ctx,
		s.broker,
		BuildMarketDataChannel("ohlc_closed", "", interval, "", s.broker.config.StreamEncoding),
		request.SymbolList(),
		func(subscription *realtime.QueueSubscription[domain.Candle], message map[string]any) {
			subscription.PublishEvent(decodeCandle(message))
		},
	)
}

func (s *MarketDataRealtimeService) SubscribeRawChannel(
	ctx context.Context,
	channel string,
	symbols []string,
) (realtime.Subscription[domain.RawPayload], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	return startMarketDataSubscription(ctx, s.broker, channel, symbols, publishRaw)
}

func (s *MarketDataRealtimeService) SubscribeTickExtra(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.RawPayload], error) {
	return s.subscribeBoardRaw(ctx, "tick_extra", request)
}

func (s *MarketDataRealtimeService) SubscribeExpectedPrice(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.RawPayload], error) {
	return s.subscribeBoardRaw(ctx, "expected_price", request)
}

func (s *MarketDataRealtimeService) SubscribeForeign(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.RawPayload], error) {
	return s.subscribeBoardRaw(ctx, "foreign", request)
}

func (s *MarketDataRealtimeService) SubscribeSecurityDefinition(
	ctx context.Context,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.RawPayload], error) {
	return s.subscribeBoardRaw(ctx, "security_definition", request)
}

func (s *MarketDataRealtimeService) SubscribeMarketIndex(
	ctx context.Context,
	indexName string,
) (realtime.Subscription[domain.RawPayload], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	channel := BuildMarketDataChannel("market_index", "", "", indexName, s.broker.config.StreamEncoding)
	return startMarketDataSubscription(ctx, s.broker, channel, nil, publishRaw)
}

func (s *MarketDataRealtimeService) subscribeBoardRaw(
	ctx context.Context,
	kind string,
	request marketdata.SubscribeSymbolRequest,
) (realtime.Subscription[domain.RawPayload], error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataRealtimeRaw); err != nil {
		return nil, err
	}
	channel := BuildMarketDataChannel(kind, s.broker.config.MarketDataBoardID, "", "", s.broker.config.StreamEncoding)
	return startMarketDataSubscription(ctx, s.broker, channel, request.SymbolList(), publishRaw)
}

func startMarketDataSubscription[T any](
	ctx context.Context,
	broker *Broker,
	channel string,
	symbols []string,
	publisher streamPublisher[T],
) (realtime.Subscription[T], error) {
	return startRealtimeSubscription(
		ctx,
		broker,
		BuildMarketDataSubscribeMessage(channel, symbols),
		isMarketDataPayload,
		publisher,
	)
}

func BuildMarketDataChannel(kind string, boardID string, resolution string, indexName string, encoding string) string {
	if encoding == "" {
		encoding = "mspack"
	}
	switch kind {
	case "tick", "tick_extra", "top_price", "expected_price", "security_definition", "foreign":
		if boardID == "" {
			boardID = "G1"
		}
		return fmt.Sprintf("%s.%s.%s", kind, boardID, encoding)
	case "ohlc", "ohlc_closed":
		if resolution == "" {
			return ""
		}
		return fmt.Sprintf("%s.%s.%s", kind, resolution, encoding)
	case "market_index":
		if indexName == "" {
			return ""
		}
		return fmt.Sprintf("market_index.%s.%s", indexName, encoding)
	default:
		return ""
	}
}

func BuildMarketDataSubscribeMessage(channel string, symbols []string) map[string]any {
	channelPayload := map[string]any{"name": channel}
	if symbols != nil {
		values := make([]any, len(symbols))
		for i, symbol := range symbols {
			values[i] = symbol
		}
		channelPayload["symbols"] = values
	}
	return map[string]any{"action": "subscribe", "channels": []any{channelPayload}}
}

func BuildMarketDataUnsubscribeMessage(channel string, symbols []string) map[string]any {
	message := BuildMarketDataSubscribeMessage(channel, symbols)
	message["action"] = "unsubscribe"
	return message
}

func BuildStreamPingMessage() map[string]string {
	return map[string]string{"action": "ping"}
}

func decodeTick(message map[string]any) domain.Tick {
	data := marketDataMessageData(message)
	return domain.Tick{
		Symbol:     stringify(data["symbol"]),
		Price:      decimalFrom(data["matchPrice"]),
		Quantity:   optionalDecimal(data["matchQtty"]),
		ReceivedAt: stringify(firstNonNil(data["sendingTime"], data["multicastReceiveTime"])),
		Raw:        rawPayload(message, nil),
	}
}

func decodeTopPrice(message map[string]any) domain.TopPrice {
	data := marketDataMessageData(message)
	bid := firstPriceLevel(data["bid"])
	offer := firstPriceLevel(data["offer"])
	return domain.TopPrice{
		Symbol:      stringify(data["symbol"]),
		BidPrice:    optionalDecimal(bid["price"]),
		BidQuantity: optionalDecimal(bid["qtty"]),
		AskPrice:    optionalDecimal(offer["price"]),
		AskQuantity: optionalDecimal(offer["qtty"]),
		ReceivedAt:  stringify(firstNonNil(data["sendingTime"], data["multicastReceiveTime"])),
		Raw:         rawPayload(message, nil),
	}
}

func decodeCandle(message map[string]any) domain.Candle {
	data := marketDataMessageData(message)
	return domain.Candle{
		Symbol:     stringify(data["symbol"]),
		Resolution: stringify(data["resolution"]),
		Time:       stringify(data["time"]),
		Open:       decimalFrom(data["open"]),
		High:       decimalFrom(data["high"]),
		Low:        decimalFrom(data["low"]),
		Close:      decimalFrom(data["close"]),
		Volume:     decimalFrom(data["volume"]),
		Raw:        rawPayload(message, nil),
	}
}

func publishRaw(subscription *realtime.QueueSubscription[domain.RawPayload], message map[string]any) {
	subscription.PublishEvent(rawPayload(message, nil))
}

func marketDataMessageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		return data
	}
	return message
}

func isMarketDataPayload(message map[string]any) bool {
	data := marketDataMessageData(message)
	return data["T"] != nil || data["symbol"] != nil
}

func firstPriceLevel(value any) map[string]any {
	levels, _ := value.([]any)
	if len(levels) == 0 {
		return map[string]any{}
	}
	level, _ := levels[0].(map[string]any)
	if level == nil {
		return map[string]any{}
	}
	return level
}
