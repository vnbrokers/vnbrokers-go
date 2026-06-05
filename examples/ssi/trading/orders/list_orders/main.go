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
		AccessToken: mustEnv("SSI_ACCESS_TOKEN"),
	})
	orders, err := broker.Trading().Accounts().Orders(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	for _, order := range orders {
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
