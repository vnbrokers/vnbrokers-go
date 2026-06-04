package entrade

import "github.com/shopspring/decimal"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type BuyingPowerRequest struct {
	InvestorID            string
	BankMarginPortfolioID string
	Symbol                string
	Price                 decimal.Decimal
	Side                  string
}

type PlaceDerivativeOrderRequest struct {
	BankMarginPortfolioID int             `json:"bankMarginPortfolioId"`
	InvestorID            int             `json:"investorId"`
	Symbol                string          `json:"symbol"`
	Price                 decimal.Decimal `json:"price"`
	OrderType             string          `json:"orderType"`
	Side                  string          `json:"side"`
	Quantity              int             `json:"quantity"`
}

type ListOrdersRequest struct {
	InvestorAccountID string
	Start             int
	End               int
	Sort              string
	Order             string
}

type ListDealsRequest struct {
	InvestorAccountID string
	Start             int
	End               int
	Sort              string
	Order             string
}

type RiskConfigRequest struct {
	CutLossRate               decimal.Decimal `json:"cutLossRate"`
	InvestorAccountID         int             `json:"investorAccountId"`
	TrailingEnabled           bool            `json:"trailingEnabled"`
	InvestorID                int             `json:"investorId"`
	AutoIncreaseDealRate      bool            `json:"autoIncreaseDealRate"`
	EnableAutoDealDepositNoti bool            `json:"enableAutoDealDepositNoti"`
}
