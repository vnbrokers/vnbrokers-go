package dto

import "github.com/shopspring/decimal"

type SubscribeSymbolsRequest struct {
	Symbols []string
	BoardID string
}
type SubscribeMarketIndexRequest struct{ IndexName string }
type SubscribeOHLCRequest struct {
	Symbols    []string
	Resolution string
}
type SubscribeTradingRequest struct{ MarketType string }
type SubscribeBrokerOrdersRequest struct {
	MarketType string
	InvestorID string
}

type ExpectedPriceEvent struct {
	MarketID            string           `json:"marketId,omitempty"`
	BoardID             string           `json:"boardId,omitempty"`
	Symbol              string           `json:"symbol,omitempty"`
	TradingSessionID    string           `json:"tradingSessionId,omitempty"`
	TransactTime        string           `json:"transactTime,omitempty"`
	ExpectedPrice       *decimal.Decimal `json:"expectedPrice,omitempty"`
	ExpectedChange      *decimal.Decimal `json:"expectedChange,omitempty"`
	ExpectedChangeRatio *decimal.Decimal `json:"expectedChangeRatio,omitempty"`
}
type ForeignEvent struct {
	MarketID                string           `json:"marketId,omitempty"`
	BoardID                 string           `json:"boardId,omitempty"`
	TradingSessionID        string           `json:"tradingSessionId,omitempty"`
	Symbol                  string           `json:"symbol,omitempty"`
	TransactTime            string           `json:"transactTime,omitempty"`
	ForeignInvestorTypeCode string           `json:"foreignInvestorTypeCode,omitempty"`
	SellVolume              int64            `json:"sellVolume,omitempty"`
	BuyVolume               int64            `json:"buyVolume,omitempty"`
	TotalSellVolume         int64            `json:"totalSellVolume,omitempty"`
	TotalBuyVolume          int64            `json:"totalBuyVolume,omitempty"`
	SellTradedAmount        *decimal.Decimal `json:"sellTradedAmount,omitempty"`
	BuyTradedAmount         *decimal.Decimal `json:"buyTradedAmount,omitempty"`
	TotalSellTradedAmount   *decimal.Decimal `json:"totalSellTradedAmount,omitempty"`
	TotalBuyTradedAmount    *decimal.Decimal `json:"totalBuyTradedAmount,omitempty"`
	CurrentRoom             *decimal.Decimal `json:"currentRoom,omitempty"`
}
type StreamTimestamp struct {
	Nanos   int64 `json:"Nanos,omitempty"`
	Seconds int64 `json:"Seconds,omitempty"`
}
type MarketIndexEvent struct {
	T                               string           `json:"T,omitempty"`
	BlockTradeAccumulatedValue      *decimal.Decimal `json:"blkTrdAccTrdVal,omitempty"`
	BlockTradeAccumulatedVolume     int64            `json:"blkTrdAccTrdVol,omitempty"`
	IndexName                       string           `json:"indexName,omitempty"`
	ChangedRatio                    *decimal.Decimal `json:"changedRatio,omitempty"`
	ChangedValue                    *decimal.Decimal `json:"changedValue,omitempty"`
	ContauctAccumulatedValue        *decimal.Decimal `json:"contauctAccTrdVal,omitempty"`
	ContauctAccumulatedVolume       int64            `json:"contauctAccTrdVol,omitempty"`
	CurrencyCode                    string           `json:"currencyCode,omitempty"`
	IndexValue                      *decimal.Decimal `json:"indexValue,omitempty"`
	TotalValue                      *decimal.Decimal `json:"totalValue,omitempty"`
	TotalVolume                     *decimal.Decimal `json:"totalVolume,omitempty"`
	FluctuationSteadinessIssueCount *int             `json:"fluctuationSteadinessIssueCount,omitempty"`
	FluctuationDownIssueCount       *int             `json:"fluctuationDownIssueCount,omitempty"`
	FluctuationDownIssueVolume      int64            `json:"fluctuationDownIssueVolume,omitempty"`
	FluctuationUpIssueCount         *int             `json:"fluctuationUpIssueCount,omitempty"`
	FluctuationUpIssueVolume        int64            `json:"fluctuationUpIssueVolume,omitempty"`
	FluctuationLowerLimitIssueCount *int             `json:"fluctuationLowerLimitIssueCount,omitempty"`
	FluctuationUpperLimitIssueCount *int             `json:"fluctuationUpperLimitIssueCount,omitempty"`
	GrossTradeAmount                *decimal.Decimal `json:"grossTradeAmount,omitempty"`
	HighestValueIndexes             *decimal.Decimal `json:"highestValueIndexes,omitempty"`
	IndexTypeCode                   string           `json:"indexTypeCode,omitempty"`
	LowestValueIndexes              *decimal.Decimal `json:"lowestValueIndexes,omitempty"`
	MarketID                        string           `json:"marketId,omitempty"`
	MarketIndexClass                string           `json:"marketIndexClass,omitempty"`
	MulticastReceiveTime            StreamTimestamp  `json:"multicastReceiveTime,omitempty"`
	PriorValueIndexes               *decimal.Decimal `json:"priorValueIndexes,omitempty"`
	TotalVolumeTraded               int64            `json:"totalVolumeTraded,omitempty"`
	TradingSessionID                string           `json:"tradingSessionId,omitempty"`
	TransactTime                    StreamTimestamp  `json:"transactTime,omitempty"`
	ValueIndexes                    *decimal.Decimal `json:"valueIndexes,omitempty"`
}
type EstimatedMarketIndexEvent struct {
	IndexName                       string           `json:"indexName,omitempty"`
	ValueIndexes                    *decimal.Decimal `json:"valueIndexes,omitempty"`
	ChangedValue                    *decimal.Decimal `json:"changedValue,omitempty"`
	ChangedRatio                    *decimal.Decimal `json:"changedRatio,omitempty"`
	FluctuationUpIssueCount         int              `json:"fluctuationUpIssueCount,omitempty"`
	FluctuationDownIssueCount       int              `json:"fluctuationDownIssueCount,omitempty"`
	FluctuationSteadinessIssueCount int              `json:"fluctuationSteadinessIssueCount,omitempty"`
	GrossTradeAmount                *decimal.Decimal `json:"grossTradeAmount,omitempty"`
	TotalVolumeTraded               int64            `json:"totalVolumeTraded,omitempty"`
	Time                            string           `json:"time,omitempty"`
}
type OHLCEvent struct {
	Symbol     string           `json:"symbol,omitempty"`
	Resolution string           `json:"resolution,omitempty"`
	Time       string           `json:"time,omitempty"`
	Open       *decimal.Decimal `json:"open,omitempty"`
	High       *decimal.Decimal `json:"high,omitempty"`
	Low        *decimal.Decimal `json:"low,omitempty"`
	Close      *decimal.Decimal `json:"close,omitempty"`
	Volume     *decimal.Decimal `json:"volume,omitempty"`
}
type QuoteEvent = Quote
type SecurityDefinitionEvent = SecurityDefinition
type TradeEvent = Trade
type TradeExtraEvent struct {
	MarketID     string           `json:"marketId,omitempty"`
	BoardID      string           `json:"boardId,omitempty"`
	ISIN         string           `json:"isin,omitempty"`
	Symbol       string           `json:"symbol,omitempty"`
	Side         string           `json:"side,omitempty"`
	TransactTime string           `json:"transactTime,omitempty"`
	Price        *decimal.Decimal `json:"price,omitempty"`
	Quantity     *decimal.Decimal `json:"quantity,omitempty"`
	AveragePrice *decimal.Decimal `json:"avgPrice,omitempty"`
	TotalVolume  *decimal.Decimal `json:"totalVolume,omitempty"`
}
type OrderEvent = Order
type BrokerOrderEvent = Order
type PositionEvent = Position
