package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		AccessToken: mustEnv("SSI_ACCESS_TOKEN"),
		PrivateKey:  mustEnv("SSI_PRIVATE_KEY"),
		DeviceID:    "FCTradingExample",
		UserAgent:   "FCTrading",
		MarketID:    "VN",
		ChannelID:   "TA",
	})
	response, err := broker.Trading().Orders().CancelWithRequest(
		context.Background(),
		ssi.CancelOrderRequest{
			AccountID: mustEnv("SSI_ACCOUNT_NO"),
			OrderID:   "12658867",
			Symbol:    "SSI",
			Side:      domain.OrderSideBuy,
			Code:      "123456",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response.Data)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
