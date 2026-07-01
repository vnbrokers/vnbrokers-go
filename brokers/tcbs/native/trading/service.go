package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityGetSubAccountInformation             core.Capability = "native.trading.get_sub_account_information"
	CapabilityTransferBetweenSubaccounts           core.Capability = "native.trading.transfer_between_subaccounts"
	CapabilityWithdrawDerivativeMargin             core.Capability = "native.trading.withdraw_derivative_margin"
	CapabilityDepositDerivativeMargin              core.Capability = "native.trading.deposit_derivative_margin"
	CapabilityPlaceStockOrder                      core.Capability = "native.trading.place_stock_order"
	CapabilityUpdateStockOrder                     core.Capability = "native.trading.update_stock_order"
	CapabilityCancelStockOrder                     core.Capability = "native.trading.cancel_stock_order"
	CapabilityGetStockOrders                       core.Capability = "native.trading.get_stock_orders"
	CapabilityGetStockOrder                        core.Capability = "native.trading.get_stock_order"
	CapabilityGetStockMatchingDetails              core.Capability = "native.trading.get_stock_matching_details"
	CapabilityGetStockPurchasingPower              core.Capability = "native.trading.get_stock_purchasing_power"
	CapabilityGetStockPurchasingPowerBySymbol      core.Capability = "native.trading.get_stock_purchasing_power_by_symbol"
	CapabilityGetStockPurchasingPowerBySymbolPrice core.Capability = "native.trading.get_stock_purchasing_power_by_symbol_price"
	CapabilityGetMarginQuota                       core.Capability = "native.trading.get_margin_quota"
	CapabilityGetMarginAccountInformation          core.Capability = "native.trading.get_margin_account_information"
	CapabilityGetSupplementaryLoanPackages         core.Capability = "native.trading.get_supplementary_loan_packages"
	CapabilityGetLoans                             core.Capability = "native.trading.get_loans"
	CapabilityGetStockAssets                       core.Capability = "native.trading.get_stock_assets"
	CapabilityGetCashInvestments                   core.Capability = "native.trading.get_cash_investments"
	CapabilityGetCashStatements                    core.Capability = "native.trading.get_cash_statements"
	CapabilityGetMarginInformation                 core.Capability = "native.trading.get_margin_information"
	CapabilityGetDerivativeCash                    core.Capability = "native.trading.get_derivative_cash"
	CapabilityGetClosedDerivativePositions         core.Capability = "native.trading.get_closed_derivative_positions"
	CapabilityGetOpenDerivativePositions           core.Capability = "native.trading.get_open_derivative_positions"
	CapabilityGetDerivativeOrders                  core.Capability = "native.trading.get_derivative_orders"
	CapabilityGetDerivativeConditionalOrders       core.Capability = "native.trading.get_derivative_conditional_orders"
	CapabilityPlaceDerivativeOrder                 core.Capability = "native.trading.place_derivative_order"
	CapabilityPlaceDerivativeConditionalOrder      core.Capability = "native.trading.place_derivative_conditional_order"
	CapabilityUpdateDerivativeOrder                core.Capability = "native.trading.update_derivative_order"
	CapabilityUpdateDerivativeConditionalOrder     core.Capability = "native.trading.update_derivative_conditional_order"
	CapabilityCancelDerivativeOrder                core.Capability = "native.trading.cancel_derivative_order"
	CapabilityCancelDerivativeConditionalOrder     core.Capability = "native.trading.cancel_derivative_conditional_order"
)

type Service interface {
	Realtime() RealtimeService
	GetSubAccountInformation(context.Context, dto.GetSubAccountInformationRequest) (*dto.GetSubAccountInformationResponse, error)
	TransferBetweenSubaccounts(context.Context, dto.TransferBetweenSubaccountsRequest) (*dto.TransferBetweenSubaccountsResponse, error)
	WithdrawDerivativeMargin(context.Context, dto.WithdrawDerivativeMarginRequest) (*dto.WithdrawDerivativeMarginResponse, error)
	DepositDerivativeMargin(context.Context, dto.DepositDerivativeMarginRequest) (*dto.DepositDerivativeMarginResponse, error)
	PlaceStockOrder(context.Context, dto.PlaceStockOrderRequest) (*dto.PlaceStockOrderResponse, error)
	UpdateStockOrder(context.Context, dto.UpdateStockOrderRequest) (*dto.UpdateStockOrderResponse, error)
	CancelStockOrder(context.Context, dto.CancelStockOrderRequest) (*dto.CancelStockOrderResponse, error)
	GetStockOrders(context.Context, dto.GetStockOrdersRequest) (*dto.GetStockOrdersResponse, error)
	GetStockOrder(context.Context, dto.GetStockOrderRequest) (*dto.GetStockOrderResponse, error)
	GetStockMatchingDetails(context.Context, dto.GetStockMatchingDetailsRequest) (*dto.GetStockMatchingDetailsResponse, error)
	GetStockPurchasingPower(context.Context, dto.GetStockPurchasingPowerRequest) (*dto.GetStockPurchasingPowerResponse, error)
	GetStockPurchasingPowerBySymbol(context.Context, dto.GetStockPurchasingPowerBySymbolRequest) (*dto.GetStockPurchasingPowerBySymbolResponse, error)
	GetStockPurchasingPowerBySymbolPrice(context.Context, dto.GetStockPurchasingPowerBySymbolPriceRequest) (*dto.GetStockPurchasingPowerBySymbolPriceResponse, error)
	GetMarginQuota(context.Context, dto.GetMarginQuotaRequest) (*dto.GetMarginQuotaResponse, error)
	GetMarginAccountInformation(context.Context, dto.GetMarginAccountInformationRequest) (*dto.GetMarginAccountInformationResponse, error)
	GetSupplementaryLoanPackages(context.Context, dto.GetSupplementaryLoanPackagesRequest) (*dto.GetSupplementaryLoanPackagesResponse, error)
	GetLoans(context.Context, dto.GetLoansRequest) (*dto.GetLoansResponse, error)
	GetStockAssets(context.Context, dto.GetStockAssetsRequest) (*dto.GetStockAssetsResponse, error)
	GetCashInvestments(context.Context, dto.GetCashInvestmentsRequest) (*dto.GetCashInvestmentsResponse, error)
	GetCashStatements(context.Context, dto.GetCashStatementsRequest) (*dto.GetCashStatementsResponse, error)
	GetMarginInformation(context.Context, dto.GetMarginInformationRequest) (*dto.GetMarginInformationResponse, error)
	GetDerivativeCash(context.Context, dto.GetDerivativeCashRequest) (*dto.GetDerivativeCashResponse, error)
	GetClosedDerivativePositions(context.Context, dto.GetClosedDerivativePositionsRequest) (*dto.GetClosedDerivativePositionsResponse, error)
	GetOpenDerivativePositions(context.Context, dto.GetOpenDerivativePositionsRequest) (*dto.GetOpenDerivativePositionsResponse, error)
	GetDerivativeOrders(context.Context, dto.GetDerivativeOrdersRequest) (*dto.GetDerivativeOrdersResponse, error)
	GetDerivativeConditionalOrders(context.Context, dto.GetDerivativeConditionalOrdersRequest) (*dto.GetDerivativeConditionalOrdersResponse, error)
	PlaceDerivativeOrder(context.Context, dto.PlaceDerivativeOrderRequest) (*dto.PlaceDerivativeOrderResponse, error)
	PlaceDerivativeConditionalOrder(context.Context, dto.PlaceDerivativeConditionalOrderRequest) (*dto.PlaceDerivativeConditionalOrderResponse, error)
	UpdateDerivativeOrder(context.Context, dto.UpdateDerivativeOrderRequest) (*dto.UpdateDerivativeOrderResponse, error)
	UpdateDerivativeConditionalOrder(context.Context, dto.UpdateDerivativeConditionalOrderRequest) (*dto.UpdateDerivativeConditionalOrderResponse, error)
	CancelDerivativeOrder(context.Context, dto.CancelDerivativeOrderRequest) (*dto.CancelDerivativeOrderResponse, error)
	CancelDerivativeConditionalOrder(context.Context, dto.CancelDerivativeConditionalOrderRequest) (*dto.CancelDerivativeConditionalOrderResponse, error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func(bool, bool) map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

func NewService(dependencies Dependencies, realtimeServices ...RealtimeService) Service {
	var realtimeService RealtimeService
	if len(realtimeServices) > 0 {
		realtimeService = realtimeServices[0]
	}
	return &service{dependencies: dependencies, realtime: realtimeService}
}

func (s *service) Realtime() RealtimeService { return s.realtime }

func do[T any](s *service, ctx context.Context, capability core.Capability, method, path string, query url.Values, body any) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method: method, URL: httpx.URL(s.dependencies.BaseURL, path, query), Headers: s.dependencies.Headers(true, body != nil), JSON: body,
	})
	if err != nil {
		return nil, err
	}
	result, err := httpx.DecodeResponse[T]("tcbs", string(capability), "decode TCBS native trading response", response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func escaped(value string) string { return url.PathEscape(value) }

func set(query url.Values, key, value string) { query.Set(key, value) }

func setOptional(query url.Values, key, value string) {
	if value != "" {
		set(query, key, value)
	}
}
