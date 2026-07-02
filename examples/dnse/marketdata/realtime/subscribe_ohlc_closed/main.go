package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:         os.Getenv("DNSE_API_KEY"),
		APISecret:      os.Getenv("DNSE_API_SECRET"),
		StreamEncoding: "msgpack",
	})
	sub, err := broker.Native().MarketData().Realtime().SubscribeClosedOHLC(context.Background(), nativedto.SubscribeOHLCRequest{Symbols: []string{"ACB"}, Resolution: "15"})
	if err != nil {
		panic(err)
	}
	defer sub.Close()
	for event := range sub.Events() {
		message, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(message))
	}
}
