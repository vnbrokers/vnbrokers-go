package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func main() {
	price := decimal.RequireFromString("23200")
	orderPrice, _ := price.Float64()
	quantity := int64(2)
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	payload, err := broker.Native().Trading().ReplaceOrder(context.Background(), nativedto.ReplaceOrderRequest{
		AccountNo:     requiredString("DNSE_ACCOUNT_NO"),
		OrderID:       "123456",
		Price:         &orderPrice,
		Quantity:      quantity,
		MarketType:    "STOCK",
		OrderCategory: "NORMAL",
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
