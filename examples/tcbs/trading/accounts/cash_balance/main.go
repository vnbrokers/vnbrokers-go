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
	response, err := broker.Trading().Accounts().CashBalance(
		context.Background(),
		mustEnv("TCBS_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	for _, balance := range response.Data {
		fmt.Printf("%+v\n", balance)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
