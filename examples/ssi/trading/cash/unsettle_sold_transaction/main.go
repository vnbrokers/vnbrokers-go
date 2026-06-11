package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		TradingToken: mustEnv("SSI_FCTRADING_TOKEN"),
	})
	response, err := broker.Native().Trading().UnsettleSoldTransaction(
		context.Background(),
		mustEnv("SSI_ACCOUNT_NO"),
		"",
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, item := range response.Data.UnsettledSoldTransactions {
		fmt.Printf("tradeDate=%s symbol=%s netSellValue=%s qty=%s price=%s\n",
			item.TradeDate, item.InstrumentID, item.NetSellValue, item.Quantity, item.Price)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
