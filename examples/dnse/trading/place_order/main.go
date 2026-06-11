package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
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
	secdefs, err := broker.Native().MarketData().GetSecurityDefinition(ctx, nativedto.GetSecurityDefinitionRequest{Symbol: symbol, BoardID: "G1"})
	if err != nil {
		panic(err)
	}
	price, ok := secdefs.FloorPrice(symbol)
	if !ok {
		panic("floor price not found for " + symbol)
	}
	price = price.Mul(decimal.NewFromInt(1000))
	fmt.Println("Floor price: ", price.String())
	orderPrice, _ := price.Float64()

	_, err = broker.Native().Trading().PlaceOrder(ctx, nativedto.PlaceOrderRequest{
		AccountNo: os.Getenv("DNSE_ACCOUNT_NO"), Symbol: symbol, Side: "NB", OrderType: "LO",
		Quantity: 1, Price: &orderPrice, MarketType: "STOCK", OrderCategory: "NORMAL",
		LoanPackageID: &loanPackageID,
	})
	if err != nil {
		panic(err)
	}
}
