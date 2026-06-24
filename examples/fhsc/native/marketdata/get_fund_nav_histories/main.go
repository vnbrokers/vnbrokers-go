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
	request := nativedto.GetFundNAVHistoriesRequest{
		FundName: env.String("FHSC_FUND_NAME", "SSISCA"),
		Period:   env.String("FHSC_PERIOD", "1Y"),
	}
	response, err := broker.Native().MarketData().GetFundNAVHistories(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
