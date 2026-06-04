package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
)

func main() {
	broker := vnbrokers.NewTCBS(tcbs.Config{
		AccessToken: mustEnv("TCBS_ACCESS_TOKEN"),
	})
	orders, err := broker.Trading().Accounts().Orders(
		context.Background(),
		mustEnv("TCBS_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	for _, order := range orders.Data {
		fmt.Printf("%+v\n", order)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
