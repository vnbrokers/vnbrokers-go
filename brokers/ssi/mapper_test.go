package ssi

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func TestMapSSIOrderConvertsKnownFieldsConservatively(t *testing.T) {
	price := decimal.NewFromInt(21000)
	quantity := decimal.NewFromInt(100)
	order := MapOrder(Order{
		OrderID:      "12658867",
		BuySell:      "B",
		Price:        &price,
		Quantity:     &quantity,
		OrderStatus:  "Qu",
		MarketID:     "VN",
		InstrumentID: "SSI",
		OrderType:    "LO",
	})

	if order.Broker != "ssi" {
		t.Fatalf("broker = %s", order.Broker)
	}
	if order.OrderID != "12658867" || order.Symbol != "SSI" {
		t.Fatalf("order = %+v", order)
	}
	if order.Side != domain.OrderSideBuy {
		t.Fatalf("side = %s", order.Side)
	}
	if order.Type != domain.OrderTypeLimit {
		t.Fatalf("type = %s", order.Type)
	}
	if order.Status != domain.OrderStatusPending {
		t.Fatalf("status = %s", order.Status)
	}
}

func TestMapStockPositionUsesOnHandAndAveragePrice(t *testing.T) {
	onHand := decimal.NewFromInt(300)
	avgPrice := decimal.NewFromInt(1529)
	marketPrice := decimal.NewFromInt(2000)

	position := MapStockPosition("0901351", StockPosition{
		InstrumentID: "AMC",
		OnHand:       &onHand,
		AveragePrice: &avgPrice,
		MarketPrice:  &marketPrice,
	})

	if position.AccountID != "0901351" || position.Symbol != "AMC" {
		t.Fatalf("position = %+v", position)
	}
	if !position.Quantity.Equal(onHand) || !position.AvailableQuantity.Equal(onHand) {
		t.Fatalf("quantity = %s available = %s", position.Quantity, position.AvailableQuantity)
	}
	if position.AveragePrice == nil || !position.AveragePrice.Equal(avgPrice) {
		t.Fatalf("avg price = %v", position.AveragePrice)
	}
	if position.MarketValue == nil || !position.MarketValue.Equal(decimal.NewFromInt(600000)) {
		t.Fatalf("market value = %v", position.MarketValue)
	}
}
