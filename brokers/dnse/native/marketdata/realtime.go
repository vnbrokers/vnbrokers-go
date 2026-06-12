package marketdata

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	nativerealtime "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/realtime"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

type realtimeService struct {
	dependencies      nativerealtime.Dependencies
	requireCapability func(core.Capability) error
}

func NewRealtimeService(dependencies nativerealtime.Dependencies, requireCapability func(core.Capability) error) RealtimeService {
	return &realtimeService{dependencies: dependencies, requireCapability: requireCapability}
}
func (s *realtimeService) SubscribeExpectedPrices(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.ExpectedPriceEvent], error) {
	return subscribe[dto.ExpectedPriceEvent](ctx, s, CapabilityRealtimeExpectedPrices, "expected_price", r.BoardID, "", "", r.Symbols)
}
func (s *realtimeService) SubscribeForeign(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.ForeignEvent], error) {
	return subscribe[dto.ForeignEvent](ctx, s, CapabilityRealtimeForeign, "foreign", r.BoardID, "", "", r.Symbols)
}
func (s *realtimeService) SubscribeMarketIndexes(ctx context.Context, r dto.SubscribeMarketIndexRequest) (realtime.Subscription[dto.MarketIndexEvent], error) {
	return subscribe[dto.MarketIndexEvent](ctx, s, CapabilityRealtimeMarketIndexes, "market_index", "", "", r.IndexName, nil)
}
func (s *realtimeService) SubscribeEstimatedMarketIndexes(ctx context.Context, r dto.SubscribeMarketIndexRequest) (realtime.Subscription[dto.EstimatedMarketIndexEvent], error) {
	return subscribe[dto.EstimatedMarketIndexEvent](ctx, s, CapabilityRealtimeEstimatedMarketIndexes, "estimated_market_index", "", "", r.IndexName, nil)
}
func (s *realtimeService) SubscribeOHLC(ctx context.Context, r dto.SubscribeOHLCRequest) (realtime.Subscription[dto.OHLCEvent], error) {
	return subscribe[dto.OHLCEvent](ctx, s, CapabilityRealtimeOHLC, "ohlc", "", r.Resolution, "", r.Symbols)
}
func (s *realtimeService) SubscribeClosedOHLC(ctx context.Context, r dto.SubscribeOHLCRequest) (realtime.Subscription[dto.OHLCEvent], error) {
	return subscribe[dto.OHLCEvent](ctx, s, CapabilityRealtimeClosedOHLC, "ohlc_closed", "", r.Resolution, "", r.Symbols)
}
func (s *realtimeService) SubscribeQuotes(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.QuoteEvent], error) {
	return subscribe[dto.QuoteEvent](ctx, s, CapabilityRealtimeQuotes, "top_price", r.BoardID, "", "", r.Symbols)
}
func (s *realtimeService) SubscribeSecurityDefinitions(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.SecurityDefinitionEvent], error) {
	return subscribe[dto.SecurityDefinitionEvent](ctx, s, CapabilityRealtimeSecurityDefinitions, "security_definition", r.BoardID, "", "", r.Symbols)
}
func (s *realtimeService) SubscribeTrades(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.TradeEvent], error) {
	return subscribe[dto.TradeEvent](ctx, s, CapabilityRealtimeTrades, "tick", r.BoardID, "", "", r.Symbols)
}
func (s *realtimeService) SubscribeTradeExtras(ctx context.Context, r dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.TradeExtraEvent], error) {
	return subscribe[dto.TradeExtraEvent](ctx, s, CapabilityRealtimeTradeExtras, "tick_extra", r.BoardID, "", "", r.Symbols)
}
func subscribe[T any](ctx context.Context, s *realtimeService, capability core.Capability, kind, board, resolution, index string, symbols []string) (realtime.Subscription[T], error) {
	if err := s.requireCapability(capability); err != nil {
		return nil, err
	}
	channel := buildChannel(kind, board, resolution, index, s.dependencies.Encoding)
	return nativerealtime.Subscribe(ctx, s.dependencies, buildSubscribeMessage(channel, symbols), isPayload, func(message map[string]any) (T, error) { return decodeEvent[T](messageData(message)) })
}
func buildChannel(kind, boardID, resolution, indexName, encoding string) string {
	if encoding == "" {
		encoding = "msgpack"
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
	case "market_index", "estimated_market_index":
		if indexName == "" {
			return ""
		}
		return fmt.Sprintf("%s.%s.%s", kind, indexName, encoding)
	default:
		return ""
	}
}
func buildSubscribeMessage(channel string, symbols []string) map[string]any {
	item := map[string]any{"name": channel}
	if symbols != nil {
		values := make([]any, len(symbols))
		for i, symbol := range symbols {
			values[i] = symbol
		}
		item["symbols"] = values
	}
	return map[string]any{"action": "subscribe", "channels": []any{item}}
}
func messageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		message = data
	}
	if marketIndex, ok := message["marketIndex"].(map[string]any); ok {
		return marketIndex
	}
	return message
}
func isPayload(message map[string]any) bool {
	data := messageData(message)
	return data["T"] != nil || data["symbol"] != nil || data["indexName"] != nil
}
func decodeEvent[T any](message map[string]any) (T, error) {
	var out T
	payload, err := json.Marshal(message)
	if err != nil {
		return out, err
	}
	err = json.Unmarshal(payload, &out)
	return out, err
}
