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
	request := nativedto.TransferBetweenSubaccountsRequest{
		TransferBetweenSubaccountsBody: nativedto.TransferBetweenSubaccountsBody{
			Amount:                   env.RequiredFloat("TCBS_AMOUNT"),
			Description:              env.RequiredFloat("TCBS_DESCRIPTION"),
			DestinationAccountNumber: env.RequiredString("TCBS_DESTINATION_ACCOUNT_NO"),
			SourceAccountNumber:      env.RequiredString("TCBS_SOURCE_ACCOUNT_NO"),
		},
	}
	response, err := broker.Native().Trading().TransferBetweenSubaccounts(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
