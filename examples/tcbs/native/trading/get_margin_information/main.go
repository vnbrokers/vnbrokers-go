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
	request := nativedto.GetMarginInformationRequest{
		AccountNo:   env.RequiredString("TCBS_ACCOUNT_NO"),
		FromDate:    env.String("TCBS_FROM_DATE", "2025-01-01"),
		ToDate:      env.String("TCBS_TO_DATE", "2026-01-01"),
		Page:        env.String("TCBS_PAGE", "1"),
		Size:        env.String("TCBS_SIZE", "20"),
		CustodyCode: env.RequiredString("TCBS_CUSTODY_CODE"),
	}
	response, err := broker.Native().Trading().GetMarginInformation(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
