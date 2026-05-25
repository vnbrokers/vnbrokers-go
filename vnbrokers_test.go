package vnbrokers

import (
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func TestNewBrokerBuildsDNSEBroker(t *testing.T) {
	broker, err := NewBroker("dnse", FactoryConfig{DNSE: &dnse.Config{}})
	if err != nil {
		t.Fatalf("new broker: %v", err)
	}
	if broker.Name() != "dnse" {
		t.Fatalf("broker name = %q", broker.Name())
	}
	if !broker.Supports(core.CapabilityMarketDataRealtimeTop) {
		t.Fatalf("expected dnse top price realtime capability")
	}
}

func TestNewBrokerRejectsUnknownBroker(t *testing.T) {
	if _, err := NewBroker("unknown", FactoryConfig{}); err == nil {
		t.Fatalf("expected unknown broker error")
	}
}
