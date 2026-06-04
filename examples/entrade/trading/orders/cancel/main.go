package main

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	payload, err := exampleutil.Broker().Trading().Orders().Cancel(
		context.Background(),
		"1110909",
	)
	if err != nil {
		panic(err)
	}
	exampleutil.PrintRaw(payload)
}
