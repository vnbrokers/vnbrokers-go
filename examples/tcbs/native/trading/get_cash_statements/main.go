package main

import (
	"context"
	"fmt"
	"time"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: env.RequiredString("TCBS_ACCESS_TOKEN")})
	request := nativedto.GetCashStatementsRequest{
		AccountNo:       env.RequiredString("TCBS_ACCOUNT_NO"),
		FromDate:        env.String("TCBS_FROM_DATE", time.Now().AddDate(0, -1, 0).Format("2006-01-02")),
		ToDate:          env.String("TCBS_TO_DATE", time.Now().Format("2006-01-02")),
		PageSize:        env.String("TCBS_PAGE_SIZE", "20"),
		PageIndex:       env.String("TCBS_PAGE_INDEX", "1"),
		TransactionCode: "",
	}
	response, err := broker.Native().Trading().GetCashStatements(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
