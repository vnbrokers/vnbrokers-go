package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func main() {
	ctx := context.Background()
	symbol := "AAA"
	loanPackageID := 1775 // GD Tien mat
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	secdef, err := broker.MarketData().Symbols().SecurityDefinition(ctx, symbol, "G1")
	if err != nil {
		panic(err)
	}
	var secdefs dnse.SecurityDefinitionList
	if err := dnse.UnmarshalRawPayload(secdef, &secdefs); err != nil {
		panic(err)
	}
	price, ok := secdefs.FloorPrice(symbol)
	if !ok {
		panic("floor price not found for " + symbol)
	}
	price = price.Mul(decimal.NewFromInt(1000))
	fmt.Println("Floor price: ", price.String())

	_, err = broker.Trading().Orders().PlaceWithRequest(ctx, dnse.PlaceOrderRequest{
		PlaceOrderRequest: domain.PlaceOrderRequest{
			AccountID: os.Getenv("DNSE_ACCOUNT_NO"),
			Symbol:    symbol,
			Side:      domain.OrderSideBuy,
			Type:      domain.OrderTypeLimit,
			Quantity:  decimal.NewFromInt(1),
			Price:     &price,
		},
		LoanPackageID: &loanPackageID,
	})
	if err != nil {
		panic(err)
	}
}
