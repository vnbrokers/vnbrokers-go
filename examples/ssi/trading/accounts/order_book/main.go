package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		TradingToken: mustEnv("SSI_FCTRADING_TOKEN"),
	})
	response, err := broker.Native().Trading().OrderBook(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, order := range response.Data.Orders {
		fmt.Printf("orderID=%s symbol=%s side=%s price=%s qty=%s status=%s\n",
			order.OrderID, order.InstrumentID, order.BuySell, order.Price, order.Quantity, order.OrderStatus)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
