package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
)

func main() {
	broker := vnbrokers.NewTCBS(tcbs.Config{
		AccessToken: mustEnv("TCBS_ACCESS_TOKEN"),
	})
	response, err := broker.Trading().Orders().Place(
		context.Background(),
		mustEnv("TCBS_ACCOUNT_NO"),
		tcbs.PlaceOrderRequest{
			ExecType:  envDefault("TCBS_EXEC_TYPE", "NB"),
			Price:     mustIntEnv("TCBS_PRICE"),
			PriceType: envDefault("TCBS_PRICE_TYPE", "LO"),
			Quantity:  mustIntEnv("TCBS_QUANTITY"),
			Symbol:    mustEnv("TCBS_SYMBOL"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}

func envDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func mustIntEnv(key string) int {
	value, err := strconv.Atoi(mustEnv(key))
	if err != nil {
		panic(err)
	}
	return value
}
