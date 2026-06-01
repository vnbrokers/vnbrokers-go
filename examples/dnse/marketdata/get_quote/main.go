package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:    os.Getenv("DNSE_API_KEY"),
		APISecret: os.Getenv("DNSE_API_SECRET"),
	})
	quote, err := broker.MarketData().Quotes().Get(
		context.Background(),
		"ACB",
		"",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(quote)
}
