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
	response, err := broker.Trading().Positions().Stock(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	for _, portfolio := range response.Data {
		fmt.Printf("%+v\n", portfolio)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
