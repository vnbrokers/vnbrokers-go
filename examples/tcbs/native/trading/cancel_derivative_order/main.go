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
	request := nativedto.CancelDerivativeOrderRequest{
		CancelDerivativeOrderBody: nativedto.CancelDerivativeOrderBody{
			AccountID: env.RequiredString("TCBS_ACCOUNT_ID"),
			Cmd:       env.RequiredString("TCBS_CMD"),
			OrderNo:   env.RequiredString("TCBS_ORDER_NO"),
			Pin:       env.RequiredString("TCBS_PIN"),
			RefID:     env.RequiredString("TCBS_REF_ID"),
		},
	}
	response, err := broker.Native().Trading().CancelDerivativeOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
