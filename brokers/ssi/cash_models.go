package ssi

import "github.com/shopspring/decimal"

type CashUnsettledSoldTransactionsRequest struct {
	AccountID  string
	SettleDate string
}

type CashTransferHistoriesRequest struct {
	AccountID string
	FromDate  string
	ToDate    string
}

type CashInAdvanceHistoriesRequest struct {
	AccountID string
	FromDate  string
	ToDate    string
}

type CashInAdvanceFeeRequest struct {
	AccountID     string
	CIAAmount     decimal.Decimal
	ReceiveAmount decimal.Decimal
}

type VSDCashDWRequest struct {
	AccountID string
	Amount    decimal.Decimal
	Type      string
	Remark    string
	Code      string
}

type CashTransferInternalRequest struct {
	AccountID          string
	BeneficiaryAccount string
	Amount             decimal.Decimal
	Remark             string
	Code               string
}

type CreateCashInAdvanceRequest struct {
	AccountID     string
	CIAAmount     decimal.Decimal
	ReceiveAmount decimal.Decimal
	Code          string
}

type CashInAdvanceAmountData struct {
	Account    string              `json:"account,omitempty"`
	CIAAmounts []CashInAdvanceItem `json:"ciaAmounts,omitempty"`
}

type CashInAdvanceItem struct {
	DueDate      string           `json:"dueDate,omitempty"`
	SellValue    *decimal.Decimal `json:"sellValue,omitempty"`
	NetSellValue *decimal.Decimal `json:"netSellValue,omitempty"`
	Advance      *decimal.Decimal `json:"advance,omitempty"`
	CashAdvance  *decimal.Decimal `json:"cashAdvance,omitempty"`
}

type UnsettledSoldTransactionsData struct {
	Account                   string                     `json:"account,omitempty"`
	UnsettledSoldTransactions []UnsettledSoldTransaction `json:"unsettledSoldTransactions,omitempty"`
}

type UnsettledSoldTransaction struct {
	TradeDate    string           `json:"tradeDate,omitempty"`
	InstrumentID string           `json:"instrumentID,omitempty"`
	NetSellValue *decimal.Decimal `json:"netSellValue,omitempty"`
	Quantity     *decimal.Decimal `json:"quantity,omitempty"`
	Price        *decimal.Decimal `json:"price,omitempty"`
}

type CashTransferHistoriesData struct {
	TransferHistories []CashTransferHistory `json:"transferHistories,omitempty"`
}

type CashInAdvanceHistoriesData struct {
	Account      string                 `json:"account,omitempty"`
	CIAHistories []CashInAdvanceHistory `json:"ciaHistories,omitempty"`
}

type CashInAdvanceHistory struct {
	TransactionID string                       `json:"transactionID,omitempty"`
	DateTime      string                       `json:"dateTime,omitempty"`
	TotalAmount   *decimal.Decimal             `json:"totalAmount,omitempty"`
	Details       []CashInAdvanceHistoryDetail `json:"details,omitempty"`
	Status        string                       `json:"status,omitempty"`
}

type CashInAdvanceHistoryDetail struct {
	Type       string           `json:"type,omitempty"`
	Value      *decimal.Decimal `json:"value,omitempty"`
	SettleDate string           `json:"settleDate,omitempty"`
}

type CashInAdvanceFeeData struct {
	Account       string           `json:"account,omitempty"`
	CIAAmount     *decimal.Decimal `json:"ciaAmount,omitempty"`
	ReceiveAmount *decimal.Decimal `json:"receiveAmount,omitempty"`
	Fee           *decimal.Decimal `json:"fee,omitempty"`
}

type TransactionResponse struct {
	Account       string `json:"account,omitempty"`
	TransactionID string `json:"transactionID,omitempty"`
}
