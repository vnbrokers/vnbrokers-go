package ssi

import "github.com/shopspring/decimal"

type RightsHistoriesRequest struct {
	AccountID string
	StartDate string
	EndDate   string
}

type RightsCreateRequest struct {
	AccountID     string
	Symbol        string
	EntitlementID string
	Quantity      decimal.Decimal
	Amount        decimal.Decimal
	Code          string
}

type DividendsData struct {
	Account   string     `json:"account,omitempty"`
	Dividends []Dividend `json:"dividends,omitempty"`
}

type Dividend struct {
	StockDividend          string           `json:"stockDividend,omitempty"`
	InstrumentID           string           `json:"instrumentID,omitempty"`
	Quantity               *decimal.Decimal `json:"quantity,omitempty"`
	ExecutedRate           string           `json:"executedRate,omitempty"`
	CloseDate              string           `json:"closeDate,omitempty"`
	PaidDate               string           `json:"paidDate,omitempty"`
	Amount                 *decimal.Decimal `json:"amount,omitempty"`
	Status                 string           `json:"status,omitempty"`
	ReceivedQuantity       *decimal.Decimal `json:"receivedQuantity,omitempty"`
	IssueInstrument        string           `json:"issueInstrument,omitempty"`
	DistributedFlag        string           `json:"distributedFlag,omitempty"`
	PayableDate            string           `json:"payableDate,omitempty"`
	SubscriptionPrice      *decimal.Decimal `json:"subscriptionPrice,omitempty"`
	SubscriptionAmount     *decimal.Decimal `json:"subscriptionAmount,omitempty"`
	SubscriptionQuantity   *decimal.Decimal `json:"subscriptionQuantity,omitempty"`
	SubscriptionPeriodFrom string           `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo   string           `json:"subscriptionPeriodTo,omitempty"`
	EntitlementID          string           `json:"entitlementID,omitempty"`
	ExchangeID             string           `json:"exchangeID,omitempty"`
}

type ExercisableQuantitiesData struct {
	Account               string                `json:"account,omitempty"`
	ExercisableQuantities []ExercisableQuantity `json:"exercisableQuantities,omitempty"`
}

type ExercisableQuantity struct {
	EntitlementID               string           `json:"entitlementID,omitempty"`
	InstrumentID                string           `json:"instrumentID,omitempty"`
	SubscriptionPrice           *decimal.Decimal `json:"subscriptionPrice,omitempty"`
	ExecutedRateFrom            *decimal.Decimal `json:"executedRateFrom,omitempty"`
	SubscriptionPeriodFrom      string           `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo        string           `json:"subscriptionPeriodTo,omitempty"`
	ExerciseableQuantity        *decimal.Decimal `json:"exerciseableQuantity,omitempty"`
	ExerciseableReceiveQuantity *decimal.Decimal `json:"exerciseableReceiveQuantity,omitempty"`
	ExercisedReceiveQuantity    *decimal.Decimal `json:"exercisedReceiveQuantity,omitempty"`
	ExecutedRateTo              *decimal.Decimal `json:"executedRateTo,omitempty"`
	ExercisedQuantity           *decimal.Decimal `json:"exercisedQuantity,omitempty"`
	PayableDate                 string           `json:"payableDate,omitempty"`
}

type RightsHistoriesData struct {
	Account                          string          `json:"account,omitempty"`
	OnlineRightSubscriptionHistories []RightsHistory `json:"onlineRightSubscriptionHistories,omitempty"`
}

type RightsHistory struct {
	TransactionID             string           `json:"transactionID,omitempty"`
	DateTime                  string           `json:"dateTime,omitempty"`
	InstrumentID              string           `json:"instrumentID,omitempty"`
	RatioFrom                 *decimal.Decimal `json:"ratioFrom,omitempty"`
	SubscriptionPrice         *decimal.Decimal `json:"subscriptionPrice,omitempty"`
	SubscriptionPeriodFrom    string           `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo      string           `json:"subscriptionPeriodTo,omitempty"`
	ExercisedReceivedQuantity *decimal.Decimal `json:"exercisedReceivedQty,omitempty"`
	Amount                    *decimal.Decimal `json:"amount,omitempty"`
	Status                    string           `json:"status,omitempty"`
	RatioTo                   *decimal.Decimal `json:"ratioTo,omitempty"`
	UnderlyingInstrumentID    string           `json:"underlyingInstrumentID,omitempty"`
}
