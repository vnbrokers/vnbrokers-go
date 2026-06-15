// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type DerivativeTicker struct {
	Avg             float64 `json:"avg"`
	BidPrice01      float64 `json:"bidPrice01"`
	BidPrice02      float64 `json:"bidPrice02"`
	BidPrice03      float64 `json:"bidPrice03"`
	BidQtty01       float64 `json:"bidQtty01"`
	BidQtty02       float64 `json:"bidQtty02"`
	BidQtty03       float64 `json:"bidQtty03"`
	BuyForeignQtty  float64 `json:"buyForeignQtty"`
	CeilPrice       float64 `json:"ceilPrice"`
	Change          float64 `json:"change"`
	ChangePercent   float64 `json:"changePercent"`
	ExpiryDate      string  `json:"expiryDate"`
	FloorPrice      float64 `json:"floorPrice"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	MatchPrice      float64 `json:"matchPrice"`
	MatchQtty       float64 `json:"matchQtty"`
	OfferPrice01    float64 `json:"offerPrice01"`
	OfferPrice02    float64 `json:"offerPrice02"`
	OfferPrice03    float64 `json:"offerPrice03"`
	OfferQtty01     float64 `json:"offerQtty01"`
	OfferQtty02     float64 `json:"offerQtty02"`
	OfferQtty03     float64 `json:"offerQtty03"`
	Open            float64 `json:"open"`
	OpenVol         float64 `json:"openVol"`
	RefPrice        float64 `json:"refPrice"`
	SellForeignQtty float64 `json:"sellForeignQtty"`
	Symbol          string  `json:"symbol"`
	TotalVol        float64 `json:"totalVol"`
}
