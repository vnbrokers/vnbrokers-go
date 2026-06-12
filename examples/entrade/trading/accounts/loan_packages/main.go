package main

import (
	"context"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	response, err := exampleutil.Broker().Native().Trading().GetDerivativeMarginPortfolios(
		context.Background(),
		nativedto.GetDerivativeMarginPortfoliosRequest{InvestorID: exampleutil.MustEnv("ENTRADE_INVESTOR_ID")},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.Print(response)
}
