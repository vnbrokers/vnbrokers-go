package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	broker := vnbrokers.NewTCBS(tcbs.Config{
		AccessToken: mustEnv("TCBS_ACCESS_TOKEN"),
	})
	subscription, err := broker.Trading().Realtime().SubscribeStockMatches(ctx)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event, ok := <-subscription.Events():
			if !ok {
				return
			}
			fmt.Printf("event: %+v\n", event)
		case err, ok := <-subscription.Errors():
			if !ok {
				return
			}
			fmt.Printf("error: %v\n", err)
		case status, ok := <-subscription.Status():
			if !ok {
				return
			}
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
