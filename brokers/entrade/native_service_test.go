package entrade_test

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
)

func TestBrokerExposesNativeServices(t *testing.T) {
	broker := entrade.NewBroker(entrade.Config{})

	if broker.Native() == nil {
		t.Fatal("Native() is nil")
	}
	if broker.Native().Trading() == nil {
		t.Fatal("Native().Trading() is nil")
	}
	if broker.Native().MarketData() == nil {
		t.Fatal("Native().MarketData() is nil")
	}
}

func TestAuthUsesTypedEntradeDTOs(t *testing.T) {
	broker := entrade.NewBroker(entrade.Config{})
	var request nativedto.LoginRequest

	response, _ := broker.Auth().Login(context.Background(), request)
	var _ *nativedto.LoginResponse = response
}
