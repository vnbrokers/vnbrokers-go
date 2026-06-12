package entrade

import (
	"context"

	nativeapi "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
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
			BrokerName:         "entrade",
			BrokerCapabilities: Capabilities,
		},
		config: config,
	}
	b.auth = &AuthService{broker: b}
	dependencies := func() (nativetrading.Dependencies, nativemarketdata.Dependencies) {
		headers := func(body bool) map[string]string { return b.headers(true, body) }
		send := func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
			return b.send(ctx, operation, false, request)
		}
		return nativetrading.Dependencies{
				BaseURL: b.config.BaseURL, Headers: headers,
				RequireCapability: b.RequireCapability, Send: send,
			}, nativemarketdata.Dependencies{
				BaseURL: b.config.BaseURL, Headers: headers,
				RequireCapability: b.RequireCapability, Send: send,
			}
	}
	tradingDependencies, marketDataDependencies := dependencies()
	b.native = nativeapi.NewService(
		nativetrading.NewService(tradingDependencies),
		nativemarketdata.NewService(marketDataDependencies),
	)
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Native() nativeapi.Service {
	return b.native
}
