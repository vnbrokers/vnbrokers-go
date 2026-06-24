package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func (s *service) GetOrderHistory(ctx context.Context, request dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error) {
	query := url.Values{}
	set(query, "fromDate", request.FromDate)
	set(query, "toDate", request.ToDate)
	if request.Page > 0 {
		set(query, "page", strconv.FormatInt(request.Page, 10))
	}
	setOptional(query, "orderStatus", request.OrderStatus)
	setOptional(query, "symbol", request.Symbol)
	path := "/trading/sub-accounts/" + escaped(request.SubAccountID) + "/orders"
	return do[dto.GetOrderHistoryResponse](s, ctx, CapabilityGetOrderHistory, "GET", path, query, nil)
}

func (s *service) GetOrderBookDetail(ctx context.Context, request dto.GetOrderBookDetailRequest) (*dto.GetOrderBookDetailResponse, error) {
	query := url.Values{}
	setOptional(query, "cache-control", request.CacheControl)
	path := "/trading/v1/accounts/" + escaped(request.SubAccountID) + "/order-book/" + escaped(request.OrderID)
	return do[dto.GetOrderBookDetailResponse](s, ctx, CapabilityGetOrderBookDetail, "GET", path, query, nil)
}

func (s *service) GetOrderBook(ctx context.Context, request dto.GetOrderBookRequest) (*dto.GetOrderBookResponse, error) {
	query := url.Values{}
	setOptional(query, "cache-control", request.CacheControl)
	path := "/trading/v1/accounts/" + escaped(request.SubAccountID) + "/order-book"
	return do[dto.GetOrderBookResponse](s, ctx, core.CapabilityTradingOrdersList, "GET", path, query, nil)
}

func (s *service) PlaceOrder(ctx context.Context, request dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, error) {
	path := "/trading/oa/sub-accounts/" + escaped(request.SubAccountID) + "/orders"
	return do[dto.PlaceOrderResponse](s, ctx, core.CapabilityTradingOrdersPlace, "POST", path, url.Values{}, request.Body)
}

func (s *service) CancelOrder(ctx context.Context, request dto.CancelOrderOperationRequest) (*dto.CancelOrderOperationResponse, error) {
	path := "/trading/oa/sub-accounts/" + escaped(request.SubAccountID) + "/orders/" + escaped(request.OrderID)
	return do[dto.CancelOrderOperationResponse](s, ctx, core.CapabilityTradingOrdersCancel, "DELETE", path, url.Values{}, request.Body)
}

func (s *service) UpdateOrder(ctx context.Context, request dto.UpdateOrderOperationRequest) (*dto.UpdateOrderOperationResponse, error) {
	path := "/trading/oa/sub-accounts/" + escaped(request.SubAccountID) + "/orders/" + escaped(request.OrderID)
	return do[dto.UpdateOrderOperationResponse](s, ctx, core.CapabilityTradingOrdersReplace, "PUT", path, url.Values{}, request.Body)
}
