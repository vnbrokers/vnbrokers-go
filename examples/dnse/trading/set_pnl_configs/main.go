package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func main() {
	takeProfitRate := 0.10
	stopLossRate := -0.05
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	payload, err := broker.Native().Trading().SetPositionPnLConfigs(
		context.Background(),
		nativedto.SetPositionPnLConfigsRequest{
			PositionID: requiredString("DNSE_POSITION_ID"),
			MarketType: "STOCK",
			Configs: nativedto.PnLConfigs{
				TakeProfit: &nativedto.PnLRule{
					Enabled:  true,
					Strategy: "RATE",
					Rate:     &takeProfitRate,
				},
				StopLoss: &nativedto.PnLRule{
					Enabled:  true,
					Strategy: "RATE",
					Rate:     &stopLossRate,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)
}

func requiredString(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(name + " is required")
	}
	return value
}
