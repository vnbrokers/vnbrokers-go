package main

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	payload, err := exampleutil.Broker().Trading().Accounts().BuyingPower(
		context.Background(),
		entrade.BuyingPowerRequest{
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
	exampleutil.PrintRaw(payload)
}
