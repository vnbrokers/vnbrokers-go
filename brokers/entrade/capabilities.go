package entrade

import (
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
)

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthTradingToken,
	nativetrading.CapabilityInvestorAccount,
	nativetrading.CapabilityAccountBalance,
	nativetrading.CapabilityDerivativeMarginPortfolios,
	nativetrading.CapabilityPPSE,
	nativetrading.CapabilityPlaceDerivativeOrder,
	nativetrading.CapabilityDerivativeOrders,
	nativetrading.CapabilityDerivativeOrder,
	nativetrading.CapabilityCancelDerivativeOrder,
	nativetrading.CapabilityDerivativeDeals,
	nativetrading.CapabilityCloseDerivativeDeal,
	nativetrading.CapabilityRiskConfig,
	nativetrading.CapabilityUpdateRiskConfig,
	nativemarketdata.CapabilityDerivatives,
}
