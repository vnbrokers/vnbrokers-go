// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type PlaceStockOrderBody struct {
	ExecType  string `json:"execType"`
	Price     int64  `json:"price"`
	PriceType string `json:"priceType"`
	Quantity  int64  `json:"quantity"`
	Symbol    string `json:"symbol"`
}

type PlaceStockOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type UpdateStockOrderBody struct {
	Price    int64 `json:"price"`
	Quantity int64 `json:"quantity"`
}

type UpdateStockOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type CancelStockOrderBody struct {
	OrdersList []StockOrderID `json:"ordersList"`
}

type CancelStockOrderResponse struct {
	Data       []CancelStockOrderResult `json:"data"`
	Object     string                   `json:"object"`
	PageIndex  decimal.Decimal          `json:"pageIndex"`
	PageSize   decimal.Decimal          `json:"pageSize"`
	TotalCount decimal.Decimal          `json:"totalCount"`
}

type CancelStockOrderDetail struct {
	Deleted     string `json:"deleted"`
	ErrorCode   string `json:"errorCode"`
	ErrorMesage string `json:"errorMesage"`
	OrderID     string `json:"orderID"`
}

type CancelStockOrderResult struct {
	Details []CancelStockOrderDetail `json:"details"`
	Object  string                   `json:"object"`
}

type StockOrderID struct {
	OrderID string `json:"orderID"`
}

type StockOrderPage struct {
	Data       []StockOrder    `json:"data"`
	Object     string          `json:"object"`
	PageIndex  decimal.Decimal `json:"pageIndex"`
	PageSize   decimal.Decimal `json:"pageSize"`
	TotalCount decimal.Decimal `json:"totalCount"`
}

type StockOrder struct {
	AccountNo    string          `json:"accountNo"`
	BRatio       decimal.Decimal `json:"bRatio"`
	CancelQtty   decimal.Decimal `json:"cancelQtty"`
	CodeID       string          `json:"codeID"`
	ExecQtty     decimal.Decimal `json:"execQtty"`
	ExecType     string          `json:"execType"`
	ExpDate      string          `json:"expDate"`
	FeeAcr       decimal.Decimal `json:"feeAcr"`
	IsAmend      string          `json:"isAmend"`
	IsCancel     string          `json:"isCancel"`
	IsDisposal   string          `json:"isDisposal"`
	IsFOOrder    string          `json:"isFOOrder"`
	LimitPrice   decimal.Decimal `json:"limitPrice"`
	MatchAmount  decimal.Decimal `json:"matchAmount"`
	MatchPrice   decimal.Decimal `json:"matchPrice"`
	MatchType    string          `json:"matchType"`
	MMType       string          `json:"mmType"`
	Object       string          `json:"object"`
	OdTimeStamp  string          `json:"odTimeStamp"`
	OrStatus     string          `json:"orStatus"`
	OrderID      string          `json:"orderID"`
	OrderQtty    decimal.Decimal `json:"orderQtty"`
	OrsOrderID   string          `json:"orsOrderID"`
	PriceType    string          `json:"priceType"`
	QuotePrice   decimal.Decimal `json:"quotePrice"`
	RemainQtty   decimal.Decimal `json:"remainQtty"`
	SecType      string          `json:"sectype"`
	Symbol       string          `json:"symbol"`
	TaxSellAmout decimal.Decimal `json:"taxSellAmout"`
	TimeType     string          `json:"timeType"`
	TradePlace   string          `json:"tradePlace"`
	TxDate       string          `json:"txdate"`
	TxTime       string          `json:"txtime"`
	UserName     string          `json:"userName"`
	Via          string          `json:"via"`
}

type StockMatchingDetails struct {
	Data       []StockMatchingDetail `json:"data"`
	Object     string                `json:"object"`
	PageIndex  decimal.Decimal       `json:"pageIndex"`
	PageSize   decimal.Decimal       `json:"pageSize"`
	TotalCount decimal.Decimal       `json:"totalCount"`
}

type StockMatchingDetail struct {
	OrderID    string          `json:"orderId"`
	Price      decimal.Decimal `json:"price"`
	Qtty       decimal.Decimal `json:"qtty"`
	QuotePrice decimal.Decimal `json:"quotePrice"`
	QuoteQtty  decimal.Decimal `json:"quoteQtty"`
	Side       string          `json:"side"`
	Symbol     string          `json:"symbol"`
	TimeExec   decimal.Decimal `json:"timeExec"`
	TradeID    string          `json:"tradeId"`
}
