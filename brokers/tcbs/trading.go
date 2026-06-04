package tcbs

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type TradingService struct {
	accounts *TradingAccountsService
	orders   *TradingOrdersService
	realtime *TradingRealtimeService
}

func (s *TradingService) Accounts() *TradingAccountsService {
	return s.accounts
}

func (s *TradingService) Orders() *TradingOrdersService {
	return s.orders
}

func (s *TradingService) Realtime() *TradingRealtimeService {
	return s.realtime
}

type TradingOrdersService struct {
	broker *Broker
}

func (s *TradingOrdersService) Place(
	ctx context.Context,
	accountNo string,
	request PlaceOrderRequest,
) (PlaceOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersPlace); err != nil {
		return PlaceOrderResponse{}, err
	}
	var response PlaceOrderResponse
	err := s.broker.sendAndDecode(ctx, "trading.orders.place", "POST", "/akhlys/v1/accounts/"+url.PathEscape(accountNo)+"/orders", request, &response)
	return response, err
}

func (s *TradingOrdersService) Update(
	ctx context.Context,
	accountNo string,
	orderID string,
	request UpdateOrderRequest,
) (UpdateOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersReplace); err != nil {
		return UpdateOrderResponse{}, err
	}
	var response UpdateOrderResponse
	path := "/akhlys/v1/accounts/" + url.PathEscape(accountNo) + "/orders/" + url.PathEscape(orderID)
	err := s.broker.sendAndDecode(ctx, "trading.orders.update", "PUT", path, request, &response)
	return response, err
}

func (s *TradingOrdersService) Cancel(
	ctx context.Context,
	accountNo string,
	request CancelOrderRequest,
) (CancelOrderResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersCancel); err != nil {
		return CancelOrderResponse{}, err
	}
	var response CancelOrderResponse
	err := s.broker.sendAndDecode(ctx, "trading.orders.cancel", "PUT", "/akhlys/v1/accounts/"+url.PathEscape(accountNo)+"/cancel-orders", request, &response)
	return response, err
}

type TradingAccountsService struct {
	broker *Broker
}

func (s *TradingAccountsService) Orders(ctx context.Context, accountNo string) (OrderSearchResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersList); err != nil {
		return OrderSearchResponse{}, err
	}
	var response OrderSearchResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.orders", "GET", "/aion/v1/accounts/"+url.PathEscape(accountNo)+"/orders", nil, &response)
	return response, err
}

func (s *TradingAccountsService) Order(ctx context.Context, accountNo string, orderID string) (OrderSearchResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersGet); err != nil {
		return OrderSearchResponse{}, err
	}
	var response OrderSearchResponse
	path := "/aion/v1/accounts/" + url.PathEscape(accountNo) + "/orders/" + url.PathEscape(orderID)
	err := s.broker.sendAndDecode(ctx, "trading.accounts.order", "GET", path, nil, &response)
	return response, err
}

func (s *TradingAccountsService) MatchingDetails(ctx context.Context, accountNo string) (CommandMatchInformationResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingOrdersExecutions); err != nil {
		return CommandMatchInformationResponse{}, err
	}
	var response CommandMatchInformationResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.matching_details", "GET", "/aion/v1/accounts/"+url.PathEscape(accountNo)+"/matching-details", nil, &response)
	return response, err
}

func (s *TradingAccountsService) PurchasingPower(ctx context.Context, accountNo string) (PurchasingPowerResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return PurchasingPowerResponse{}, err
	}
	var response PurchasingPowerResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.ppse", "GET", "/aion/v1/accounts/"+url.PathEscape(accountNo)+"/ppse", nil, &response)
	return response, err
}

func (s *TradingAccountsService) PurchasingPowerBySymbol(ctx context.Context, accountNo string, symbol string) (PurchasingPowerResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return PurchasingPowerResponse{}, err
	}
	var response PurchasingPowerResponse
	path := "/aion/v1/accounts/" + url.PathEscape(accountNo) + "/ppse/" + url.PathEscape(symbol)
	err := s.broker.sendAndDecode(ctx, "trading.accounts.ppse_symbol", "GET", path, nil, &response)
	return response, err
}

func (s *TradingAccountsService) PurchasingPowerBySymbolPrice(ctx context.Context, accountNo string, symbol string, price string) (PurchasingPowerResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingBuyingPower); err != nil {
		return PurchasingPowerResponse{}, err
	}
	var response PurchasingPowerResponse
	path := "/aion/v1/accounts/" + url.PathEscape(accountNo) + "/ppse/" + url.PathEscape(symbol) + "/" + url.PathEscape(price)
	err := s.broker.sendAndDecode(ctx, "trading.accounts.ppse_symbol_price", "GET", path, nil, &response)
	return response, err
}

func (s *TradingAccountsService) MarginQuota(ctx context.Context, custodyID string) ([]MarginQuotaResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return nil, err
	}
	var response []MarginQuotaResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.margin_quota", "GET", "/aion/v1/customers/"+url.PathEscape(custodyID)+"/accounts", nil, &response)
	return response, err
}

func (s *TradingAccountsService) MarginAccountInfo(ctx context.Context, accountNo string) ([]MarginAccountInfoResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return nil, err
	}
	var response []MarginAccountInfoResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.margin_account_info", "GET", "/hydros/v1/account/"+url.PathEscape(accountNo)+"/risk", nil, &response)
	return response, err
}

func (s *TradingAccountsService) SupplementaryLoanPackages(ctx context.Context, accountNo string) (SupplementaryLoanPackageResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return SupplementaryLoanPackageResponse{}, err
	}
	var response SupplementaryLoanPackageResponse
	path := "/campaign-management/v1/margin/subscription/" + url.PathEscape(accountNo) + "/addons/detail"
	err := s.broker.sendAndDecode(ctx, "trading.accounts.supplementary_loan_packages", "GET", path, nil, &response)
	return response, err
}

func (s *TradingAccountsService) Loans(ctx context.Context, accountNo string) (LoanResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingLoanPackages); err != nil {
		return LoanResponse{}, err
	}
	var response LoanResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.loans", "GET", "/khaos/v1/loan/"+url.PathEscape(accountNo), nil, &response)
	return response, err
}

func (s *TradingAccountsService) StockAssets(ctx context.Context, accountNo string) (SeInfoDTO, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingPositionsList); err != nil {
		return SeInfoDTO{}, err
	}
	var response SeInfoDTO
	err := s.broker.sendAndDecode(ctx, "trading.accounts.stock_assets", "GET", "/aion/v1/accounts/"+url.PathEscape(accountNo)+"/se", nil, &response)
	return response, err
}

func (s *TradingAccountsService) CashBalance(ctx context.Context, accountNo string) (CashInvestmentResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAccountsBalance); err != nil {
		return CashInvestmentResponse{}, err
	}
	var response CashInvestmentResponse
	err := s.broker.sendAndDecode(ctx, "trading.accounts.cash_balance", "GET", "/aion/v1/accounts/"+url.PathEscape(accountNo)+"/cashInvestments", nil, &response)
	return response, err
}

func (b *Broker) sendAndDecode(
	ctx context.Context,
	operation string,
	method string,
	path string,
	body any,
	out any,
) error {
	response, err := b.send(ctx, operation, true, transport.HTTPRequest{
		Method:  method,
		URL:     b.url(path),
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
		return sdkerrors.Decode("tcbs", operation, "decode TCBS response", response.Body, err)
	}
	return nil
}
