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
	request := nativedto.GetFinancialStatementV2Request{
		Symbol: env.String("FHSC_SYMBOL", "HPG"),
		Type:   env.String("FHSC_FINANCIAL_STATEMENT_TYPE", "income-statement"),
		Period: env.String("FHSC_PERIOD", "quarter"),
	}
	response, err := broker.Native().MarketData().GetFinancialStatementV2(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
