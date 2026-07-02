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
	sub, err := broker.Native().MarketData().Realtime().SubscribeForeign(context.Background(), nativedto.SubscribeSymbolsRequest{Symbols: []string{"ACB", "BID", "BSR", "CTG", "FPT", "GAS", "GVR", "HDB", "HPG", "LPB", "MBB", "MSN", "MWG", "PLX", "SAB", "SHB", "SSB", "SSI", "STB", "TCB", "TPB", "VCB", "VHM", "VIB", "VIC", "VJC", "VNM", "VPB", "VPL", "VRE"}, BoardID: "G1"})
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
