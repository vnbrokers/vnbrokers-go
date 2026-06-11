package marketdata

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
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
