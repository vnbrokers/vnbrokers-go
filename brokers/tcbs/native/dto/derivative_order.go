// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type DerivativeOrder struct {
	AccountCode  string          `json:"accountCode"`
	AutoType     string          `json:"autoType"`
	CancelTime   string          `json:"cancelTime"`
	Channel      string          `json:"channel"`
	Group        string          `json:"group"`
	Info         string          `json:"info"`
	IsAmend      decimal.Decimal `json:"isAmend"`
	IsCancel     decimal.Decimal `json:"isCancel"`
	MatchPriceBQ decimal.Decimal `json:"matchPriceBQ"`
	MatchValue   decimal.Decimal `json:"matchValue"`
	MatchVolume  decimal.Decimal `json:"matchVolume"`
	MaxPrice     decimal.Decimal `json:"maxPrice"`
	OrderNo      string          `json:"orderNo"`
	OrderStatus  string          `json:"orderStatus"`
	OrderTime    string          `json:"orderTime"`
	OrderType    string          `json:"orderType"`
	PkOrderNo    string          `json:"pk_orderNo"`
	Product      string          `json:"product"`
	Quote        string          `json:"quote"`
	RefID        string          `json:"refId"`
	ShowPrice    decimal.Decimal `json:"showPrice"`
	Side         string          `json:"side"`
	Source       string          `json:"source"`
	Status       string          `json:"status"`
	Symbol       string          `json:"symbol"`
	Volume       decimal.Decimal `json:"volume"`
}

type DerivativeConditionalOrder struct {
	AccountCode string          `json:"accountCode"`
	ActiveTime  string          `json:"active_time"`
	CancelTime  string          `json:"cancel_time"`
	Channel     string          `json:"channel"`
	Condition   string          `json:"condition"`
	Details     string          `json:"details"`
	ExpTime     string          `json:"exp_time"`
	FromTime    string          `json:"from_time"`
	Group       string          `json:"group"`
	GroupOrder  string          `json:"groupOrder"`
	MaxPrice    string          `json:"maxPrice"`
	Notes       string          `json:"notes"`
	OrderNo     string          `json:"orderNo"`
	OrderType   string          `json:"orderType"`
	PkOrderNo   string          `json:"pk_orderNo"`
	Result      string          `json:"result"`
	SendTime    string          `json:"send_time"`
	ShowPrice   decimal.Decimal `json:"showPrice"`
	Side        string          `json:"side"`
	SoPrice     decimal.Decimal `json:"soPrice"`
	Status      string          `json:"status"`
	Symbol      string          `json:"symbol"`
	Volume      decimal.Decimal `json:"volume"`
}

type PlaceDerivativeOrderBody struct {
	AccountID    string  `json:"accountId"`
	Advance      string  `json:"advance"`
	OrderType    string  `json:"orderType"`
	Pin          string  `json:"pin"`
	Price        float64 `json:"price"`
	RefID        string  `json:"refId"`
	Side         string  `json:"side"`
	SubAccountID string  `json:"subAccountId"`
	Symbol       string  `json:"symbol"`
	Volume       int64   `json:"volume"`
}

type UpdateDerivativeOrderBody struct {
	AccountID    string  `json:"accountId"`
	Nprice       float64 `json:"nprice"`
	Nvol         float64 `json:"nvol"`
	OrderNo      string  `json:"orderNo"`
	RefID        string  `json:"refId"`
	SubAccountID string  `json:"subAccountId"`
}

type UpdateDerivativeConditionalOrderBody struct {
	AccountID string  `json:"accountId"`
	Cmd       string  `json:"cmd"`
	PkOrderNo string  `json:"pkOrderNo"`
	RefID     string  `json:"refId"`
	SoPrice   float64 `json:"soPrice"`
	Type      string  `json:"type"`
}

type PlaceDerivativeConditionalOrderBody struct {
	AccountID       string  `json:"accountId"`
	ActivationPrice float64 `json:"activationPrice"`
	Advance         string  `json:"advance"`
	CallbackPoint   float64 `json:"callbackPoint"`
	Cmd             string  `json:"cmd"`
	OrderType       string  `json:"orderType"`
	Pin             string  `json:"pin"`
	Price           float64 `json:"price"`
	RefID           string  `json:"refId"`
	Side            string  `json:"side"`
	SoPrice         float64 `json:"soPrice"`
	SubAccountID    string  `json:"subAccountId"`
	Symbol          string  `json:"symbol"`
	Type            string  `json:"type"`
	Volume          float64 `json:"volume"`
}

type PlaceDerivativeOrderResult struct {
	AccType     string          `json:"accType"`
	AccountCode string          `json:"accountCode"`
	AutoType    string          `json:"autoType"`
	Channel     string          `json:"channel"`
	Group       string          `json:"group"`
	Market      string          `json:"market"`
	MatchVolume decimal.Decimal `json:"matchVolume"`
	MsgType     string          `json:"msg_type"`
	OrderNo     string          `json:"orderNo"`
	OrderTime   string          `json:"orderTime"`
	PkOrderNo   string          `json:"pk_orderNo"`
	Product     string          `json:"product"`
	Quote       string          `json:"quote"`
	RefID       string          `json:"refID"`
	ShareStatus string          `json:"shareStatus"`
	ShowPrice   decimal.Decimal `json:"showPrice"`
	Side        string          `json:"side"`
	Status      string          `json:"status"`
	Symbol      string          `json:"symbol"`
	Type        string          `json:"type"`
	Volume      decimal.Decimal `json:"volume"`
}

type UpdateDerivativeOrderResult struct {
	MsgType   string          `json:"msg_type"`
	OrderNo   string          `json:"orderNo"`
	PkOrderNo string          `json:"pk_orderNo"`
	ShowPrice decimal.Decimal `json:"showPrice"`
	Status    string          `json:"status"`
	Volume    decimal.Decimal `json:"volume"`
}

type UpdateDerivativeConditionalOrderResult struct {
	Notes     string `json:"notes"`
	PkOrderNo string `json:"pkOrderNo"`
	Price     string `json:"price"`
	ShowPrice string `json:"showPrice"`
	Volume    string `json:"volume"`
}

type PlaceDerivativeConditionalOrderResult struct {
	AccType     string          `json:"accType"`
	AccountCode string          `json:"accountCode"`
	AutoType    string          `json:"autoType"`
	Channel     string          `json:"channel"`
	Group       string          `json:"group"`
	Market      string          `json:"market"`
	MatchVolume decimal.Decimal `json:"matchVolume"`
	MsgType     string          `json:"msg_type"`
	OrderNo     decimal.Decimal `json:"orderNo"`
	OrderTime   string          `json:"orderTime"`
	PkOrderNo   string          `json:"pk_orderNo"`
	Product     string          `json:"product"`
	Quote       string          `json:"quote"`
	RefID       string          `json:"refID"`
	ShareStatus string          `json:"shareStatus"`
	ShowPrice   decimal.Decimal `json:"showPrice"`
	Side        string          `json:"side"`
	Status      string          `json:"status"`
	Symbol      string          `json:"symbol"`
	Type        string          `json:"type"`
	Volume      decimal.Decimal `json:"volume"`
}

type CancelDerivativeOrderResult struct {
	CancelTime string `json:"cancelTime"`
	MsgType    string `json:"msg_type"`
	OrderNo    string `json:"orderNo"`
	PkOrderNo  string `json:"pk_orderNo"`
	Status     string `json:"status"`
}

type CancelDerivativeConditionalOrderResult struct {
}

type CancelDerivativeOrderBody struct {
	AccountID string `json:"accountId"`
	Cmd       string `json:"cmd"`
	OrderNo   string `json:"orderNo"`
	Pin       string `json:"pin"`
	RefID     string `json:"refId"`
}

type CancelDerivativeConditionalOrderBody struct {
	AccountID    string `json:"accountId"`
	OrderNo      string `json:"orderNo"`
	SubAccountID string `json:"subAccountId"`
}
