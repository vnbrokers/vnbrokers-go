package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	payload, err := broker.Native().Trading().CancelOrder(context.Background(), nativedto.CancelOrderRequest{
		AccountNo:     requiredString("DNSE_ACCOUNT_NO"),
		OrderID:       "123456",
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
