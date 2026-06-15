package tcbs

import (
	"context"

	nativeapi "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Broker struct {
	core.BaseBroker
	config      Config
	accessToken string
	auth        *AuthService
	native      nativeapi.Service
}

func NewBroker(config Config) *Broker {
	config = config.withDefaults()
	b := &Broker{
		BaseBroker: core.BaseBroker{
			BrokerName:         "tcbs",
			BrokerCapabilities: Capabilities,
		},
		config:      config,
		accessToken: config.AccessToken,
	}
	b.auth = &AuthService{broker: b}
	send := func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
		return b.send(ctx, operation, true, request)
	}
	realtimeDependencies := nativetrading.RealtimeDependencies{
		BaseURL: b.config.BaseURL, AccessToken: func() string { return b.accessToken }, Headers: b.headers,
		RequireCapability: b.RequireCapability, WebSocketFactory: b.config.WebSocketFactory,
	}
	b.native = nativeapi.NewService(
		nativetrading.NewService(nativetrading.Dependencies{
			BaseURL: b.config.BaseURL, Headers: b.headers, RequireCapability: b.RequireCapability, Send: send,
		}, nativetrading.NewRealtimeService(realtimeDependencies)),
		nativemarketdata.NewService(nativemarketdata.Dependencies{
			BaseURL: b.config.BaseURL, Headers: b.headers, RequireCapability: b.RequireCapability, Send: send,
		}, nativemarketdata.NewRealtimeService(nativemarketdata.RealtimeDependencies{
			BaseURL: b.config.BaseURL, AccessToken: func() string { return b.accessToken }, Headers: b.headers,
			RequireCapability: b.RequireCapability, WebSocketFactory: b.config.WebSocketFactory,
		})),
	)
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Native() nativeapi.Service {
	return b.native
}
