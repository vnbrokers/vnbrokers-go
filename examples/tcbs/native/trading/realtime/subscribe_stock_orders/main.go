package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, envDuration("TCBS_STREAM_DURATION", time.Minute))
	defer cancel()

	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: mustEnv("TCBS_ACCESS_TOKEN")})
	request := nativedto.SubscribeStockOrdersRequest{}
	subscription, err := broker.Native().Trading().Realtime().SubscribeStockOrders(ctx, request)
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
		case status, ok := <-subscription.Status():
			if !ok {
				return
			}
			fmt.Printf("status: %s\n", status)
		case err, ok := <-subscription.Errors():
			if !ok {
				return
			}
			fmt.Printf("error: %v\n", err)
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

func envList(key, fallback string) []string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	items := strings.Split(value, ",")
	for index := range items {
		items[index] = strings.TrimSpace(items[index])
	}
	return items
}

func envDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return duration
}
