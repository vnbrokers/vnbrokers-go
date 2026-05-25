package dnse

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/trading"
)

type TradingRealtimeService struct {
	broker *Broker
}

func (s *TradingRealtimeService) SubscribeOrders(
	ctx context.Context,
	request trading.SubscribeOrdersRequest,
) (realtime.Subscription[domain.OrderEvent], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimeOrders); err != nil {
		return nil, err
	}
	marketType := request.MarketType
	if marketType == "" {
		marketType = "STOCK"
	}
	return startRealtimeSubscription(
		ctx,
		s.broker,
		BuildStreamSubscribeOrdersMessage(marketType, s.broker.config.StreamEncoding),
		isTradingStreamPayload,
		func(subscription *realtime.QueueSubscription[domain.OrderEvent], message map[string]any) {
			subscription.PublishEvent(MapOrderEvent(message))
		},
	)
}

func (s *TradingRealtimeService) SubscribePositions(
	ctx context.Context,
	request trading.SubscribePositionsRequest,
) (realtime.Subscription[domain.Position], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimePosition); err != nil {
		return nil, err
	}
	marketType := request.MarketType
	if marketType == "" {
		marketType = "STOCK"
	}
	return startRealtimeSubscription(
		ctx,
		s.broker,
		BuildStreamSubscribePositionsMessage(marketType, s.broker.config.StreamEncoding),
		isTradingStreamPayload,
		func(subscription *realtime.QueueSubscription[domain.Position], message map[string]any) {
			subscription.PublishEvent(MapPosition(tradingMessageData(message)))
		},
	)
}

func BuildStreamSubscribeOrdersMessage(marketType string, encoding string) map[string]any {
	if encoding == "" {
		encoding = "json"
	}
	return map[string]any{
		"action": "subscribe",
		"channels": []any{
			map[string]any{"name": "order." + marketType + "." + encoding, "symbols": []any{}},
		},
	}
}

func BuildStreamSubscribePositionsMessage(marketType string, encoding string) map[string]any {
	if encoding == "" {
		encoding = "json"
	}
	return map[string]any{
		"action": "subscribe",
		"channels": []any{
			map[string]any{"name": "position." + marketType + "." + encoding, "symbols": []any{}},
		},
	}
}

func tradingMessageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		return data
	}
	if message["T"] == "eo" {
		if order, ok := message["order"].(map[string]any); ok {
			return order
		}
	}
	if message["T"] == "ep" {
		if position, ok := message["position"].(map[string]any); ok {
			return position
		}
	}
	if order, ok := message["order"].(map[string]any); ok {
		return order
	}
	if position, ok := message["position"].(map[string]any); ok {
		return position
	}
	return message
}

func isTradingStreamPayload(message map[string]any) bool {
	if message["T"] == "eo" {
		_, ok := message["order"].(map[string]any)
		return ok
	}
	if message["T"] == "ep" {
		_, ok := message["position"].(map[string]any)
		return ok
	}
	data := tradingMessageData(message)
	return data["accountNo"] != nil && (data["id"] != nil || data["orderId"] != nil || data["symbol"] != nil)
}
