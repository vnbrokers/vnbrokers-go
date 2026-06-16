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
	request := nativedto.UpdateDerivativeOrderRequest{
		UpdateDerivativeOrderBody: nativedto.UpdateDerivativeOrderBody{
			AccountID:    env.RequiredString("TCBS_ACCOUNT_ID"),
			Nprice:       env.RequiredFloat("TCBS_NEW_PRICE"),
			Nvol:         env.RequiredFloat("TCBS_NEW_VOLUME"),
			OrderNo:      env.RequiredString("TCBS_ORDER_NO"),
			RefID:        env.RequiredString("TCBS_REF_ID"),
			SubAccountID: env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
		},
	}
	response, err := broker.Native().Trading().UpdateDerivativeOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
