package dnse

import (
	"context"
	"net/url"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

var (
	_ sdktrading.AccountsService  = (*TradingAccountsService)(nil)
	_ sdktrading.OrdersService    = (*TradingOrdersService)(nil)
	_ sdktrading.PositionsService = (*TradingPositionsService)(nil)
)

type TradingAccountsService struct {
	broker *Broker
}

func (s *TradingAccountsService) List(ctx context.Context) ([]domain.Account, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "trading.accounts.list", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts"),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapAccounts(expectObject(response.Body)), nil
}

func (s *TradingAccountsService) Balance(ctx context.Context, accountID string) (domain.Balance, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsBalance); err != nil {
		return domain.Balance{}, err
	}
	response, err := s.broker.send(ctx, "trading.accounts.balance", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/balances"),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.Balance{}, err
	}
	return MapBalance(accountID, expectObject(response.Body)), nil
}

func (s *TradingAccountsService) Orders(ctx context.Context, accountID string) ([]domain.Order, error) {
	return s.OrdersWithRequest(ctx, sdktrading.ListOrdersRequest{AccountID: accountID})
}

func (s *TradingAccountsService) OrdersWithRequest(
	ctx context.Context,
	request sdktrading.ListOrdersRequest,
) ([]domain.Order, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersList); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "trading.accounts.orders", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(request.AccountID) + "/orders?" + s.broker.marketOrderQueryFrom(request.MarketType, request.OrderCategory)),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapOrders(expectObject(response.Body)), nil
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
	params.Set("marketType", s.broker.marketType(request.MarketType))
	params.Set("from", request.FromDate)
	params.Set("to", request.ToDate)
	params.Set("pageIndex", decimal.NewFromInt(int64(request.PageIndex)).String())
	if request.PageSize > 0 {
		params.Set("pageSize", decimal.NewFromInt(int64(request.PageSize)).String())
	}
	response, err := s.broker.send(ctx, "trading.accounts.order_history", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(request.AccountID) + "/orders/history?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapOrderHistory(expectObject(response.Body)), nil
}

func (s *TradingAccountsService) Executions(
	ctx context.Context,
	accountID string,
	orderID string,
) (domain.RawPayload, error) {
	return s.ExecutionsWithRequest(ctx, sdktrading.ExecutionsRequest{
		AccountID: accountID,
		OrderID:   orderID,
	})
}

func (s *TradingAccountsService) ExecutionsWithRequest(
	ctx context.Context,
	request sdktrading.ExecutionsRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersExecutions); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/accounts/" + url.PathEscape(request.AccountID) + "/executions/" + url.PathEscape(request.OrderID)
	response, err := s.broker.send(ctx, "trading.accounts.executions", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(path + "?" + s.broker.marketOrderQueryFrom(request.MarketType, request.OrderCategory)),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *TradingAccountsService) PPSE(
	ctx context.Context,
	accountID string,
	symbol string,
	price decimal.Decimal,
	loanPackageID *int,
) (domain.RawPayload, error) {
	return s.PPSEWithRequest(ctx, sdktrading.BuyingPowerRequest{
		AccountID:     accountID,
		Symbol:        symbol,
		Price:         price,
		LoanPackageID: loanPackageID,
	})
}

func (s *TradingAccountsService) PPSEWithRequest(
	ctx context.Context,
	request sdktrading.BuyingPowerRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.marketType(request.MarketType))
	params.Set("symbol", request.Symbol)
	if request.LoanPackageID != nil {
		params.Set("loanPackageId", decimal.NewFromInt(int64(*request.LoanPackageID)).String())
	} else if s.broker.config.LoanPackageID != nil {
		params.Set("loanPackageId", decimal.NewFromInt(int64(*s.broker.config.LoanPackageID)).String())
	} else {
		params.Set("loanPackageId", "0")
	}
	params.Set("price", request.Price.String())
	response, err := s.broker.send(ctx, "trading.accounts.ppse", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(request.AccountID) + "/ppse?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *TradingAccountsService) LoanPackages(
	ctx context.Context,
	accountID string,
	symbol string,
) (domain.RawPayload, error) {
	return s.LoanPackagesWithRequest(ctx, sdktrading.LoanPackagesRequest{
		AccountID: accountID,
		Symbol:    symbol,
	})
}

func (s *TradingAccountsService) LoanPackagesWithRequest(
	ctx context.Context,
	request sdktrading.LoanPackagesRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.marketType(request.MarketType))
	params.Set("symbol", request.Symbol)
	response, err := s.broker.send(ctx, "trading.accounts.loan_packages", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(request.AccountID) + "/loan-packages?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

type TradingPositionsService struct {
	broker *Broker
}

func (s *TradingPositionsService) List(ctx context.Context, accountID string) ([]domain.Position, error) {
	return s.ListWithRequest(ctx, sdktrading.ListPositionsRequest{AccountID: accountID})
}

func (s *TradingPositionsService) ListWithRequest(
	ctx context.Context,
	request sdktrading.ListPositionsRequest,
) ([]domain.Position, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return nil, err
	}
	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = s.broker.config.PositionsPageSize
	}
	params := url.Values{}
	params.Set("marketType", s.broker.marketType(request.MarketType))
	params.Set("pageSize", decimal.NewFromInt(int64(pageSize)).String())
	response, err := s.broker.send(ctx, "trading.positions.list", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(request.AccountID) + "/positions?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapPositions(expectObject(response.Body)), nil
}

func (s *TradingPositionsService) Get(ctx context.Context, positionID string) (domain.Position, error) {
	return s.GetWithRequest(ctx, sdktrading.GetPositionRequest{PositionID: positionID})
}

func (s *TradingPositionsService) GetWithRequest(
	ctx context.Context,
	request sdktrading.GetPositionRequest,
) (domain.Position, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsGet); err != nil {
		return domain.Position{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.marketType(request.MarketType))
	response, err := s.broker.send(ctx, "trading.positions.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/positions/" + url.PathEscape(request.PositionID) + "?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.Position{}, err
	}
	return MapPosition(expectObject(response.Body)), nil
}

func (s *TradingPositionsService) Close(ctx context.Context, positionID string) (domain.RawPayload, error) {
	return s.CloseWithRequest(ctx, sdktrading.ClosePositionRequest{PositionID: positionID})
}

func (s *TradingPositionsService) CloseWithRequest(
	ctx context.Context,
	request sdktrading.ClosePositionRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsClose); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.marketType(request.MarketType))
	response, err := s.broker.send(ctx, "trading.positions.close", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/accounts/positions/" + url.PathEscape(request.PositionID) + "/close?" + params.Encode()),
		Headers: s.broker.tradingHeaders(false),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

type TradingOrdersService struct {
	broker *Broker
}

type PlaceOrderRequest struct {
	domain.PlaceOrderRequest
	MarketType    sdktrading.MarketType
	OrderCategory sdktrading.OrderCategory
	LoanPackageID *int
}

func (s *TradingOrdersService) Place(
	ctx context.Context,
	request domain.PlaceOrderRequest,
) (domain.PlaceOrderResponse, error) {
	return s.PlaceWithRequest(ctx, PlaceOrderRequest{PlaceOrderRequest: request})
}

func (s *TradingOrdersService) PlaceWithRequest(
	ctx context.Context,
	request PlaceOrderRequest,
) (domain.PlaceOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.PlaceOrderResponse{}, err
	}
	order := request.PlaceOrderRequest
	body := map[string]any{
		"accountNo": order.AccountID,
		"orderType": dnseOrderType(order.Type),
		"price":     numberValue(order.Price),
		"quantity":  order.Quantity.IntPart(),
		"side":      dnseSide(order.Side),
		"symbol":    order.Symbol,
	}
	loanPackageID := request.LoanPackageID
	if loanPackageID == nil {
		loanPackageID = s.broker.config.LoanPackageID
	}
	if loanPackageID != nil {
		body["loanPackageId"] = *loanPackageID
	}
	response, err := s.broker.send(ctx, "trading.orders.place", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/accounts/orders?" + s.broker.marketOrderQueryFrom(request.MarketType, request.OrderCategory)),
		Headers: s.broker.tradingHeaders(true),
		JSON:    body,
	})
	if err != nil {
		return domain.PlaceOrderResponse{}, err
	}
	return MapPlaceOrderResponse(expectObject(response.Body)), nil
}

func (s *TradingOrdersService) Cancel(ctx context.Context, accountID string, orderID string) error {
	_, err := s.CancelWithRequest(ctx, sdktrading.CancelOrderRequest{
		AccountID: accountID,
		OrderID:   orderID,
	})
	return err
}

func (s *TradingOrdersService) CancelWithRequest(
	ctx context.Context,
	request sdktrading.CancelOrderRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersCancel); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "trading.orders.cancel", transport.HTTPRequest{
		Method:  "DELETE",
		URL:     s.broker.url(s.broker.orderPathFrom(request.AccountID, request.OrderID, request.MarketType, request.OrderCategory)),
		Headers: s.broker.tradingHeaders(false),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *TradingOrdersService) Get(
	ctx context.Context,
	accountID string,
	orderID string,
) (domain.Order, error) {
	return s.GetWithRequest(ctx, sdktrading.GetOrderRequest{
		AccountID: accountID,
		OrderID:   orderID,
	})
}

func (s *TradingOrdersService) GetWithRequest(
	ctx context.Context,
	request sdktrading.GetOrderRequest,
) (domain.Order, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersGet); err != nil {
		return domain.Order{}, err
	}
	response, err := s.broker.send(ctx, "trading.orders.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(s.broker.orderPathFrom(request.AccountID, request.OrderID, request.MarketType, request.OrderCategory)),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.Order{}, err
	}
	return MapOrder(expectObject(response.Body)), nil
}

func (s *TradingOrdersService) Update(
	ctx context.Context,
	accountID string,
	orderID string,
	price decimal.Decimal,
	quantity int,
) (domain.RawPayload, error) {
	return s.Replace(ctx, sdktrading.ReplaceOrderRequest{
		AccountID: accountID,
		OrderID:   orderID,
		Price:     price,
		Quantity:  quantity,
	})
}

func (s *TradingOrdersService) Replace(
	ctx context.Context,
	request sdktrading.ReplaceOrderRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersReplace); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "trading.orders.update", transport.HTTPRequest{
		Method:  "PUT",
		URL:     s.broker.url(s.broker.orderPathFrom(request.AccountID, request.OrderID, request.MarketType, request.OrderCategory)),
		Headers: s.broker.tradingHeaders(true),
		JSON: map[string]any{
			"price":    numberValue(&request.Price),
			"quantity": request.Quantity,
		},
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (b *Broker) marketOrderQuery() string {
	return b.marketOrderQueryFrom("", "")
}

func (b *Broker) marketOrderQueryFrom(
	marketType sdktrading.MarketType,
	orderCategory sdktrading.OrderCategory,
) string {
	params := url.Values{}
	params.Set("marketType", b.marketType(marketType))
	params.Set("orderCategory", b.orderCategory(orderCategory))
	return params.Encode()
}

func (b *Broker) orderPath(accountID string, orderID string) string {
	return b.orderPathFrom(accountID, orderID, "", "")
}

func (b *Broker) orderPathFrom(
	accountID string,
	orderID string,
	marketType sdktrading.MarketType,
	orderCategory sdktrading.OrderCategory,
) string {
	return "/accounts/" + url.PathEscape(accountID) + "/orders/" + url.PathEscape(orderID) + "?" + b.marketOrderQueryFrom(marketType, orderCategory)
}

func (b *Broker) marketType(marketType sdktrading.MarketType) string {
	if marketType != "" {
		return string(marketType)
	}
	return b.config.MarketType
}

func (b *Broker) orderCategory(orderCategory sdktrading.OrderCategory) string {
	if orderCategory != "" {
		return string(orderCategory)
	}
	return b.config.OrderCategory
}
