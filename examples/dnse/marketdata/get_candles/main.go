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
	candles, err := broker.MarketData().Candles().Get(
		context.Background(),
		"ACB",
		"15",
		0,
		0,
		"STOCK",
	)
	if err != nil {
		panic(err)
	}
	for _, candle := range candles {
		fmt.Println(candle)
	}
}
