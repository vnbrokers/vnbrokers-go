package main

import (
	"context"
	"strconv"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	investorID, err := strconv.Atoi(exampleutil.MustEnv("ENTRADE_INVESTOR_ID"))
	if err != nil {
		panic(err)
	}
	payload, err := exampleutil.Broker().Trading().Orders().Place(
		context.Background(),
		entrade.PlaceDerivativeOrderRequest{
			BankMarginPortfolioID: 34,
			InvestorID:            investorID,
			Symbol:                "VN30F2512",
			Price:                 decimal.RequireFromString("1920.9"),
			OrderType:             "LO",
			Side:                  "NB",
			Quantity:              1,
		},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
