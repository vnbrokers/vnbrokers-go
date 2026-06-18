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
	request := nativedto.GetStockTradeHistoryRequest{
		Ticker:    env.String("TCBS_TICKER", "TCX"),
		Page:      env.Float("TCBS_PAGE", 0),
		Size:      env.Float("TCBS_SIZE", 20),
		HeadIndex: env.Float("TCBS_HEAD_INDEX", 0),
	}
	response, err := broker.Native().MarketData().GetStockTradeHistory(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
