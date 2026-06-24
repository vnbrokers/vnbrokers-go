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
		APIKey:         env.RequiredString("FHSC_API_KEY"),
		APISecret:      env.RequiredString("FHSC_API_SECRET"),
		UserID:         env.RequiredInt("FHSC_USER_ID"),
		TwoFactorToken: env.String("FHSC_2FA_TOKEN", ""),
	})
	request := nativedto.GetOrderBookDetailRequest{
		SubAccountID: env.RequiredString("FHSC_SUB_ACCOUNT_ID"),
		OrderID:      env.RequiredString("FHSC_ORDER_ID"),
	}
	response, err := broker.Native().Trading().GetOrderBookDetail(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
