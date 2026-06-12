package main

import (
	"context"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	response, err := exampleutil.Broker().Native().Trading().GetDerivativeDeals(
		context.Background(),
		nativedto.GetDerivativeDealsRequest{
			InvestorAccountID: exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"),
			Start:             0,
			End:               20,
		},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.Print(response)
}
