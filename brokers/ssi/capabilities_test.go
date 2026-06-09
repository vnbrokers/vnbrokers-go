package ssi

import (
	"testing"

	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestSSIMarketDataCapabilitiesMatchTypedRealtimeAPI(t *testing.T) {
	broker := NewBroker(Config{})

	if !broker.Supports(core.CapabilityMarketDataRealtimeRaw) {
		t.Fatal("expected raw market data realtime capability")
	}
	if broker.Supports(core.CapabilityMarketDataRealtimeTicks) {
		t.Fatal("did not expect removed ticks capability")
	}
	if broker.Supports(core.CapabilityMarketDataRealtimeTop) {
		t.Fatal("did not expect removed top price capability")
	}
}
