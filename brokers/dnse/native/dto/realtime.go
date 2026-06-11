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
type MarketIndexEvent struct {
	T                               string           `json:"T,omitempty"`
	IndexName                       string           `json:"indexName,omitempty"`
	ChangedRatio                    *decimal.Decimal `json:"changedRatio,omitempty"`
	ChangedValue                    *decimal.Decimal `json:"changedValue,omitempty"`
	IndexValue                      *decimal.Decimal `json:"indexValue,omitempty"`
	TotalValue                      *decimal.Decimal `json:"totalValue,omitempty"`
	TotalVolume                     *decimal.Decimal `json:"totalVolume,omitempty"`
	FluctuationSteadinessIssueCount *int             `json:"fluctuationSteadinessIssueCount,omitempty"`
	FluctuationDownIssueCount       *int             `json:"fluctuationDownIssueCount,omitempty"`
	FluctuationUpIssueCount         *int             `json:"fluctuationUpIssueCount,omitempty"`
	FluctuationLowerLimitIssueCount *int             `json:"fluctuationLowerLimitIssueCount,omitempty"`
	FluctuationUpperLimitIssueCount *int             `json:"fluctuationUpperLimitIssueCount,omitempty"`
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
type PositionEvent = Position
