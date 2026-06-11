package trading

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
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
