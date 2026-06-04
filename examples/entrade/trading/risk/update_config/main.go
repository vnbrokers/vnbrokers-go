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
	investorAccountID, err := strconv.Atoi(exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}
	payload, err := exampleutil.Broker().Trading().Risk().UpdateConfig(
		context.Background(),
		exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"),
		entrade.RiskConfigRequest{
			CutLossRate:               decimal.RequireFromString("0.24"),
			InvestorAccountID:         investorAccountID,
			TrailingEnabled:           false,
			InvestorID:                investorID,
			AutoIncreaseDealRate:      true,
			EnableAutoDealDepositNoti: true,
		},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
