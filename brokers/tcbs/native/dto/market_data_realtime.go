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

type SubscribeDerivativeBidPricesRequest struct{ Symbols []string }
type SubscribeDerivativeOfferPricesRequest struct{ Symbols []string }
type SubscribeDerivativeForeignTradingRequest struct{ Symbols []string }
type SubscribeDerivativeBasePricesRequest struct{ Symbols []string }
type SubscribeDerivativeMatchedPricesRequest struct{ Symbols []string }
type SubscribeDerivativeTickerMatchesRequest struct{ Symbols []string }
type SubscribeMarketIndexesRequest struct{ Indexes []string }

type BidPriceEvent struct {
	Symbol     string          `json:"symbol"`
	BidPrice01 decimal.Decimal `json:"bidPrice01"`
	BidPrice02 decimal.Decimal `json:"bidPrice02"`
	BidPrice03 decimal.Decimal `json:"bidPrice03"`
	BidQtty01  decimal.Decimal `json:"bidQtty01"`
	BidQtty02  decimal.Decimal `json:"bidQtty02"`
	BidQtty03  decimal.Decimal `json:"bidQtty03"`
}

type OfferPriceEvent struct {
	Symbol       string `json:"symbol"`
	OfferPrice01 string `json:"offerPrice01"`
	OfferPrice02 string `json:"offerPrice02"`
	OfferPrice03 string `json:"offerPrice03"`
	OfferQtty01  string `json:"offerQtty01"`
	OfferQtty02  string `json:"offerQtty02"`
	OfferQtty03  string `json:"offerQtty03"`
}

type DerivativeForeignTradingEvent struct {
	Symbol          string `json:"symbol"`
	BuyForeignQtty  string `json:"buyForeignQtty"`
	SellForeignQtty string `json:"sellForeignQtty"`
	Room            string `json:"room"`
}

type DerivativeBasePriceEvent struct {
	Symbol     string  `json:"symbol"`
	CeilPrice  float64 `json:"ceilPrice"`
	FloorPrice float64 `json:"floorPrice"`
	RefPrice   float64 `json:"refPrice"`
}

type DerivativeMatchedPriceEvent struct {
	Symbol  string  `json:"symbol"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Average float64 `json:"avg"`
	Open    string  `json:"open"`
}

type DerivativeTickerMatchEvent struct {
	Symbol        string  `json:"symbol"`
	MatchPrice    string  `json:"matchPrice"`
	MatchQtty     string  `json:"matchQtty"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	TotalVolume   string  `json:"totalVolume"`
	TotalValue    string  `json:"totalValue"`
}

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
