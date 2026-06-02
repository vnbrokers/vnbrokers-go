package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/trading"
)

func main() {
	price := decimal.RequireFromString("23200")
	quantity := 2
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	payload, err := broker.Trading().Orders().Replace(context.Background(), trading.ReplaceOrderRequest{
		AccountID:     requiredString("DNSE_ACCOUNT_NO"),
		OrderID:       "123456",
		Price:         price,
		Quantity:      quantity,
		MarketType:    trading.MarketTypeStock,
		OrderCategory: trading.OrderCategoryNormal,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)
}

func requiredString(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(name + " is required")
	}
	return value
}
