package tcbs

import "github.com/vnbrokers/vnbrokers-go/core"

type Broker struct {
	core.BaseBroker
	config      Config
	accessToken string
	auth        *AuthService
	account     *AccountService
	trading     *TradingService
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
	b.account = &AccountService{broker: b}
	b.trading = &TradingService{
		accounts: &TradingAccountsService{broker: b},
		orders:   &TradingOrdersService{broker: b},
		realtime: &TradingRealtimeService{broker: b},
	}
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Account() *AccountService {
	return b.account
}

func (b *Broker) Trading() *TradingService {
	return b.trading
}
