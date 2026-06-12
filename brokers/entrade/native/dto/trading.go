package dto

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type GetInvestorAccountRequest struct {
	InvestorID string
}

type GetInvestorAccountResponse struct {
	ID                int64  `json:"id"`
	InvestorID        int64  `json:"investorId,omitempty"`
	InvestorAccountID string `json:"investorAccountId,omitempty"`
	AccountNo         string `json:"accountNo,omitempty"`
}

type GetAccountBalanceRequest struct {
	InvestorID string
}

type GetAccountBalanceResponse struct {
	AvailableCash decimal.Decimal `json:"availableCash"`
	TotalCash     decimal.Decimal `json:"totalCash,omitempty"`
	Equity        decimal.Decimal `json:"equity,omitempty"`
	Margin        decimal.Decimal `json:"margin,omitempty"`
}

type GetDerivativeMarginPortfoliosRequest struct {
	InvestorID string
}

type DerivativeMarginPortfolio struct {
	ID           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	InvestorID   int64  `json:"investorId,omitempty"`
	InvestorType string `json:"investorType,omitempty"`
}

type GetDerivativeMarginPortfoliosResponse struct {
	Data  []DerivativeMarginPortfolio `json:"data,omitempty"`
	Total int                         `json:"total,omitempty"`
}

type GetPPSERequest struct {
	InvestorID            string
	BankMarginPortfolioID string
	Symbol                string
	Price                 decimal.Decimal
	Side                  string
}

type GetPPSEResponse struct {
	QMax        decimal.Decimal `json:"qmax"`
	BuyingPower decimal.Decimal `json:"buyingPower,omitempty"`
}

type PlaceDerivativeOrderRequest struct {
	BankMarginPortfolioID int             `json:"bankMarginPortfolioId"`
	InvestorID            int             `json:"investorId"`
	Symbol                string          `json:"symbol"`
	Price                 decimal.Decimal `json:"-"`
	OrderType             string          `json:"orderType"`
	Side                  string          `json:"side"`
	Quantity              int             `json:"quantity"`
}

func (r PlaceDerivativeOrderRequest) MarshalJSON() ([]byte, error) {
	type request PlaceDerivativeOrderRequest
	return json.Marshal(struct {
		request
		Price json.Number `json:"price"`
	}{request: request(r), Price: json.Number(r.Price.String())})
}

type DerivativeOrder struct {
	ID          int64           `json:"id"`
	Symbol      string          `json:"symbol,omitempty"`
	Price       decimal.Decimal `json:"price,omitempty"`
	OrderType   string          `json:"orderType,omitempty"`
	Side        string          `json:"side,omitempty"`
	Quantity    int             `json:"quantity,omitempty"`
	Filled      int             `json:"filled,omitempty"`
	Status      string          `json:"status,omitempty"`
	CreatedDate string          `json:"createdDate,omitempty"`
}

type PlaceDerivativeOrderResponse DerivativeOrder

type GetDerivativeOrdersRequest struct {
	InvestorAccountID string
	Start             int
	End               int
	Sort              string
	Order             string
}

type GetDerivativeOrdersResponse struct {
	Data  []DerivativeOrder `json:"data,omitempty"`
	Total int               `json:"total,omitempty"`
}

type GetDerivativeOrderRequest struct {
	OrderID string
}

type GetDerivativeOrderResponse DerivativeOrder

type CancelDerivativeOrderRequest struct {
	OrderID string
}

type CancelDerivativeOrderResponse DerivativeOrder

type DerivativeDeal struct {
	ID        int64           `json:"id"`
	Symbol    string          `json:"symbol,omitempty"`
	Price     decimal.Decimal `json:"price,omitempty"`
	Side      string          `json:"side,omitempty"`
	Quantity  int             `json:"quantity,omitempty"`
	Status    string          `json:"status,omitempty"`
	OpenPrice decimal.Decimal `json:"openPrice,omitempty"`
	Profit    decimal.Decimal `json:"profit,omitempty"`
}

type GetDerivativeDealsRequest struct {
	InvestorAccountID string
	Start             int
	End               int
	Sort              string
	Order             string
}

type GetDerivativeDealsResponse struct {
	Data  []DerivativeDeal `json:"data,omitempty"`
	Total int              `json:"total,omitempty"`
}

type CloseDerivativeDealRequest struct {
	DealID    string
	OrderType string
}

type CloseDerivativeDealResponse struct {
	Data []DerivativeDeal `json:"data,omitempty"`
}

type GetRiskConfigRequest struct {
	InvestorAccountID string
}

type RiskConfig struct {
	CutLossRate               decimal.Decimal `json:"cutLossRate"`
	InvestorAccountID         int             `json:"investorAccountId,omitempty"`
	TrailingEnabled           bool            `json:"trailingEnabled,omitempty"`
	InvestorID                int             `json:"investorId,omitempty"`
	AutoIncreaseDealRate      bool            `json:"autoIncreaseDealRate,omitempty"`
	EnableAutoDealDepositNoti bool            `json:"enableAutoDealDepositNoti,omitempty"`
}

type GetRiskConfigResponse struct {
	Data  []RiskConfig `json:"data,omitempty"`
	Total int          `json:"total,omitempty"`
}

type UpdateRiskConfigRequest struct {
	PathInvestorAccountID     string          `json:"-"`
	CutLossRate               decimal.Decimal `json:"-"`
	InvestorAccountID         int             `json:"investorAccountId"`
	TrailingEnabled           bool            `json:"trailingEnabled"`
	InvestorID                int             `json:"investorId"`
	AutoIncreaseDealRate      bool            `json:"autoIncreaseDealRate"`
	EnableAutoDealDepositNoti bool            `json:"enableAutoDealDepositNoti"`
}

func (r UpdateRiskConfigRequest) MarshalJSON() ([]byte, error) {
	type request UpdateRiskConfigRequest
	return json.Marshal(struct {
		request
		CutLossRate json.Number `json:"cutLossRate"`
	}{request: request(r), CutLossRate: json.Number(r.CutLossRate.String())})
}

type UpdateRiskConfigResponse RiskConfig
