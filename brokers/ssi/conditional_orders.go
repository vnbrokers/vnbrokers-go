package ssi

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/core"
)

type TradingConditionalOrdersService struct {
	broker *Broker
}

func (s *TradingConditionalOrdersService) NewWithRequest(
	ctx context.Context,
	request ConditionalOrderNewRequest,
) (ConditionalOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return ConditionalOrderResponse{}, err
	}
	body := map[string]any{
		"instrumentID": request.Symbol,
		"side":         request.Side,
		"type":         request.Type,
		"price":        request.Price,
		"quantity":     numberValue(&request.Quantity),
		"account":      request.AccountID,
		"fromDate":     request.FromDate,
		"toDate":       request.ToDate,
		"tpPrice":      request.TPPrice,
		"slPrice":      request.SLPrice,
		"operator":     request.Operator,
		"code":         request.Code,
		"userAgent":    firstString(request.UserAgent, s.broker.config.UserAgent),
		"deviceId":     firstString(request.DeviceID, s.broker.config.DeviceID),
	}
	setOptionalDecimalBody(body, "priceSlip", request.PriceSlip)
	setOptionalDecimalBody(body, "stopPrice", request.StopPrice)
	setOptionalDecimalBody(body, "activePrice", request.ActivePrice)
	setOptionalDecimalBody(body, "trailingAmount", request.TrailingAmount)
	setOptionalDecimalBody(body, "tpActivePrice", request.TPActivePrice)
	setOptionalDecimalBody(body, "slActivePrice", request.SLActivePrice)
	setOptionalDecimalBody(body, "tpSlip", request.TPSlip)
	setOptionalDecimalBody(body, "slSlip", request.SLSlip)
	var response ConditionalOrderResponse
	err := s.broker.postAndDecode(ctx, "trading.conditional_orders.new", "/api/v2/fco/neworder", body, &response)
	return response, err
}

func (s *TradingConditionalOrdersService) CancelWithRequest(
	ctx context.Context,
	request ConditionalOrderCancelRequest,
) (ConditionalOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return ConditionalOrderResponse{}, err
	}
	body := map[string]any{
		"fcoId":     request.FCOID,
		"code":      request.Code,
		"userAgent": firstString(request.UserAgent, s.broker.config.UserAgent),
		"deviceId":  firstString(request.DeviceID, s.broker.config.DeviceID),
	}
	var response ConditionalOrderResponse
	err := s.broker.postAndDecode(ctx, "trading.conditional_orders.cancel", "/api/v2/fco/cancelorder", body, &response)
	return response, err
}

func (s *TradingConditionalOrdersService) OrderBook(
	ctx context.Context,
	request ConditionalOrderBookRequest,
) (ConditionalOrderPage[ConditionalTriggeredOrder], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return ConditionalOrderPage[ConditionalTriggeredOrder]{}, err
	}
	params := fcoPageParams(request.FCOID, request.PageIndex, request.PageSize)
	var response ConditionalOrderPage[ConditionalTriggeredOrder]
	err := s.broker.getAndDecode(ctx, "trading.conditional_orders.order_book", "/api/v2/fco/orderbook", params, &response)
	return response, err
}

func (s *TradingConditionalOrdersService) StatusHistory(
	ctx context.Context,
	request ConditionalOrderStatusHistoryRequest,
) (ConditionalOrderPage[ConditionalOrderStatus], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return ConditionalOrderPage[ConditionalOrderStatus]{}, err
	}
	params := fcoPageParams(request.FCOID, request.PageIndex, request.PageSize)
	var response ConditionalOrderPage[ConditionalOrderStatus]
	err := s.broker.getAndDecode(ctx, "trading.conditional_orders.status_history", "/api/v2/fco/statusHistory", params, &response)
	return response, err
}

func (s *TradingConditionalOrdersService) List(
	ctx context.Context,
	request ConditionalOrderListRequest,
) (ConditionalOrderPage[ConditionalOrder], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingConditionalOrders); err != nil {
		return ConditionalOrderPage[ConditionalOrder]{}, err
	}
	params := url.Values{}
	setOptionalString(params, "fcoId", request.FCOID)
	setOptionalString(params, "account", request.AccountID)
	setOptionalString(params, "type", request.Type)
	setOptionalString(params, "processStatus", request.ProcessStatus)
	setOptionalString(params, "instrumentID", request.Symbol)
	setOptionalString(params, "side", request.Side)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalInt(params, "pageIndex", request.PageIndex)
	setOptionalInt(params, "pageSize", request.PageSize)
	var response ConditionalOrderPage[ConditionalOrder]
	err := s.broker.getAndDecode(ctx, "trading.conditional_orders.list", "/api/v2/fco/list", params, &response)
	return response, err
}

func fcoPageParams(fcoID string, pageIndex int, pageSize int) url.Values {
	params := url.Values{}
	setOptionalString(params, "fcoId", fcoID)
	setOptionalInt(params, "pageIndex", pageIndex)
	setOptionalInt(params, "pageSize", pageSize)
	return params
}
