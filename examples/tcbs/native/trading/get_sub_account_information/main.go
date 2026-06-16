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
	request := nativedto.GetSubAccountInformationRequest{
		CustodyCode: env.RequiredString("TCBS_CUSTODY_CODE"),
		Fields:      env.String("TCBS_FIELDS", "basicInfo,personalInfo,bankAccounts,bankSubAccounts"),
	}
	response, err := broker.Native().Trading().GetSubAccountInformation(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
