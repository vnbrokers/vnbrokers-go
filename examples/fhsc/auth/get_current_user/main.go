package main

import (
	"context"
	"fmt"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewFHSC(fhsc.Config{
		APIKey:         env.RequiredString("FHSC_API_KEY"),
		APISecret:      env.RequiredString("FHSC_API_SECRET"),
		UserID:         env.RequiredInt("FHSC_USER_ID"),
		TwoFactorToken: env.String("FHSC_2FA_TOKEN", ""),
	})
	response, err := broker.Auth().GetCurrentUser(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
