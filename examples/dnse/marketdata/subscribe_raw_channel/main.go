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
		APIKey:         os.Getenv("DNSE_API_KEY"),
		APISecret:      os.Getenv("DNSE_API_SECRET"),
		StreamEncoding: "msgpack",
	})
	sub, err := broker.MarketData().Realtime().SubscribeRawChannel(
		context.Background(),
		"top_price.G1.msgpack",
		[]string{"ACB"},
	)
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		fmt.Printf("%+v\n", event)
	}
}
