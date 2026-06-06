package ssi

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
	_ trading.SubscribeOrdersRequest,
) (realtime.Subscription[domain.OrderEvent], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimeOrders); err != nil {
		return nil, err
	}
	return startSSISignalRSubscription(
		ctx,
		s.broker,
		s.broker.config.TradingStreamURL,
		ssiTradingHub,
		func(client SignalRClient, subscription *realtime.QueueSubscription[domain.OrderEvent]) {
			client.SetQuery("notify_id", "-1")
			registerSSIBroadcastHandler(
				client,
				subscription,
				[]string{"orderUpdate", "orderMatchEvent", "orderError"},
				func(message ssiRealtimeMessage) domain.OrderEvent {
					if message.DataType == "orderError" {
						return mapSSIOrderErrorEvent(message)
					}
					return mapSSIOrderEvent(message)
				},
			)
		},
		nil,
	)
}

func (s *TradingRealtimeService) SubscribePositions(
	ctx context.Context,
	_ trading.SubscribePositionsRequest,
) (realtime.Subscription[domain.Position], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimePosition); err != nil {
		return nil, err
	}
	return startSSISignalRSubscription(
		ctx,
		s.broker,
		s.broker.config.TradingStreamURL,
		ssiTradingHub,
		func(client SignalRClient, subscription *realtime.QueueSubscription[domain.Position]) {
			client.SetQuery("notify_id", "-1")
			registerSSIBroadcastHandler(client, subscription, []string{"clientPortfolioEvent"}, mapSSIPositionEvent)
		},
		nil,
	)
}

func (s *TradingRealtimeService) SubscribeFCOEvents(
	ctx context.Context,
) (realtime.Subscription[FCOEvent], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return nil, err
	}
	return startSSISignalRSubscription(
		ctx,
		s.broker,
		s.broker.config.TradingStreamURL,
		ssiTradingHub,
		func(client SignalRClient, subscription *realtime.QueueSubscription[FCOEvent]) {
			client.SetQuery("notify_id", "-1")
			registerSSIBroadcastHandler(client, subscription, []string{"fcoEvent"}, mapSSIFCOEvent)
		},
		nil,
	)
}

func (s *TradingRealtimeService) SubscribeConditionalOrders(
	ctx context.Context,
) (realtime.Subscription[FCOEvent], error) {
	return s.SubscribeFCOEvents(ctx)
}

func mapSSIOrderEvent(message ssiRealtimeMessage) domain.OrderEvent {
	status := ssiString(message.Data, "OrderStatus", "orderStatus", "Status", "status")
	return domain.OrderEvent{
		Broker:         "ssi",
		AccountID:      ssiString(message.Data, "Account", "AccountNo", "account"),
		OrderID:        ssiString(message.Data, "OrderID", "OrderId", "orderID", "orderId"),
		Symbol:         ssiString(message.Data, "InstrumentID", "InstrumentId", "Symbol", "symbol"),
		Status:         MapOrderStatus(status),
		RawStatus:      status,
		FilledQuantity: ssiString(message.Data, "FilledQty", "FilledQuantity", "filledQty", "matchedQuantity"),
		ReceivedAt:     ssiString(message.Data, "ModifiedTime", "InputTime", "Time", "time"),
		Raw:            ssiRawPayload(message),
	}
}

func mapSSIOrderErrorEvent(message ssiRealtimeMessage) domain.OrderEvent {
	event := mapSSIOrderEvent(message)
	if event.Status == domain.OrderStatusUnknown {
		event.Status = domain.OrderStatusRejected
	}
	return event
}

func mapSSIPositionEvent(message ssiRealtimeMessage) domain.Position {
	quantity := decimalFrom(ssiValue(message.Data, "OnHand", "Quantity", "quantity"))
	available := decimalFrom(ssiValue(message.Data, "SellableQty", "AvailableQuantity", "availableQuantity"))
	marketPrice := optionalDecimal(ssiValue(message.Data, "MarketPrice", "marketPrice"))
	var marketValue = marketPrice
	if marketPrice != nil {
		value := quantity.Mul(*marketPrice)
		marketValue = &value
	}
	return domain.Position{
		AccountID:         ssiString(message.Data, "Account", "AccountNo", "account"),
		Symbol:            ssiString(message.Data, "InstrumentID", "InstrumentId", "Symbol", "symbol"),
		Quantity:          quantity,
		AvailableQuantity: available,
		AveragePrice:      optionalDecimal(ssiValue(message.Data, "AvgPrice", "AveragePrice", "avgPrice")),
		MarketValue:       marketValue,
		Raw:               ssiRawPayload(message),
	}
}

func mapSSIFCOEvent(message ssiRealtimeMessage) FCOEvent {
	return FCOEvent{
		FCOID:           ssiString(message.Data, "fcoId"),
		NotifyID:        ssiInt64(message.Data, "notifyID"),
		Data:            ssiValue(message.Data, "data"),
		ProcessStatus:   ssiString(message.Data, "processStatus"),
		LastAction:      ssiString(message.Data, "lastAction"),
		UniqueID:        ssiString(message.Data, "uniqueID"),
		MatchedQuantity: decimalFrom(ssiValue(message.Data, "matchedQuantity")),
		IsPlaceOrder:    ssiBool(message.Data, "isPlaceOrder"),
		IPAddress:       ssiString(message.Data, "ipAddress"),
		Symbol:          ssiString(message.Data, "instrumentID"),
		Prefix:          ssiString(message.Data, "prefix"),
		Quantity:        decimalFrom(ssiValue(message.Data, "quantity")),
		BrokerID:        ssiString(message.Data, "brokerId"),
		Price:           ssiString(message.Data, "price"),
		AccountID:       ssiString(message.Data, "account"),
		BrokerIDUpdate:  ssiString(message.Data, "brokerIdUpdate"),
		UpdatedTime:     ssiString(message.Data, "updatedTime"),
		Status:          ssiString(message.Data, "status"),
		Message:         ssiString(message.Data, "message"),
		Username:        ssiString(message.Data, "username"),
		Raw:             ssiRawPayload(message),
	}
}
