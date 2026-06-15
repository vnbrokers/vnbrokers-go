package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func main() {
	apiKey := mustEnv("TCBS_API_KEY")
	otp := mustEnv("TCBS_OTP")

	broker := vnbrokers.NewTCBS(tcbs.Config{})
	response, err := broker.Auth().GetToken(context.Background(), nativedto.GetTokenRequest{APIKey: apiKey, OTP: otp})
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", response.Token)
	printJWTPart("header", response.Token, 0)
	printJWTPart("payload", response.Token, 1)
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
