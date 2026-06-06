package ssi

import (
	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

type ConditionalOrderNewRequest struct {
	AccountID      string
	Symbol         string
	Side           string
	Type           string
	Price          string
	PriceSlip      decimal.Decimal
	Quantity       decimal.Decimal
	FromDate       string
	ToDate         string
	StopPrice      decimal.Decimal
	ActivePrice    decimal.Decimal
	TrailingAmount decimal.Decimal
	TPActivePrice  decimal.Decimal
	SLActivePrice  decimal.Decimal
	TPPrice        string
	SLPrice        string
	TPSlip         decimal.Decimal
	SLSlip         decimal.Decimal
	Operator       string
	Code           string
	UserAgent      string
	DeviceID       string
}

type ConditionalOrderCancelRequest struct {
	FCOID     string
	Code      string
	UserAgent string
	DeviceID  string
}

type ConditionalOrderBookRequest struct {
	FCOID     string
	PageIndex int
	PageSize  int
}

type ConditionalOrderStatusHistoryRequest struct {
	FCOID     string
	PageIndex int
	PageSize  int
}

type ConditionalOrderListRequest struct {
	FCOID         string
	AccountID     string
	Type          string
	ProcessStatus string
	Symbol        string
	Side          string
	FromDate      string
	ToDate        string
	PageIndex     int
	PageSize      int
}

type ConditionalOrderResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	FCOID   string `json:"fcoid,omitempty"`
	DCOID   string `json:"dcoid,omitempty"`
}

type ConditionalOrderPage[T any] struct {
	Message       string `json:"message,omitempty"`
	Status        int    `json:"status,omitempty"`
	Data          []T    `json:"data,omitempty"`
	FCOList       []T    `json:"fcoList,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	PageNumber    int    `json:"pageNumber,omitempty"`
	ItemsCount    int    `json:"itemsCount,omitempty"`
	PageCount     any    `json:"pageCount,omitempty"`
	TotalElements int    `json:"totalElements,omitempty"`
	TotalPages    int    `json:"totalPages,omitempty"`
	Page          int    `json:"page,omitempty"`
	Size          int    `json:"size,omitempty"`
}

type ConditionalTriggeredOrder struct {
	FCOID           string           `json:"fcoId,omitempty"`
	Account         string           `json:"account,omitempty"`
	Quantity        *decimal.Decimal `json:"quantity,omitempty"`
	Price           string           `json:"price,omitempty"`
	InstrumentID    string           `json:"instrumentId,omitempty"`
	Side            string           `json:"side,omitempty"`
	OrderType       string           `json:"orderType,omitempty"`
	IsMainOrder     bool             `json:"isMainOrder,omitempty"`
	IsAttachedOrder bool             `json:"isAttachedOrder,omitempty"`
	CreatedTime     string           `json:"createdTime,omitempty"`
	UpdatedTime     string           `json:"updatedTime,omitempty"`
	UniqueID        string           `json:"uniqueId,omitempty"`
	OrderID         string           `json:"orderId,omitempty"`
	MatchedQuantity *decimal.Decimal `json:"matchedQuantity,omitempty"`
	OSQuantity      *decimal.Decimal `json:"osQuantity,omitempty"`
	AveragePrice    *decimal.Decimal `json:"avgPrice,omitempty"`
	Status          string           `json:"status,omitempty"`
	Detail          string           `json:"detail,omitempty"`
}

type ConditionalOrderStatus struct {
	State  string `json:"state,omitempty"`
	Time   string `json:"time,omitempty"`
	Code   int    `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type FCOEvent struct {
	FCOID           string
	NotifyID        int64
	Data            any
	ProcessStatus   string
	LastAction      string
	UniqueID        string
	MatchedQuantity decimal.Decimal
	IsPlaceOrder    bool
	IPAddress       string
	Symbol          string
	Prefix          string
	Quantity        decimal.Decimal
	BrokerID        string
	Price           string
	AccountID       string
	BrokerIDUpdate  string
	UpdatedTime     string
	Status          string
	Message         string
	Username        string
	Raw             domain.RawPayload
}
