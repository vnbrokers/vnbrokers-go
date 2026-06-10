package ssi

import (
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/core"
)

const (
	CapabilityNativeMarketDataSecurities        = nativemarketdata.CapabilitySecurities
	CapabilityNativeMarketDataSecuritiesDetails = nativemarketdata.CapabilitySecuritiesDetails
	CapabilityNativeMarketDataIndexComponents   = nativemarketdata.CapabilityIndexComponents
	CapabilityNativeMarketDataIndexList         = nativemarketdata.CapabilityIndexList
	CapabilityNativeMarketDataDailyOhlc         = nativemarketdata.CapabilityDailyOhlc
	CapabilityNativeMarketDataIntradayOhlc      = nativemarketdata.CapabilityIntradayOhlc
	CapabilityNativeMarketDataDailyIndex        = nativemarketdata.CapabilityDailyIndex
	CapabilityNativeMarketDataDailyStockPrice   = nativemarketdata.CapabilityDailyStockPrice
)

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthSendOTP,
	core.CapabilityTradingAuthTradingToken,
	core.CapabilityTradingAccountsBalance,
	core.CapabilityTradingBuyingPower,
	core.CapabilityTradingOrdersList,
	core.CapabilityTradingOrdersHistory,
	core.CapabilityTradingOrdersGet,
	core.CapabilityTradingOrdersPlace,
	core.CapabilityTradingOrdersCancel,
	core.CapabilityTradingOrdersReplace,
	core.CapabilityTradingPositionsList,
	core.CapabilityTradingCashTransfer,
	core.CapabilityTradingStockTransfer,
	core.CapabilityTradingConditionalOrders,
	core.CapabilityTradingRealtimeOrders,
	core.CapabilityTradingRealtimePosition,
	CapabilityNativeMarketDataSecurities,
	CapabilityNativeMarketDataSecuritiesDetails,
	CapabilityNativeMarketDataIndexComponents,
	CapabilityNativeMarketDataIndexList,
	CapabilityNativeMarketDataDailyOhlc,
	CapabilityNativeMarketDataIntradayOhlc,
	CapabilityNativeMarketDataDailyIndex,
	CapabilityNativeMarketDataDailyStockPrice,
	core.CapabilityMarketDataRealtimeRaw,
}
