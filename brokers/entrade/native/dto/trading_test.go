package dto

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
)

func TestDecimalRequestFieldsMarshalAsJSONNumbers(t *testing.T) {
	order, err := json.Marshal(PlaceDerivativeOrderRequest{Price: decimal.RequireFromString("1920.9")})
	if err != nil {
		t.Fatal(err)
	}
	if string(order) != `{"bankMarginPortfolioId":0,"investorId":0,"symbol":"","orderType":"","side":"","quantity":0,"price":1920.9}` {
		t.Fatalf("order json=%s", order)
	}

	risk, err := json.Marshal(UpdateRiskConfigRequest{CutLossRate: decimal.RequireFromString("0.24")})
	if err != nil {
		t.Fatal(err)
	}
	if string(risk) != `{"investorAccountId":0,"trailingEnabled":false,"investorId":0,"autoIncreaseDealRate":false,"enableAutoDealDepositNoti":false,"cutLossRate":0.24}` {
		t.Fatalf("risk json=%s", risk)
	}
}

func TestNullableAndAdditiveResponseFieldsDecode(t *testing.T) {
	var order GetDerivativeOrderResponse
	if err := json.Unmarshal([]byte(`{"id":1,"price":null,"unknown":"kept-compatible"}`), &order); err != nil {
		t.Fatal(err)
	}
	if order.ID != 1 || !order.Price.IsZero() {
		t.Fatalf("order=%#v", order)
	}
}
