package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	broker := vnbrokers.NewSSI(ssi.Config{
		AccessToken: mustEnv("SSI_ACCESS_TOKEN"),
	})
	subscription, err := broker.MarketData().Realtime().SubscribeRawChannel(ctx, "X:ALL")
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case payload := <-subscription.Events():
			fmt.Printf("market data: %+v\n", payload.Data)
		case err := <-subscription.Errors():
			fmt.Printf("stream error: %v\n", err)
		case status := <-subscription.Status():
			fmt.Printf("status: %s\n", status)
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
