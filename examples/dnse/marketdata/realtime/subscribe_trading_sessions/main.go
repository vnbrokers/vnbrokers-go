package main

import (
	"context"
	"encoding/json"
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
	sub, err := broker.Native().MarketData().Realtime().SubscribeTradingSessions(
		context.Background(),
		nativedto.SubscribeTradingSessionRequest{
			TSCProdGrpID: "STO",
			BoardID:      "G1",
		},
	)
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		message, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(message))
	}
}
