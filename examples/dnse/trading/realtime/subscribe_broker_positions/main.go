package main

import (
	"context"
	"fmt"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:    env.RequiredString("DNSE_API_KEY"),
		APISecret: env.RequiredString("DNSE_API_SECRET"),
	})
	sub, err := broker.Native().Trading().Realtime().SubscribeBrokerPositions(
		context.Background(),
		nativedto.SubscribeBrokerPositionsRequest{
			MarketType: env.String("DNSE_MARKET_TYPE", "STOCK"),
			InvestorID: env.RequiredString("DNSE_INVESTOR_ID"),
		},
	)
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		fmt.Println(event)
	}
}
