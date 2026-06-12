package main

import (
	"context"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	response, err := exampleutil.Broker().Native().Trading().CancelDerivativeOrder(
		context.Background(),
		nativedto.CancelDerivativeOrderRequest{OrderID: "1110909"},
	)
	if err != nil {
		panic(err)
	}
	exampleutil.Print(response)
}
