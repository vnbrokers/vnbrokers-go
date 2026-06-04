package main

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	symbols, err := exampleutil.Broker().MarketData().Derivatives().List(context.Background())
	if err != nil {
		panic(err)
	}
	exampleutil.Print(symbols)
}
