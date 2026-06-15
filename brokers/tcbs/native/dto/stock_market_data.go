// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type StockTicker struct {
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
	FloorPrice      float64 `json:"floorPrice"`
	High            float64 `json:"high"`
	IndexNumber     float64 `json:"indexNumber"`
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
	Room            string  `json:"room"`
	SellForeignQtty float64 `json:"sellForeignQtty"`
	Symbol          string  `json:"symbol"`
	TotalVal        float64 `json:"totalVal"`
	TotalVol        float64 `json:"totalVol"`
}

type StockForeignRoom struct {
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
	RefPrice        float64 `json:"refPrice"`
	Room            string  `json:"room"`
	SellForeignQtty float64 `json:"sellForeignQtty"`
	Symbol          string  `json:"symbol"`
	TotalValue      float64 `json:"totalValue"`
	TotalVolume     float64 `json:"totalVolume"`
}

type PutThroughAdvertisement struct {
	Color   float64 `json:"color"`
	OrderID string  `json:"orderId"`
	Price   float64 `json:"price"`
	Side    string  `json:"side"`
	Status  float64 `json:"status"`
	Symbol  string  `json:"symbol"`
	Time    string  `json:"time"`
	Vol     float64 `json:"vol"`
}

type PutThroughMatch struct {
	AccumulatedValue float64 `json:"accumulatedValue"`
	Price            float64 `json:"price"`
	Symbol           string  `json:"symbol"`
	Time             string  `json:"time"`
	Val              float64 `json:"val"`
	Vol              float64 `json:"vol"`
}

type StockTradeHistoryEntry struct {
	A   string  `json:"a"`
	Ba  float64 `json:"ba"`
	Cp  float64 `json:"cp"`
	Hl  bool    `json:"hl"`
	P   float64 `json:"p"`
	Pcp float64 `json:"pcp"`
	Rcp float64 `json:"rcp"`
	Sa  float64 `json:"sa"`
	T   string  `json:"t"`
	V   float64 `json:"v"`
}

type StockSupplyDemand15MinutesEntry struct {
	Bms float64 `json:"bms"`
	Bsr float64 `json:"bsr"`
	Bu  float64 `json:"bu"`
	Bup float64 `json:"bup"`
	S   float64 `json:"s"`
	Sd  float64 `json:"sd"`
	Sdp float64 `json:"sdp"`
	Sms string  `json:"sms"`
	T   string  `json:"t"`
}

type StockSupplyDemandEntry struct {
	Bsr float64 `json:"bsr"`
	Bup float64 `json:"bup"`
	Sdp float64 `json:"sdp"`
	T   string  `json:"t"`
}
