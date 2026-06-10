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
	response, err := broker.MarketData().DailyStockPrice(context.Background(), dto.DailyStockPriceRequest{
		Symbol:    "SSI",
		Market:    "HOSE",
		FromDate:  "19/07/2023",
		ToDate:    "19/07/2023",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, item := range response.Data {
		fmt.Printf("%s %s close=%s totalValue=%s\n", item.Symbol, item.TradingDate, item.ClosePrice, item.TotalTradedValue)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
