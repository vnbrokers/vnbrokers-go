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
	response, err := broker.Trading().Orders().Update(
		context.Background(),
		mustEnv("TCBS_ACCOUNT_NO"),
		mustEnv("TCBS_ORDER_ID"),
		tcbs.UpdateOrderRequest{
			Price:    mustIntEnv("TCBS_PRICE"),
			Quantity: mustIntEnv("TCBS_QUANTITY"),
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

func mustIntEnv(key string) int {
	value, err := strconv.Atoi(mustEnv(key))
	if err != nil {
		panic(err)
	}
	return value
}
