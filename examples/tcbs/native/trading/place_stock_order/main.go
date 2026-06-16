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
	request := nativedto.PlaceStockOrderRequest{
		AccountNo: env.RequiredString("TCBS_ACCOUNT_NO"),
		PlaceStockOrderBody: nativedto.PlaceStockOrderBody{
			ExecType:  env.String("TCBS_EXEC_TYPE", "NB"),
			Price:     env.RequiredInt("TCBS_PRICE"),
			PriceType: env.String("TCBS_PRICE_TYPE", "LO"),
			Quantity:  env.RequiredInt("TCBS_QUANTITY"),
			Symbol:    env.RequiredString("TCBS_SYMBOL"),
		},
	}
	response, err := broker.Native().Trading().PlaceStockOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
