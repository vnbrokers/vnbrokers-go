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
	request := nativedto.UpdateDerivativeConditionalOrderRequest{
		UpdateDerivativeConditionalOrderBody: nativedto.UpdateDerivativeConditionalOrderBody{
			AccountID: env.RequiredString("TCBS_ACCOUNT_ID"),
			Cmd:       env.RequiredString("TCBS_CMD"),
			PkOrderNo: env.RequiredString("TCBS_PK_ORDER_NO"),
			RefID:     env.RequiredString("TCBS_REF_ID"),
			SoPrice:   env.RequiredFloat("TCBS_SO_PRICE"),
			Type:      env.RequiredString("TCBS_TYPE"),
		},
	}
	response, err := broker.Native().Trading().UpdateDerivativeConditionalOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
