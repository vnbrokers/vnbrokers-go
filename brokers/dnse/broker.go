package dnse

import "github.com/vnbrokers/vnbrokers-go/core"

type Broker struct {
	core.BaseBroker
	config     Config
	auth       *AuthService
	brokerage  *BrokerageService
	trading    *TradingService
	marketData *MarketDataService
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
	b.brokerage = &BrokerageService{broker: b}
	b.trading = &TradingService{
		accounts:  &TradingAccountsService{broker: b},
		orders:    &TradingOrdersService{broker: b},
		positions: &TradingPositionsService{broker: b},
		realtime:  &TradingRealtimeService{broker: b},
	}
	b.marketData = &MarketDataService{
		symbols:  &MarketDataSymbolsService{broker: b},
		quotes:   &MarketDataQuotesService{broker: b},
		candles:  &MarketDataCandlesService{broker: b},
		realtime: &MarketDataRealtimeService{broker: b},
	}
	return b
}

func (b *Broker) Auth() *AuthService {
	return b.auth
}

func (b *Broker) Brokerage() *BrokerageService {
	return b.brokerage
}

func (b *Broker) Trading() *TradingService {
	return b.trading
}

func (b *Broker) MarketData() *MarketDataService {
	return b.marketData
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

type MarketDataService struct {
	symbols  *MarketDataSymbolsService
	quotes   *MarketDataQuotesService
	candles  *MarketDataCandlesService
	realtime *MarketDataRealtimeService
}

func (s *MarketDataService) Symbols() *MarketDataSymbolsService {
	return s.symbols
}

func (s *MarketDataService) Quotes() *MarketDataQuotesService {
	return s.quotes
}

func (s *MarketDataService) Candles() *MarketDataCandlesService {
	return s.candles
}

func (s *MarketDataService) Realtime() *MarketDataRealtimeService {
	return s.realtime
}
