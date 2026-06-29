package vnbrokers

import (
	"reflect"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativemarketdata "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	fhscmarketdata "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/marketdata"
)

func TestNewBrokerBuildsDNSEBroker(t *testing.T) {
	broker, err := NewBroker("dnse", dnse.Config{})
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
	broker, err := NewBroker("entrade", entrade.Config{})
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

func TestNewBrokerBuildsFHSCBroker(t *testing.T) {
	broker, err := NewBroker("fhsc", &fhsc.Config{})
	if err != nil {
		t.Fatalf("new broker: %v", err)
	}
	if broker.Name() != "fhsc" {
		t.Fatalf("broker name = %q", broker.Name())
	}
	if !broker.Supports(fhscmarketdata.CapabilityGetStockRealtime) {
		t.Fatalf("expected fhsc stock realtime capability")
	}
}

func TestRegisteredBrokersIncludesBuiltIns(t *testing.T) {
	got := RegisteredBrokers()
	want := []string{"dnse", "entrade", "fhsc", "ssi", "tcbs"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("registered brokers = %#v, want %#v", got, want)
	}
}

func TestNewBrokerRejectsWrongConfigType(t *testing.T) {
	if _, err := NewBroker("dnse", entrade.Config{}); err == nil {
		t.Fatal("expected wrong config type error")
	}
}

func TestNewBrokerRejectsUnknownBroker(t *testing.T) {
	if _, err := NewBroker("unknown", nil); err == nil {
		t.Fatalf("expected unknown broker error")
	}
}

func TestNewBrokersBuildsMultipleBrokerInstances(t *testing.T) {
	brokers, err := NewBrokers([]BrokerConfig{
		{ID: "dnse-main", Config: dnse.Config{}},
		{ID: "dnse-alt", Config: &dnse.Config{}},
		{ID: "fhsc", Config: fhsc.Config{}},
	})
	if err != nil {
		t.Fatalf("new brokers: %v", err)
	}
	if len(brokers) != 3 {
		t.Fatalf("brokers length = %d", len(brokers))
	}
	if broker, ok := brokers.Get("dnse-main"); !ok || broker.Name() != "dnse" {
		t.Fatalf("dnse-main broker = %#v, %t", broker, ok)
	}
	if broker, ok := brokers.Get("dnse-alt"); !ok || broker.Name() != "dnse" {
		t.Fatalf("dnse-alt broker = %#v, %t", broker, ok)
	}
	if broker, ok := brokers.Get("fhsc"); !ok || broker.Name() != "fhsc" {
		t.Fatalf("fhsc broker = %#v, %t", broker, ok)
	}
}

func TestNewBrokersRejectsDuplicateInstanceID(t *testing.T) {
	_, err := NewBrokers([]BrokerConfig{
		{ID: "main", Config: dnse.Config{}},
		{ID: "main", Config: fhsc.Config{}},
	})
	if err == nil {
		t.Fatal("expected duplicate id error")
	}
}

func TestNewBrokersRejectsMissingConfig(t *testing.T) {
	_, err := NewBrokers([]BrokerConfig{{ID: "main"}})
	if err == nil {
		t.Fatal("expected missing config error")
	}
}
