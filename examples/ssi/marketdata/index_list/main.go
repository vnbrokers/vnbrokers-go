package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{DataToken: mustEnv("SSI_FCDATA_TOKEN")})
	response, err := broker.Native().MarketData().GetIndexList(context.Background(), nativedto.GetIndexListRequest{
		Exchange:  "HOSE",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, item := range response.Data {
		fmt.Printf("%s: %s (%s)\n", item.IndexCode, item.IndexName, item.Exchange)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
