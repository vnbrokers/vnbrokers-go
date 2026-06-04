package main

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	payload, err := exampleutil.Broker().Trading().Risk().Config(
		context.Background(),
		exampleutil.MustEnv("ENTRADE_INVESTOR_ACCOUNT_ID"),
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
