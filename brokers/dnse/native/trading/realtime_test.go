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
	if event.T != "eo" || event.Action != "order_update" || event.Event != "canceled" || event.Sequence != 4 || event.Timestamp != 1781234866949 {
		t.Fatalf("event = %+v", event)
	}
	if event.Order.ID != 0 || event.Order.Symbol != "AAA" || event.Order.OrderStatus != "Canceled" {
		t.Fatalf("event = %+v", event)
	}
}

func TestPositionUpdatePreservesWrapper(t *testing.T) {
	var message map[string]any
	if err := json.Unmarshal([]byte(`{
		"T":"dp",
		"action":"position_update",
		"event":"pendingcancel",
		"position":{"id":177796763592657,"accountNo":"0001179019","symbol":"41I1G5000","status":"OPEN"},
		"sequence":1,
		"timestamp":1776054245274
	}`), &message); err != nil {
		t.Fatal(err)
	}
	if !isPayload(message) {
		t.Fatal("position update was ignored")
	}
	event, err := decodeEvent[dto.PositionEvent](messageData(message))
	if err != nil {
		t.Fatal(err)
	}
	if event.T != "dp" || event.Action != "position_update" || event.Position.ID != 177796763592657 {
		t.Fatalf("event = %+v", event)
	}
}

func TestTradingWrapperDiscriminatorAliases(t *testing.T) {
	tests := []struct {
		name  string
		input string
		field string
	}{
		{name: "do order", input: `{"T":"do","order":{"id":596}}`, field: "order"},
		{name: "eo order", input: `{"T":"eo","order":{"id":596}}`, field: "order"},
		{name: "dp position", input: `{"T":"dp","position":{"id":177796763592657}}`, field: "position"},
		{name: "ep position", input: `{"T":"ep","position":{"id":177796763592657}}`, field: "position"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var message map[string]any
			if err := json.Unmarshal([]byte(tt.input), &message); err != nil {
				t.Fatal(err)
			}
			if !isPayload(message) {
				t.Fatal("wrapper was ignored")
			}
			data := messageData(message)
			if data["T"] != message["T"] || data[tt.field] == nil {
				t.Fatalf("data = %+v", data)
			}
		})
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
		func() error {
			_, err := service.SubscribeBrokerPositions(ctx, dto.SubscribeBrokerPositionsRequest{})
			return err
		},
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

func TestBrokerPositionsSubscriptionUsesInvestorChannel(t *testing.T) {
	var required core.Capability
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "json"}, func(capability core.Capability) error {
		required = capability
		return context.Canceled
	})
	_, err := service.SubscribeBrokerPositions(context.Background(), dto.SubscribeBrokerPositionsRequest{
		MarketType: "STOCK",
		InvestorID: "1000009250",
	})
	if err != context.Canceled {
		t.Fatalf("error = %v", err)
	}
	if required != CapabilityRealtimeBrokerPositions {
		t.Fatalf("capability = %q", required)
	}

	message := buildBrokerPositionsSubscribeMessage("STOCK", "1000009250", "json")
	channel := message["channels"].([]any)[0].(map[string]any)
	if channel["name"] != "position.broker.STOCK.1000009250.json" {
		t.Fatalf("channel = %v", channel["name"])
	}
}

func TestBrokerPositionEventDecodesDirectPayloadIntoWrapper(t *testing.T) {
	var message map[string]any
	if err := json.Unmarshal([]byte(`{
		"id": 177796763592657,
		"accountNo": "0001179019",
		"symbol": "41I1G5000",
		"status": "OPEN",
		"loanPackageId": 2278,
		"side": "NB",
		"accumulateQuantity": 247,
		"tradeQuantity": null,
		"closedQuantity": 236,
		"costPrice": 2057.72425,
		"marketPrice": 2070.0,
		"breakEvenPrice": 2058.21911,
		"openQuantity": 11,
		"overNightQuantity": 0,
		"averageClosePrice": 2094.28941,
		"marketType": "DERIVATIVE",
		"createdDate": "2026-05-05T09:17:50.457893Z",
		"modifiedDate": "2026-05-07T04:19:20.901188117Z"
	}`), &message); err != nil {
		t.Fatal(err)
	}
	if !isPayload(message) {
		t.Fatal("broker position payload was ignored")
	}
	event, err := decodeBrokerPositionEvent(message)
	if err != nil {
		t.Fatal(err)
	}
	if event.Position.ID != 177796763592657 || event.Position.AccountNo != "0001179019" || event.Position.Symbol != "41I1G5000" {
		t.Fatalf("event = %+v", event)
	}
	if event.Position.CostPrice == nil || event.Position.CostPrice.String() != "2057.72425" {
		t.Fatalf("cost price = %v", event.Position.CostPrice)
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
