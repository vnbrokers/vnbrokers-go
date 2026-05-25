package dnse

import "github.com/vnbrokers/vnbrokers-go/core"

var Capabilities = []core.Capability{
	core.CapabilityTradingAccountsList,
	core.CapabilityTradingOrdersPlace,
	core.CapabilityTradingOrdersCancel,
	core.CapabilityTradingPositionsList,
	core.CapabilityTradingRealtimeOrders,
	core.CapabilityTradingRealtimePosition,
	core.CapabilityBrokerageCareBy,
	core.CapabilityMarketDataSymbolsList,
	core.CapabilityMarketDataQuotes,
	core.CapabilityMarketDataCandles,
	core.CapabilityMarketDataRealtimeTicks,
	core.CapabilityMarketDataRealtimeTop,
	core.CapabilityMarketDataRealtimeCandle,
	core.CapabilityMarketDataRealtimeRaw,
}
