package exampleutil

import (
	"encoding/json"
	"fmt"
	"os"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
)

func Broker() *entrade.Broker {
	return vnbrokers.NewEntrade(entrade.Config{
		Token: os.Getenv("ENTRADE_TOKEN"),
	})
}

func MustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}

func Print(value any) {
	pretty, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pretty))
}
