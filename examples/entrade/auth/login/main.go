package main

import (
	"context"
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	broker := exampleutil.Broker()
	response, err := broker.Auth().Login(
		context.Background(),
		exampleutil.MustEnv("ENTRADE_USERNAME"),
		exampleutil.MustEnv("ENTRADE_PASSWORD"),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", response.Token)
}
