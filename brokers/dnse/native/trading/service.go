package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityAccounts                core.Capability = "native.trading.accounts"
	CapabilityAccountBalances         core.Capability = "native.trading.account_balances"
	CapabilityCorporateActionHistory  core.Capability = "native.trading.corporate_action_history"
	CapabilityExecutions              core.Capability = "native.trading.executions"
	CapabilityLoanPackages            core.Capability = "native.trading.loan_packages"
	CapabilityOrderHistory            core.Capability = "native.trading.order_history"
	CapabilityOrders                  core.Capability = "native.trading.orders"
	CapabilityPositionPnLConfigs      core.Capability = "native.trading.position_pnl_configs"
	CapabilityPosition                core.Capability = "native.trading.position"
	CapabilityPositions               core.Capability = "native.trading.positions"
	CapabilityPPSE                    core.Capability = "native.trading.ppse"
	CapabilityCancelOrder             core.Capability = "native.trading.cancel_order"
	CapabilityClosePosition           core.Capability = "native.trading.close_position"
	CapabilityOrder                   core.Capability = "native.trading.order"
	CapabilityPlaceOrder              core.Capability = "native.trading.place_order"
	CapabilitySetPositionPnLConfigs   core.Capability = "native.trading.set_position_pnl_configs"
	CapabilityReplaceOrder            core.Capability = "native.trading.replace_order"
	CapabilityRealtimeOrders          core.Capability = "native.trading.realtime.orders"
	CapabilityRealtimeBrokerOrders    core.Capability = "native.trading.realtime.broker_orders"
	CapabilityRealtimePositions       core.Capability = "native.trading.realtime.positions"
	CapabilityRealtimeBrokerPositions core.Capability = "native.trading.realtime.broker_positions"
)

type Service interface {
	Realtime() RealtimeService
	GetAccounts(context.Context, dto.GetAccountsRequest) (*dto.AccountsResponse, error)
	GetAccountBalances(context.Context, dto.GetAccountBalancesRequest) (*dto.AccountBalancesResponse, error)
	GetCorporateActionHistory(context.Context, dto.GetCorporateActionHistoryRequest) (*dto.CorporateActionHistoryResponse, error)
	GetExecutions(context.Context, dto.GetExecutionsRequest) (*dto.ExecutionsResponse, error)
	GetLoanPackages(context.Context, dto.GetLoanPackagesRequest) (*dto.LoanPackagesResponse, error)
	GetOrderHistory(context.Context, dto.GetOrderHistoryRequest) (*dto.OrdersHistoryResponse, error)
	GetOrders(context.Context, dto.GetOrdersRequest) (*dto.OrdersResponse, error)
	GetPositionPnLConfigs(context.Context, dto.GetPositionPnLConfigsRequest) (*dto.PositionPnLConfigsResponse, error)
	GetPosition(context.Context, dto.GetPositionRequest) (*dto.PositionByIDResponse, error)
	GetPositions(context.Context, dto.GetPositionsRequest) (*dto.PositionsResponse, error)
	GetPPSE(context.Context, dto.GetPPSERequest) (*dto.PPSECredit, error)
	CancelOrder(context.Context, dto.CancelOrderRequest) (*dto.CancelOrderResponse, error)
	ClosePosition(context.Context, dto.ClosePositionRequest) (*dto.ClosePositionResponse, error)
	GetOrder(context.Context, dto.GetOrderRequest) (*dto.OrderDetailResponse, error)
	PlaceOrder(context.Context, dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, error)
	SetPositionPnLConfigs(context.Context, dto.SetPositionPnLConfigsRequest) (*dto.PositionPnLConfigsResponse, error)
	ReplaceOrder(context.Context, dto.ReplaceOrderRequest) (*dto.ReplaceOrderResponse, error)
}

type RealtimeService interface {
	SubscribeOrders(context.Context, dto.SubscribeTradingRequest) (realtime.Subscription[dto.OrderEvent], error)
	SubscribeBrokerOrders(context.Context, dto.SubscribeBrokerOrdersRequest) (realtime.Subscription[dto.BrokerOrderEvent], error)
	SubscribePositions(context.Context, dto.SubscribeTradingRequest) (realtime.Subscription[dto.PositionEvent], error)
	SubscribeBrokerPositions(context.Context, dto.SubscribeBrokerPositionsRequest) (realtime.Subscription[dto.BrokerPositionEvent], error)
}

type Dependencies struct {
	BaseURL                    string
	APIHeaders, TradingHeaders func(bool) map[string]string
	RequireCapability          func(core.Capability) error
	Send                       func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}
type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

func NewService(d Dependencies, realtimeServices ...RealtimeService) Service {
	var rt RealtimeService
	if len(realtimeServices) > 0 {
		rt = realtimeServices[0]
	}
	return &service{dependencies: d, realtime: rt}
}
func (s *service) Realtime() RealtimeService { return s.realtime }

func (s *service) GetAccounts(ctx context.Context, _ dto.GetAccountsRequest) (*dto.AccountsResponse, error) {
	return do[dto.AccountsResponse](s, ctx, CapabilityAccounts, "GET", "/accounts", nil, nil, false)
}
func (s *service) GetAccountBalances(ctx context.Context, r dto.GetAccountBalancesRequest) (*dto.AccountBalancesResponse, error) {
	return do[dto.AccountBalancesResponse](s, ctx, CapabilityAccountBalances, "GET", "/accounts/"+esc(r.AccountNo)+"/balances", nil, nil, false)
}
func (s *service) GetCorporateActionHistory(ctx context.Context, r dto.GetCorporateActionHistoryRequest) (*dto.CorporateActionHistoryResponse, error) {
	q := url.Values{}
	set(q, "symbol", r.Symbol)
	set(q, "caType", r.CAType)
	set(q, "caStatus", r.CAStatus)
	setInt(q, "pageIndex", r.PageIndex, true)
	setInt(q, "pageSize", r.PageSize, false)
	return do[dto.CorporateActionHistoryResponse](s, ctx, CapabilityCorporateActionHistory, "GET", "/accounts/"+esc(r.AccountNo)+"/corporate-action-history", q, nil, false)
}
func (s *service) GetExecutions(ctx context.Context, r dto.GetExecutionsRequest) (*dto.ExecutionsResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	return do[dto.ExecutionsResponse](s, ctx, CapabilityExecutions, "GET", "/accounts/"+esc(r.AccountNo)+"/executions/"+esc(r.OrderID), q, nil, false)
}
func (s *service) GetLoanPackages(ctx context.Context, r dto.GetLoanPackagesRequest) (*dto.LoanPackagesResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	set(q, "symbol", r.Symbol)
	return do[dto.LoanPackagesResponse](s, ctx, CapabilityLoanPackages, "GET", "/accounts/"+esc(r.AccountNo)+"/loan-packages", q, nil, false)
}
func (s *service) GetOrderHistory(ctx context.Context, r dto.GetOrderHistoryRequest) (*dto.OrdersHistoryResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	set(q, "from", r.From)
	set(q, "to", r.To)
	return do[dto.OrdersHistoryResponse](s, ctx, CapabilityOrderHistory, "GET", "/accounts/"+esc(r.AccountNo)+"/orders/history", q, nil, false)
}
func (s *service) GetOrders(ctx context.Context, r dto.GetOrdersRequest) (*dto.OrdersResponse, error) {
	return do[dto.OrdersResponse](s, ctx, CapabilityOrders, "GET", "/accounts/"+esc(r.AccountNo)+"/orders", orderQuery(r.MarketType, r.OrderCategory), nil, false)
}
func (s *service) GetPositionPnLConfigs(ctx context.Context, r dto.GetPositionPnLConfigsRequest) (*dto.PositionPnLConfigsResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	return do[dto.PositionPnLConfigsResponse](s, ctx, CapabilityPositionPnLConfigs, "GET", "/accounts/positions/"+esc(r.PositionID)+"/pnl-configs", q, nil, false)
}
func (s *service) GetPosition(ctx context.Context, r dto.GetPositionRequest) (*dto.PositionByIDResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	return do[dto.PositionByIDResponse](s, ctx, CapabilityPosition, "GET", "/accounts/positions/"+esc(r.PositionID), q, nil, false)
}
func (s *service) GetPositions(ctx context.Context, r dto.GetPositionsRequest) (*dto.PositionsResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	setInt(q, "pageSize", r.PageSize, false)
	return do[dto.PositionsResponse](s, ctx, CapabilityPositions, "GET", "/accounts/"+esc(r.AccountNo)+"/positions", q, nil, false)
}
func (s *service) GetPPSE(ctx context.Context, r dto.GetPPSERequest) (*dto.PPSECredit, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	set(q, "symbol", r.Symbol)
	if r.LoanPackageID != nil {
		q.Set("loanPackageId", strconv.Itoa(*r.LoanPackageID))
	}
	if r.Price != nil {
		q.Set("price", r.Price.String())
	}
	return do[dto.PPSECredit](s, ctx, CapabilityPPSE, "GET", "/accounts/"+esc(r.AccountNo)+"/ppse", q, nil, false)
}
func (s *service) CancelOrder(ctx context.Context, r dto.CancelOrderRequest) (*dto.CancelOrderResponse, error) {
	return do[dto.CancelOrderResponse](s, ctx, CapabilityCancelOrder, "DELETE", orderPath(r.AccountNo, r.OrderID), orderQuery(r.MarketType, r.OrderCategory), nil, true)
}
func (s *service) ClosePosition(ctx context.Context, r dto.ClosePositionRequest) (*dto.ClosePositionResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	return do[dto.ClosePositionResponse](s, ctx, CapabilityClosePosition, "POST", "/accounts/positions/"+esc(r.PositionID)+"/close", q, nil, true)
}
func (s *service) GetOrder(ctx context.Context, r dto.GetOrderRequest) (*dto.OrderDetailResponse, error) {
	return do[dto.OrderDetailResponse](s, ctx, CapabilityOrder, "GET", orderPath(r.AccountNo, r.OrderID), orderQuery(r.MarketType, r.OrderCategory), nil, false)
}
func (s *service) PlaceOrder(ctx context.Context, r dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, error) {
	body := map[string]any{"accountNo": r.AccountNo, "orderType": r.OrderType, "price": r.Price, "quantity": r.Quantity, "side": r.Side, "symbol": r.Symbol}
	if r.LoanPackageID != nil {
		body["loanPackageId"] = *r.LoanPackageID
	}
	return do[dto.PlaceOrderResponse](s, ctx, CapabilityPlaceOrder, "POST", "/accounts/orders", orderQuery(r.MarketType, r.OrderCategory), body, true)
}
func (s *service) SetPositionPnLConfigs(ctx context.Context, r dto.SetPositionPnLConfigsRequest) (*dto.PositionPnLConfigsResponse, error) {
	q := url.Values{}
	set(q, "marketType", r.MarketType)
	return do[dto.PositionPnLConfigsResponse](s, ctx, CapabilitySetPositionPnLConfigs, "POST", "/accounts/positions/"+esc(r.PositionID)+"/pnl-configs", q, r.Configs, true)
}
func (s *service) ReplaceOrder(ctx context.Context, r dto.ReplaceOrderRequest) (*dto.ReplaceOrderResponse, error) {
	body := map[string]any{"price": r.Price, "quantity": r.Quantity}
	return do[dto.ReplaceOrderResponse](s, ctx, CapabilityReplaceOrder, "PUT", orderPath(r.AccountNo, r.OrderID), orderQuery(r.MarketType, r.OrderCategory), body, true)
}

func do[T any](s *service, ctx context.Context, capability core.Capability, method, path string, q url.Values, body any, trading bool) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	headers := map[string]string{}
	if trading && s.dependencies.TradingHeaders != nil {
		headers = s.dependencies.TradingHeaders(body != nil)
	} else if s.dependencies.APIHeaders != nil {
		headers = s.dependencies.APIHeaders(false)
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{Method: method, URL: httpx.URL(s.dependencies.BaseURL, path, q), Headers: headers, JSON: body})
	if err != nil {
		return nil, err
	}
	result, err := httpx.DecodeResponse[T]("dnse", string(capability), "decode DNSE native trading response", response)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func esc(v string) string { return url.PathEscape(v) }
func set(q url.Values, k, v string) {
	if v != "" {
		q.Set(k, v)
	}
}
func setInt(q url.Values, k string, v int, includeZero bool) {
	if v != 0 || includeZero {
		q.Set(k, strconv.Itoa(v))
	}
}
func orderQuery(m, c string) url.Values {
	q := url.Values{}
	set(q, "marketType", m)
	set(q, "orderCategory", c)
	return q
}
func orderPath(a, o string) string { return "/accounts/" + esc(a) + "/orders/" + esc(o) }
