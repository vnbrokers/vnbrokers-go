// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type TransferBetweenSubaccountsBody struct {
	Amount                   float64 `json:"amount"`
	Description              float64 `json:"description"`
	DestinationAccountNumber string  `json:"destinationAccountNumber"`
	SourceAccountNumber      string  `json:"sourceAccountNumber"`
}

type WithdrawDerivativeMarginBody struct {
	AccountID      string  `json:"accountId"`
	Amount         float64 `json:"amount"`
	PaymentContent float64 `json:"paymentContent"`
	SubAccountID   string  `json:"subAccountId"`
}

type DepositDerivativeMarginBody struct {
	AccountID      string  `json:"accountId"`
	Amount         float64 `json:"amount"`
	PaymentContent float64 `json:"paymentContent"`
	SubAccountID   string  `json:"subAccountId"`
}

type DerivativeMarginTransaction struct {
	TransactionID string `json:"transactionId"`
}
