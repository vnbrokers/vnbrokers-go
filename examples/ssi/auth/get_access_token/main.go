package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		ConsumerID:            mustEnv("SSI_FCTRADING_CONSUMER_ID"),
		DataConsumerSecret:    mustEnv("SSI_FCDATA_CONSUMER_SECRET"),
		TradingConsumerSecret: mustEnv("SSI_FCTRADING_CONSUMER_SECRET"),
	})
	ctx := context.Background()
	dataResponse, err := broker.Auth().GetAccessToken(ctx)
	if err != nil {
		panic(err)
	}
	printToken("data", dataResponse.AccessToken)

	tradingResponse, err := broker.Auth().GetTradingToken(ctx, ssi.TradingTokenRequest{
		TwoFactorType: 1,
		Code:          mustEnv("SSI_FCTRADING_OTP"),
		IsSave:        false,
	})
	if err != nil {
		panic(err)
	}
	printToken("trading", tradingResponse.AccessToken)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}

func printToken(label string, token string) {
	fmt.Printf("%s token: %s\n", label, token)
	printJWTPart(label+" header", token, 0)
	printJWTPart(label+" payload", token, 1)
}

func printJWTPart(label string, token string, index int) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		panic("invalid jwt token")
	}
	decoded, err := base64.RawURLEncoding.DecodeString(parts[index])
	if err != nil {
		panic(err)
	}
	var data any
	if err := json.Unmarshal(decoded, &data); err != nil {
		panic(err)
	}
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s: %s\n", label, pretty)
}
