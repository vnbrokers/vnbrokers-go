package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	ctx, cancel := context.WithTimeout(ctx, env.Duration("TCBS_STREAM_DURATION", time.Minute))
	defer cancel()

	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: env.RequiredString("TCBS_ACCESS_TOKEN")})
	request := nativedto.SubscribeStockPricesRequest{}
	subscription, err := broker.Native().MarketData().Realtime().SubscribeStockPrices(ctx, request)
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
			fmt.Printf("event: %s\n", event)
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
