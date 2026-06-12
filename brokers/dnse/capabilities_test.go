package dnse

import (
	"testing"

	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestCapabilitiesIncludeRealtimeEstimatedMarketIndexes(t *testing.T) {
	assertCapability(t, nativemarketdata.CapabilityRealtimeEstimatedMarketIndexes)
}

func TestCapabilitiesIncludeRealtimeBrokerOrders(t *testing.T) {
	assertCapability(t, nativetrading.CapabilityRealtimeBrokerOrders)
}

func assertCapability(t *testing.T, expected core.Capability) {
	t.Helper()
	for _, capability := range Capabilities {
		if capability == expected {
			return
		}
	}
	t.Fatalf("capability %q is not registered", expected)
}
