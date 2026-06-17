package dto

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type SubscribeStockPricesRequest struct{}

type SubscribeStockTradeHistoryRequest struct{ Tickers []string }
type SubscribeStockSupplyDemandOneMinuteRequest struct{ Tickers []string }
type SubscribeStockSupplyDemandFifteenMinutesRequest struct{ Tickers []string }

type StockTradeHistoryEvent struct {
	Symbol              string  `json:"symbol"`
	ClosePrice          float64 `json:"closePrice"`
	CloseVolume         float64 `json:"closeVol"`
	Change              float64 `json:"change"`
	Reference           float64 `json:"reference"`
	TotalTrading        float64 `json:"totalTrading"`
	TotalTradingValue   float64 `json:"totalTradingValue"`
	TimeSeconds         string  `json:"timeSec"`
	Action              string  `json:"action"`
	UnitTimeFrame       string  `json:"unitTimeFrame"`
	TradingValue        float64 `json:"tradingValue"`
	BuyUpAccumulated    float64 `json:"buyUpAcc"`
	SellDownAccumulated float64 `json:"sellDownAcc"`
	BidPrice            float64 `json:"bidPrice"`
	BidVolume           float64 `json:"bidVol"`
	AskPrice            float64 `json:"askPrice"`
	AskVolume           float64 `json:"askVol"`
	PreviousChange      float64 `json:"prevChange"`
}

type RawEvent = json.RawMessage
type SubscribeRawRequest struct {
	Tickers []string `json:"tickers"`
}

type StockSupplyDemandEvent struct {
	Symbol                         string  `json:"symbol"`
	TimeSeconds                    string  `json:"timeSec"`
	TotalBuyUpVolume               float64 `json:"totalBUVol"`
	TotalBuyUpVolumeAccumulated    float64 `json:"totalBUVolAcc"`
	TotalSellDownVolume            float64 `json:"totalSDVol"`
	TotalSellDownVolumeAccumulated float64 `json:"totalSDVolAcc"`
	BuySellAccumulatedRatio        float64 `json:"bsAccRatio"`
	UnitTimeFrame                  string  `json:"unitTimeFrame"`
}

type SubscribeBidPricesRequest struct{ Symbols []string }
type SubscribeOfferPricesRequest struct{ Symbols []string }
type SubscribeForeignTradingRequest struct{ Symbols []string }
type SubscribeBasePricesRequest struct{ Symbols []string }
type SubscribeMatchedPricesRequest struct{ Symbols []string }
type SubscribeTickerMatchesRequest struct{ Symbols []string }
type SubscribeMarketIndexesRequest struct{ Indexes []string }

// bi: STOCK (s|1) - DE (s|23) - CW (s|28)  # (d|s|ro|bi|1,2 or d|s|tk|bi|AAA,ACB)
type BidPriceEvent struct {
	Symbol     string          `json:"symbol"`
	BidPrice01 decimal.Decimal `json:"bidPrice01"`
	BidPrice02 decimal.Decimal `json:"bidPrice02"`
	BidPrice03 decimal.Decimal `json:"bidPrice03"`
	BidQtty01  decimal.Decimal `json:"bidQtty01"`
	BidQtty02  decimal.Decimal `json:"bidQtty02"`
	BidQtty03  decimal.Decimal `json:"bidQtty03"`
}

// op: STOCK (s|2) - DE (s|24) - CW (s|29)  # (d|s|ro|op|1,2 or d|s|tk|op|AAA,ACB)
type OfferPriceEvent struct {
	Symbol       string `json:"symbol"`
	OfferPrice01 string `json:"offerPrice01"`
	OfferPrice02 string `json:"offerPrice02"`
	OfferPrice03 string `json:"offerPrice03"`
	OfferQtty01  string `json:"offerQtty01"`
	OfferQtty02  string `json:"offerQtty02"`
	OfferQtty03  string `json:"offerQtty03"`
}

// fe: STOCK (s|3) - DE (s|25) - CW (s|30)  # (d|s|ro|fe|1,2 or d|s|tk|fe|AAA,ACB)
type ForeignTradingEvent struct {
	Symbol          string `json:"symbol"`
	BuyForeignQtty  string `json:"buyForeignQtty"`
	SellForeignQtty string `json:"sellForeignQtty"`
	Room            string `json:"room"`
}

// bp: STOCK (s|4) - DE (s|26) - CW (s|31)  # (d|s|ro|bp|1,2 or d|s|tk|bp|AAA,ACB)
type BasePriceEvent struct {
	Symbol     string  `json:"symbol"`
	CeilPrice  float64 `json:"ceilPrice"`
	FloorPrice float64 `json:"floorPrice"`
	RefPrice   float64 `json:"refPrice"`
}

// mp: STOCK (s|5) - DE (s|32) - CW (s|27)  # (d|s|ro|mp|1,2 or d|s|tk|mp|AAA,ACB)
type MatchedPriceEvent struct {
	Symbol             string          `json:"symbol"`
	High               decimal.Decimal `json:"high"`
	Low                decimal.Decimal `json:"low"`
	Average            decimal.Decimal `json:"avg"`
	Open               decimal.Decimal `json:"open"`
	HighestFromListing decimal.Decimal `json:"highestFromListing,omitempty"` // CW (s|32)
	LowestFromListing  decimal.Decimal `json:"lowestFromListing,omitempty"`  // CW (s|32)
}

// tm: STOCK (s|6)  - DE (s|21) - CW (s|18) # (d|s|ro|mp|1,2 or d|s|tk|mp|AAA,ACB)
type TickerMatchEvent struct {
	Symbol        string          `json:"symbol"`
	MatchPrice    decimal.Decimal `json:"matchPrice"`
	MatchQtty     decimal.Decimal `json:"matchQtty"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
	TotalVolume   decimal.Decimal `json:"totalVolume"`
	TotalValue    decimal.Decimal `json:"totalValue"`
}

// rt: INDEX (s|8) # (d|s|si|rt|1,2)
type MarketIndexEvent struct {
	IndexNumber   int             `json:"indexNumber"`
	Index         decimal.Decimal `json:"index"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
	Volume        decimal.Decimal `json:"volume"`
	Value         decimal.Decimal `json:"value"`
	Increase      int             `json:"increase"`
	Decrease      int             `json:"decrease"`
	NotChange     int             `json:"notChange"`
	Session       string          `json:"session"`
	CeilIncrease  int             `json:"ceilIncrease"`
	FloorDecrease int             `json:"floorDecrease"`
}

// d|s|pt|ptm+ptb+pts|1,2,3???
// pta (16): STOCK (s|?) - DE (s|?) - CW (s|16)  # (d|s|ro|pta|1,2 or d|s|tk|pta|AAA,ACB)
type PutThroughAdvertisementEvent struct {
	Symbol  string  `json:"symbol"`
	Price   float64 `json:"price"`
	Volume  float64 `json:"vol"`
	Time    string  `json:"time"`
	Status  int     `json:"status"`
	Color   int     `json:"color"`
	OrderID string  `json:"orderId"`
	Side    string  `json:"side"`
}

// ptm (17): STOCK (s|?) - DE (s|?) - CW (s|17)  # (d|s|pt|ptm+ptb+pts|1,2,3)
type PutThroughMatchedEvent struct {
	Symbol           string  `json:"symbol"`
	Price            float64 `json:"price"`
	Volume           float64 `json:"vol"`
	Value            float64 `json:"val"`
	Time             string  `json:"time"`
	Color            int     `json:"color"`
	AccumulatedValue float64 `json:"accumulatedValue"`
}
