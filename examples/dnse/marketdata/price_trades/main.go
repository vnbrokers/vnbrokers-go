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
	payload, err := broker.MarketData().Quotes().PriceTrades(
		context.Background(),
		"ACB",
		0,
		0,
		"G1",
		100,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
