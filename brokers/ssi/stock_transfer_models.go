package ssi

import "github.com/shopspring/decimal"

type StockTransferHistoriesRequest struct {
	AccountID string
	StartDate string
	EndDate   string
}

type StockTransferRequest struct {
	AccountID          string
	BeneficiaryAccount string
	ExchangeID         string
	Symbol             string
	Quantity           decimal.Decimal
	Code               string
}
