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
	var fundCompanyID *int64
	if value := env.Int("FHSC_FUND_COMPANY_ID", 0); value > 0 {
		fundCompanyID = &value
	}
	request := nativedto.GetFundCertificatesRequest{
		FundType:      env.String("FHSC_FUND_TYPE", "MUTUAL_FUND"),
		FundCompanyID: fundCompanyID,
	}
	response, err := broker.Native().MarketData().GetFundCertificates(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
