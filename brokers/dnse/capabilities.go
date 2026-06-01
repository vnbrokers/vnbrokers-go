package dnse

import "github.com/vnbrokers/vnbrokers-go/core"

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthSendOTP,
	core.CapabilityTradingAuthTradingToken,
	core.CapabilityTradingAccountsList,
	core.CapabilityTradingAccountsBalance,
	core.CapabilityTradingBuyingPower,
	core.CapabilityTradingLoanPackages,
	core.CapabilityTradingOrdersList,
	core.CapabilityTradingOrdersHistory,
	core.CapabilityTradingOrdersGet,
	core.CapabilityTradingOrdersPlace,
	core.CapabilityTradingOrdersCancel,
	core.CapabilityTradingOrdersReplace,
	core.CapabilityTradingOrdersExecutions,
	core.CapabilityTradingPositionsList,
	core.CapabilityTradingPositionsGet,
	core.CapabilityTradingPositionsClose,
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
