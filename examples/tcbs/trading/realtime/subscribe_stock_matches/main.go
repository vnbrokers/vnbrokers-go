package main

import (
	"context"
	stderrors "errors"
	"fmt"
	"os"
	"os/signal"

	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	broker := vnbrokers.NewTCBS(tcbs.Config{
		AccessToken: mustEnv("TCBS_ACCESS_TOKEN"),
	})
	subscription, err := broker.Trading().Realtime().SubscribeStockMatches(ctx)
	if err != nil {
		panic(err)
	}
	defer subscription.Close()

	for {
		select {
		case event, ok := <-subscription.Events():
			if !ok {
				return
			}
			fmt.Printf("event: %+v\n", event)
		case err, ok := <-subscription.Errors():
			if !ok {
				return
			}
			printError(err)
		case status, ok := <-subscription.Status():
			if !ok {
				return
			}
			fmt.Printf("status: %s\n", status)
		case <-ctx.Done():
			return
		}
	}
}

func printError(err error) {
	fmt.Printf("error: %v\n", err)
	var brokerErr *sdkerrors.BrokerError
	if !stderrors.As(err, &brokerErr) {
		return
	}
	fmt.Printf("error category: %s\n", brokerErr.Category)
	if brokerErr.Operation != "" {
		fmt.Printf("error operation: %s\n", brokerErr.Operation)
	}
	if brokerErr.Cause != nil {
		fmt.Printf("error cause: %v\n", brokerErr.Cause)
	}
	if brokerErr.Raw != nil {
		fmt.Printf("error raw: %s\n", rawString(brokerErr.Raw))
	}
}

func rawString(raw any) string {
	switch value := raw.(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case transport.WebSocketMessage:
		return string(value)
	default:
		return fmt.Sprintf("%#v", raw)
	}
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " is required")
	}
	return value
}
