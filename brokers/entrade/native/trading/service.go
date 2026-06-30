package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityInvestorAccount            core.Capability = "native.trading.investor_account"
	CapabilityAccountBalance             core.Capability = "native.trading.account_balance"
	CapabilityDerivativeMarginPortfolios core.Capability = "native.trading.derivative_margin_portfolios"
	CapabilityPPSE                       core.Capability = "native.trading.ppse"
	CapabilityPlaceDerivativeOrder       core.Capability = "native.trading.place_derivative_order"
	CapabilityDerivativeOrders           core.Capability = "native.trading.derivative_orders"
	CapabilityDerivativeOrder            core.Capability = "native.trading.derivative_order"
	CapabilityCancelDerivativeOrder      core.Capability = "native.trading.cancel_derivative_order"
	CapabilityDerivativeDeals            core.Capability = "native.trading.derivative_deals"
	CapabilityCloseDerivativeDeal        core.Capability = "native.trading.close_derivative_deal"
	CapabilityRiskConfig                 core.Capability = "native.trading.risk_config"
	CapabilityUpdateRiskConfig           core.Capability = "native.trading.update_risk_config"
)

type Service interface {
	GetInvestorAccount(context.Context, dto.GetInvestorAccountRequest) (*dto.GetInvestorAccountResponse, error)
	GetAccountBalance(context.Context, dto.GetAccountBalanceRequest) (*dto.GetAccountBalanceResponse, error)
	GetDerivativeMarginPortfolios(context.Context, dto.GetDerivativeMarginPortfoliosRequest) (*dto.GetDerivativeMarginPortfoliosResponse, error)
	GetPPSE(context.Context, dto.GetPPSERequest) (*dto.GetPPSEResponse, error)
	PlaceDerivativeOrder(context.Context, dto.PlaceDerivativeOrderRequest) (*dto.PlaceDerivativeOrderResponse, error)
	GetDerivativeOrders(context.Context, dto.GetDerivativeOrdersRequest) (*dto.GetDerivativeOrdersResponse, error)
	GetDerivativeOrder(context.Context, dto.GetDerivativeOrderRequest) (*dto.GetDerivativeOrderResponse, error)
	CancelDerivativeOrder(context.Context, dto.CancelDerivativeOrderRequest) (*dto.CancelDerivativeOrderResponse, error)
	GetDerivativeDeals(context.Context, dto.GetDerivativeDealsRequest) (*dto.GetDerivativeDealsResponse, error)
	CloseDerivativeDeal(context.Context, dto.CloseDerivativeDealRequest) (*dto.CloseDerivativeDealResponse, error)
	GetRiskConfig(context.Context, dto.GetRiskConfigRequest) (*dto.GetRiskConfigResponse, error)
	UpdateRiskConfig(context.Context, dto.UpdateRiskConfigRequest) (*dto.UpdateRiskConfigResponse, error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func(bool) map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct{ dependencies Dependencies }

func NewService(dependencies Dependencies) Service { return &service{dependencies: dependencies} }

func (s *service) GetInvestorAccount(ctx context.Context, r dto.GetInvestorAccountRequest) (*dto.GetInvestorAccountResponse, error) {
	return send[dto.GetInvestorAccountResponse](ctx, s, CapabilityInvestorAccount, "get_investor_account", "GET", "/investors/"+url.PathEscape(r.InvestorID)+"/investor_account", nil)
}

func (s *service) GetAccountBalance(ctx context.Context, r dto.GetAccountBalanceRequest) (*dto.GetAccountBalanceResponse, error) {
	return send[dto.GetAccountBalanceResponse](ctx, s, CapabilityAccountBalance, "get_account_balance", "GET", "/account_balances/"+url.PathEscape(r.InvestorID), nil)
}

func (s *service) GetDerivativeMarginPortfolios(ctx context.Context, r dto.GetDerivativeMarginPortfoliosRequest) (*dto.GetDerivativeMarginPortfoliosResponse, error) {
	return send[dto.GetDerivativeMarginPortfoliosResponse](ctx, s, CapabilityDerivativeMarginPortfolios, "get_derivative_margin_portfolios", "GET", "/investors/"+url.PathEscape(r.InvestorID)+"/derivative_margin_portfolios", nil)
}

func (s *service) GetPPSE(ctx context.Context, r dto.GetPPSERequest) (*dto.GetPPSEResponse, error) {
	query := url.Values{}
	query.Set("bankMarginPortfolioId", r.BankMarginPortfolioID)
	query.Set("price", r.Price.String())
	query.Set("side", r.Side)
	query.Set("symbol", r.Symbol)
	path := "/derivative/investors/" + url.PathEscape(r.InvestorID) + "/ppse?" + query.Encode()
	return send[dto.GetPPSEResponse](ctx, s, CapabilityPPSE, "get_ppse", "GET", path, nil)
}

func (s *service) PlaceDerivativeOrder(ctx context.Context, r dto.PlaceDerivativeOrderRequest) (*dto.PlaceDerivativeOrderResponse, error) {
	return send[dto.PlaceDerivativeOrderResponse](ctx, s, CapabilityPlaceDerivativeOrder, "place_derivative_order", "POST", "/derivative/orders", r)
}

func (s *service) GetDerivativeOrders(ctx context.Context, r dto.GetDerivativeOrdersRequest) (*dto.GetDerivativeOrdersResponse, error) {
	return send[dto.GetDerivativeOrdersResponse](ctx, s, CapabilityDerivativeOrders, "get_derivative_orders", "GET", "/derivative/orders?"+rangeQuery(r.InvestorAccountID, r.Start, r.End, r.Sort, r.Order).Encode(), nil)
}

func (s *service) GetDerivativeOrder(ctx context.Context, r dto.GetDerivativeOrderRequest) (*dto.GetDerivativeOrderResponse, error) {
	return send[dto.GetDerivativeOrderResponse](ctx, s, CapabilityDerivativeOrder, "get_derivative_order", "GET", "/derivative/orders/"+url.PathEscape(r.OrderID), nil)
}

func (s *service) CancelDerivativeOrder(ctx context.Context, r dto.CancelDerivativeOrderRequest) (*dto.CancelDerivativeOrderResponse, error) {
	return send[dto.CancelDerivativeOrderResponse](ctx, s, CapabilityCancelDerivativeOrder, "cancel_derivative_order", "DELETE", "/derivative/orders/"+url.PathEscape(r.OrderID), nil)
}

func (s *service) GetDerivativeDeals(ctx context.Context, r dto.GetDerivativeDealsRequest) (*dto.GetDerivativeDealsResponse, error) {
	return send[dto.GetDerivativeDealsResponse](ctx, s, CapabilityDerivativeDeals, "get_derivative_deals", "GET", "/derivative/deals?"+rangeQuery(r.InvestorAccountID, r.Start, r.End, r.Sort, r.Order).Encode(), nil)
}

func (s *service) CloseDerivativeDeal(ctx context.Context, r dto.CloseDerivativeDealRequest) (*dto.CloseDerivativeDealResponse, error) {
	body := map[string]any{"orderType": r.OrderType, "triggeredBy": "close-deal"}
	return send[dto.CloseDerivativeDealResponse](ctx, s, CapabilityCloseDerivativeDeal, "close_derivative_deal", "POST", "/derivative/deals/"+url.PathEscape(r.DealID)+"/_close_deal", body)
}

func (s *service) GetRiskConfig(ctx context.Context, r dto.GetRiskConfigRequest) (*dto.GetRiskConfigResponse, error) {
	query := url.Values{"investorAccountId": []string{r.InvestorAccountID}}
	return send[dto.GetRiskConfigResponse](ctx, s, CapabilityRiskConfig, "get_risk_config", "GET", "/risk_configs?"+query.Encode(), nil)
}

func (s *service) UpdateRiskConfig(ctx context.Context, r dto.UpdateRiskConfigRequest) (*dto.UpdateRiskConfigResponse, error) {
	return send[dto.UpdateRiskConfigResponse](ctx, s, CapabilityUpdateRiskConfig, "update_risk_config", "PATCH", "/risk_configs/"+url.PathEscape(r.PathInvestorAccountID), r)
}

func send[T any](ctx context.Context, s *service, capability core.Capability, operation, method, path string, body any) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	response, err := s.dependencies.Send(ctx, "native.trading."+operation, transport.HTTPRequest{
		Method: method, URL: httpx.URL(s.dependencies.BaseURL, path, nil),
		Headers: s.dependencies.Headers(body != nil), JSON: body,
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[T]("entrade", "native.trading."+operation, "decode response", response)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func rangeQuery(accountID string, start, end int, sort, order string) url.Values {
	query := url.Values{}
	query.Set("investorAccountId", accountID)
	query.Set("_start", strconv.Itoa(start))
	if end == 0 {
		end = 20
	}
	query.Set("_end", strconv.Itoa(end))
	if sort != "" {
		query.Set("_sort", sort)
	}
	if order != "" {
		query.Set("_order", order)
	}
	return query
}
