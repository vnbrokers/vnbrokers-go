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
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:     os.Getenv("DNSE_API_KEY"),
		APISecret:  os.Getenv("DNSE_API_SECRET"),
		MarketType: "STOCK",
	})
	payload, err := broker.Native().Trading().GetPositionPnLConfigs(
		context.Background(),
		nativedto.GetPositionPnLConfigsRequest{
			PositionID: requiredString("DNSE_POSITION_ID"),
			MarketType: "STOCK",
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
