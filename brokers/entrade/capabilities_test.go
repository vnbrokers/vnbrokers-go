package entrade

import (
	"testing"

	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestCapabilitiesExposeOnlyAuthAndNativeEndpoints(t *testing.T) {
	want := []core.Capability{
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
	if len(Capabilities) != len(want) {
		t.Fatalf("capabilities=%v", Capabilities)
	}
	for i := range want {
		if Capabilities[i] != want[i] {
			t.Fatalf("capability %d=%q want %q", i, Capabilities[i], want[i])
		}
	}
}
