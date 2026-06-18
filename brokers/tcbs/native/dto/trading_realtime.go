package dto

import "github.com/shopspring/decimal"

type SubscribeStockOrdersRequest struct{}

type StockOrderEvent struct {
	Object      string          `json:"object"`
	AccountNo   string          `json:"accountNo"`
	OrderID     string          `json:"orderId"`
	ExecType    string          `json:"execType"`
	OrderQtty   decimal.Decimal `json:"orderQtty"`
	Symbol      string          `json:"symbol"`
	PriceType   string          `json:"priceType"`
	TxTime      string          `json:"txTime"`
	TxDate      string          `json:"txDate"`
	ExpDate     string          `json:"expDate"`
	TimeType    string          `json:"timeType"`
	OrStatus    string          `json:"orStatus"`
	LimitPrice  decimal.Decimal `json:"limitPrice"`
	RemainQtty  decimal.Decimal `json:"remainQtty"`
	Via         string          `json:"via"`
	QuotePrice  decimal.Decimal `json:"quotePrice"`
	TradePlace  string          `json:"tradePlace"`
	MatchType   string          `json:"matchType"`
	IsDisposal  string          `json:"isDisposal"`
	IsCancel    string          `json:"isCancel"`
	IsAmend     string          `json:"isAmend"`
	UserName    string          `json:"userName"`
	ORSOrderID  string          `json:"orsOrderId"`
	SecType     string          `json:"secType"`
	IsFOOrder   string          `json:"isFOOrder"`
	ODTimestamp string          `json:"odTimeStamp"`
}

type SubscribeDerivativeOrdersRequest struct{}

type DerivativeOrderEvent struct {
	SubAccount   string          `json:"subAccount"`
	OrderNo      string          `json:"orderNo"`
	PKOrderNo    string          `json:"pkOrderNo"`
	OrderTime    string          `json:"orderTime"`
	AccountCode  string          `json:"accountCode"`
	Side         string          `json:"side"`
	Symbol       string          `json:"symbol"`
	Volume       decimal.Decimal `json:"volume"`
	ShowPrice    string          `json:"showPrice"`
	MatchVolume  decimal.Decimal `json:"matchVolume"`
	MatchPriceBQ decimal.Decimal `json:"matchPriceBQ"`
	Status       string          `json:"status"`
	OrderStatus  string          `json:"orderStatus"`
	Channel      string          `json:"channel"`
	Group        string          `json:"group"`
	CancelTime   string          `json:"cancelTime"`
	IsCancel     string          `json:"isCancel"`
	IsAmend      string          `json:"isAmend"`
	Info         string          `json:"info"`
	MaxPrice     decimal.Decimal `json:"maxPrice"`
	MatchValue   decimal.Decimal `json:"matchValue"`
	Quote        string          `json:"quote"`
	AutoType     string          `json:"autoType"`
	Product      string          `json:"product"`
	OrderType    string          `json:"orderType"`
	Source       string          `json:"source"`
	TraderCode   string          `json:"traderCode"`
}

type SubscribeDerivativeOpenPositionsRequest struct{}

type DerivativeOpenPositionEvent struct {
	SubAccount    string          `json:"subAccount"`
	Symbol        string          `json:"symbol"`
	Side          string          `json:"side"`
	Account       string          `json:"account"`
	LastPrice     decimal.Decimal `json:"lastPrice"`
	AvgRemain     decimal.Decimal `json:"avgRemain"`
	IMValue       decimal.Decimal `json:"imValue"`
	Net           decimal.Decimal `json:"net"`
	StopLoss      decimal.Decimal `json:"stoploss"`
	TakeProfit    decimal.Decimal `json:"takeprofit"`
	PCRemain      decimal.Decimal `json:"pcRemain"`
	VMRemain      decimal.Decimal `json:"vmRemain"`
	DueDate       string          `json:"duedate"`
	NetOffVolume  decimal.Decimal `json:"netoffvol"`
	TriggerType   string          `json:"triggerType"`
	CallbackPoint decimal.Decimal `json:"callBackPoint"`
	TrailingPrice decimal.Decimal `json:"trailingPrice"`
	TotalVMValue  decimal.Decimal `json:"totalVmValue"`
}

type SubscribeDerivativeConditionalOrdersRequest struct{}

type DerivativeConditionalOrderEvent struct {
	SubAccount       string          `json:"subAccount"`
	OrderNo          string          `json:"orderNo"`
	GroupOrder       string          `json:"groupOrder"`
	PKOrderNo        string          `json:"pkOrderNo"`
	AccountCode      string          `json:"accountCode"`
	Side             string          `json:"side"`
	Symbol           string          `json:"symbol"`
	Volume           decimal.Decimal `json:"volume"`
	ShowPrice        string          `json:"showPrice"`
	Condition        string          `json:"condition"`
	Result           string          `json:"result"`
	ActiveTime       string          `json:"activeTime"`
	SendTime         string          `json:"sendTime"`
	CancelTime       string          `json:"cancelTime"`
	Group            string          `json:"group"`
	Channel          string          `json:"channel"`
	SOPrice          decimal.Decimal `json:"soPrice"`
	OrderType        string          `json:"orderType"`
	FromTime         string          `json:"fromTime"`
	ExpirationTime   string          `json:"expTime"`
	Status           string          `json:"status"`
	Details          string          `json:"details"`
	Message          string          `json:"message"`
	Notes            string          `json:"notes"`
	CallbackPoint    string          `json:"callBackPoint"`
	TrailingPrice    string          `json:"trailingPrice"`
	TriggerCondition string          `json:"triggerCondition"`
}
