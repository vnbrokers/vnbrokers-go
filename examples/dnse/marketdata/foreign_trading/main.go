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
		APIKey:    env.String("DNSE_API_KEY", ""),
		APISecret: env.String("DNSE_API_SECRET", ""),
	})
	request := nativedto.GetForeignTradingRequest{
		Symbol:  env.String("DNSE_SYMBOL", "ACB"),
		BoardID: env.String("DNSE_BOARD_ID", "G1"),
		From:    env.Int("DNSE_FOREIGN_TRADING_FROM", 1781139600),
		To:      env.Int("DNSE_FOREIGN_TRADING_TO", 1781172000),
		Limit:   int(env.Int("DNSE_LIMIT", 100)),
		Order:   env.String("DNSE_ORDER", ""),
	}
	payload, err := broker.Native().MarketData().GetForeignTrading(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
