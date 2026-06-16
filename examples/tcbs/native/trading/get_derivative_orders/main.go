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
	request := nativedto.GetDerivativeOrdersRequest{
		PageNo:    env.Float("TCBS_PAGE_NO", 1),
		PageSize:  env.Float("TCBS_PAGE_SIZE", 20),
		AccountID: env.RequiredString("TCBS_ACCOUNT_ID"),
		Symbol:    env.String("TCBS_SYMBOL", ""),
		RefID:     env.String("TCBS_REF_ID", ""),
		OrderType: env.String("TCBS_ORDER_TYPE", ""),
		Status:    env.String("TCBS_STATUS", ""),
	}
	response, err := broker.Native().Trading().GetDerivativeOrders(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
