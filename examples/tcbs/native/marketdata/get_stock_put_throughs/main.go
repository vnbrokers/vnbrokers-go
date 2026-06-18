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
	request := nativedto.GetStockPutThroughsRequest{Floor: env.Float("TCBS_FLOOR", 1)}
	response, err := broker.Native().MarketData().GetStockPutThroughs(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
