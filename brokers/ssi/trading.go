package ssi

import (
	"context"
	"net/url"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type TradingAccountsService struct {
	broker *Broker
}

func (s *TradingAccountsService) Balance(ctx context.Context, accountID string) (domain.Balance, error) {
	response, err := s.StockBalance(ctx, accountID)
	if err != nil {
		return domain.Balance{}, err
	}
	if len(response.Data) == 0 {
		return domain.Balance{AccountID: accountID, Currency: "VND", Raw: rawPayload(response, nil)}, nil
	}
	return MapStockBalance(response.Data[0]), nil
}

func (s *TradingAccountsService) Orders(ctx context.Context, accountID string) ([]domain.Order, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersList); err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("account", accountID)
	var response TradingResponse[[]Order]
	if err := s.broker.sendAndDecode(ctx, "trading.accounts.orders", "GET", s.broker.query("/api/v2/Trading/OrderBook", params), nil, false, &response); err != nil {
		return nil, err
	}
	return MapOrders(response.Data), nil
}

func (s *TradingAccountsService) OrderHistory(
	ctx context.Context,
	accountID string,
	fromDate string,
	toDate string,
	pageIndex int,
) ([]domain.Order, error) {
	return s.OrderHistoryWithRequest(ctx, sdktrading.OrderHistoryRequest{
		AccountID: accountID,
		FromDate:  fromDate,
		ToDate:    toDate,
		PageIndex: pageIndex,
	})
}

func (s *TradingAccountsService) OrderHistoryWithRequest(
	ctx context.Context,
	request sdktrading.OrderHistoryRequest,
) ([]domain.Order, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersHistory); err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("account", request.AccountID)
	params.Set("startDate", request.FromDate)
	params.Set("endDate", request.ToDate)
	var response TradingResponse[OrderHistoryData]
	if err := s.broker.sendAndDecode(ctx, "trading.accounts.order_history", "GET", s.broker.query("/api/v2/Trading/orderHistory", params), nil, false, &response); err != nil {
		return nil, err
	}
	return MapOrders(response.Data.OrderHistories), nil
}

func (s *TradingAccountsService) PPSE(
	ctx context.Context,
	accountID string,
	symbol string,
	price decimal.Decimal,
	_ *int,
) (domain.RawPayload, error) {
	response, err := s.MaxBuyQuantity(ctx, MaxBuyQuantityRequest{
		AccountID: accountID,
		Symbol:    symbol,
		Price:     price,
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response, nil), nil
}

func (s *TradingAccountsService) StockBalance(
	ctx context.Context,
	accountID string,
) (TradingResponse[[]StockAccountBalance], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsBalance); err != nil {
		return TradingResponse[[]StockAccountBalance]{}, err
	}
	params := url.Values{}
	params.Set("account", accountID)
	var response TradingResponse[[]StockAccountBalance]
	err := s.broker.sendAndDecode(ctx, "trading.accounts.stock_balance", "GET", s.broker.query("/api/v2/Trading/cashAcctBal", params), nil, false, &response)
	return response, err
}

func (s *TradingAccountsService) DerivativeBalance(
	ctx context.Context,
	accountID string,
) (TradingResponse[[]DerivativeAccountBalance], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsBalance); err != nil {
		return TradingResponse[[]DerivativeAccountBalance]{}, err
	}
	params := url.Values{}
	params.Set("account", accountID)
	var response TradingResponse[[]DerivativeAccountBalance]
	err := s.broker.sendAndDecode(ctx, "trading.accounts.derivative_balance", "GET", s.broker.query("/api/v2/Trading/derivAcctBal", params), nil, false, &response)
	return response, err
}

func (s *TradingAccountsService) MaxBuyQuantity(
	ctx context.Context,
	request MaxBuyQuantityRequest,
) (TradingResponse[[]MaxBuyQuantity], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return TradingResponse[[]MaxBuyQuantity]{}, err
	}
	params := url.Values{}
	params.Set("account", request.AccountID)
	params.Set("instrumentID", request.Symbol)
	params.Set("price", request.Price.String())
	var response TradingResponse[[]MaxBuyQuantity]
	err := s.broker.sendAndDecode(ctx, "trading.accounts.max_buy_quantity", "GET", s.broker.query("/api/v2/Trading/maxBuyQty", params), nil, false, &response)
	return response, err
}

type TradingPositionsService struct {
	broker *Broker
}

func (s *TradingPositionsService) List(ctx context.Context, accountID string) ([]domain.Position, error) {
	response, err := s.Stock(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return MapStockPortfolios(accountID, response.Data), nil
}

func (s *TradingPositionsService) Stock(
	ctx context.Context,
	accountID string,
) (TradingResponse[[]StockPortfolio], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return TradingResponse[[]StockPortfolio]{}, err
	}
	params := url.Values{}
	params.Set("account", accountID)
	var response TradingResponse[[]StockPortfolio]
	err := s.broker.sendAndDecode(ctx, "trading.positions.stock", "GET", s.broker.query("/api/v2/Trading/stockPosition", params), nil, false, &response)
	return response, err
}

func (s *TradingPositionsService) Derivative(
	ctx context.Context,
	accountID string,
) (TradingResponse[[]DerivativePositions], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return TradingResponse[[]DerivativePositions]{}, err
	}
	params := url.Values{}
	params.Set("account", accountID)
	var response TradingResponse[[]DerivativePositions]
	err := s.broker.sendAndDecode(ctx, "trading.positions.derivative", "GET", s.broker.query("/api/v2/Trading/derivPosition", params), nil, false, &response)
	return response, err
}

type TradingOrdersService struct {
	broker *Broker
}

func (s *TradingOrdersService) Place(
	ctx context.Context,
	request domain.PlaceOrderRequest,
) (domain.PlaceOrderResponse, error) {
	response, err := s.PlaceWithRequest(ctx, PlaceOrderRequest{PlaceOrderRequest: request})
	if err != nil {
		return domain.PlaceOrderResponse{}, err
	}
	return MapPlaceOrderResponse(response.Data), nil
}

func (s *TradingOrdersService) PlaceWithRequest(
	ctx context.Context,
	request PlaceOrderRequest,
) (TradingResponse[OrderRequestResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return TradingResponse[OrderRequestResponse]{}, err
	}
	body := s.broker.placeOrderBody(request)
	var response TradingResponse[OrderRequestResponse]
	err := s.broker.sendAndDecode(ctx, "trading.orders.place", "POST", s.broker.url("/api/v2/Trading/NewOrder"), body, true, &response)
	return response, err
}

func (s *TradingOrdersService) Cancel(ctx context.Context, accountID string, orderID string) error {
	_, err := s.CancelWithRequest(ctx, CancelOrderRequest{AccountID: accountID, OrderID: orderID})
	return err
}

func (s *TradingOrdersService) CancelWithRequest(
	ctx context.Context,
	request CancelOrderRequest,
) (TradingResponse[OrderRequestResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersCancel); err != nil {
		return TradingResponse[OrderRequestResponse]{}, err
	}
	body := s.broker.cancelOrderBody(request)
	var response TradingResponse[OrderRequestResponse]
	err := s.broker.sendAndDecode(ctx, "trading.orders.cancel", "POST", s.broker.url("/api/v2/Trading/CancelOrder"), body, true, &response)
	return response, err
}

func (s *TradingOrdersService) Get(ctx context.Context, accountID string, orderID string) (domain.Order, error) {
	orders, err := s.broker.Trading().Accounts().Orders(ctx, accountID)
	if err != nil {
		return domain.Order{}, err
	}
	for _, order := range orders {
		if order.OrderID == orderID {
			return order, nil
		}
	}
	return domain.Order{}, sdkerrors.BrokerRejected("ssi", "trading.orders.get", "not_found", "order not found", nil)
}

func (s *TradingOrdersService) Update(
	ctx context.Context,
	accountID string,
	orderID string,
	price decimal.Decimal,
	quantity int,
) (domain.RawPayload, error) {
	response, err := s.ModifyWithRequest(ctx, ModifyOrderRequest{
		AccountID: accountID,
		OrderID:   orderID,
		Price:     price,
		Quantity:  quantity,
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response, nil), nil
}

func (s *TradingOrdersService) ModifyWithRequest(
	ctx context.Context,
	request ModifyOrderRequest,
) (TradingResponse[OrderRequestResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersReplace); err != nil {
		return TradingResponse[OrderRequestResponse]{}, err
	}
	body := s.broker.modifyOrderBody(request)
	var response TradingResponse[OrderRequestResponse]
	err := s.broker.sendAndDecode(ctx, "trading.orders.modify", "POST", s.broker.url("/api/v2/Trading/ModifyOrder"), body, true, &response)
	return response, err
}

func (b *Broker) sendAndDecode(
	ctx context.Context,
	operation string,
	method string,
	url string,
	body any,
	sign bool,
	out any,
) error {
	response, err := b.send(ctx, operation, true, sign, transport.HTTPRequest{
		Method:  method,
		URL:     url,
		Headers: b.headers(true, body != nil),
		JSON:    body,
	})
	if err != nil {
		return err
	}
	if out == nil {
		return nil
	}
	if err := decode(response, out); err != nil {
		return sdkerrors.Decode("ssi", operation, "decode SSI response", response.Body, err)
	}
	return nil
}

func (b *Broker) placeOrderBody(request PlaceOrderRequest) map[string]any {
	order := request.PlaceOrderRequest
	return map[string]any{
		"account":      order.AccountID,
		"instrumentID": order.Symbol,
		"marketID":     firstString(request.MarketID, b.config.MarketID),
		"buySell":      ssiSide(order.Side),
		"orderType":    ssiOrderType(order.Type),
		"price":        numberValue(order.Price),
		"quantity":     order.Quantity.IntPart(),
		"requestID":    firstString(request.RequestID, b.config.RequestID()),
		"channelID":    firstString(request.ChannelID, b.config.ChannelID),
		"code":         request.Code,
		"deviceId":     firstString(request.DeviceID, b.config.DeviceID),
		"userAgent":    firstString(request.UserAgent, b.config.UserAgent),
	}
}

func (b *Broker) cancelOrderBody(request CancelOrderRequest) map[string]any {
	return map[string]any{
		"account":      request.AccountID,
		"orderID":      request.OrderID,
		"marketID":     firstString(request.MarketID, b.config.MarketID),
		"instrumentID": request.Symbol,
		"buySell":      ssiSide(request.Side),
		"requestID":    firstString(request.RequestID, b.config.RequestID()),
		"code":         request.Code,
		"deviceId":     firstString(request.DeviceID, b.config.DeviceID),
		"userAgent":    firstString(request.UserAgent, b.config.UserAgent),
	}
}

func (b *Broker) modifyOrderBody(request ModifyOrderRequest) map[string]any {
	return map[string]any{
		"account":      request.AccountID,
		"orderID":      request.OrderID,
		"marketID":     firstString(request.MarketID, b.config.MarketID),
		"instrumentID": request.Symbol,
		"buySell":      ssiSide(request.Side),
		"price":        numberValue(&request.Price),
		"quantity":     request.Quantity,
		"requestID":    firstString(request.RequestID, b.config.RequestID()),
		"code":         request.Code,
		"deviceId":     firstString(request.DeviceID, b.config.DeviceID),
		"userAgent":    firstString(request.UserAgent, b.config.UserAgent),
	}
}

func firstString(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}
