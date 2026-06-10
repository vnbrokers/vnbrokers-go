package ssi

import (
	"testing"

	"github.com/vnbrokers/vnbrokers-go/core"
)

var nativeMarketDataCapabilities = []core.Capability{
	CapabilityNativeMarketDataSecurities,
	CapabilityNativeMarketDataSecuritiesDetails,
	CapabilityNativeMarketDataIndexComponents,
	CapabilityNativeMarketDataIndexList,
	CapabilityNativeMarketDataDailyOhlc,
	CapabilityNativeMarketDataIntradayOhlc,
	CapabilityNativeMarketDataDailyIndex,
	CapabilityNativeMarketDataDailyStockPrice,
}

func TestSSINativeMarketDataCapabilities(t *testing.T) {
	broker := NewBroker(Config{})
	for _, capability := range nativeMarketDataCapabilities {
		if !broker.Supports(capability) {
			t.Errorf("missing capability %q", capability)
		}
	}
	if broker.Supports(core.CapabilityMarketDataSymbolsList) {
		t.Fatal("did not expect normalized symbols capability")
	}
}

func TestSSIMarketDataCapabilities(t *testing.T) {
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
