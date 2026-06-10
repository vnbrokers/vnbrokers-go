package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/dto"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{DataToken: mustEnv("SSI_FCDATA_TOKEN")})
	response, err := broker.MarketData().DailyOhlc(context.Background(), dto.DailyOhlcRequest{
		Symbol:    "SSI",
		FromDate:  "10/08/2023",
		ToDate:    "13/08/2023",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, item := range response.Data {
		fmt.Printf("%s %s O=%s H=%s L=%s C=%s V=%s\n", item.Symbol, item.TradingDate, item.Open, item.High, item.Low, item.Close, item.Volume)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
