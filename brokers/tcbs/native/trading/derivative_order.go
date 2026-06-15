package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetDerivativeOrders(ctx context.Context, request dto.GetDerivativeOrdersRequest) (*dto.GetDerivativeOrdersResponse, error) {
	query := url.Values{}
	set(query, "pageNo", strconv.FormatFloat(request.PageNo, 'f', -1, 64))
	set(query, "pageSize", strconv.FormatFloat(request.PageSize, 'f', -1, 64))
	set(query, "accountId", request.AccountID)
	set(query, "symbol", request.Symbol)
	set(query, "refId", request.RefID)
	set(query, "orderType", request.OrderType)
	set(query, "status", request.Status)
	return do[dto.GetDerivativeOrdersResponse](s, ctx, CapabilityGetDerivativeOrders, "GET", "/khronos/v1/order/in-day", query, nil)
}

func (s *service) GetDerivativeConditionalOrders(ctx context.Context, request dto.GetDerivativeConditionalOrdersRequest) (*dto.GetDerivativeConditionalOrdersResponse, error) {
	query := url.Values{}
	set(query, "pageNo", request.PageNo)
	set(query, "PageSize", request.PageSize)
	set(query, "accountId", request.AccountID)
	set(query, "subAccountID", request.SubAccountID)
	set(query, "orderStatus", request.OrderStatus)
	set(query, "orderType", request.OrderType)
	set(query, "Symbol", request.Symbol)
	return do[dto.GetDerivativeConditionalOrdersResponse](s, ctx, CapabilityGetDerivativeConditionalOrders, "GET", "/khronos/v1/order/condition/detail", query, nil)
}

func (s *service) PlaceDerivativeOrder(ctx context.Context, request dto.PlaceDerivativeOrderRequest) (*dto.PlaceDerivativeOrderResponse, error) {
	return do[dto.PlaceDerivativeOrderResponse](s, ctx, CapabilityPlaceDerivativeOrder, "POST", "/khronos/v1/order/place", url.Values{}, request)
}

func (s *service) PlaceDerivativeConditionalOrder(ctx context.Context, request dto.PlaceDerivativeConditionalOrderRequest) (*dto.PlaceDerivativeConditionalOrderResponse, error) {
	return do[dto.PlaceDerivativeConditionalOrderResponse](s, ctx, CapabilityPlaceDerivativeConditionalOrder, "POST", "/khronos/v1/order/condition/place", url.Values{}, request)
}

func (s *service) UpdateDerivativeOrder(ctx context.Context, request dto.UpdateDerivativeOrderRequest) (*dto.UpdateDerivativeOrderResponse, error) {
	return do[dto.UpdateDerivativeOrderResponse](s, ctx, CapabilityUpdateDerivativeOrder, "POST", "/khronos/v1/order/change", url.Values{}, request)
}

func (s *service) UpdateDerivativeConditionalOrder(ctx context.Context, request dto.UpdateDerivativeConditionalOrderRequest) (*dto.UpdateDerivativeConditionalOrderResponse, error) {
	return do[dto.UpdateDerivativeConditionalOrderResponse](s, ctx, CapabilityUpdateDerivativeConditionalOrder, "POST", "/khronos/v2/order/condition/change", url.Values{}, request)
}

func (s *service) CancelDerivativeOrder(ctx context.Context, request dto.CancelDerivativeOrderRequest) (*dto.CancelDerivativeOrderResponse, error) {
	return do[dto.CancelDerivativeOrderResponse](s, ctx, CapabilityCancelDerivativeOrder, "POST", "/khronos/v1/order/cancel", url.Values{}, request)
}

func (s *service) CancelDerivativeConditionalOrder(ctx context.Context, request dto.CancelDerivativeConditionalOrderRequest) (*dto.CancelDerivativeConditionalOrderResponse, error) {
	return do[dto.CancelDerivativeConditionalOrderResponse](s, ctx, CapabilityCancelDerivativeConditionalOrder, "POST", "/khronos/v1/order/condition/cancel", url.Values{}, request)
}
