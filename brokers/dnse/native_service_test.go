package dnse_test

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func TestBrokerExposesNativeServices(t *testing.T) {
	broker := dnse.NewBroker(dnse.Config{})

	if broker.Native() == nil {
		t.Fatal("Native() is nil")
	}
	if broker.Native().MarketData() == nil {
		t.Fatal("Native().MarketData() is nil")
	}
	if broker.Native().Trading() == nil {
		t.Fatal("Native().Trading() is nil")
	}
	if broker.Native().Brokerage() == nil {
		t.Fatal("Native().Brokerage() is nil")
	}
}

func TestAuthUsesTypedDNSEDTOs(t *testing.T) {
	broker := dnse.NewBroker(dnse.Config{})
	var request nativedto.GetTradingTokenRequest

	_, _ = broker.Auth().GetTradingToken(context.Background(), request)
}
