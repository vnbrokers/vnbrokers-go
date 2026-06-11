package trading

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
		func() error { _, err := service.SubscribeOrders(ctx, dto.SubscribeTradingRequest{}); return err },
		func() error { _, err := service.SubscribePositions(ctx, dto.SubscribeTradingRequest{}); return err },
	}
	_ = compile
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
