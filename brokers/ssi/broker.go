package ssi

import "github.com/vnbrokers/vnbrokers-go/core"

type Broker struct {
	core.BaseBroker
	config      Config
	accessToken string
	auth        *AuthService
	trading     *TradingService
	marketData  *MarketDataService
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
		accounts:          &TradingAccountsService{broker: b},
		orders:            &TradingOrdersService{broker: b},
		positions:         &TradingPositionsService{broker: b},
		cash:              &TradingCashService{broker: b},
		stockTransfers:    &TradingStockTransferService{broker: b},
		rights:            &TradingRightsService{broker: b},
		conditionalOrders: &TradingConditionalOrdersService{broker: b},
		realtime:          &TradingRealtimeService{broker: b},
	}
	b.marketData = &MarketDataService{
		realtime: &MarketDataRealtimeService{broker: b},
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

type TradingService struct {
	accounts          *TradingAccountsService
	orders            *TradingOrdersService
	positions         *TradingPositionsService
	cash              *TradingCashService
	stockTransfers    *TradingStockTransferService
	rights            *TradingRightsService
	conditionalOrders *TradingConditionalOrdersService
	realtime          *TradingRealtimeService
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

func (s *TradingService) Cash() *TradingCashService {
	return s.cash
}

func (s *TradingService) StockTransfers() *TradingStockTransferService {
	return s.stockTransfers
}

func (s *TradingService) Rights() *TradingRightsService {
	return s.rights
}

func (s *TradingService) ConditionalOrders() *TradingConditionalOrdersService {
	return s.conditionalOrders
}

func (s *TradingService) Realtime() *TradingRealtimeService {
	return s.realtime
}

type MarketDataService struct {
	realtime *MarketDataRealtimeService
}

func (s *MarketDataService) Realtime() *MarketDataRealtimeService {
	return s.realtime
}
