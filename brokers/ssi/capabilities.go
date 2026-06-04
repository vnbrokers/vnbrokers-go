package ssi

import "github.com/vnbrokers/vnbrokers-go/core"

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
}
