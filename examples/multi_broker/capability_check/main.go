package main

import (
	"fmt"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func main() {
	broker := vnbrokers.NewDNSE(dnse.Config{})
	fmt.Println(broker.Supports(core.CapabilityMarketDataRealtimeTop))
}
