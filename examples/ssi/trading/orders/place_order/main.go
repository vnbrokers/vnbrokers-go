package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
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
	price := decimal.NewFromInt(21000)
	response, err := broker.Trading().Orders().PlaceWithRequest(
		context.Background(),
		ssi.PlaceOrderRequest{
			PlaceOrderRequest: domain.PlaceOrderRequest{
				AccountID: mustEnv("SSI_ACCOUNT_NO"),
				Symbol:    "SSI",
				Side:      domain.OrderSideBuy,
				Type:      domain.OrderTypeLimit,
				Quantity:  decimal.NewFromInt(100),
				Price:     &price,
			},
			Code: "123456",
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
