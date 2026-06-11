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
		PrivateKey:   mustEnv("SSI_FCTRADING_PRIVATE_KEY"),
		DeviceID:     "Example",
		UserAgent:    "FCTrading",
		MarketID:     "VN",
		ChannelID:    "TA",
	})
	response, err := broker.Native().Trading().NewOrder(
		context.Background(),
		map[string]any{
			"account":      mustEnv("SSI_ACCOUNT_NO"),
			"instrumentID": "SSI",
			"marketID":     "VN",
			"buySell":      "B",
			"orderType":    "LO",
			"price":        21000,
			"quantity":     100,
			"code":         "123456",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	fmt.Printf("requestID=%s\n", response.Data.RequestID)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
