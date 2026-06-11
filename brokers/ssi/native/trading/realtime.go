package trading

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
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
	Raw      []byte
}

func (s *realtimeService) SubscribeOrders(ctx context.Context, _ sdktrading.SubscribeOrdersRequest) (realtime.Subscription[domain.OrderEvent], error) {
	if err := s.deps.RequireCapability(core.CapabilityTradingRealtimeOrders); err != nil {
		return nil, err
	}
	return startTradingSubscription(ctx, s.deps,
		func(client SignalRClient, subscription *realtime.QueueSubscription[domain.OrderEvent]) {
			client.SetQuery("notify_id", "-1")
			registerBroadcastHandler(client, subscription, []string{"orderUpdate", "orderMatchEvent", "orderError"},
				func(message realtimeMessage) domain.OrderEvent {
					if message.DataType == "orderError" {
						return mapOrderErrorEvent(message)
					}
					return mapOrderEvent(message)
				},
			)
		},
	)
}

func (s *realtimeService) SubscribePositions(ctx context.Context, _ sdktrading.SubscribePositionsRequest) (realtime.Subscription[domain.Position], error) {
	if err := s.deps.RequireCapability(core.CapabilityTradingRealtimePosition); err != nil {
		return nil, err
	}
	return startTradingSubscription(ctx, s.deps,
		func(client SignalRClient, subscription *realtime.QueueSubscription[domain.Position]) {
			client.SetQuery("notify_id", "-1")
			registerBroadcastHandler(client, subscription, []string{"clientPortfolioEvent"}, mapPositionEvent)
		},
	)
}

func (s *realtimeService) SubscribeFCOEvents(ctx context.Context) (realtime.Subscription[dto.FCOEvent], error) {
	if err := s.deps.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return nil, err
	}
	return startTradingSubscription(ctx, s.deps,
		func(client SignalRClient, subscription *realtime.QueueSubscription[dto.FCOEvent]) {
			client.SetQuery("notify_id", "-1")
			registerBroadcastHandler(client, subscription, []string{"fcoEvent"}, mapFCOEvent)
		},
	)
}

func (s *realtimeService) SubscribeConditionalOrders(ctx context.Context) (realtime.Subscription[dto.FCOEvent], error) {
	return s.SubscribeFCOEvents(ctx)
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
	return subscription, nil
}

func registerBroadcastHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	eventTypes []string,
	mapEvent func(realtimeMessage) T,
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
			subscription.PublishEvent(mapEvent(message))
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

func containsFold(values []string, target string) bool {
	for _, value := range values {
		if strings.EqualFold(value, target) {
			return true
		}
	}
	return false
}

// ── Value extractors ──

func ssiValue(data map[string]any, keys ...string) any {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			return value
		}
	}
	for existing, value := range data {
		for _, key := range keys {
			if strings.EqualFold(existing, key) {
				return value
			}
		}
	}
	return nil
}

func stringify(value any) string {
	if value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	case fmt.Stringer:
		return typed.String()
	}
	return fmt.Sprint(value)
}

func ssiString(data map[string]any, keys ...string) string {
	return stringify(ssiValue(data, keys...))
}

func ssiInt64(data map[string]any, keys ...string) int64 {
	value := ssiValue(data, keys...)
	switch typed := value.(type) {
	case float64:
		return int64(typed)
	case int64:
		return typed
	case int:
		return int64(typed)
	case json.Number:
		out, _ := typed.Int64()
		return out
	default:
		out, _ := strconv.ParseInt(fmt.Sprint(value), 10, 64)
		return out
	}
}

func ssiBool(data map[string]any, keys ...string) bool {
	value := ssiValue(data, keys...)
	if typed, ok := value.(bool); ok {
		return typed
	}
	out, _ := strconv.ParseBool(fmt.Sprint(value))
	return out
}

func ssiRawPayload(message realtimeMessage) domain.RawPayload {
	return domain.RawPayload{Source: "ssi", Data: message.Data, Bytes: message.Raw}
}

func decimalFrom(value any) decimal.Decimal {
	switch typed := value.(type) {
	case float64:
		return decimal.NewFromFloat(typed)
	case int64:
		return decimal.NewFromInt(typed)
	case int:
		return decimal.NewFromInt(int64(typed))
	case string:
		out, _ := decimal.NewFromString(typed)
		return out
	case json.Number:
		out, _ := decimal.NewFromString(typed.String())
		return out
	}
	return decimal.Zero
}

func optionalDecimal(value any) *decimal.Decimal {
	if value == nil {
		return nil
	}
	out := decimalFrom(value)
	return &out
}

// ── Event mappers ──

func mapOrderEvent(message realtimeMessage) domain.OrderEvent {
	status := ssiString(message.Data, "OrderStatus", "orderStatus", "Status", "status")
	return domain.OrderEvent{
		Broker:         "ssi",
		AccountID:      ssiString(message.Data, "Account", "AccountNo", "account"),
		OrderID:        ssiString(message.Data, "OrderID", "OrderId", "orderID", "orderId"),
		Symbol:         ssiString(message.Data, "InstrumentID", "InstrumentId", "Symbol", "symbol"),
		Status:         mapOrderStatus(status),
		RawStatus:      status,
		FilledQuantity: ssiString(message.Data, "FilledQty", "FilledQuantity", "filledQty", "matchedQuantity"),
		ReceivedAt:     ssiString(message.Data, "ModifiedTime", "InputTime", "Time", "time"),
		Raw:            ssiRawPayload(message),
	}
}

func mapOrderErrorEvent(message realtimeMessage) domain.OrderEvent {
	event := mapOrderEvent(message)
	if event.Status == domain.OrderStatusUnknown {
		event.Status = domain.OrderStatusRejected
	}
	return event
}

func mapPositionEvent(message realtimeMessage) domain.Position {
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

func mapFCOEvent(message realtimeMessage) dto.FCOEvent {
	return dto.FCOEvent{
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

func mapOrderStatus(raw string) domain.OrderStatus {
	normalized := strings.NewReplacer("_", "", "-", "", " ", "").Replace(strings.ToUpper(raw))
	switch normalized {
	case "QU", "QUEUE", "PENDING":
		return domain.OrderStatusPending
	case "RS", "ACCEPTED", "NEW":
		return domain.OrderStatusAccepted
	case "PF", "PARTIALLYFILLED":
		return domain.OrderStatusPartiallyFilled
	case "FF", "FILLED", "FULLFILLED":
		return domain.OrderStatusFilled
	case "PC", "PENDINGCANCEL":
		return domain.OrderStatusPendingCancel
	case "CA", "CANCELLED", "CANCELED":
		return domain.OrderStatusCancelled
	case "RJ", "REJECTED":
		return domain.OrderStatusRejected
	default:
		return domain.OrderStatusUnknown
	}
}
