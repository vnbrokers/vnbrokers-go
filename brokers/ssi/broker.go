package ssi

import (
	"context"
	"sync"

	nativeapi "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Broker struct {
	core.BaseBroker
	config             Config
	tokenMu            sync.RWMutex
	dataAccessToken    string
	tradingAccessToken string
	auth               *AuthService
	native             nativeapi.Service
}

func (b *Broker) dataToken() string {
	b.tokenMu.RLock()
	defer b.tokenMu.RUnlock()
	return b.dataAccessToken
}

func (b *Broker) setDataToken(token string) {
	b.tokenMu.Lock()
	defer b.tokenMu.Unlock()
	b.dataAccessToken = token
}

func (b *Broker) tradingToken() string {
	b.tokenMu.RLock()
	defer b.tokenMu.RUnlock()
	return b.tradingAccessToken
}

func (b *Broker) setTradingToken(token string) {
	b.tokenMu.Lock()
	defer b.tokenMu.Unlock()
	b.tradingAccessToken = token
}

func NewBroker(config Config) *Broker {
	config = config.withDefaults()
	b := &Broker{
		BaseBroker: core.BaseBroker{
			BrokerName:         "ssi",
			BrokerCapabilities: Capabilities,
		},
		config:             config,
		dataAccessToken:    config.DataToken,
		tradingAccessToken: config.TradingToken,
	}
	b.auth = &AuthService{broker: b}
	b.native = nativeapi.NewService(
		nativemarketdata.NewService(nativemarketdata.Dependencies{
			BaseURL:   b.config.DataBaseURL,
			DataToken: b.dataToken,
			RequireCapability: func(capability core.Capability) error {
				return b.RequireCapability(capability)
			},
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, false, false, request)
			},
		}, nativemarketdata.NewRealtimeService(nativemarketdata.RealtimeDependencies{
			DataToken:           b.dataToken,
			MarketDataStreamURL: b.config.MarketDataStreamURL,
			RequireCapability: func(capability core.Capability) error {
				return b.RequireCapability(capability)
			},
			NewSignalRClient: func(baseURL string, hubs []string) nativemarketdata.SignalRClient {
				return b.config.SignalRFactory(baseURL, hubs)
			},
		})),
		nativetrading.NewService(nativetrading.Dependencies{
			BaseURL:      b.config.BaseURL,
			TradingToken: b.tradingToken,
			RequireCapability: func(capability core.Capability) error {
				return b.RequireCapability(capability)
			},
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, true, true, request)
			},
		}, nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
			TradingToken:      b.tradingToken,
			TradingStreamURL:  b.config.TradingStreamURL,
			RequireCapability: func(capability core.Capability) error {
				return b.RequireCapability(capability)
			},
			NewSignalRClient: func(baseURL string, hubs []string) nativetrading.SignalRClient {
				return b.config.SignalRFactory(baseURL, hubs)
			},
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
