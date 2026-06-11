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
	response, err := broker.Native().Trading().CashInAdvanceAmount(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, item := range response.Data.CIAAmounts {
		fmt.Printf("dueDate=%s sellValue=%s advance=%s\n", item.DueDate, item.SellValue, item.Advance)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
