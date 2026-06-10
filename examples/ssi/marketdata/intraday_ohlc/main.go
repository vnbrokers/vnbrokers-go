package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{DataToken: mustEnv("SSI_FCDATA_TOKEN")})
	response, err := broker.Native().MarketData().GetIntradayOhlc(context.Background(), nativedto.GetIntradayOhlcRequest{
		Symbol:     "SSI",
		FromDate:   "14/08/2023",
		ToDate:     "14/08/2023",
		PageIndex:  1,
		PageSize:   10,
		Resolution: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, item := range response.Data {
		fmt.Printf("%s %v O=%s H=%s L=%s C=%s V=%s\n", item.Symbol, item.Time, item.Open, item.High, item.Low, item.Close, item.Volume)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
