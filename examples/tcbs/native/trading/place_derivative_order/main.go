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
	request := nativedto.PlaceDerivativeOrderRequest{PlaceDerivativeOrderBody: nativedto.PlaceDerivativeOrderBody{AccountID: mustEnv("TCBS_ACCOUNT_ID"), Advance: os.Getenv("TCBS_ADVANCE"), OrderType: envDefault("TCBS_ORDER_TYPE", "LO"), Pin: mustEnv("TCBS_PIN"), Price: mustFloatEnv("TCBS_PRICE"), RefID: mustEnv("TCBS_REF_ID"), Side: mustEnv("TCBS_SIDE"), SubAccountID: mustEnv("TCBS_SUB_ACCOUNT_ID"), Symbol: mustEnv("TCBS_SYMBOL"), Volume: mustIntEnv("TCBS_VOLUME")}}
	response, err := broker.Native().Trading().PlaceDerivativeOrder(context.Background(), request)
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
