package dto

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type SubscribeSymbolsRequest struct {
	Symbols []string `json:"symbols,omitempty"`
	BoardID string   `json:"boardId,omitempty"`
}
type SubscribeMarketIndexRequest struct {
	IndexName string `json:"indexName,omitempty"`
}
type SubscribeOHLCRequest struct {
	Symbols    []string `json:"symbols,omitempty"`
	Resolution string   `json:"resolution,omitempty"`
}
type SubscribeTradingSessionRequest struct {
	TSCProdGrpID string `json:"tscProdGrpId,omitempty"`
	BoardID      string `json:"boardId,omitempty"`
}
type SubscribeTradingRequest struct {
	MarketType string `json:"marketType,omitempty"`
}
type SubscribeBrokerOrdersRequest struct {
	MarketType string `json:"marketType,omitempty"`
	InvestorID string `json:"investorId,omitempty"`
}
type SubscribeBrokerPositionsRequest struct {
	MarketType string `json:"marketType,omitempty"`
	InvestorID string `json:"investorId,omitempty"`
}

type ExpectedPriceEvent = ExpectedPrice
type ForeignEvent = Foreign
type PbTimestamp struct {
	Nanos   int64 `json:"Nanos,omitempty"`
	Seconds int64 `json:"Seconds,omitempty"`
}

func (t PbTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Unix(t.Seconds, t.Nanos).UTC().Format(time.RFC3339Nano))
}

func (t *PbTimestamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		*t = PbTimestamp{}
		return nil
	}

	type timestampAlias PbTimestamp
	var object timestampAlias
	if err := json.Unmarshal(data, &object); err == nil {
		*t = PbTimestamp(object)
		return nil
	}

	var seconds int64
	if err := json.Unmarshal(data, &seconds); err == nil {
		t.Seconds = seconds
		t.Nanos = 0
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	t.Seconds = parsed
	t.Nanos = 0
	return nil
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
	MulticastReceiveTime            PbTimestamp      `json:"multicastReceiveTime,omitempty"`
	PriorValueIndexes               *decimal.Decimal `json:"priorValueIndexes,omitempty"`
	TotalVolumeTraded               int64            `json:"totalVolumeTraded,omitempty"`
	TradingSessionID                string           `json:"tradingSessionId,omitempty"`
	TransactTime                    PbTimestamp      `json:"transactTime,omitempty"`
	ValueIndexes                    *decimal.Decimal `json:"valueIndexes,omitempty"`
}
type EstimatedMarketIndexEvent = EstimatedMarketIndex
type OHLCEvent = OHLC
type OHLCClosedEvent = OHLCClosed
type QuoteEvent = Quote
type SecurityDefinitionEvent = SecurityDefinition
type TradingSessionEvent = TradingSession
type TradeEvent = Trade
type TradeExtraEvent = TradeExtra
type OrderEvent = WrapperOrderEvent
type BrokerOrderEvent = WrapperOrderEvent
type PositionEvent = WrapperPositionEvent
type BrokerPositionEvent = WrapperPositionEvent
