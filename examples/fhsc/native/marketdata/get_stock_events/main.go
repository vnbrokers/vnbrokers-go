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
	request := nativedto.GetStockEventsRequest{
		Stock:    env.String("FHSC_STOCK", ""),
		Stocks:   env.String("FHSC_STOCKS", "HPG,FPT"),
		FromDate: env.String("FHSC_FROM_DATE", ""),
		ToDate:   env.String("FHSC_TO_DATE", ""),
	}
	response, err := broker.Native().MarketData().GetStockEvents(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
