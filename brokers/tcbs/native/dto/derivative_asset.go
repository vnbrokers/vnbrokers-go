// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type DerivativeCash struct {
	Assets           decimal.Decimal `json:"assets"`
	AvaiCash         decimal.Decimal `json:"avaiCash"`
	AvaiColla        decimal.Decimal `json:"avaiColla"`
	Cash             decimal.Decimal `json:"cash"`
	CashOut          decimal.Decimal `json:"cashOut"`
	CashWithdraw     decimal.Decimal `json:"cashWithdraw"`
	Cashavaiwithdraw decimal.Decimal `json:"cashavaiwithdraw"`
	Collateral       decimal.Decimal `json:"collateral"`
	Color            string          `json:"color"`
	Debt             string          `json:"debt"`
	DM               decimal.Decimal `json:"dm"`
	FeeCTCK          decimal.Decimal `json:"feeCTCK"`
	FeeHNX           decimal.Decimal `json:"feeHNX"`
	FeeMan           decimal.Decimal `json:"feeMan"`
	FeePos           decimal.Decimal `json:"feePos"`
	IM               decimal.Decimal `json:"im"`
	Info             string          `json:"info"`
	Limit            decimal.Decimal `json:"limit"`
	MR               decimal.Decimal `json:"mr"`
	NAV              decimal.Decimal `json:"nav"`
	Net              string          `json:"net"`
	Others           decimal.Decimal `json:"others"`
	Package          string          `json:"package"`
	Product          string          `json:"product"`
	Status           string          `json:"status"`
	Stock            decimal.Decimal `json:"stock"`
	Tax              decimal.Decimal `json:"tax"`
	Tienbosung       decimal.Decimal `json:"tienbosung"`
	Tyle             string          `json:"tyle"`
	Type             string          `json:"type"`
	UnrelizeVM       decimal.Decimal `json:"unrelizeVM"`
	VM               decimal.Decimal `json:"vm"`
	VmEod            string          `json:"vm_eod"`
	Vmunpay          decimal.Decimal `json:"vmunpay"`
	W1               decimal.Decimal `json:"w1"`
	W2               decimal.Decimal `json:"w2"`
}

type ClosedDerivativePosition struct {
	ClosePC       string `json:"closePC"`
	ClosePosition string `json:"closePosition"`
	ClosePrice    string `json:"closePrice"`
	CloseVM       string `json:"closeVM"`
	Fee           string `json:"fee"`
	OpenPrice     string `json:"openPrice"`
	Side          string `json:"side"`
	Symbol        string `json:"symbol"`
	Tax           string `json:"tax"`
	Time          string `json:"time"`
	Unrealize     string `json:"unrealize"`
}

type OpenDerivativePosition struct {
	Account    string          `json:"account"`
	AvgRemain  decimal.Decimal `json:"avg_remain"`
	Deliver    decimal.Decimal `json:"deliver"`
	Duedate    string          `json:"duedate"`
	IM         string          `json:"im"`
	ImValue    decimal.Decimal `json:"imValue"`
	LastPrice  decimal.Decimal `json:"lastPrice"`
	MrValue    decimal.Decimal `json:"mrValue"`
	Net        decimal.Decimal `json:"net"`
	Netoffvol  decimal.Decimal `json:"netoffvol"`
	PcRemain   decimal.Decimal `json:"pc_remain"`
	Receive    decimal.Decimal `json:"receive"`
	Side       string          `json:"side"`
	Stoploss   string          `json:"stoploss"`
	Symbol     string          `json:"symbol"`
	Takeprofit string          `json:"takeprofit"`
	VmValue    decimal.Decimal `json:"vmValue"`
	VmRemain   decimal.Decimal `json:"vm_remain"`
	Wapb       decimal.Decimal `json:"wapb"`
	Wasp       decimal.Decimal `json:"wasp"`
}
