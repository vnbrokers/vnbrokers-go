package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	broker := vnbrokers.NewSSI(ssi.Config{DataToken: mustEnv("SSI_FCDATA_TOKEN")})
	subscription, err := broker.MarketData().Realtime().SubscribeTopPrice(
		ctx,
		marketdata.SubscribeSymbolRequest{Symbol: "ALL"},
	)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event := <-subscription.Events():
			fmt.Printf("%+v\n", event)
		case err := <-subscription.Errors():
			fmt.Printf("stream error: %v\n", err)
		case <-ctx.Done():
			return
		}
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
