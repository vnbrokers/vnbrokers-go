package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

// ──────────────────────────── Accounts ────────────────────────────

func (s *service) CashAcctBal(ctx context.Context, accountID string) (*dto.TradingResponse[[]dto.StockAccountBalance], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[[]dto.StockAccountBalance]](ctx, s, CapabilityCashAcctBal, "/api/v2/Trading/cashAcctBal", params)
}

func (s *service) DerivAcctBal(ctx context.Context, accountID string) (*dto.TradingResponse[[]dto.DerivativeAccountBalance], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[[]dto.DerivativeAccountBalance]](ctx, s, CapabilityDerivAcctBal, "/api/v2/Trading/derivAcctBal", params)
}

func (s *service) MaxBuyQty(ctx context.Context, accountID string, symbol string, price string) (*dto.TradingResponse[[]dto.MaxBuyQuantityData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("instrumentID", symbol)
	params.Set("price", price)
	return get[dto.TradingResponse[[]dto.MaxBuyQuantityData]](ctx, s, CapabilityMaxBuyQty, "/api/v2/Trading/maxBuyQty", params)
}

func (s *service) MaxSellQty(ctx context.Context, accountID string, symbol string, price string) (*dto.TradingResponse[dto.MaxSellQuantityData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("instrumentID", symbol)
	setOptionalString(params, "price", price)
	return get[dto.TradingResponse[dto.MaxSellQuantityData]](ctx, s, CapabilityMaxSellQty, "/api/v2/Trading/maxSellQty", params)
}

func (s *service) PpmrAccount(ctx context.Context, accountID string) (*dto.TradingResponse[dto.AccountAssetData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.AccountAssetData]](ctx, s, CapabilityPpmrAccount, "/api/v2/Trading/ppmmraccount", params)
}

func (s *service) RateLimit(ctx context.Context) (*dto.TradingResponse[[]dto.APILimitData], error) {
	return get[dto.TradingResponse[[]dto.APILimitData]](ctx, s, CapabilityRateLimit, "/api/v2/Trading/rateLimit", nil)
}

// ──────────────────────────── Orders ────────────────────────────

func (s *service) OrderBook(ctx context.Context, accountID string) (*dto.TradingResponse[dto.OrderBookData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.OrderBookData]](ctx, s, CapabilityOrderBook, "/api/v2/Trading/OrderBook", params)
}

func (s *service) OrderHistory(ctx context.Context, accountID string, fromDate string, toDate string, pageIndex int) (*dto.TradingResponse[dto.OrderHistoryData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", fromDate)
	params.Set("endDate", toDate)
	params.Set("pageIndex", strconv.Itoa(pageIndex))
	return get[dto.TradingResponse[dto.OrderHistoryData]](ctx, s, CapabilityOrderHistory, "/api/v2/Trading/orderHistory", params)
}

func (s *service) NewOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityNewOrder, "/api/v2/Trading/NewOrder", body)
}

func (s *service) CancelOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityCancelOrder, "/api/v2/Trading/CancelOrder", body)
}

func (s *service) ModifyOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityModifyOrder, "/api/v2/Trading/ModifyOrder", body)
}

func (s *service) DerNewOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityDerNewOrder, "/api/v2/Trading/derNewOrder", body)
}

func (s *service) DerCancelOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityDerCancelOrder, "/api/v2/Trading/derCancelOrder", body)
}

func (s *service) DerModifyOrder(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.PlaceOrderResponse], error) {
	return post[dto.TradingResponse[dto.PlaceOrderResponse]](ctx, s, CapabilityDerModifyOrder, "/api/v2/Trading/derModifyOrder", body)
}

func (s *service) AuditOrderBook(ctx context.Context, accountID string) (*dto.TradingResponse[dto.OrderBookData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.OrderBookData]](ctx, s, CapabilityAuditOrderBook, "/api/v2/Trading/auditOrderBook", params)
}

// ──────────────────────────── Positions ────────────────────────────

func (s *service) StockPosition(ctx context.Context, accountID string) (*dto.TradingResponse[[]dto.StockPortfolioData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[[]dto.StockPortfolioData]](ctx, s, CapabilityStockPosition, "/api/v2/Trading/stockPosition", params)
}

func (s *service) DerivPosition(ctx context.Context, accountID string, querySummary bool) (*dto.TradingResponse[[]dto.DerivativePositionsData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("querySummary", strconv.FormatBool(querySummary))
	return get[dto.TradingResponse[[]dto.DerivativePositionsData]](ctx, s, CapabilityDerivPosition, "/api/v2/Trading/derivPosition", params)
}
