package marketdata

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	nativerealtime "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/realtime"
	"github.com/vnbrokers/vnbrokers-go/core"
)

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

func TestNewRealtimeServiceBuildsMarketDataChannels(t *testing.T) {
	service := NewRealtimeService(nativerealtime.Dependencies{Encoding: "msgpack"}, func(core.Capability) error { return nil })
	if service == nil {
		t.Fatal("service is nil")
	}
	if channel := buildChannel("top_price", "G1", "", "", ""); channel != "top_price.G1.msgpack" {
		t.Fatalf("channel = %q", channel)
	}
}
