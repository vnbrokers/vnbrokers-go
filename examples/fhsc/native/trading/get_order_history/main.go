package main

import (
	"context"
	"fmt"
	"time"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewFHSC(fhsc.Config{
		APIKey:         env.RequiredString("FHSC_API_KEY"),
		APISecret:      env.RequiredString("FHSC_API_SECRET"),
		UserID:         env.RequiredInt("FHSC_USER_ID"),
		TwoFactorToken: env.String("FHSC_2FA_TOKEN", ""),
	})
	request := nativedto.GetOrderHistoryRequest{
		SubAccountID: env.RequiredString("FHSC_SUB_ACCOUNT_ID"),
		FromDate:     env.String("FHSC_FROM_DATE", time.Now().AddDate(0, 0, -30).Format("2006-01-02")),
		ToDate:       env.String("FHSC_TO_DATE", time.Now().Format("2006-01-02")),
		Page:         env.Int("FHSC_PAGE", 1),
		OrderStatus:  env.String("FHSC_ORDER_STATUS", ""),
		Symbol:       env.String("FHSC_SYMBOL", ""),
	}
	response, err := broker.Native().Trading().GetOrderHistory(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
