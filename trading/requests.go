package trading

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

type AssetClass string

const (
	AssetClassStock      AssetClass = "STOCK"
	AssetClassDerivative AssetClass = "DERIVATIVE"
)

type MarketType string

const (
	MarketTypeStock      MarketType = "STOCK"
	MarketTypeDerivative MarketType = "DERIVATIVE"
)

type OrderCategory string

const (
	OrderCategoryNormal      OrderCategory = "NORMAL"
	OrderCategoryConditional OrderCategory = "CONDITIONAL"
)

type OTPType string

const (
	OTPTypeEmail OTPType = "email_otp"
	OTPTypeSmart OTPType = "smart_otp"
)

type ListOrdersRequest struct {
	AccountID     string
	MarketType    MarketType
	OrderCategory OrderCategory
	PageIndex     int
	PageSize      int
}

type OrderHistoryRequest struct {
	AccountID     string
	FromDate      string
	ToDate        string
	MarketType    MarketType
	OrderCategory OrderCategory
	PageIndex     int
	PageSize      int
}

type GetOrderRequest struct {
	AccountID     string
	OrderID       string
	MarketType    MarketType
	OrderCategory OrderCategory
}

type CancelOrderRequest struct {
	AccountID     string
	OrderID       string
	MarketType    MarketType
	OrderCategory OrderCategory
}

type ReplaceOrderRequest struct {
	AccountID     string
	OrderID       string
	Price         decimal.Decimal
	Quantity      int
	MarketType    MarketType
	OrderCategory OrderCategory
}

type ExecutionsRequest struct {
	AccountID     string
	OrderID       string
	MarketType    MarketType
	OrderCategory OrderCategory
}

type BuyingPowerRequest struct {
	AccountID     string
	Symbol        string
	Price         decimal.Decimal
	Side          domain.OrderSide
	MarketType    MarketType
	LoanPackageID *int
}

type LoanPackagesRequest struct {
	AccountID  string
	Symbol     string
	MarketType MarketType
}

type ListPositionsRequest struct {
	AccountID  string
	MarketType MarketType
	PageSize   int
}

type GetPositionRequest struct {
	PositionID string
	MarketType MarketType
}

type ClosePositionRequest struct {
	PositionID string
	MarketType MarketType
}

type TradingTokenRequest struct {
	OTPType  OTPType
	Passcode string
}

type AccountsService interface {
	List(context.Context) ([]domain.Account, error)
	Balance(context.Context, string) (domain.Balance, error)
	Orders(context.Context, string) ([]domain.Order, error)
	OrderHistory(context.Context, string, string, string, int) ([]domain.Order, error)
	Executions(context.Context, string, string) (domain.RawPayload, error)
	PPSE(context.Context, string, string, decimal.Decimal, *int) (domain.RawPayload, error)
	LoanPackages(context.Context, string, string) (domain.RawPayload, error)
}

type OrdersService interface {
	Place(context.Context, domain.PlaceOrderRequest) (domain.PlaceOrderResponse, error)
	Cancel(context.Context, string, string) error
	Get(context.Context, string, string) (domain.Order, error)
	Update(context.Context, string, string, decimal.Decimal, int) (domain.RawPayload, error)
}

type PositionsService interface {
	List(context.Context, string) ([]domain.Position, error)
	Get(context.Context, string) (domain.Position, error)
	Close(context.Context, string) (domain.RawPayload, error)
}

type SubscribeOrdersRequest struct {
	MarketType string
}

type SubscribePositionsRequest struct {
	MarketType string
}
