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
	quote, err := broker.Native().MarketData().GetLatestQuotes(context.Background(), nativedto.GetLatestQuotesRequest{Symbol: "ACB", BoardID: "G1"})
	if err != nil {
		panic(err)
	}
	fmt.Println(quote)
}
