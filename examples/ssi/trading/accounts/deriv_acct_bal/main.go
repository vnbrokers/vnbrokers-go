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
	response, err := broker.Native().Trading().DerivAcctBal(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, bal := range response.Data {
		fmt.Printf("account=%s balance=%s withdrawable=%s\n", bal.Account, bal.AccountBalance, bal.Withdrawable)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
