package entrade

import "github.com/vnbrokers/vnbrokers-go/core"

var Capabilities = []core.Capability{
	core.CapabilityTradingAuthTradingToken,
	core.CapabilityTradingAccountsList,
	core.CapabilityTradingAccountsBalance,
	core.CapabilityTradingBuyingPower,
	core.CapabilityTradingLoanPackages,
	core.CapabilityTradingOrdersList,
	core.CapabilityTradingOrdersGet,
	core.CapabilityTradingOrdersPlace,
	core.CapabilityTradingOrdersCancel,
	core.CapabilityTradingPositionsList,
	core.CapabilityTradingPositionsClose,
	core.CapabilityMarketDataSymbolsList,
}
