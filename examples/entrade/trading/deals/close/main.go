package main

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	payload, err := exampleutil.Broker().Trading().Deals().Close(
		context.Background(),
		"1000546",
		"LO",
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
