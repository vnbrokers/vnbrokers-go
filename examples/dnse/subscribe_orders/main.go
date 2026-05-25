package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/trading"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:         os.Getenv("DNSE_API_KEY"),
		APISecret:      os.Getenv("DNSE_API_SECRET"),
		StreamEncoding: os.Getenv("DNSE_STREAM_ENCODING"),
	})
	sub, err := broker.Trading().Realtime().SubscribeOrders(context.Background(), trading.SubscribeOrdersRequest{
		MarketType: os.Getenv("DNSE_MARKET_TYPE"),
	})
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		fmt.Println(event)
	}
}
