package dnse

import (
	"testing"
)

func TestMapPositionUnwrapsDataPayload(t *testing.T) {
	position := MapPosition(map[string]any{
		"data": map[string]any{
			"accountNo":        "0001179019",
			"symbol":           "41I1G5000",
			"openQuantity":     23,
			"averageCostPrice": 2036.78986,
			"marketPrice":      1915.8,
		},
	})

	if position.AccountID != "0001179019" {
		t.Fatalf("account id = %q", position.AccountID)
	}
	if position.Symbol != "41I1G5000" {
		t.Fatalf("symbol = %q", position.Symbol)
	}
	if position.Quantity.String() != "23" {
		t.Fatalf("quantity = %s", position.Quantity)
	}
	if position.MarketValue == nil || position.MarketValue.String() != "44063.4" {
		t.Fatalf("market value = %v", position.MarketValue)
	}
}

func TestMapOrderEventUsesTEventShape(t *testing.T) {
	event := MapOrderEvent(map[string]any{
		"T": "eo",
		"order": map[string]any{
			"accountNo":    "0001179019",
			"id":           16701,
			"symbol":       "ACB",
			"orderStatus":  "Canceled",
			"fillQuantity": 0,
		},
	})

	if event.OrderID != "16701" {
		t.Fatalf("order id = %q", event.OrderID)
	}
	if event.Status != "CANCELLED" {
		t.Fatalf("status = %q", event.Status)
	}
}

func TestMapOrderStatusHandlesPendingCancel(t *testing.T) {
	if got := MapOrderStatus("PendingCancel"); got != "PENDING_CANCEL" {
		t.Fatalf("status = %q", got)
	}
	if got := MapOrderStatus("PARTIALLY_FILLED"); got != "PARTIALLY_FILLED" {
		t.Fatalf("partial status = %q", got)
	}
}
