// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type DerivativeTicker struct {
	Avg             decimal.Decimal `json:"avg"`
	BidPrice01      decimal.Decimal `json:"bidPrice01"`
	BidPrice02      decimal.Decimal `json:"bidPrice02"`
	BidPrice03      decimal.Decimal `json:"bidPrice03"`
	BidQtty01       decimal.Decimal `json:"bidQtty01"`
	BidQtty02       decimal.Decimal `json:"bidQtty02"`
	BidQtty03       decimal.Decimal `json:"bidQtty03"`
	BuyForeignQtty  decimal.Decimal `json:"buyForeignQtty"`
	CeilPrice       decimal.Decimal `json:"ceilPrice"`
	Change          decimal.Decimal `json:"change"`
	ChangePercent   decimal.Decimal `json:"changePercent"`
	ExpiryDate      string          `json:"expiryDate"`
	FloorPrice      decimal.Decimal `json:"floorPrice"`
	High            decimal.Decimal `json:"high"`
	Low             decimal.Decimal `json:"low"`
	MatchPrice      decimal.Decimal `json:"matchPrice"`
	MatchQtty       decimal.Decimal `json:"matchQtty"`
	OfferPrice01    decimal.Decimal `json:"offerPrice01"`
	OfferPrice02    decimal.Decimal `json:"offerPrice02"`
	OfferPrice03    decimal.Decimal `json:"offerPrice03"`
	OfferQtty01     decimal.Decimal `json:"offerQtty01"`
	OfferQtty02     decimal.Decimal `json:"offerQtty02"`
	OfferQtty03     decimal.Decimal `json:"offerQtty03"`
	Open            decimal.Decimal `json:"open"`
	OpenVol         decimal.Decimal `json:"openVol"`
	RefPrice        decimal.Decimal `json:"refPrice"`
	SellForeignQtty decimal.Decimal `json:"sellForeignQtty"`
	Symbol          string          `json:"symbol"`
	TotalVol        decimal.Decimal `json:"totalVol"`
}
