package domain

import "github.com/shopspring/decimal"

type RawPayload struct {
	Source string
	Data   any
	Bytes  []byte
}

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type OrderType string

const (
	OrderTypeLimit   OrderType = "LIMIT"
	OrderTypeMarket  OrderType = "MARKET"
	OrderTypeStop    OrderType = "STOP"
	OrderTypeUnknown OrderType = "UNKNOWN"
)

type OrderStatus string

const (
	OrderStatusUnknown         OrderStatus = "UNKNOWN"
	OrderStatusPending         OrderStatus = "PENDING"
	OrderStatusAccepted        OrderStatus = "ACCEPTED"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusPendingCancel   OrderStatus = "PENDING_CANCEL"
	OrderStatusCancelled       OrderStatus = "CANCELLED"
	OrderStatusRejected        OrderStatus = "REJECTED"
)

type Account struct {
	Broker      string
	AccountID   string
	DisplayName string
	Raw         RawPayload
}

type Balance struct {
	AccountID     string
	CashAvailable *decimal.Decimal
	CashTotal     *decimal.Decimal
	BuyingPower   *decimal.Decimal
	Currency      string
	Raw           RawPayload
}

type Order struct {
	Broker    string
	AccountID string
	OrderID   string
	Symbol    string
	Side      OrderSide
	Type      OrderType
	Quantity  decimal.Decimal
	Status    OrderStatus
	Price     *decimal.Decimal
	Raw       RawPayload
}

type PlaceOrderRequest struct {
	AccountID string
	Symbol    string
	Side      OrderSide
	Type      OrderType
	Quantity  decimal.Decimal
	Price     *decimal.Decimal
}

type PlaceOrderResponse struct {
	OrderID string
	Status  OrderStatus
	Raw     RawPayload
}

type OrderEvent struct {
	Broker         string
	AccountID      string
	OrderID        string
	Symbol         string
	Status         OrderStatus
	RawStatus      string
	FilledQuantity string
	ReceivedAt     string
	Raw            RawPayload
}

type Position struct {
	AccountID         string
	Symbol            string
	Quantity          decimal.Decimal
	AvailableQuantity decimal.Decimal
	AveragePrice      *decimal.Decimal
	MarketValue       *decimal.Decimal
	Raw               RawPayload
}

type Symbol struct {
	Symbol      string
	Exchange    string
	DisplayName string
	Raw         RawPayload
}

type Quote struct {
	Symbol     string
	LastPrice  *decimal.Decimal
	ReceivedAt string
	Raw        RawPayload
}

type Candle struct {
	Symbol   string
	Interval string
	OpenedAt string
	Open     decimal.Decimal
	High     decimal.Decimal
	Low      decimal.Decimal
	Close    decimal.Decimal
	Volume   decimal.Decimal
	Raw      RawPayload
}

type Tick struct {
	Symbol     string
	Price      decimal.Decimal
	Quantity   *decimal.Decimal
	ReceivedAt string
	Raw        RawPayload
}

type TopPrice struct {
	Symbol      string
	BidPrice    *decimal.Decimal
	BidQuantity *decimal.Decimal
	AskPrice    *decimal.Decimal
	AskQuantity *decimal.Decimal
	ReceivedAt  string
	Raw         RawPayload
}
