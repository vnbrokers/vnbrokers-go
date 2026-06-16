package main

import (
	"context"
	"fmt"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: env.RequiredString("TCBS_ACCESS_TOKEN")})
	request := nativedto.DepositDerivativeMarginRequest{
		DepositDerivativeMarginBody: nativedto.DepositDerivativeMarginBody{
			AccountID:      env.RequiredString("TCBS_ACCOUNT_ID"),
			Amount:         env.RequiredFloat("TCBS_AMOUNT"),
			PaymentContent: env.RequiredFloat("TCBS_PAYMENT_CONTENT"),
			SubAccountID:   env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
		},
	}
	response, err := broker.Native().Trading().DepositDerivativeMargin(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
