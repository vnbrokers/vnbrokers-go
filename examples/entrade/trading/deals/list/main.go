package main

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	payload, err := exampleutil.Broker().Trading().Deals().List(
		context.Background(),
		entrade.ListDealsRequest{
			InvestorAccountID: exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"),
			Start:             0,
			End:               20,
		},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
