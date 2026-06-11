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

	broker := vnbrokers.NewSSI(ssi.Config{TradingToken: mustEnv("SSI_FCTRADING_TOKEN")})
	subscription, err := broker.Native().Trading().Realtime().SubscribeFCOEvents(ctx)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event := <-subscription.Events():
			fmt.Printf("fco event: %+v\n", event)
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
