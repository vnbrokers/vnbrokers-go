package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:    os.Getenv("DNSE_API_KEY"),
		APISecret: os.Getenv("DNSE_API_SECRET"),
	})
	payload, err := broker.Native().MarketData().GetSecurityDefinition(context.Background(), nativedto.GetSecurityDefinitionRequest{Symbol: "ACB", BoardID: "G1"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", payload)
}
