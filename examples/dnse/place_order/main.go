package main

import (
	"context"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func main() {
	price := decimal.RequireFromString("25000")
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: "replace-with-trading-token",
		MarketType:   "STOCK",
	})
	_, err := broker.Trading().Orders().Place(context.Background(), domain.PlaceOrderRequest{
		AccountID: os.Getenv("DNSE_ACCOUNT_NO"),
		Symbol:    "ACB",
		Side:      domain.OrderSideBuy,
		Type:      domain.OrderTypeLimit,
		Quantity:  decimal.NewFromInt(1),
		Price:     &price,
	})
	if err != nil {
		panic(err)
	}
}
