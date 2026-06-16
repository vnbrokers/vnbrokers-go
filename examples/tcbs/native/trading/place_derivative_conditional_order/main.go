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
	request := nativedto.PlaceDerivativeConditionalOrderRequest{
		PlaceDerivativeConditionalOrderBody: nativedto.PlaceDerivativeConditionalOrderBody{
			AccountID:       env.RequiredString("TCBS_ACCOUNT_ID"),
			ActivationPrice: env.RequiredFloat("TCBS_ACTIVATION_PRICE"),
			Advance:         env.String("TCBS_ADVANCE", ""),
			CallbackPoint:   env.Float("TCBS_CALLBACK_POINT", 0),
			Cmd:             env.RequiredString("TCBS_CMD"),
			OrderType:       env.String("TCBS_ORDER_TYPE", "LO"),
			Pin:             env.RequiredString("TCBS_PIN"),
			Price:           env.RequiredFloat("TCBS_PRICE"),
			RefID:           env.RequiredString("TCBS_REF_ID"),
			Side:            env.RequiredString("TCBS_SIDE"),
			SoPrice:         env.Float("TCBS_SO_PRICE", 0),
			SubAccountID:    env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
			Symbol:          env.RequiredString("TCBS_SYMBOL"),
			Type:            env.RequiredString("TCBS_TYPE"),
			Volume:          env.RequiredFloat("TCBS_VOLUME"),
		},
	}
	response, err := broker.Native().Trading().PlaceDerivativeConditionalOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
