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
	symbols, err := broker.MarketData().Symbols().List(
		context.Background(),
		"ACB",
		"STO",
		"",
		"",
		20,
		0,
	)
	if err != nil {
		panic(err)
	}
	for _, symbol := range symbols {
		fmt.Println(symbol)
	}
}
