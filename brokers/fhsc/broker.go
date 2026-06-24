package fhsc

import (
	"context"

	nativeapi "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/trading"
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
			BrokerName:         "fhsc",
			BrokerCapabilities: Capabilities,
		},
		config: config,
	}
	b.auth = &AuthService{broker: b}
	send := func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
		return b.send(ctx, operation, true, request)
	}
	b.native = nativeapi.NewService(
		nativemarketdata.NewService(nativemarketdata.Dependencies{
			BaseURL:           b.config.BaseURL,
			Headers:           b.headers,
			RequireCapability: b.RequireCapability,
			Send:              send,
		}),
		nativetrading.NewService(nativetrading.Dependencies{
			BaseURL:           b.config.BaseURL,
			Headers:           b.headers,
			RequireCapability: b.RequireCapability,
			Send:              send,
		}),
	)
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Native() nativeapi.Service {
	return b.native
}
