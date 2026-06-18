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
	request := nativedto.GetClosedDerivativePositionsRequest{
		AccountID:    env.RequiredString("TCBS_ACCOUNT_ID"),
		SubAccountID: env.RequiredString("TCBS_SUB_ACCOUNT_ID"),
		Symbol:       env.String("TCBS_SYMBOL", "41I1G6000"),
		PageNo:       env.Int("TCBS_PAGE_NO", 1),
		PageSize:     env.Int("TCBS_PAGE_SIZE", 20),
	}
	response, err := broker.Native().Trading().GetClosedDerivativePositions(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
