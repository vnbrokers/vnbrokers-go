package trading

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityCashInAdvanceAmount     core.Capability = "native.trading.cash_in_advance_amount"
	CapabilityUnsettleSoldTransaction core.Capability = "native.trading.unsettle_sold_transaction"
	CapabilityTransferHistories       core.Capability = "native.trading.transfer_histories"
	CapabilityCashInAdvanceHistories  core.Capability = "native.trading.cash_in_advance_histories"
	CapabilityEstCashInAdvanceFee     core.Capability = "native.trading.est_cash_in_advance_fee"
	CapabilityVSDCashDW               core.Capability = "native.trading.vsd_cash_dw"
	CapabilityTransferInternal        core.Capability = "native.trading.transfer_internal"
	CapabilityCreateCashInAdvance     core.Capability = "native.trading.create_cash_in_advance"
	CapabilityCashAcctBal             core.Capability = "native.trading.cash_acct_bal"
	CapabilityDerivAcctBal            core.Capability = "native.trading.deriv_acct_bal"
	CapabilityMaxBuyQty               core.Capability = "native.trading.max_buy_qty"
	CapabilityMaxSellQty              core.Capability = "native.trading.max_sell_qty"
	CapabilityOrderBook               core.Capability = "native.trading.order_book"
	CapabilityOrderHistory            core.Capability = "native.trading.order_history"
	CapabilityStockPosition           core.Capability = "native.trading.stock_position"
	CapabilityDerivPosition           core.Capability = "native.trading.deriv_position"
	CapabilityNewOrder                core.Capability = "native.trading.new_order"
	CapabilityCancelOrder             core.Capability = "native.trading.cancel_order"
	CapabilityModifyOrder             core.Capability = "native.trading.modify_order"
	CapabilityDerNewOrder             core.Capability = "native.trading.der_new_order"
	CapabilityDerCancelOrder          core.Capability = "native.trading.der_cancel_order"
	CapabilityDerModifyOrder          core.Capability = "native.trading.der_modify_order"
	CapabilityAuditOrderBook          core.Capability = "native.trading.audit_order_book"
	CapabilityPpmrAccount             core.Capability = "native.trading.ppmr_account"
	CapabilityRateLimit               core.Capability = "native.trading.rate_limit"
	CapabilityTransferable            core.Capability = "native.trading.transferable"
	CapabilityStockTransferHistories  core.Capability = "native.trading.stock_transfer_histories"
	CapabilityStockTransfer           core.Capability = "native.trading.stock_transfer"
	CapabilityDividend                core.Capability = "native.trading.dividend"
	CapabilityExercisableQuantity     core.Capability = "native.trading.exercisable_quantity"
	CapabilityRightsHistories         core.Capability = "native.trading.rights_histories"
	CapabilityCreateRight             core.Capability = "native.trading.create_right"
	CapabilityFcoNewOrder             core.Capability = "native.trading.fco_new_order"
	CapabilityFcoCancelOrder          core.Capability = "native.trading.fco_cancel_order"
	CapabilityFcoOrderBook            core.Capability = "native.trading.fco_order_book"
	CapabilityFcoStatusHistory        core.Capability = "native.trading.fco_status_history"
	CapabilityFcoList                 core.Capability = "native.trading.fco_list"
)

type Service interface {
	Realtime() RealtimeService

	// Cash
	CashInAdvanceAmount(context.Context, string) (*dto.CashInAdvanceAmountResponse, error)
	UnsettleSoldTransaction(context.Context, string, string) (*dto.UnsettleSoldTransactionResponse, error)
	TransferHistories(context.Context, string, string, string) (*dto.CashTransferHistoriesResponse, error)
	CashInAdvanceHistories(context.Context, string, string, string) (*dto.CashInAdvanceHistoriesResponse, error)
	EstCashInAdvanceFee(context.Context, string, string, string) (*dto.EstCashInAdvanceFeeResponse, error)
	VSDCashDW(context.Context, string, string, string, string, string) (*dto.VSDCashDWResponse, error)
	TransferInternal(context.Context, string, string, string, string, string) (*dto.TransferInternalResponse, error)
	CreateCashInAdvance(context.Context, string, string, string, string) (*dto.CreateCashInAdvanceResponse, error)

	// Core Trading - Accounts
	CashAcctBal(context.Context, string) (*dto.CashAccountBalanceResponse, error)
	DerivAcctBal(context.Context, string) (*dto.DerivativeAccountBalanceResponse, error)
	MaxBuyQty(context.Context, string, string, string) (*dto.MaxBuyQtyResponse, error)
	MaxSellQty(context.Context, string, string, string) (*dto.MaxSellQtyResponse, error)
	PpmrAccount(context.Context, string) (*dto.PpmmrAccountResponse, error)
	RateLimit(context.Context) (*dto.RateLimitResponse, error)

	// Core Trading - Orders
	OrderBook(context.Context, string) (*dto.OrderBookResponse, error)
	OrderHistory(context.Context, string, string, string, int) (*dto.OrderHistoryResponse, error)
	NewOrder(context.Context, map[string]any) (*dto.NewOrderResponse, error)
	CancelOrder(context.Context, map[string]any) (*dto.CancelOrderResponse, error)
	ModifyOrder(context.Context, map[string]any) (*dto.ModifyOrderResponse, error)
	DerNewOrder(context.Context, map[string]any) (*dto.DerNewOrderResponse, error)
	DerCancelOrder(context.Context, map[string]any) (*dto.DerCancelOrderResponse, error)
	DerModifyOrder(context.Context, map[string]any) (*dto.DerModifyOrderResponse, error)
	AuditOrderBook(context.Context, string) (*dto.AuditOrderBookResponse, error)

	// Core Trading - Positions
	StockPosition(context.Context, string) (*dto.StockPositionResponse, error)
	DerivPosition(context.Context, string, bool) (*dto.DerivativePositionResponse, error)

	// Stock Transfer
	Transferable(context.Context, string) (*dto.TransferableResponse, error)
	StockTransferHistories(context.Context, string, string, string) (*dto.StockTransferHistoriesResponse, error)
	StockTransfer(context.Context, map[string]any) (*dto.StockTransferResponse, error)

	// Rights
	Dividend(context.Context, string) (*dto.DividendResponse, error)
	ExercisableQuantity(context.Context, string) (*dto.ExercisableQuantityResponse, error)
	RightsHistories(context.Context, string, string, string) (*dto.RightsHistoriesResponse, error)
	CreateRight(context.Context, map[string]any) (*dto.RightsCreateResponse, error)

	// Conditional Orders (FCO)
	FcoNewOrder(context.Context, map[string]any) (*dto.FCONewOrderResponse, error)
	FcoCancelOrder(context.Context, map[string]any) (*dto.FCOCancelOrderResponse, error)
	FcoOrderBook(context.Context, string, int, int) (*dto.FCOOrderBookResponse, error)
	FcoStatusHistory(context.Context, string, int, int) (*dto.FCOStatusHistoryResponse, error)
	FcoList(context.Context, url.Values) (*dto.FCOListResponse, error)
}

type RealtimeService interface {
	SubscribeOrderEvents(context.Context) (realtime.Subscription[dto.OrderEvent], error)
	SubscribeOrderErrors(context.Context) (realtime.Subscription[dto.OrderError], error)
	SubscribeOrderMatchEvents(context.Context) (realtime.Subscription[dto.OrderMatchEvent], error)
	SubscribeClientPortfolioEvents(context.Context) (realtime.Subscription[dto.ClientPortfolioEvent], error)
	SubscribeFCOEvents(context.Context) (realtime.Subscription[dto.FCOEvent], error)
}

type Dependencies struct {
	BaseURL           string
	TradingToken      func() string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

func NewService(dependencies Dependencies, realtime RealtimeService) Service {
	return &service{dependencies: dependencies, realtime: realtime}
}

func (s *service) Realtime() RealtimeService {
	return s.realtime
}

func get[T any](ctx context.Context, s *service, capability core.Capability, path string, params url.Values) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}

	endpoint := strings.TrimRight(s.dependencies.BaseURL, "/") + path
	if encoded := params.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	headers := map[string]string{"Accept": "application/json"}
	if token := s.dependencies.TradingToken(); token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method:  "GET",
		URL:     endpoint,
		Headers: headers,
	})
	if err != nil {
		return nil, err
	}

	result := new(T)
	if err := decode(response, result); err != nil {
		return nil, sdkerrors.Decode("ssi", string(capability), "decode SSI native trading response", response.Body, err)
	}
	return result, nil
}

func post[T any](ctx context.Context, s *service, capability core.Capability, path string, body any) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}

	endpoint := strings.TrimRight(s.dependencies.BaseURL, "/") + path
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	if token := s.dependencies.TradingToken(); token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method:  "POST",
		URL:     endpoint,
		Headers: headers,
		JSON:    body,
	})
	if err != nil {
		return nil, err
	}

	result := new(T)
	if err := decode(response, result); err != nil {
		return nil, sdkerrors.Decode("ssi", string(capability), "decode SSI native trading response", response.Body, err)
	}
	return result, nil
}

func decode(response transport.HTTPResponse, out any) error {
	if len(response.Raw) > 0 {
		return json.Unmarshal(response.Raw, out)
	}
	payload, err := json.Marshal(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, out)
}

func setOptionalString(params url.Values, key string, value string) {
	if value != "" {
		params.Set(key, value)
	}
}

func setOptionalInt(params url.Values, key string, value int) {
	if value > 0 {
		params.Set(key, strconv.Itoa(value))
	}
}
