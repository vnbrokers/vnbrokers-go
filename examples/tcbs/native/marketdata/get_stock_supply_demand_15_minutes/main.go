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
	request := nativedto.GetStockSupplyDemand15MinutesRequest{
		Ticker:     env.String("TCBS_TICKER", "TCB"),
		TimeWindow: env.String("TCBS_TIME_WINDOW", "15"),
		TWindow:    env.String("TCBS_T_WINDOW", "15"),
		Type:       env.String("TCBS_TYPE", "all"),
	}
	response, err := broker.Native().MarketData().GetStockSupplyDemand15Minutes(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
