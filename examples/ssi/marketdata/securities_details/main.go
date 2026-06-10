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
	response, err := broker.MarketData().SecuritiesDetails(context.Background(), dto.SecuritiesDetailsRequest{
		Market:    "HOSE",
		Symbol:    "SSI",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, group := range response.Data {
		fmt.Printf("report=%s totalSymbols=%s\n", group.ReportDate, group.TotalNoSym)
		for _, item := range group.RepeatedInfo {
			fmt.Printf("%s: %s (%s)\n", item.Symbol, item.SymbolName, item.SymbolEngName)
		}
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
