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
	request := nativedto.PlaceDerivativeOrderRequest{
		PlaceDerivativeOrderBody: nativedto.PlaceDerivativeOrderBody{
			AccountID:    env.RequiredString("TCBS_ACCOUNT_ID"),
			Advance:      env.String("TCBS_ADVANCE", ""),
			OrderType:    env.String("TCBS_ORDER_TYPE", "LO"),
			Pin:          env.RequiredString("TCBS_PIN"),
			Price:        env.RequiredFloat("TCBS_PRICE"),
			RefID:        env.RequiredString("TCBS_REF_ID"),
			Side:         env.RequiredString("TCBS_SIDE"),
			SubAccountID: env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
			Symbol:       env.RequiredString("TCBS_SYMBOL"),
			Volume:       env.RequiredInt("TCBS_VOLUME"),
		},
	}
	response, err := broker.Native().Trading().PlaceDerivativeOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
