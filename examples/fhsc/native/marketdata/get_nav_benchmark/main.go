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
	request := nativedto.GetNAVBenchmarkRequest{
		FundNames: env.String("FHSC_FUND_NAMES", "SSISCA,SSISCA-E"),
		Period:    env.String("FHSC_PERIOD", "1Y"),
		FromMonth: env.String("FHSC_FROM_MONTH", ""),
		ToMonth:   env.String("FHSC_TO_MONTH", ""),
	}
	response, err := broker.Native().MarketData().GetNAVBenchmark(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
