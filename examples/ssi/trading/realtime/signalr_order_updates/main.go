package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/trading"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	broker := vnbrokers.NewSSI(ssi.Config{
		AccessToken: mustEnv("SSI_ACCESS_TOKEN"),
	})
	subscription, err := broker.Trading().Realtime().SubscribeOrders(
		ctx,
		trading.SubscribeOrdersRequest{},
	)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event := <-subscription.Events():
			fmt.Printf("order event: %+v\n", event)
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
