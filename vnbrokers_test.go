package vnbrokers

import (
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
)

func TestNewBrokerBuildsDNSEBroker(t *testing.T) {
	broker, err := NewBroker("dnse", FactoryConfig{DNSE: &dnse.Config{}})
	if err != nil {
		t.Fatalf("new broker: %v", err)
	}
	if broker.Name() != "dnse" {
		t.Fatalf("broker name = %q", broker.Name())
	}
	if !broker.Supports(nativemarketdata.CapabilityRealtimeQuotes) {
		t.Fatalf("expected dnse top price realtime capability")
	}
}

func TestNewBrokerBuildsEntradeBroker(t *testing.T) {
	broker, err := NewBroker("entrade", FactoryConfig{Entrade: &entrade.Config{}})
	if err != nil {
		t.Fatalf("new broker: %v", err)
	}
	if broker.Name() != "entrade" {
		t.Fatalf("broker name = %q", broker.Name())
	}
	if !broker.Supports(nativetrading.CapabilityPlaceDerivativeOrder) {
		t.Fatalf("expected entrade place order capability")
	}
}

func TestNewBrokerRejectsUnknownBroker(t *testing.T) {
	if _, err := NewBroker("unknown", FactoryConfig{}); err == nil {
		t.Fatalf("expected unknown broker error")
	}
}
