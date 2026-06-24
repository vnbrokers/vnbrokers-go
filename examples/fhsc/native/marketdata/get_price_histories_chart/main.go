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
		APIKey:    env.RequiredString("FHSC_API_KEY"),
		APISecret: env.RequiredString("FHSC_API_SECRET"),
	})
	to := time.Now().Unix()
	from := time.Now().AddDate(0, 0, -30).Unix()
	request := nativedto.GetPriceHistoriesChartRequest{
		Symbol:     env.String("FHSC_SYMBOL", "HPG"),
		Resolution: env.String("FHSC_RESOLUTION", "D"),
		From:       env.Int("FHSC_FROM_UNIX", from),
		To:         env.Int("FHSC_TO_UNIX", to),
	}
	response, err := broker.Native().MarketData().GetPriceHistoriesChart(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
