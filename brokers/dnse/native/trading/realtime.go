package trading

import (
	"context"
	"encoding/json"

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
func (s *realtimeService) SubscribeOrders(ctx context.Context, r dto.SubscribeTradingRequest) (realtime.Subscription[dto.OrderEvent], error) {
	return subscribeTrading[dto.OrderEvent](ctx, s, CapabilityRealtimeOrders, "order", r)
}
func (s *realtimeService) SubscribeBrokerOrders(ctx context.Context, r dto.SubscribeBrokerOrdersRequest) (realtime.Subscription[dto.BrokerOrderEvent], error) {
	if err := s.requireCapability(CapabilityRealtimeBrokerOrders); err != nil {
		return nil, err
	}
	marketType := r.MarketType
	if marketType == "" {
		marketType = "STOCK"
	}
	return nativerealtime.Subscribe(ctx, s.dependencies, buildBrokerOrdersSubscribeMessage(marketType, r.InvestorID, s.dependencies.Encoding), isPayload, func(message map[string]any) (dto.BrokerOrderEvent, error) {
		return decodeEvent[dto.BrokerOrderEvent](messageData(message))
	})
}
func (s *realtimeService) SubscribePositions(ctx context.Context, r dto.SubscribeTradingRequest) (realtime.Subscription[dto.PositionEvent], error) {
	return subscribeTrading[dto.PositionEvent](ctx, s, CapabilityRealtimePositions, "position", r)
}
func subscribeTrading[T any](ctx context.Context, s *realtimeService, capability core.Capability, kind string, r dto.SubscribeTradingRequest) (realtime.Subscription[T], error) {
	if err := s.requireCapability(capability); err != nil {
		return nil, err
	}
	marketType := r.MarketType
	if marketType == "" {
		marketType = "STOCK"
	}
	return nativerealtime.Subscribe(ctx, s.dependencies, buildSubscribeMessage(kind, marketType, s.dependencies.Encoding), isPayload, func(message map[string]any) (T, error) { return decodeEvent[T](messageData(message)) })
}
func buildSubscribeMessage(kind, marketType, encoding string) map[string]any {
	if encoding == "" {
		encoding = "msgpack"
	}
	return map[string]any{"action": "subscribe", "channels": []any{map[string]any{"name": kind + "." + marketType + "." + encoding, "symbols": []any{}}}}
}
func buildBrokerOrdersSubscribeMessage(marketType, investorID, encoding string) map[string]any {
	if encoding == "" {
		encoding = "msgpack"
	}
	channel := "order.broker." + marketType + "." + investorID + "." + encoding
	return map[string]any{"action": "subscribe", "channels": []any{map[string]any{"name": channel, "symbols": []any{}}}}
}
func messageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		message = data
	}
	if _, ok := wrapperPayloadField(message["T"]); ok {
		return message
	}
	if order, ok := message["order"].(map[string]any); ok {
		return order
	}
	if position, ok := message["position"].(map[string]any); ok {
		return position
	}
	return message
}
func isPayload(message map[string]any) bool {
	data := messageData(message)
	if field, wrapped := wrapperPayloadField(data["T"]); wrapped {
		_, ok := data[field].(map[string]any)
		return ok
	}
	return data["accountNo"] != nil && (data["id"] != nil || data["orderId"] != nil || data["symbol"] != nil)
}
func wrapperPayloadField(eventType any) (string, bool) {
	switch eventType {
	case "do", "eo":
		return "order", true
	case "dp", "ep":
		return "position", true
	default:
		return "", false
	}
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
