package core

type Capability string

const (
	CapabilityTradingAccountsList      Capability = "trading.accounts.list"
	CapabilityTradingOrdersPlace       Capability = "trading.orders.place"
	CapabilityTradingOrdersCancel      Capability = "trading.orders.cancel"
	CapabilityTradingPositionsList     Capability = "trading.positions.list"
	CapabilityTradingRealtimeOrders    Capability = "trading.realtime.orders"
	CapabilityTradingRealtimePosition  Capability = "trading.realtime.positions"
	CapabilityBrokerageCareBy          Capability = "brokerage.accounts.care_by"
	CapabilityMarketDataSymbolsList    Capability = "marketdata.symbols.list"
	CapabilityMarketDataQuotes         Capability = "marketdata.quotes.snapshot"
	CapabilityMarketDataCandles        Capability = "marketdata.candles.history"
	CapabilityMarketDataRealtimeTicks  Capability = "marketdata.realtime.ticks"
	CapabilityMarketDataRealtimeTop    Capability = "marketdata.realtime.top_price"
	CapabilityMarketDataRealtimeCandle Capability = "marketdata.realtime.candles"
	CapabilityMarketDataRealtimeRaw    Capability = "marketdata.realtime.raw"
)
