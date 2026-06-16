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
	request := nativedto.GetDerivativeConditionalOrdersRequest{
		PageNo:       env.String("TCBS_PAGE_NO", "0"),
		PageSize:     env.String("TCBS_PAGE_SIZE", "20"),
		AccountID:    env.RequiredString("TCBS_ACCOUNT_ID"),
		SubAccountID: env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
		OrderStatus:  env.String("TCBS_ORDER_STATUS", ""),
		OrderType:    env.String("TCBS_ORDER_TYPE", ""),
		Symbol:       env.String("TCBS_SYMBOL", ""),
	}
	response, err := broker.Native().Trading().GetDerivativeConditionalOrders(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
