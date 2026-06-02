package main

import (
	"context"
	"os"

	"github.com/shopspring/decimal"
	vnbrokers "github.com/vnbrokers/vnbrokers-go"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func main() {
	symbol := "ACB"
	price := decimal.RequireFromString("23200")
	loanPackageID := 1775 // requiredInt("DNSE_LOAN_PACKAGE_ID")
	broker := vnbrokers.NewDNSE(dnse.Config{
		APIKey:       os.Getenv("DNSE_API_KEY"),
		APISecret:    os.Getenv("DNSE_API_SECRET"),
		TradingToken: os.Getenv("DNSE_TRADING_TOKEN"),
		MarketType:   "STOCK",
	})
	_, err := broker.Trading().Orders().PlaceWithRequest(context.Background(), dnse.PlaceOrderRequest{
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
