package core

import "testing"

func TestBaseBrokerRequiresCapability(t *testing.T) {
	broker := BaseBroker{
		BrokerName:         "test",
		BrokerCapabilities: []Capability{CapabilityTradingAccountsList},
	}

	if !broker.Supports(CapabilityTradingAccountsList) {
		t.Fatalf("expected broker to support accounts list")
	}
	if err := broker.RequireCapability(CapabilityTradingAccountsList); err != nil {
		t.Fatalf("expected supported capability: %v", err)
	}
	if err := broker.RequireCapability(CapabilityTradingOrdersPlace); err == nil {
		t.Fatalf("expected unsupported capability error")
	}
}
