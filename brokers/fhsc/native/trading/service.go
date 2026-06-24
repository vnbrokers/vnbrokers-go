package trading

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityGetAccountSummary    core.Capability = "native.trading.get_account_summary"
	CapabilityGetUserAssetsSummary core.Capability = "native.trading.get_user_assets_summary"
	CapabilityGetPnLToday          core.Capability = "native.trading.get_pnl_today"
	CapabilityGetPortfolio         core.Capability = "native.trading.get_portfolio"
	CapabilityGetAvailableTrade    core.Capability = "native.trading.get_available_trade"
	CapabilityGetOrderHistory      core.Capability = "native.trading.get_order_history"
	CapabilityGetOrderBookDetail   core.Capability = "native.trading.get_order_book_detail"
	CapabilityGetUserRights        core.Capability = "native.trading.get_user_rights"
	CapabilityGetMarketSession     core.Capability = "native.trading.get_market_session"
)

type RealtimeService interface{}

type Service interface {
	Realtime() RealtimeService
	GetAccountSummary(context.Context, dto.GetAccountSummaryRequest) (*dto.SummaryAccountResponse, error)
	GetUserAssetsSummary(context.Context, dto.GetUserAssetsSummaryRequest) (*dto.UserAssetsSummaryResponse, error)
	GetPnLToday(context.Context, dto.GetPnLTodayRequest) (*dto.PnLTodayResponse, error)
	GetPortfolio(context.Context, dto.GetPortfolioRequest) (*dto.GetPortfolioResponse, error)
	GetAvailableTrade(context.Context, dto.GetAvailableTradeRequest) (*dto.AvailableTradeResult, error)
	GetUserRights(context.Context, dto.GetUserRightsRequest) (*[]dto.UserRight, error)
	GetMarketSession(context.Context, dto.GetMarketSessionRequest) (*dto.ExchangeSessionInfo, error)
	GetOrderHistory(context.Context, dto.GetOrderHistoryRequest) (*dto.GetOrderHistoryResponse, error)
	GetOrderBookDetail(context.Context, dto.GetOrderBookDetailRequest) (*dto.GetOrderBookDetailResponse, error)
	GetOrderBook(context.Context, dto.GetOrderBookRequest) (*dto.GetOrderBookResponse, error)
	PlaceOrder(context.Context, dto.PlaceOrderRequest) (*dto.PlaceOrderResponse, error)
	CancelOrder(context.Context, dto.CancelOrderOperationRequest) (*dto.CancelOrderOperationResponse, error)
	UpdateOrder(context.Context, dto.UpdateOrderOperationRequest) (*dto.UpdateOrderOperationResponse, error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func(authenticated bool, hasBody bool) map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

type envelope[T any] struct {
	Status       *int64  `json:"status,omitempty"`
	ErrorCode    *string `json:"error_code,omitempty"`
	Message      *string `json:"message,omitempty"`
	PopupMessage *string `json:"popup_message,omitempty"`
	Title        *string `json:"title,omitempty"`
	Data         *T      `json:"data,omitempty"`
	Result       *T      `json:"result,omitempty"`
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
	endpoint := strings.TrimRight(s.dependencies.BaseURL, "/") + path
	if encoded := query.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method:  method,
		URL:     endpoint,
		Headers: s.dependencies.Headers(true, body != nil),
		JSON:    body,
	})
	if err != nil {
		return nil, err
	}

	var env envelope[T]
	if err := decode(response, &env); err != nil {
		return nil, sdkerrors.Decode("fhsc", string(capability), "decode FHSC trading response", responsePayload(response), err)
	}
	if code := stringValue(env.ErrorCode); code != "" && code != "0" {
		return nil, sdkerrors.BrokerRejected("fhsc", string(capability), code, firstNonEmpty(env.Message, env.PopupMessage, env.Title), responsePayload(response))
	}
	if env.Result != nil {
		return env.Result, nil
	}
	if env.Data != nil {
		return env.Data, nil
	}
	result := new(T)
	return result, nil
}

func escaped(value string) string { return url.PathEscape(value) }

func set(query url.Values, key, value string) { query.Set(key, value) }

func setOptional(query url.Values, key, value string) {
	if value != "" {
		query.Set(key, value)
	}
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func firstNonEmpty(values ...*string) string {
	for _, value := range values {
		if value != nil && *value != "" {
			return *value
		}
	}
	return ""
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

func responsePayload(response transport.HTTPResponse) any {
	if len(response.Raw) > 0 {
		return response.Raw
	}
	return response.Body
}
