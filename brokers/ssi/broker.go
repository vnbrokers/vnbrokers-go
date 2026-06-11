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
	trading            *TradingService
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
	b.trading = &TradingService{
		accounts:  &TradingAccountsService{broker: b},
		orders:    &TradingOrdersService{broker: b},
		positions: &TradingPositionsService{broker: b},
		realtime:  &TradingRealtimeService{broker: b},
	}
	realtimeMarketData := &MarketDataRealtimeService{broker: b}
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
		}, realtimeMarketData),
		nativetrading.NewService(nativetrading.Dependencies{
			BaseURL:      b.config.BaseURL,
			TradingToken: b.tradingToken,
			RequireCapability: func(capability core.Capability) error {
				return b.RequireCapability(capability)
			},
			Send: func(ctx context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
				return b.send(ctx, operation, true, true, request)
			},
		}, nil),
	)
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Trading() *TradingService {
	return b.trading
}

func (b *Broker) Native() nativeapi.Service {
	return b.native
}

type TradingService struct {
	accounts  *TradingAccountsService
	orders    *TradingOrdersService
	positions *TradingPositionsService
	realtime  *TradingRealtimeService
}

func (s *TradingService) Accounts() *TradingAccountsService {
	return s.accounts
}

func (s *TradingService) Orders() *TradingOrdersService {
	return s.orders
}

func (s *TradingService) Positions() *TradingPositionsService {
	return s.positions
}

func (s *TradingService) Realtime() *TradingRealtimeService {
	return s.realtime
}
