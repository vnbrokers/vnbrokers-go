package main

import (
	"context"
	"fmt"
	"strings"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/internal/env"
)

func main() {
	broker := vnbrokers.NewFHSC(fhsc.Config{
		APIKey:         env.RequiredString("FHSC_API_KEY"),
		APISecret:      env.RequiredString("FHSC_API_SECRET"),
		UserID:         env.RequiredInt("FHSC_USER_ID"),
		TwoFactorToken: env.String("FHSC_2FA_TOKEN", ""),
	})
	orderType := strings.ToUpper(env.String("FHSC_ORDER_TYPE", "LIMIT"))
	body := nativedto.CreateOrderRequest{
		SubAccount: env.RequiredString("FHSC_SUB_ACCOUNT_EXT"),
		Side:       strings.ToUpper(env.String("FHSC_ORDER_SIDE", "BUY")),
		Symbol:     env.String("FHSC_SYMBOL", "HPG"),
		Quantity:   env.Int("FHSC_ORDER_QTY", 100),
		TypeValue:  orderType,
		StockType:  strings.ToUpper(env.String("FHSC_STOCK_TYPE", "STOCK")),
	}
	if orderType == "LIMIT" {
		limitPrice := env.RequiredInt("FHSC_LIMIT_PRICE")
		body.LimitPrice = &limitPrice
	} else {
		marketPrice := strings.ToUpper(env.String("FHSC_MARKET_PRICE", "ATO"))
		body.MarketPrice = &marketPrice
	}
	request := nativedto.PlaceOrderRequest{
		SubAccountID: env.RequiredString("FHSC_SUB_ACCOUNT_ID"),
		Body:         body,
	}
	response, err := broker.Native().Trading().PlaceOrder(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", response)
}
