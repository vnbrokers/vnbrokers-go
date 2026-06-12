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

type ExpectedPriceEvent = ExpectedPrice
type ForeignEvent = Foreign
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
type EstimatedMarketIndexEvent = EstimatedMarketIndex
type OHLCEvent = OHLC
type OHLCClosedEvent = OHLCClosed
type QuoteEvent = Quote
type SecurityDefinitionEvent = SecurityDefinition
type TradeEvent = Trade
type TradeExtraEvent = TradeExtra
type OrderEvent = Order
type BrokerOrderEvent = Order
type PositionEvent = Position
