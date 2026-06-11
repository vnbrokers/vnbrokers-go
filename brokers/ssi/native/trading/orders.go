package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

// ──────────────────────────── Accounts ────────────────────────────

func (s *service) CashAcctBal(ctx context.Context, accountID string) (*dto.CashAccountBalanceResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.CashAccountBalanceResponse](ctx, s, CapabilityCashAcctBal, "/api/v2/Trading/cashAcctBal", params)
}

func (s *service) DerivAcctBal(ctx context.Context, accountID string) (*dto.DerivativeAccountBalanceResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.DerivativeAccountBalanceResponse](ctx, s, CapabilityDerivAcctBal, "/api/v2/Trading/derivAcctBal", params)
}

func (s *service) MaxBuyQty(ctx context.Context, accountID string, symbol string, price string) (*dto.MaxBuyQtyResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("instrumentID", symbol)
	params.Set("price", price)
	return get[dto.MaxBuyQtyResponse](ctx, s, CapabilityMaxBuyQty, "/api/v2/Trading/maxBuyQty", params)
}

func (s *service) MaxSellQty(ctx context.Context, accountID string, symbol string, price string) (*dto.MaxSellQtyResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("instrumentID", symbol)
	setOptionalString(params, "price", price)
	return get[dto.MaxSellQtyResponse](ctx, s, CapabilityMaxSellQty, "/api/v2/Trading/maxSellQty", params)
}

func (s *service) PpmrAccount(ctx context.Context, accountID string) (*dto.PpmmrAccountResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.PpmmrAccountResponse](ctx, s, CapabilityPpmrAccount, "/api/v2/Trading/ppmmraccount", params)
}

func (s *service) RateLimit(ctx context.Context) (*dto.RateLimitResponse, error) {
	return get[dto.RateLimitResponse](ctx, s, CapabilityRateLimit, "/api/v2/Trading/rateLimit", nil)
}

// ──────────────────────────── Orders ────────────────────────────

func (s *service) OrderBook(ctx context.Context, accountID string) (*dto.OrderBookResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.OrderBookResponse](ctx, s, CapabilityOrderBook, "/api/v2/Trading/OrderBook", params)
}

func (s *service) OrderHistory(ctx context.Context, accountID string, fromDate string, toDate string, pageIndex int) (*dto.OrderHistoryResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", fromDate)
	params.Set("endDate", toDate)
	params.Set("pageIndex", strconv.Itoa(pageIndex))
	return get[dto.OrderHistoryResponse](ctx, s, CapabilityOrderHistory, "/api/v2/Trading/orderHistory", params)
}

func (s *service) NewOrder(ctx context.Context, body map[string]any) (*dto.NewOrderResponse, error) {
	return post[dto.NewOrderResponse](ctx, s, CapabilityNewOrder, "/api/v2/Trading/NewOrder", body)
}

func (s *service) CancelOrder(ctx context.Context, body map[string]any) (*dto.CancelOrderResponse, error) {
	return post[dto.CancelOrderResponse](ctx, s, CapabilityCancelOrder, "/api/v2/Trading/CancelOrder", body)
}

func (s *service) ModifyOrder(ctx context.Context, body map[string]any) (*dto.ModifyOrderResponse, error) {
	return post[dto.ModifyOrderResponse](ctx, s, CapabilityModifyOrder, "/api/v2/Trading/ModifyOrder", body)
}

func (s *service) DerNewOrder(ctx context.Context, body map[string]any) (*dto.DerNewOrderResponse, error) {
	return post[dto.DerNewOrderResponse](ctx, s, CapabilityDerNewOrder, "/api/v2/Trading/derNewOrder", body)
}

func (s *service) DerCancelOrder(ctx context.Context, body map[string]any) (*dto.DerCancelOrderResponse, error) {
	return post[dto.DerCancelOrderResponse](ctx, s, CapabilityDerCancelOrder, "/api/v2/Trading/derCancelOrder", body)
}

func (s *service) DerModifyOrder(ctx context.Context, body map[string]any) (*dto.DerModifyOrderResponse, error) {
	return post[dto.DerModifyOrderResponse](ctx, s, CapabilityDerModifyOrder, "/api/v2/Trading/derModifyOrder", body)
}

func (s *service) AuditOrderBook(ctx context.Context, accountID string) (*dto.AuditOrderBookResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.AuditOrderBookResponse](ctx, s, CapabilityAuditOrderBook, "/api/v2/Trading/auditOrderBook", params)
}

// ──────────────────────────── Positions ────────────────────────────

func (s *service) StockPosition(ctx context.Context, accountID string) (*dto.StockPositionResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.StockPositionResponse](ctx, s, CapabilityStockPosition, "/api/v2/Trading/stockPosition", params)
}

func (s *service) DerivPosition(ctx context.Context, accountID string, querySummary bool) (*dto.DerivativePositionResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("querySummary", strconv.FormatBool(querySummary))
	return get[dto.DerivativePositionResponse](ctx, s, CapabilityDerivPosition, "/api/v2/Trading/derivPosition", params)
}
