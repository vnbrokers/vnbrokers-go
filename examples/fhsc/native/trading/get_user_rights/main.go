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
	request := nativedto.GetUserRightsRequest{
		SubAccountID: env.RequiredString("FHSC_SUB_ACCOUNT_ID"),
		FromDate:     env.String("FHSC_FROM_DATE", time.Now().AddDate(0, -1, 0).Format("2006-01-02")),
		ToDate:       env.String("FHSC_TO_DATE", time.Now().Format("2006-01-02")),
		CatType:      env.String("FHSC_RIGHTS_CAT_TYPE", ""),
		IsCom:        env.String("FHSC_RIGHTS_IS_COM", ""),
		Symbol:       env.String("FHSC_SYMBOL", ""),
		Status:       env.String("FHSC_RIGHTS_STATUS", ""),
	}
	response, err := broker.Native().Trading().GetUserRights(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
