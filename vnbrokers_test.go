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
