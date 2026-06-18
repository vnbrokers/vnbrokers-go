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
	ctx, cancel := context.WithTimeout(ctx, env.Duration("TCBS_STREAM_DURATION", 8*time.Hour))
	defer cancel()

	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: env.RequiredString("TCBS_ACCESS_TOKEN")})
	// `1` VN-INDEX, `2` VN30-INDEX, `3` HNX, `4` HNX30-INDEX, `5` UPCOM,
	request := nativedto.SubscribeMarketIndexesRequest{Indexes: env.List("TCBS_INDEXES", "1,2,3,4,5")}
	subscription, err := broker.Native().MarketData().Realtime().SubscribeMarketIndexes(ctx, request)
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
