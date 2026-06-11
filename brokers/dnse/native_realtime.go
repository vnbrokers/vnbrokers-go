package dnse

import (
	"context"
	"encoding/json"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

type nativeMarketDataRealtimeService struct{ broker *Broker }

func (s *nativeMarketDataRealtimeService) SubscribeExpectedPrices(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.ExpectedPriceEvent], error) {
	return subscribeNativeMarketData[nativedto.ExpectedPriceEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeExpectedPrices, "expected_price", r.BoardID, "", "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeForeign(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.ForeignEvent], error) {
	return subscribeNativeMarketData[nativedto.ForeignEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeForeign, "foreign", r.BoardID, "", "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeMarketIndexes(ctx context.Context, r nativedto.SubscribeMarketIndexRequest) (realtime.Subscription[nativedto.MarketIndexEvent], error) {
	return subscribeNativeMarketData[nativedto.MarketIndexEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeMarketIndexes, "market_index", "", "", r.IndexName, nil)
}
func (s *nativeMarketDataRealtimeService) SubscribeOHLC(ctx context.Context, r nativedto.SubscribeOHLCRequest) (realtime.Subscription[nativedto.OHLCEvent], error) {
	return subscribeNativeMarketData[nativedto.OHLCEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeOHLC, "ohlc", "", r.Resolution, "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeClosedOHLC(ctx context.Context, r nativedto.SubscribeOHLCRequest) (realtime.Subscription[nativedto.OHLCEvent], error) {
	return subscribeNativeMarketData[nativedto.OHLCEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeClosedOHLC, "ohlc_closed", "", r.Resolution, "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeQuotes(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.QuoteEvent], error) {
	return subscribeNativeMarketData[nativedto.QuoteEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeQuotes, "top_price", r.BoardID, "", "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeSecurityDefinitions(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.SecurityDefinitionEvent], error) {
	return subscribeNativeMarketData[nativedto.SecurityDefinitionEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeSecurityDefinitions, "security_definition", r.BoardID, "", "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeTrades(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.TradeEvent], error) {
	return subscribeNativeMarketData[nativedto.TradeEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeTrades, "tick", r.BoardID, "", "", r.Symbols)
}
func (s *nativeMarketDataRealtimeService) SubscribeTradeExtras(ctx context.Context, r nativedto.SubscribeSymbolsRequest) (realtime.Subscription[nativedto.TradeExtraEvent], error) {
	return subscribeNativeMarketData[nativedto.TradeExtraEvent](ctx, s.broker, nativemarketdata.CapabilityRealtimeTradeExtras, "tick_extra", r.BoardID, "", "", r.Symbols)
}

func subscribeNativeMarketData[T any](ctx context.Context, b *Broker, capability core.Capability, kind, board, resolution, index string, symbols []string) (realtime.Subscription[T], error) {
	if err := b.RequireCapability(capability); err != nil {
		return nil, err
	}
	channel := buildMarketDataChannel(kind, board, resolution, index, b.config.StreamEncoding)
	return startMarketDataSubscription(ctx, b, channel, symbols, func(subscription *realtime.QueueSubscription[T], message map[string]any) {
		event, err := decodeNativeEvent[T](marketDataMessageData(message))
		if err != nil {
			subscription.PublishError(err)
			return
		}
		subscription.PublishEvent(event)
	})
}

type nativeTradingRealtimeService struct{ broker *Broker }

func (s *nativeTradingRealtimeService) SubscribeOrders(ctx context.Context, r nativedto.SubscribeTradingRequest) (realtime.Subscription[nativedto.OrderEvent], error) {
	return subscribeNativeTrading[nativedto.OrderEvent](ctx, s.broker, nativetrading.CapabilityRealtimeOrders, false, r)
}
func (s *nativeTradingRealtimeService) SubscribePositions(ctx context.Context, r nativedto.SubscribeTradingRequest) (realtime.Subscription[nativedto.PositionEvent], error) {
	return subscribeNativeTrading[nativedto.PositionEvent](ctx, s.broker, nativetrading.CapabilityRealtimePositions, true, r)
}
func subscribeNativeTrading[T any](ctx context.Context, b *Broker, capability core.Capability, position bool, r nativedto.SubscribeTradingRequest) (realtime.Subscription[T], error) {
	if err := b.RequireCapability(capability); err != nil {
		return nil, err
	}
	marketType := r.MarketType
	if marketType == "" {
		marketType = "STOCK"
	}
	message := buildStreamSubscribeOrdersMessage(marketType, b.config.StreamEncoding)
	if position {
		message = buildStreamSubscribePositionsMessage(marketType, b.config.StreamEncoding)
	}
	return startRealtimeSubscription(ctx, b, message, isTradingStreamPayload, func(subscription *realtime.QueueSubscription[T], message map[string]any) {
		event, err := decodeNativeEvent[T](tradingMessageData(message))
		if err != nil {
			subscription.PublishError(err)
			return
		}
		subscription.PublishEvent(event)
	})
}

func decodeNativeEvent[T any](message map[string]any) (T, error) {
	var out T
	payload, err := json.Marshal(message)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(payload, &out)
	return out, err
}
