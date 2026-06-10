package main

import (
	"context"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/dto"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{DataToken: mustEnv("SSI_FCDATA_TOKEN")})
	response, err := broker.MarketData().IndexComponents(context.Background(), dto.IndexComponentsRequest{
		IndexCode: "VN30",
		PageIndex: 1,
		PageSize:  50,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("status=%s message=%s total=%d\n", response.Status, response.Message, response.TotalRecord)
	for _, group := range response.Data {
		fmt.Printf("%s %s (%s)\n", group.IndexCode, group.IndexName, group.Exchange)
		for _, item := range group.IndexComponent {
			fmt.Printf("- %s\n", item.StockSymbol)
		}
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
