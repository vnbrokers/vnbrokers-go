package ssi

import "github.com/vnbrokers/vnbrokers-go/core"

type Broker struct {
	core.BaseBroker
	config      Config
	accessToken string
	auth        *AuthService
	trading     *TradingService
}

func NewBroker(config Config) *Broker {
	config = config.withDefaults()
	b := &Broker{
		BaseBroker: core.BaseBroker{
			BrokerName:         "ssi",
			BrokerCapabilities: Capabilities,
		},
		config:      config,
		accessToken: config.AccessToken,
	}
	b.auth = &AuthService{broker: b}
	b.trading = &TradingService{
		accounts:  &TradingAccountsService{broker: b},
		orders:    &TradingOrdersService{broker: b},
		positions: &TradingPositionsService{broker: b},
	}
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Trading() *TradingService {
	return b.trading
}

type TradingService struct {
	accounts  *TradingAccountsService
	orders    *TradingOrdersService
	positions *TradingPositionsService
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
