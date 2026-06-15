package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func main() {
	broker := vnbrokers.NewTCBS(tcbs.Config{AccessToken: mustEnv("TCBS_ACCESS_TOKEN")})
	request := nativedto.GetClosedDerivativePositionsRequest{AccountID: mustEnv("TCBS_ACCOUNT_ID"), SubAccountID: mustEnv("TCBS_SUB_ACCOUNT_ID"), Symbol: os.Getenv("TCBS_SYMBOL"), PageNo: envFloat("TCBS_PAGE_NO", 0), PageSize: envFloat("TCBS_PAGE_SIZE", 20)}
	response, err := broker.Native().Trading().GetClosedDerivativePositions(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}

func envDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func mustIntEnv(key string) int64 {
	value, err := strconv.ParseInt(mustEnv(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func mustFloatEnv(key string) float64 {
	value, err := strconv.ParseFloat(mustEnv(key), 64)
	if err != nil {
		panic(err)
	}
	return value
}

func envFloat(key string, fallback float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return parsed
}
