package main

import (
	"context"
	"fmt"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewFHSC(fhsc.Config{
		APIKey:    env.RequiredString("FHSC_API_KEY"),
		APISecret: env.RequiredString("FHSC_API_SECRET"),
	})
	request := nativedto.GetMarketDataRequest{
		Type:  env.String("FHSC_MARKET_DATA_TYPE", "top-gainers"),
		Limit: env.Int("FHSC_MARKET_DATA_LIMIT", 10),
	}
	response, err := broker.Native().MarketData().GetMarketData(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
