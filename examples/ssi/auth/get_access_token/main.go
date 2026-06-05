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
		ConsumerID:     mustEnv("SSI_FCTRADING_CONSUMER_ID"),
		ConsumerSecret: mustEnv("SSI_FCTRADING_CONSUMER_SECRET"),
	})
	response, err := broker.Auth().GetAccessToken(context.Background(), ssi.AccessTokenRequest{
		TwoFactorType: 1,
		Code:          mustEnv("SSI_FCTRADING_OTP"),
		IsSave:        false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", response.AccessToken)
	printJWTPart("header", response.AccessToken, 0)
	printJWTPart("payload", response.AccessToken, 1)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
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
