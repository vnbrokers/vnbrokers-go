package dnse

import (
	"context"
	"net/url"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/transport"
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
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
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
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "trading.accounts.orders", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/orders?" + s.broker.marketOrderQuery()),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	params.Set("from", fromDate)
	params.Set("to", toDate)
	params.Set("pageIndex", decimal.NewFromInt(int64(pageIndex)).String())
	response, err := s.broker.send(ctx, "trading.accounts.order_history", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/orders/history?" + params.Encode()),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.RawPayload{}, err
	}
	path := "/accounts/" + url.PathEscape(accountID) + "/executions/" + url.PathEscape(orderID)
	response, err := s.broker.send(ctx, "trading.accounts.executions", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(path + "?" + s.broker.marketOrderQuery()),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	params.Set("symbol", symbol)
	if loanPackageID != nil {
		params.Set("loanPackageId", decimal.NewFromInt(int64(*loanPackageID)).String())
	} else if s.broker.config.LoanPackageID != nil {
		params.Set("loanPackageId", decimal.NewFromInt(int64(*s.broker.config.LoanPackageID)).String())
	} else {
		params.Set("loanPackageId", "0")
	}
	params.Set("price", price.String())
	response, err := s.broker.send(ctx, "trading.accounts.ppse", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/ppse?" + params.Encode()),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsList); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	params.Set("symbol", symbol)
	response, err := s.broker.send(ctx, "trading.accounts.loan_packages", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/loan-packages?" + params.Encode()),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	params.Set("pageSize", decimal.NewFromInt(int64(s.broker.config.PositionsPageSize)).String())
	response, err := s.broker.send(ctx, "trading.positions.list", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/" + url.PathEscape(accountID) + "/positions?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return nil, err
	}
	return MapPositions(expectObject(response.Body)), nil
}

func (s *TradingPositionsService) Get(ctx context.Context, positionID string) (domain.Position, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return domain.Position{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	response, err := s.broker.send(ctx, "trading.positions.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/accounts/positions/" + url.PathEscape(positionID) + "?" + params.Encode()),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.Position{}, err
	}
	return MapPosition(expectObject(response.Body)), nil
}

func (s *TradingPositionsService) Close(ctx context.Context, positionID string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.RawPayload{}, err
	}
	params := url.Values{}
	params.Set("marketType", s.broker.config.MarketType)
	response, err := s.broker.send(ctx, "trading.positions.close", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/accounts/positions/" + url.PathEscape(positionID) + "/close?" + params.Encode()),
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

func (s *TradingOrdersService) Place(
	ctx context.Context,
	request domain.PlaceOrderRequest,
) (domain.PlaceOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.PlaceOrderResponse{}, err
	}
	body := map[string]any{
		"accountNo": request.AccountID,
		"orderType": dnseOrderType(request.Type),
		"price":     numberValue(request.Price),
		"quantity":  request.Quantity.IntPart(),
		"side":      dnseSide(request.Side),
		"symbol":    request.Symbol,
	}
	if s.broker.config.LoanPackageID != nil {
		body["loanPackageId"] = *s.broker.config.LoanPackageID
	}
	response, err := s.broker.send(ctx, "trading.orders.place", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/accounts/orders?" + s.broker.marketOrderQuery()),
		Headers: s.broker.tradingHeaders(true),
		JSON:    body,
	})
	if err != nil {
		return domain.PlaceOrderResponse{}, err
	}
	return MapPlaceOrderResponse(expectObject(response.Body)), nil
}

func (s *TradingOrdersService) Cancel(ctx context.Context, accountID string, orderID string) error {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersCancel); err != nil {
		return err
	}
	_, err := s.broker.send(ctx, "trading.orders.cancel", transport.HTTPRequest{
		Method:  "DELETE",
		URL:     s.broker.url(s.broker.orderPath(accountID, orderID)),
		Headers: s.broker.tradingHeaders(false),
	})
	return err
}

func (s *TradingOrdersService) Get(
	ctx context.Context,
	accountID string,
	orderID string,
) (domain.Order, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.Order{}, err
	}
	response, err := s.broker.send(ctx, "trading.orders.get", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(s.broker.orderPath(accountID, orderID)),
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
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "trading.orders.update", transport.HTTPRequest{
		Method:  "PUT",
		URL:     s.broker.url(s.broker.orderPath(accountID, orderID)),
		Headers: s.broker.tradingHeaders(true),
		JSON: map[string]any{
			"price":    numberValue(&price),
			"quantity": quantity,
		},
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (b *Broker) marketOrderQuery() string {
	params := url.Values{}
	params.Set("marketType", b.config.MarketType)
	params.Set("orderCategory", b.config.OrderCategory)
	return params.Encode()
}

func (b *Broker) orderPath(accountID string, orderID string) string {
	return "/accounts/" + url.PathEscape(accountID) + "/orders/" + url.PathEscape(orderID) + "?" + b.marketOrderQuery()
}
