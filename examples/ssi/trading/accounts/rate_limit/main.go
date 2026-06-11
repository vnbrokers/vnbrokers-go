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
	response, err := broker.Native().Trading().RateLimit(
		context.Background(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, item := range response.Data {
		fmt.Printf("endpoint=%s period=%s limit=%d\n", item.Endpoint, item.Period, item.Limit)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
