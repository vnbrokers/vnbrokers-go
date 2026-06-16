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
	request := nativedto.UpdateStockOrderRequest{
		AccountNo: env.RequiredString("TCBS_ACCOUNT_NO"),
		OrderID:   env.RequiredString("TCBS_ORDER_ID"),
		UpdateStockOrderBody: nativedto.UpdateStockOrderBody{
			Price:    env.RequiredInt("TCBS_PRICE"),
			Quantity: env.RequiredInt("TCBS_QUANTITY"),
		},
	}
	response, err := broker.Native().Trading().UpdateStockOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
