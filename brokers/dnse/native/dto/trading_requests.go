package dto

import "github.com/shopspring/decimal"

type GetAccountsRequest struct{}
type GetAccountBalancesRequest struct{ AccountNo string }
type GetCorporateActionHistoryRequest struct {
	AccountNo, Symbol, CAType, CAStatus string
	PageIndex, PageSize                 int
}
type GetExecutionsRequest struct{ AccountNo, OrderID, MarketType string }
type GetLoanPackagesRequest struct{ AccountNo, MarketType, Symbol string }
type GetOrderHistoryRequest struct{ AccountNo, MarketType, From, To string }
type GetOrdersRequest struct{ AccountNo, MarketType, OrderCategory string }
type GetPositionPnLConfigsRequest struct{ PositionID, MarketType string }
type GetPositionRequest struct{ PositionID, MarketType string }
type GetPositionsRequest struct {
	AccountNo, MarketType string
	PageSize              int
}
type GetPPSERequest struct {
	AccountNo, MarketType, Symbol string
	LoanPackageID                 *int
	Price                         *decimal.Decimal
}
type CancelOrderRequest struct{ AccountNo, OrderID, MarketType, OrderCategory string }
type ClosePositionRequest struct{ PositionID, MarketType string }
type GetOrderRequest = CancelOrderRequest
type PlaceOrderRequest struct {
	AccountNo, OrderType, Side, Symbol, MarketType, OrderCategory string
	Price                                                         *float64
	Quantity                                                      int64
	LoanPackageID                                                 *int
}
type ReplaceOrderRequest struct {
	AccountNo, OrderID, MarketType, OrderCategory string
	Price                                         *float64
	Quantity                                      int64
}

type PnLRule struct {
	Enabled         bool     `json:"enabled"`
	Strategy        string   `json:"strategy,omitempty"`
	Rate            *float64 `json:"rate,omitempty"`
	DeltaPrice      *float64 `json:"deltaPrice,omitempty"`
	OrderMethod     string   `json:"orderMethod,omitempty"`
	OrderDeltaPrice *float64 `json:"orderDeltaPrice,omitempty"`
	TrailingEnabled bool     `json:"trailingEnabled,omitempty"`
}
type PnLConfigs struct {
	TakeProfit *PnLRule `json:"takeProfit,omitempty"`
	StopLoss   *PnLRule `json:"stopLoss,omitempty"`
}
type SetPositionPnLConfigsRequest struct {
	PositionID, MarketType string
	Configs                PnLConfigs
}
type PositionPnLConfigsResponse struct {
	AccountNo    string     `json:"accountNo,omitempty"`
	PositionID   int64      `json:"positionId,omitempty"`
	Configs      PnLConfigs `json:"configs,omitempty"`
	CreatedDate  string     `json:"createdDate,omitempty"`
	ModifiedDate string     `json:"modifiedDate,omitempty"`
}
