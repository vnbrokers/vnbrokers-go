package main

import (
	"context"
	"strconv"

	"github.com/shopspring/decimal"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
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
	response, err := exampleutil.Broker().Native().Trading().UpdateRiskConfig(
		context.Background(),
		nativedto.UpdateRiskConfigRequest{
			PathInvestorAccountID:     exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"),
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
	exampleutil.Print(response)
}
