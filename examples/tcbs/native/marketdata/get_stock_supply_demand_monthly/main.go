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
	request := nativedto.GetStockSupplyDemandMonthlyRequest{
		Ticker:     env.String("TCBS_TICKER", "TCBS"),
		TimeWindow: env.String("TCBS_TIME_WINDOW", "1M"),
		Type:       env.String("TCBS_TYPE", "stock"),
	}
	response, err := broker.Native().MarketData().GetStockSupplyDemandMonthly(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
