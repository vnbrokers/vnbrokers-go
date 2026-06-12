package main

import (
	"context"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	response, err := exampleutil.Broker().Native().Trading().CloseDerivativeDeal(
		context.Background(),
		nativedto.CloseDerivativeDealRequest{DealID: "1000546", OrderType: "LO"},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.Print(response)
}
