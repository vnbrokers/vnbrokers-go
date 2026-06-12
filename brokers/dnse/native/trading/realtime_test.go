package trading

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	nativerealtime "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/realtime"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestOrderUpdateAcceptsEmptyOrderID(t *testing.T) {
	var message map[string]any
	if err := json.Unmarshal([]byte(`{
		"T":"eo",
		"action":"order_update",
		"event":"canceled",
		"order":{
			"accountNo":"",
			"averagePrice":0,
			"canceledQuantity":1,
			"createdDate":"2026-06-12T03:14:26.158Z",
			"fillQuantity":0,
			"id":"",
			"investorId":"",
			"leaveQuantity":0,
			"loanPackageId":1775,
			"marketType":"STOCK",
			"modifiedDate":"2026-06-12T03:14:26.945Z",
			"orderStatus":"Canceled",
			"orderType":"LO",
			"price":7370,
			"priceSecure":7370,
			"quantity":1,
			"side":"NS",
			"symbol":"AAA",
			"transDate":"2026-06-12T00:00:00Z"
		},
		"sequence":4,
		"timestamp":1781234866949
	}`), &message); err != nil {
		t.Fatal(err)
	}
	if !isPayload(message) {
		t.Fatal("order update was ignored")
	}
	event, err := decodeEvent[dto.OrderEvent](messageData(message))
	if err != nil {
		t.Fatal(err)
	}
	if event.ID != 0 || event.Symbol != "AAA" || event.OrderStatus != "Canceled" {
		t.Fatalf("event = %+v", event)
	}
}

func TestRealtimeServiceExposesTypedChannels(t *testing.T) {
	var service RealtimeService
	ctx := context.Background()
	compile := []func() error{
		func() error { _, err := service.SubscribeOrders(ctx, dto.SubscribeTradingRequest{}); return err },
		func() error {
			_, err := service.SubscribeBrokerOrders(ctx, dto.SubscribeBrokerOrdersRequest{})
			return err
		},
		func() error { _, err := service.SubscribePositions(ctx, dto.SubscribeTradingRequest{}); return err },
	}
	_ = compile
}

func TestBrokerOrdersSubscriptionUsesInvestorChannel(t *testing.T) {
	var required core.Capability
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "json"}, func(capability core.Capability) error {
		required = capability
		return context.Canceled
	})
	_, err := service.SubscribeBrokerOrders(context.Background(), dto.SubscribeBrokerOrdersRequest{
		MarketType: "DERIVATIVE",
		InvestorID: "investor-123",
	})
	if err != context.Canceled {
		t.Fatalf("error = %v", err)
	}
	if required != CapabilityRealtimeBrokerOrders {
		t.Fatalf("capability = %q", required)
	}

	message := buildBrokerOrdersSubscribeMessage("DERIVATIVE", "investor-123", "json")
	channel := message["channels"].([]any)[0].(map[string]any)
	if channel["name"] != "order.broker.DERIVATIVE.investor-123.json" {
		t.Fatalf("channel = %v", channel["name"])
	}
}

func TestNewRealtimeServiceBuildsTradingChannels(t *testing.T) {
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "msgpack"}, func(core.Capability) error { return nil })
	if service == nil {
		t.Fatal("service is nil")
	}
	message := buildSubscribeMessage("order", "STOCK", "")
	channel := message["channels"].([]any)[0].(map[string]any)
	if channel["name"] != "order.STOCK.msgpack" {
		t.Fatalf("channel = %v", channel["name"])
	}
}
