package entrade

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

type TradingService struct {
	accounts *TradingAccountsService
	orders   *TradingOrdersService
	deals    *TradingDealsService
	risk     *TradingRiskService
}

func (s *TradingService) Accounts() *TradingAccountsService {
	return s.accounts
}

func (s *TradingService) Orders() *TradingOrdersService {
	return s.orders
}

func (s *TradingService) Deals() *TradingDealsService {
	return s.deals
}

func (s *TradingService) Risk() *TradingRiskService {
	return s.risk
}

type TradingAccountsService struct {
	broker *Broker
}

func (s *TradingAccountsService) Master(ctx context.Context, investorID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/investors/" + url.PathEscape(investorID) + "/investor_account"
	return s.broker.sendRaw(ctx, "trading.accounts.master", "GET", path, nil)
}

func (s *TradingAccountsService) Balance(ctx context.Context, investorID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsBalance); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/account_balances/" + url.PathEscape(investorID)
	return s.broker.sendRaw(ctx, "trading.accounts.balance", "GET", path, nil)
}

func (s *TradingAccountsService) LoanPackages(ctx context.Context, investorID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/investors/" + url.PathEscape(investorID) + "/derivative_margin_portfolios"
	return s.broker.sendRaw(ctx, "trading.accounts.loan_packages", "GET", path, nil)
}

func (s *TradingAccountsService) BuyingPower(ctx context.Context, request BuyingPowerRequest) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("bankMarginPortfolioId", request.BankMarginPortfolioID)
	params.Set("price", request.Price.String())
	params.Set("side", request.Side)
	params.Set("symbol", request.Symbol)
	path := "/derivative/investors/" + url.PathEscape(request.InvestorID) + "/ppse?" + params.Encode()
	return s.broker.sendRaw(ctx, "trading.accounts.buying_power", "GET", path, nil)
}

type TradingOrdersService struct {
	broker *Broker
}

func (s *TradingOrdersService) Place(ctx context.Context, request PlaceDerivativeOrderRequest) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.RawPayload{}, err
	}
	body := map[string]any{
		"bankMarginPortfolioId": request.BankMarginPortfolioID,
		"investorId":            request.InvestorID,
		"symbol":                request.Symbol,
		"price":                 numberValue(&request.Price),
		"orderType":             request.OrderType,
		"side":                  request.Side,
		"quantity":              request.Quantity,
	}
	return s.broker.sendRaw(ctx, "trading.orders.place", "POST", "/derivative/orders", body)
}

func (s *TradingOrdersService) List(ctx context.Context, request ListOrdersRequest) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersList); err != nil {
		return domain.RawPayload{}, err
	}
	params := rangeParams(request.InvestorAccountID, request.Start, request.End, request.Sort, request.Order)
	return s.broker.sendRaw(ctx, "trading.orders.list", "GET", "/derivative/orders?"+params.Encode(), nil)
}

func (s *TradingOrdersService) Get(ctx context.Context, orderID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersGet); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/derivative/orders/" + url.PathEscape(orderID)
	return s.broker.sendRaw(ctx, "trading.orders.get", "GET", path, nil)
}

func (s *TradingOrdersService) Cancel(ctx context.Context, orderID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersCancel); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/derivative/orders/" + url.PathEscape(orderID)
	return s.broker.sendRaw(ctx, "trading.orders.cancel", "DELETE", path, nil)
}

type TradingDealsService struct {
	broker *Broker
}

func (s *TradingDealsService) List(ctx context.Context, request ListDealsRequest) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return domain.RawPayload{}, err
	}
	params := rangeParams(request.InvestorAccountID, request.Start, request.End, request.Sort, request.Order)
	return s.broker.sendRaw(ctx, "trading.deals.list", "GET", "/derivative/deals?"+params.Encode(), nil)
}

func (s *TradingDealsService) Close(ctx context.Context, dealID string, orderType string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsClose); err != nil {
		return domain.RawPayload{}, err
	}
	body := map[string]any{
		"orderType":   orderType,
		"triggeredBy": "close-deal",
	}
	path := "/derivative/deals/" + url.PathEscape(dealID) + "/_close_deal"
	return s.broker.sendRaw(ctx, "trading.deals.close", "POST", path, body)
}

type TradingRiskService struct {
	broker *Broker
}

func (s *TradingRiskService) Config(ctx context.Context, investorAccountID string) (domain.RawPayload, error) {
	params := url.Values{}
	params.Set("investorAccountId", investorAccountID)
	return s.broker.sendRaw(ctx, "trading.risk.config", "GET", "/risk_configs?"+params.Encode(), nil)
}

func (s *TradingRiskService) UpdateConfig(
	ctx context.Context,
	investorAccountID string,
	request RiskConfigRequest,
) (domain.RawPayload, error) {
	body := map[string]any{
		"cutLossRate":               numberValue(&request.CutLossRate),
		"investorAccountId":         request.InvestorAccountID,
		"trailingEnabled":           request.TrailingEnabled,
		"investorId":                request.InvestorID,
		"autoIncreaseDealRate":      request.AutoIncreaseDealRate,
		"enableAutoDealDepositNoti": request.EnableAutoDealDepositNoti,
	}
	path := "/risk_configs/" + url.PathEscape(investorAccountID)
	return s.broker.sendRaw(ctx, "trading.risk.update_config", "PATCH", path, body)
}

func rangeParams(investorAccountID string, start int, end int, sort string, order string) url.Values {
	params := url.Values{}
	params.Set("investorAccountId", investorAccountID)
	params.Set("_start", stringify(start))
	if end == 0 {
		end = 20
	}
	params.Set("_end", stringify(end))
	if sort != "" {
		params.Set("_sort", sort)
	}
	if order != "" {
		params.Set("_order", order)
	}
	return params
}
