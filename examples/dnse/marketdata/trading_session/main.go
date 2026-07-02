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
	payload, err := broker.Native().MarketData().GetTradingSession(
		context.Background(),
		nativedto.GetTradingSessionRequest{
			BoardID:      "G1",
			TSCProdGrpID: "STO",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
