package ssi

import (
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/trading"
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

	CapabilityNativeTradingCashInAdvanceAmount      = nativetrading.CapabilityCashInAdvanceAmount
	CapabilityNativeTradingUnsettleSoldTransaction   = nativetrading.CapabilityUnsettleSoldTransaction
	CapabilityNativeTradingTransferHistories         = nativetrading.CapabilityTransferHistories
	CapabilityNativeTradingCashInAdvanceHistories    = nativetrading.CapabilityCashInAdvanceHistories
	CapabilityNativeTradingEstCashInAdvanceFee       = nativetrading.CapabilityEstCashInAdvanceFee
	CapabilityNativeTradingVSDCashDW                 = nativetrading.CapabilityVSDCashDW
	CapabilityNativeTradingTransferInternal          = nativetrading.CapabilityTransferInternal
	CapabilityNativeTradingCreateCashInAdvance       = nativetrading.CapabilityCreateCashInAdvance
	CapabilityNativeTradingCashAcctBal               = nativetrading.CapabilityCashAcctBal
	CapabilityNativeTradingDerivAcctBal              = nativetrading.CapabilityDerivAcctBal
	CapabilityNativeTradingMaxBuyQty                 = nativetrading.CapabilityMaxBuyQty
	CapabilityNativeTradingMaxSellQty                = nativetrading.CapabilityMaxSellQty
	CapabilityNativeTradingOrderBook                 = nativetrading.CapabilityOrderBook
	CapabilityNativeTradingOrderHistory              = nativetrading.CapabilityOrderHistory
	CapabilityNativeTradingStockPosition             = nativetrading.CapabilityStockPosition
	CapabilityNativeTradingDerivPosition             = nativetrading.CapabilityDerivPosition
	CapabilityNativeTradingNewOrder                  = nativetrading.CapabilityNewOrder
	CapabilityNativeTradingCancelOrder               = nativetrading.CapabilityCancelOrder
	CapabilityNativeTradingModifyOrder               = nativetrading.CapabilityModifyOrder
	CapabilityNativeTradingDerNewOrder               = nativetrading.CapabilityDerNewOrder
	CapabilityNativeTradingDerCancelOrder            = nativetrading.CapabilityDerCancelOrder
	CapabilityNativeTradingDerModifyOrder            = nativetrading.CapabilityDerModifyOrder
	CapabilityNativeTradingAuditOrderBook            = nativetrading.CapabilityAuditOrderBook
	CapabilityNativeTradingPpmrAccount               = nativetrading.CapabilityPpmrAccount
	CapabilityNativeTradingRateLimit                 = nativetrading.CapabilityRateLimit
	CapabilityNativeTradingTransferable              = nativetrading.CapabilityTransferable
	CapabilityNativeTradingStockTransferHistories    = nativetrading.CapabilityStockTransferHistories
	CapabilityNativeTradingStockTransfer             = nativetrading.CapabilityStockTransfer
	CapabilityNativeTradingDividend                  = nativetrading.CapabilityDividend
	CapabilityNativeTradingExercisableQuantity       = nativetrading.CapabilityExercisableQuantity
	CapabilityNativeTradingRightsHistories           = nativetrading.CapabilityRightsHistories
	CapabilityNativeTradingCreateRight               = nativetrading.CapabilityCreateRight
	CapabilityNativeTradingFcoNewOrder               = nativetrading.CapabilityFcoNewOrder
	CapabilityNativeTradingFcoCancelOrder            = nativetrading.CapabilityFcoCancelOrder
	CapabilityNativeTradingFcoOrderBook              = nativetrading.CapabilityFcoOrderBook
	CapabilityNativeTradingFcoStatusHistory          = nativetrading.CapabilityFcoStatusHistory
	CapabilityNativeTradingFcoList                   = nativetrading.CapabilityFcoList
)

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthSendOTP,
	core.CapabilityTradingAuthTradingToken,
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
	CapabilityNativeTradingCashInAdvanceAmount,
	CapabilityNativeTradingUnsettleSoldTransaction,
	CapabilityNativeTradingTransferHistories,
	CapabilityNativeTradingCashInAdvanceHistories,
	CapabilityNativeTradingEstCashInAdvanceFee,
	CapabilityNativeTradingVSDCashDW,
	CapabilityNativeTradingTransferInternal,
	CapabilityNativeTradingCreateCashInAdvance,
	CapabilityNativeTradingCashAcctBal,
	CapabilityNativeTradingDerivAcctBal,
	CapabilityNativeTradingMaxBuyQty,
	CapabilityNativeTradingMaxSellQty,
	CapabilityNativeTradingOrderBook,
	CapabilityNativeTradingOrderHistory,
	CapabilityNativeTradingStockPosition,
	CapabilityNativeTradingDerivPosition,
	CapabilityNativeTradingNewOrder,
	CapabilityNativeTradingCancelOrder,
	CapabilityNativeTradingModifyOrder,
	CapabilityNativeTradingDerNewOrder,
	CapabilityNativeTradingDerCancelOrder,
	CapabilityNativeTradingDerModifyOrder,
	CapabilityNativeTradingAuditOrderBook,
	CapabilityNativeTradingPpmrAccount,
	CapabilityNativeTradingRateLimit,
	CapabilityNativeTradingTransferable,
	CapabilityNativeTradingStockTransferHistories,
	CapabilityNativeTradingStockTransfer,
	CapabilityNativeTradingDividend,
	CapabilityNativeTradingExercisableQuantity,
	CapabilityNativeTradingRightsHistories,
	CapabilityNativeTradingCreateRight,
	CapabilityNativeTradingFcoNewOrder,
	CapabilityNativeTradingFcoCancelOrder,
	CapabilityNativeTradingFcoOrderBook,
	CapabilityNativeTradingFcoStatusHistory,
	CapabilityNativeTradingFcoList,
	core.CapabilityMarketDataRealtimeRaw,
}
