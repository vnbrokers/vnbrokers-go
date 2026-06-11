package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
)

func main() {
	broker := vnbrokers.NewSSI(ssi.Config{
		TradingToken: mustEnv("SSI_FCTRADING_TOKEN"),
	})
	response, err := broker.Native().Trading().FcoList(
		context.Background(),
		url.Values{
			"account":   {mustEnv("SSI_ACCOUNT_NO")},
			"pageIndex": {"1"},
			"pageSize":  {"50"},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status=%d message=%s\n", response.Status, response.Message)
	for _, item := range response.Data {
		fmt.Printf("fcoId=%s symbol=%s side=%s type=%s processStatus=%s created=%s\n",
			item.FCOID, item.InstrumentID, item.Side, item.Type, item.ProcessStatus, item.CreatedDate)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
