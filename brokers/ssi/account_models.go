package ssi

import "github.com/shopspring/decimal"

type MaxSellQuantityRequest struct {
	AccountID string
	Symbol    string
	Price     decimal.Decimal
}

type DerivativePositionRequest struct {
	AccountID    string
	QuerySummary bool
}
