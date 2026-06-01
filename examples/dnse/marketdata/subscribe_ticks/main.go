package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:         os.Getenv("DNSE_API_KEY"),
		APISecret:      os.Getenv("DNSE_API_SECRET"),
		StreamEncoding: "msgpack",
	})
	sub, err := broker.MarketData().Realtime().SubscribeTicks(context.Background(), marketdata.SubscribeSymbolRequest{
		Symbol: "ACB",
	})
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		fmt.Println(event)
	}
}
