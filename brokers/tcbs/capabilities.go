package tcbs

import (
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
)

const (
	CapabilityNativeTradingGetSubAccountInformation              = nativetrading.CapabilityGetSubAccountInformation
	CapabilityNativeTradingTransferBetweenSubaccounts            = nativetrading.CapabilityTransferBetweenSubaccounts
	CapabilityNativeTradingWithdrawDerivativeMargin              = nativetrading.CapabilityWithdrawDerivativeMargin
	CapabilityNativeTradingDepositDerivativeMargin               = nativetrading.CapabilityDepositDerivativeMargin
	CapabilityNativeTradingPlaceStockOrder                       = nativetrading.CapabilityPlaceStockOrder
	CapabilityNativeTradingUpdateStockOrder                      = nativetrading.CapabilityUpdateStockOrder
	CapabilityNativeTradingCancelStockOrder                      = nativetrading.CapabilityCancelStockOrder
	CapabilityNativeTradingGetStockOrders                        = nativetrading.CapabilityGetStockOrders
	CapabilityNativeTradingGetStockOrder                         = nativetrading.CapabilityGetStockOrder
	CapabilityNativeTradingGetStockMatchingDetails               = nativetrading.CapabilityGetStockMatchingDetails
	CapabilityNativeTradingGetStockPurchasingPower               = nativetrading.CapabilityGetStockPurchasingPower
	CapabilityNativeTradingGetStockPurchasingPowerBySymbol       = nativetrading.CapabilityGetStockPurchasingPowerBySymbol
	CapabilityNativeTradingGetStockPurchasingPowerBySymbolPrice  = nativetrading.CapabilityGetStockPurchasingPowerBySymbolPrice
	CapabilityNativeTradingGetMarginQuota                        = nativetrading.CapabilityGetMarginQuota
	CapabilityNativeTradingGetMarginAccountInformation           = nativetrading.CapabilityGetMarginAccountInformation
	CapabilityNativeTradingGetSupplementaryLoanPackages          = nativetrading.CapabilityGetSupplementaryLoanPackages
	CapabilityNativeTradingGetLoans                              = nativetrading.CapabilityGetLoans
	CapabilityNativeTradingGetStockAssets                        = nativetrading.CapabilityGetStockAssets
	CapabilityNativeTradingGetCashInvestments                    = nativetrading.CapabilityGetCashInvestments
	CapabilityNativeTradingGetCashStatements                     = nativetrading.CapabilityGetCashStatements
	CapabilityNativeTradingGetMarginInformation                  = nativetrading.CapabilityGetMarginInformation
	CapabilityNativeTradingGetDerivativeCash                     = nativetrading.CapabilityGetDerivativeCash
	CapabilityNativeTradingGetClosedDerivativePositions          = nativetrading.CapabilityGetClosedDerivativePositions
	CapabilityNativeTradingGetOpenDerivativePositions            = nativetrading.CapabilityGetOpenDerivativePositions
	CapabilityNativeTradingGetDerivativeOrders                   = nativetrading.CapabilityGetDerivativeOrders
	CapabilityNativeTradingGetDerivativeConditionalOrders        = nativetrading.CapabilityGetDerivativeConditionalOrders
	CapabilityNativeTradingPlaceDerivativeOrder                  = nativetrading.CapabilityPlaceDerivativeOrder
	CapabilityNativeTradingPlaceDerivativeConditionalOrder       = nativetrading.CapabilityPlaceDerivativeConditionalOrder
	CapabilityNativeTradingUpdateDerivativeOrder                 = nativetrading.CapabilityUpdateDerivativeOrder
	CapabilityNativeTradingUpdateDerivativeConditionalOrder      = nativetrading.CapabilityUpdateDerivativeConditionalOrder
	CapabilityNativeTradingCancelDerivativeOrder                 = nativetrading.CapabilityCancelDerivativeOrder
	CapabilityNativeTradingCancelDerivativeConditionalOrder      = nativetrading.CapabilityCancelDerivativeConditionalOrder
	CapabilityNativeTradingRealtimeStockOrders                   = nativetrading.CapabilityRealtimeStockOrders
	CapabilityNativeTradingRealtimeDerivativeOrders              = nativetrading.CapabilityRealtimeDerivativeOrders
	CapabilityNativeTradingRealtimeDerivativeOpenPositions       = nativetrading.CapabilityRealtimeDerivativeOpenPositions
	CapabilityNativeTradingRealtimeDerivativeConditionalOrders   = nativetrading.CapabilityRealtimeDerivativeConditionalOrders
	CapabilityNativeMarketDataGetDerivativeTickers               = nativemarketdata.CapabilityGetDerivativeTickers
	CapabilityNativeMarketDataGetStockTickers                    = nativemarketdata.CapabilityGetStockTickers
	CapabilityNativeMarketDataGetStockForeignRooms               = nativemarketdata.CapabilityGetStockForeignRooms
	CapabilityNativeMarketDataGetStockPutThroughs                = nativemarketdata.CapabilityGetStockPutThroughs
	CapabilityNativeMarketDataGetStockTradeHistory               = nativemarketdata.CapabilityGetStockTradeHistory
	CapabilityNativeMarketDataGetStockSupplyDemand15Minutes      = nativemarketdata.CapabilityGetStockSupplyDemand15Minutes
	CapabilityNativeMarketDataGetStockSupplyDemandDaily          = nativemarketdata.CapabilityGetStockSupplyDemandDaily
	CapabilityNativeMarketDataGetStockSupplyDemandMonthly        = nativemarketdata.CapabilityGetStockSupplyDemandMonthly
	CapabilityNativeMarketDataRealtimeStockPrices                = nativemarketdata.CapabilityRealtimeStockPrices
	CapabilityNativeMarketDataRealtimeStockTradeHistory          = nativemarketdata.CapabilityRealtimeStockTradeHistory
	CapabilityNativeMarketDataRealtimeStockSupplyDemandOneMinute = nativemarketdata.CapabilityRealtimeStockSupplyDemandOneMinute
	CapabilityNativeMarketDataRealtimeStockSupplyDemand15Minutes = nativemarketdata.CapabilityRealtimeStockSupplyDemandFifteenMinutes
	CapabilityNativeMarketDataRealtimeDerivativeBidPrices        = nativemarketdata.CapabilityRealtimeDerivativeBidPrices
	CapabilityNativeMarketDataRealtimeDerivativeOfferPrices      = nativemarketdata.CapabilityRealtimeDerivativeOfferPrices
	CapabilityNativeMarketDataRealtimeDerivativeForeignTrading   = nativemarketdata.CapabilityRealtimeDerivativeForeignTrading
	CapabilityNativeMarketDataRealtimeDerivativeBasePrices       = nativemarketdata.CapabilityRealtimeDerivativeBasePrices
	CapabilityNativeMarketDataRealtimeDerivativeMatchedPrices    = nativemarketdata.CapabilityRealtimeDerivativeMatchedPrices
	CapabilityNativeMarketDataRealtimeDerivativeTickerMatches    = nativemarketdata.CapabilityRealtimeDerivativeTickerMatches
	CapabilityNativeMarketDataRealtimeDerivativeIndexes          = nativemarketdata.CapabilityRealtimeDerivativeIndexes
)

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthTradingToken,
	CapabilityNativeTradingGetSubAccountInformation,
	CapabilityNativeTradingTransferBetweenSubaccounts,
	CapabilityNativeTradingWithdrawDerivativeMargin,
	CapabilityNativeTradingDepositDerivativeMargin,
	CapabilityNativeTradingPlaceStockOrder,
	CapabilityNativeTradingUpdateStockOrder,
	CapabilityNativeTradingCancelStockOrder,
	CapabilityNativeTradingGetStockOrders,
	CapabilityNativeTradingGetStockOrder,
	CapabilityNativeTradingGetStockMatchingDetails,
	CapabilityNativeTradingGetStockPurchasingPower,
	CapabilityNativeTradingGetStockPurchasingPowerBySymbol,
	CapabilityNativeTradingGetStockPurchasingPowerBySymbolPrice,
	CapabilityNativeTradingGetMarginQuota,
	CapabilityNativeTradingGetMarginAccountInformation,
	CapabilityNativeTradingGetSupplementaryLoanPackages,
	CapabilityNativeTradingGetLoans,
	CapabilityNativeTradingGetStockAssets,
	CapabilityNativeTradingGetCashInvestments,
	CapabilityNativeTradingGetCashStatements,
	CapabilityNativeTradingGetMarginInformation,
	CapabilityNativeTradingGetDerivativeCash,
	CapabilityNativeTradingGetClosedDerivativePositions,
	CapabilityNativeTradingGetOpenDerivativePositions,
	CapabilityNativeTradingGetDerivativeOrders,
	CapabilityNativeTradingGetDerivativeConditionalOrders,
	CapabilityNativeTradingPlaceDerivativeOrder,
	CapabilityNativeTradingPlaceDerivativeConditionalOrder,
	CapabilityNativeTradingUpdateDerivativeOrder,
	CapabilityNativeTradingUpdateDerivativeConditionalOrder,
	CapabilityNativeTradingCancelDerivativeOrder,
	CapabilityNativeTradingCancelDerivativeConditionalOrder,
	CapabilityNativeTradingRealtimeStockOrders,
	CapabilityNativeTradingRealtimeDerivativeOrders,
	CapabilityNativeTradingRealtimeDerivativeOpenPositions,
	CapabilityNativeTradingRealtimeDerivativeConditionalOrders,
	CapabilityNativeMarketDataGetDerivativeTickers,
	CapabilityNativeMarketDataGetStockTickers,
	CapabilityNativeMarketDataGetStockForeignRooms,
	CapabilityNativeMarketDataGetStockPutThroughs,
	CapabilityNativeMarketDataGetStockTradeHistory,
	CapabilityNativeMarketDataGetStockSupplyDemand15Minutes,
	CapabilityNativeMarketDataGetStockSupplyDemandDaily,
	CapabilityNativeMarketDataGetStockSupplyDemandMonthly,
	CapabilityNativeMarketDataRealtimeStockPrices,
	CapabilityNativeMarketDataRealtimeStockTradeHistory,
	CapabilityNativeMarketDataRealtimeStockSupplyDemandOneMinute,
	CapabilityNativeMarketDataRealtimeStockSupplyDemand15Minutes,
	CapabilityNativeMarketDataRealtimeDerivativeBidPrices,
	CapabilityNativeMarketDataRealtimeDerivativeOfferPrices,
	CapabilityNativeMarketDataRealtimeDerivativeForeignTrading,
	CapabilityNativeMarketDataRealtimeDerivativeBasePrices,
	CapabilityNativeMarketDataRealtimeDerivativeMatchedPrices,
	CapabilityNativeMarketDataRealtimeDerivativeTickerMatches,
	CapabilityNativeMarketDataRealtimeDerivativeIndexes,
}
