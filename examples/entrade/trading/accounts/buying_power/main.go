package main

import (
	"context"

	"github.com/shopspring/decimal"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	response, err := exampleutil.Broker().Native().Trading().GetPPSE(
		context.Background(),
		nativedto.GetPPSERequest{
			InvestorID:            exampleutil.MustEnv("ENTRADE_INVESTOR_ID"),
			BankMarginPortfolioID: "34",
			Symbol:                "VN30F2512",
			Price:                 decimal.RequireFromString("1922.8"),
			Side:                  "NB",
		},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.Print(response)
}
