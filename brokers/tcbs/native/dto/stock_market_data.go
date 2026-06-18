// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type StockTicker struct {
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
	FloorPrice      decimal.Decimal `json:"floorPrice"`
	High            decimal.Decimal `json:"high"`
	IndexNumber     decimal.Decimal `json:"indexNumber"`
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
	Room            decimal.Decimal `json:"room"`
	SellForeignQtty decimal.Decimal `json:"sellForeignQtty"`
	Symbol          string          `json:"symbol"`
	TotalVal        decimal.Decimal `json:"totalVal"`
	TotalVol        decimal.Decimal `json:"totalVol"`
}

type StockForeignRoom struct {
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
	RefPrice        decimal.Decimal `json:"refPrice"`
	Room            decimal.Decimal `json:"room"`
	SellForeignQtty decimal.Decimal `json:"sellForeignQtty"`
	Symbol          string          `json:"symbol"`
	TotalValue      decimal.Decimal `json:"totalValue"`
	TotalVolume     decimal.Decimal `json:"totalVolume"`
}

type PutThroughAdvertisement struct {
	Color   decimal.Decimal `json:"color"`
	OrderID string          `json:"orderId"`
	Price   decimal.Decimal `json:"price"`
	Side    string          `json:"side"`
	Status  decimal.Decimal `json:"status"`
	Symbol  string          `json:"symbol"`
	Time    string          `json:"time"`
	Vol     decimal.Decimal `json:"vol"`
}

type PutThroughMatch struct {
	AccumulatedValue decimal.Decimal `json:"accumulatedValue"`
	Price            decimal.Decimal `json:"price"`
	Symbol           string          `json:"symbol"`
	Time             string          `json:"time"`
	Val              decimal.Decimal `json:"val"`
	Vol              decimal.Decimal `json:"vol"`
}

type StockTradeHistoryEntry struct {
	A   string          `json:"a"`
	Ba  decimal.Decimal `json:"ba"`
	Cp  decimal.Decimal `json:"cp"`
	Hl  bool            `json:"hl"`
	P   decimal.Decimal `json:"p"`
	Pcp decimal.Decimal `json:"pcp"`
	Rcp decimal.Decimal `json:"rcp"`
	Sa  decimal.Decimal `json:"sa"`
	T   string          `json:"t"`
	V   decimal.Decimal `json:"v"`
}

type StockSupplyDemand15MinutesEntry struct {
	Bms decimal.Decimal `json:"bms"`
	Bsr decimal.Decimal `json:"bsr"`
	Bu  decimal.Decimal `json:"bu"`
	Bup decimal.Decimal `json:"bup"`
	S   decimal.Decimal `json:"s"`
	Sd  decimal.Decimal `json:"sd"`
	Sdp decimal.Decimal `json:"sdp"`
	Sms decimal.Decimal `json:"sms"`
	T   string          `json:"t"`
}

type StockSupplyDemandEntry struct {
	Bsr decimal.Decimal `json:"bsr"`
	Bup decimal.Decimal `json:"bup"`
	Sdp decimal.Decimal `json:"sdp"`
	T   string          `json:"t"`
}
