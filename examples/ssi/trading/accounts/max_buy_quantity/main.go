package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		AccessToken: mustEnv("SSI_ACCESS_TOKEN"),
	})
	response, err := broker.Trading().Accounts().MaxBuyQuantity(
		context.Background(),
		ssi.MaxBuyQuantityRequest{
			AccountID: mustEnv("SSI_ACCOUNT_NO"),
			Symbol:    "SSI",
			Price:     decimal.NewFromInt(21000),
		},
	)
	if err != nil {
		panic(err)
	}
	for _, item := range response.Data {
		fmt.Printf("%+v\n", item)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
