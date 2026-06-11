package dnse

import (
	"context"

	nativeapi "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native"
	nativebrokerage "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/brokerage"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Broker struct {
	core.BaseBroker
	config Config
	auth   *AuthService
	native nativeapi.Service
}

func NewBroker(config Config) *Broker {
	config = config.withDefaults()
	b := &Broker{
		BaseBroker: core.BaseBroker{
			BrokerName:         "dnse",
			BrokerCapabilities: Capabilities,
		},
		config: config,
	}
	b.auth = &AuthService{broker: b}
	b.native = nativeapi.NewService(
		nativemarketdata.NewService(nativemarketdata.Dependencies{
			BaseURL:           b.config.BaseURL,
			Headers:           b.apiHeaders,
			RequireCapability: b.RequireCapability,
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, request)
			},
		}, &nativeMarketDataRealtimeService{broker: b}),
		nativetrading.NewService(nativetrading.Dependencies{
			BaseURL:           b.config.BaseURL,
			APIHeaders:        func(bool) map[string]string { return b.apiHeaders() },
			TradingHeaders:    b.tradingHeaders,
			RequireCapability: b.RequireCapability,
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, request)
			},
		}, &nativeTradingRealtimeService{broker: b}),
		nativebrokerage.NewService(nativebrokerage.Dependencies{
			BaseURL:           b.config.BaseURL,
			Headers:           b.apiHeaders,
			RequireCapability: b.RequireCapability,
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, request)
			},
		}),
	)
	return b
}

func (b *Broker) Native() nativeapi.Service {
	return b.native
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}
