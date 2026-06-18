package dto

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type SubscribeStockTradeHistoryRequest struct{ Tickers []string }
type SubscribeStockSupplyDemandOneMinuteRequest struct{ Tickers []string }
type SubscribeStockSupplyDemandFifteenMinutesRequest struct{ Tickers []string }

type StockTradeHistoryEvent struct {
	Symbol              string          `json:"symbol"`
	ClosePrice          decimal.Decimal `json:"closePrice"`
	CloseVolume         decimal.Decimal `json:"closeVol"`
	Change              decimal.Decimal `json:"change"`
	Reference           decimal.Decimal `json:"reference"`
	TotalTrading        decimal.Decimal `json:"totalTrading"`
	TotalTradingValue   decimal.Decimal `json:"totalTradingValue"`
	TimeSeconds         string          `json:"timeSec"`
	Action              string          `json:"action"`
	UnitTimeFrame       string          `json:"unitTimeFrame"`
	TradingValue        decimal.Decimal `json:"tradingValue"`
	BuyUpAccumulated    decimal.Decimal `json:"buyUpAcc"`
	SellDownAccumulated decimal.Decimal `json:"sellDownAcc"`
	BidPrice            decimal.Decimal `json:"bidPrice"`
	BidVolume           decimal.Decimal `json:"bidVol"`
	AskPrice            decimal.Decimal `json:"askPrice"`
	AskVolume           decimal.Decimal `json:"askVol"`
	PreviousChange      decimal.Decimal `json:"prevChange"`
}

type RawEvent = json.RawMessage
type SubscribeRawRequest struct {
	Tickers []string `json:"tickers"`
}

type StockSupplyDemandEvent struct {
	Symbol                         string          `json:"symbol"`
	TimeSeconds                    string          `json:"timeSec"`
	TotalBuyUpVolume               decimal.Decimal `json:"totalBUVol"`
	TotalBuyUpVolumeAccumulated    decimal.Decimal `json:"totalBUVolAcc"`
	TotalSellDownVolume            decimal.Decimal `json:"totalSDVol"`
	TotalSellDownVolumeAccumulated decimal.Decimal `json:"totalSDVolAcc"`
	BuySellAccumulatedRatio        decimal.Decimal `json:"bsAccRatio"`
	UnitTimeFrame                  string          `json:"unitTimeFrame"`
}

type SubscribeBidPricesRequest struct{ Symbols []string }
type SubscribeOfferPricesRequest struct{ Symbols []string }
type SubscribeForeignTradingRequest struct{ Symbols []string }
type SubscribeBasePricesRequest struct{ Symbols []string }
type SubscribeMatchedPricesRequest struct{ Symbols []string }
type SubscribeTickerMatchesRequest struct{ Symbols []string }
type SubscribePutThroughAdvertisementsRequest struct{ Symbols []string }
type SubscribePutThroughMatchesRequest struct{ Symbols []string }
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
	Symbol       string          `json:"symbol"`
	OfferPrice01 decimal.Decimal `json:"offerPrice01"`
	OfferPrice02 decimal.Decimal `json:"offerPrice02"`
	OfferPrice03 decimal.Decimal `json:"offerPrice03"`
	OfferQtty01  decimal.Decimal `json:"offerQtty01"`
	OfferQtty02  decimal.Decimal `json:"offerQtty02"`
	OfferQtty03  decimal.Decimal `json:"offerQtty03"`
}

// fe: STOCK (s|3) - DE (s|25) - CW (s|30)  # (d|s|ro|fe|1,2 or d|s|tk|fe|AAA,ACB)
type ForeignTradingEvent struct {
	Symbol          string          `json:"symbol"`
	BuyForeignQtty  decimal.Decimal `json:"buyForeignQtty"`
	SellForeignQtty decimal.Decimal `json:"sellForeignQtty"`
	Room            decimal.Decimal `json:"room"`
}

// bp: STOCK (s|4) - DE (s|26) - CW (s|31)  # (d|s|ro|bp|1,2 or d|s|tk|bp|AAA,ACB)
type BasePriceEvent struct {
	Symbol     string          `json:"symbol"`
	CeilPrice  decimal.Decimal `json:"ceilPrice"`
	FloorPrice decimal.Decimal `json:"floorPrice"`
	RefPrice   decimal.Decimal `json:"refPrice"`
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
	IndexNumber   decimal.Decimal `json:"indexNumber"`
	Index         decimal.Decimal `json:"index"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
	Volume        decimal.Decimal `json:"volume"`
	Value         decimal.Decimal `json:"value"`
	Increase      decimal.Decimal `json:"increase"`
	Decrease      decimal.Decimal `json:"decrease"`
	NotChange     decimal.Decimal `json:"notChange"`
	Session       string          `json:"session"`
	CeilIncrease  decimal.Decimal `json:"ceilIncrease"`
	FloorDecrease decimal.Decimal `json:"floorDecrease"`
}

// d|s|pt|ptm+ptb+pts|1,2,3???
// pta (16): STOCK (s|?) - DE (s|?) - CW (s|16)  # (d|s|ro|pta|1,2 or d|s|tk|pta|AAA,ACB)
type PutThroughAdvertisementEvent struct {
	Symbol  string          `json:"symbol"`
	Price   decimal.Decimal `json:"price"`
	Volume  decimal.Decimal `json:"vol"`
	Time    string          `json:"time"`
	Status  decimal.Decimal `json:"status"`
	Color   decimal.Decimal `json:"color"`
	OrderID string          `json:"orderId"`
	Side    string          `json:"side"`
}

// ptm (17): STOCK (s|?) - DE (s|?) - CW (s|17)  # (d|s|pt|ptm+ptb+pts|1,2,3)
type PutThroughMatchedEvent struct {
	Symbol           string          `json:"symbol"`
	Price            decimal.Decimal `json:"price"`
	Volume           decimal.Decimal `json:"vol"`
	Value            decimal.Decimal `json:"val"`
	Time             string          `json:"time"`
	Color            decimal.Decimal `json:"color"`
	AccumulatedValue decimal.Decimal `json:"accumulatedValue"`
}
