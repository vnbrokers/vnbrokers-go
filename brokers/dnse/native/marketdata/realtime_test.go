package marketdata

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	nativerealtime "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/realtime"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestEstimatedMarketIndexUpdatePreservesWrapper(t *testing.T) {
	var message map[string]any
	if err := json.Unmarshal([]byte(`{
		"T":"emi",
		"action":"estimated_market_index_update",
		"marketIndex":{
			"changedRatio":0.98,
			"changedValue":19.07,
			"fluctuationDownIssueCount":3,
			"fluctuationSteadinessIssueCount":0,
			"fluctuationUpIssueCount":27,
			"grossTradeAmount":2292.3,
			"indexName":"VN30",
			"time":"2026-06-12 10:19:54.396",
			"totalVolumeTraded":82215100,
			"valueIndexes":1966.35
		},
		"timestamp":1781234394402
	}`), &message); err != nil {
		t.Fatal(err)
	}
	if !isPayload(message) {
		t.Fatal("estimated market index update was ignored")
	}
	event, err := decodeEvent[dto.EstimatedMarketIndexEvent](messageData(message))
	if err != nil {
		t.Fatal(err)
	}
	if event.T != "emi" || event.Action != "estimated_market_index_update" || event.Timestamp != 1781234394402 {
		t.Fatalf("event = %+v", event)
	}
	if event.MarketIndex.IndexName != "VN30" || event.MarketIndex.ValueIndexes == nil || event.MarketIndex.ValueIndexes.String() != "1966.35" {
		t.Fatalf("event = %+v", event)
	}
}

func TestRealtimeServiceExposesTypedChannels(t *testing.T) {
	var service RealtimeService
	ctx := context.Background()
	compile := []func() error{
		func() error {
			_, err := service.SubscribeExpectedPrices(ctx, dto.SubscribeSymbolsRequest{})
			return err
		},
		func() error { _, err := service.SubscribeForeign(ctx, dto.SubscribeSymbolsRequest{}); return err },
		func() error {
			_, err := service.SubscribeMarketIndexes(ctx, dto.SubscribeMarketIndexRequest{})
			return err
		},
		func() error {
			_, err := service.SubscribeEstimatedMarketIndexes(ctx, dto.SubscribeMarketIndexRequest{})
			return err
		},
		func() error { _, err := service.SubscribeOHLC(ctx, dto.SubscribeOHLCRequest{}); return err },
		func() error { _, err := service.SubscribeClosedOHLC(ctx, dto.SubscribeOHLCRequest{}); return err },
		func() error { _, err := service.SubscribeQuotes(ctx, dto.SubscribeSymbolsRequest{}); return err },
		func() error {
			_, err := service.SubscribeSecurityDefinitions(ctx, dto.SubscribeSymbolsRequest{})
			return err
		},
		func() error { _, err := service.SubscribeTrades(ctx, dto.SubscribeSymbolsRequest{}); return err },
		func() error { _, err := service.SubscribeTradeExtras(ctx, dto.SubscribeSymbolsRequest{}); return err },
	}
	_ = compile
}

func TestEstimatedMarketIndexSubscriptionUsesNativeChannel(t *testing.T) {
	var required core.Capability
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "json"}, func(capability core.Capability) error {
		required = capability
		return context.Canceled
	})
	_, err := service.SubscribeEstimatedMarketIndexes(context.Background(), dto.SubscribeMarketIndexRequest{IndexName: "VN30"})
	if err != context.Canceled {
		t.Fatalf("error = %v", err)
	}
	if required != CapabilityRealtimeEstimatedMarketIndexes {
		t.Fatalf("capability = %q", required)
	}

	channel := buildChannel("estimated_market_index", "", "", "VN30", "json")
	if channel != "estimated_market_index.VN30.json" {
		t.Fatalf("channel = %q", channel)
	}
	message := buildSubscribeMessage(channel, nil)
	item := message["channels"].([]any)[0].(map[string]any)
	if item["name"] != channel {
		t.Fatalf("channel name = %v", item["name"])
	}
	if _, ok := item["symbols"]; ok {
		t.Fatalf("unexpected symbols = %v", item["symbols"])
	}
	if !isPayload(map[string]any{"indexName": "VN30", "valueIndexes": 1948.57}) {
		t.Fatal("estimated market index payload was ignored")
	}
}

func TestNewRealtimeServiceBuildsMarketDataChannels(t *testing.T) {
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "msgpack"}, func(core.Capability) error { return nil })
	if service == nil {
		t.Fatal("service is nil")
	}
	if channel := buildChannel("top_price", "G1", "", "", ""); channel != "top_price.G1.msgpack" {
		t.Fatalf("channel = %q", channel)
	}
}
