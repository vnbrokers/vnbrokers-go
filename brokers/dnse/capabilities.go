package dnse

import (
	nativebrokerage "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/brokerage"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
)

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthSendOTP,
	core.CapabilityTradingAuthTradingToken,
	nativemarketdata.CapabilityTradeHistory,
	nativemarketdata.CapabilityInstrumentDetails,
	nativemarketdata.CapabilityInstruments,
	nativemarketdata.CapabilityLatestQuotes,
	nativemarketdata.CapabilityLatestTrades,
	nativemarketdata.CapabilityOHLC,
	nativemarketdata.CapabilityClosePrice,
	nativemarketdata.CapabilityQuoteHistory,
	nativemarketdata.CapabilityForeignTrading,
	nativemarketdata.CapabilitySecurityDefinition,
	nativemarketdata.CapabilityWorkingDates,
	nativemarketdata.CapabilityRealtimeExpectedPrices,
	nativemarketdata.CapabilityRealtimeForeign,
	nativemarketdata.CapabilityRealtimeMarketIndexes,
	nativemarketdata.CapabilityRealtimeEstimatedMarketIndexes,
	nativemarketdata.CapabilityRealtimeOHLC,
	nativemarketdata.CapabilityRealtimeClosedOHLC,
	nativemarketdata.CapabilityRealtimeQuotes,
	nativemarketdata.CapabilityRealtimeSecurityDefinitions,
	nativemarketdata.CapabilityRealtimeTrades,
	nativemarketdata.CapabilityRealtimeTradeExtras,
	nativetrading.CapabilityAccounts,
	nativetrading.CapabilityAccountBalances,
	nativetrading.CapabilityCorporateActionHistory,
	nativetrading.CapabilityExecutions,
	nativetrading.CapabilityLoanPackages,
	nativetrading.CapabilityOrderHistory,
	nativetrading.CapabilityOrders,
	nativetrading.CapabilityPositionPnLConfigs,
	nativetrading.CapabilityPosition,
	nativetrading.CapabilityPositions,
	nativetrading.CapabilityPPSE,
	nativetrading.CapabilityCancelOrder,
	nativetrading.CapabilityClosePosition,
	nativetrading.CapabilityOrder,
	nativetrading.CapabilityPlaceOrder,
	nativetrading.CapabilitySetPositionPnLConfigs,
	nativetrading.CapabilityReplaceOrder,
	nativetrading.CapabilityRealtimeOrders,
	nativetrading.CapabilityRealtimeBrokerOrders,
	nativetrading.CapabilityRealtimePositions,
	nativebrokerage.CapabilityCareByAccounts,
}
