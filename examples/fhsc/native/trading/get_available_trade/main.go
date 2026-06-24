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
	request := nativedto.GetAvailableTradeRequest{
		SubAccountID: env.RequiredString("FHSC_SUB_ACCOUNT_ID"),
		OrderSide:    env.String("FHSC_ORDER_SIDE", "BUY"),
		Symbol:       env.String("FHSC_SYMBOL", "HPG"),
		QuotePrice:   env.Int("FHSC_QUOTE_PRICE", 25000),
	}
	response, err := broker.Native().Trading().GetAvailableTrade(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
