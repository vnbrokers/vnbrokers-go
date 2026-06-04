package entrade

import "github.com/vnbrokers/vnbrokers-go/core"

type Broker struct {
	core.BaseBroker
	config     Config
	auth       *AuthService
	trading    *TradingService
	marketData *MarketDataService
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
	b.trading = &TradingService{
		accounts: &TradingAccountsService{broker: b},
		orders:   &TradingOrdersService{broker: b},
		deals:    &TradingDealsService{broker: b},
		risk:     &TradingRiskService{broker: b},
	}
	b.marketData = &MarketDataService{
		derivatives: &MarketDataDerivativesService{broker: b},
	}
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Trading() *TradingService {
	return b.trading
}

func (b *Broker) MarketData() *MarketDataService {
	return b.marketData
}
