package main

import (
	"context"
	"fmt"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/examples/entrade/internal/exampleutil"
)

func main() {
	broker := exampleutil.Broker()
	response, err := broker.Auth().Login(
		context.Background(),
		nativedto.LoginRequest{
			Username: exampleutil.MustEnv("ENTRADE_USERNAME"),
			Password: exampleutil.MustEnv("ENTRADE_PASSWORD"),
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("token:", response.Token)
}
