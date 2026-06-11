package marketdata

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityTradeHistory                core.Capability = "native.marketdata.trade_history"
	CapabilityInstrumentDetails           core.Capability = "native.marketdata.instrument_details"
	CapabilityInstruments                 core.Capability = "native.marketdata.instruments"
	CapabilityLatestQuotes                core.Capability = "native.marketdata.latest_quotes"
	CapabilityLatestTrades                core.Capability = "native.marketdata.latest_trades"
	CapabilityOHLC                        core.Capability = "native.marketdata.ohlc"
	CapabilityClosePrice                  core.Capability = "native.marketdata.close_price"
	CapabilityQuoteHistory                core.Capability = "native.marketdata.quote_history"
	CapabilitySecurityDefinition          core.Capability = "native.marketdata.security_definition"
	CapabilityWorkingDates                core.Capability = "native.marketdata.working_dates"
	CapabilityRealtimeExpectedPrices      core.Capability = "native.marketdata.realtime.expected_prices"
	CapabilityRealtimeForeign             core.Capability = "native.marketdata.realtime.foreign"
	CapabilityRealtimeMarketIndexes       core.Capability = "native.marketdata.realtime.market_indexes"
	CapabilityRealtimeOHLC                core.Capability = "native.marketdata.realtime.ohlc"
	CapabilityRealtimeClosedOHLC          core.Capability = "native.marketdata.realtime.closed_ohlc"
	CapabilityRealtimeQuotes              core.Capability = "native.marketdata.realtime.quotes"
	CapabilityRealtimeSecurityDefinitions core.Capability = "native.marketdata.realtime.security_definitions"
	CapabilityRealtimeTrades              core.Capability = "native.marketdata.realtime.trades"
	CapabilityRealtimeTradeExtras         core.Capability = "native.marketdata.realtime.trade_extras"
)

type Service interface {
	Realtime() RealtimeService
	GetTradeHistory(context.Context, dto.GetTradeHistoryRequest) (*dto.TradesResponse, error)
	GetInstrumentDetails(context.Context, dto.GetInstrumentDetailsRequest) (*dto.InstrumentsByFilterResponse, error)
	GetInstruments(context.Context, dto.GetInstrumentsRequest) (*dto.InstrumentsResponse, error)
	GetLatestQuotes(context.Context, dto.GetLatestQuotesRequest) (*dto.QuotesResponse, error)
	GetLatestTrades(context.Context, dto.GetLatestTradesRequest) (*dto.LatestTradesResponse, error)
	GetOHLC(context.Context, dto.GetOHLCRequest) (*dto.OhlcResponse, error)
	GetClosePrice(context.Context, dto.GetClosePriceRequest) (*dto.PriceSymbolCloseResponse, error)
	GetQuoteHistory(context.Context, dto.GetQuoteHistoryRequest) (*dto.QuotesResponse, error)
	GetSecurityDefinition(context.Context, dto.GetSecurityDefinitionRequest) (*dto.SecdefResponse, error)
	GetWorkingDates(context.Context, dto.GetWorkingDatesRequest) (*dto.WorkingDatesResponse, error)
}

type RealtimeService interface {
	SubscribeExpectedPrices(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.ExpectedPriceEvent], error)
	SubscribeForeign(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.ForeignEvent], error)
	SubscribeMarketIndexes(context.Context, dto.SubscribeMarketIndexRequest) (realtime.Subscription[dto.MarketIndexEvent], error)
	SubscribeOHLC(context.Context, dto.SubscribeOHLCRequest) (realtime.Subscription[dto.OHLCEvent], error)
	SubscribeClosedOHLC(context.Context, dto.SubscribeOHLCRequest) (realtime.Subscription[dto.OHLCEvent], error)
	SubscribeQuotes(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.QuoteEvent], error)
	SubscribeSecurityDefinitions(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.SecurityDefinitionEvent], error)
	SubscribeTrades(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.TradeEvent], error)
	SubscribeTradeExtras(context.Context, dto.SubscribeSymbolsRequest) (realtime.Subscription[dto.TradeExtraEvent], error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func() map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct {
	dependencies Dependencies
	realtime     RealtimeService
}

func NewService(dependencies Dependencies, realtimeServices ...RealtimeService) Service {
	var rt RealtimeService
	if len(realtimeServices) > 0 {
		rt = realtimeServices[0]
	}
	return &service{dependencies: dependencies, realtime: rt}
}
func (s *service) Realtime() RealtimeService { return s.realtime }

func (s *service) GetTradeHistory(ctx context.Context, r dto.GetTradeHistoryRequest) (*dto.TradesResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	setInt64(q, "from", r.From)
	setInt64(q, "to", r.To)
	setInt(q, "limit", r.Limit)
	return get[dto.TradesResponse](ctx, s, CapabilityTradeHistory, "/price/"+url.PathEscape(r.Symbol)+"/trades", q)
}
func (s *service) GetInstrumentDetails(ctx context.Context, r dto.GetInstrumentDetailsRequest) (*dto.InstrumentsByFilterResponse, error) {
	q := url.Values{}
	set(q, "symbol", r.Symbol)
	set(q, "marketId", r.MarketID)
	set(q, "securityGroupId", r.SecurityGroupID)
	set(q, "indexName", r.IndexName)
	setInt(q, "limit", r.Limit)
	setInt(q, "page", r.Page)
	return get[dto.InstrumentsByFilterResponse](ctx, s, CapabilityInstrumentDetails, "/instruments", q)
}
func (s *service) GetInstruments(ctx context.Context, r dto.GetInstrumentsRequest) (*dto.InstrumentsResponse, error) {
	q := url.Values{}
	set(q, "symbol", r.Symbol)
	setInt(q, "limit", r.Limit)
	setInt(q, "page", r.Page)
	return get[dto.InstrumentsResponse](ctx, s, CapabilityInstruments, "/instruments", q)
}
func (s *service) GetLatestQuotes(ctx context.Context, r dto.GetLatestQuotesRequest) (*dto.QuotesResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	return get[dto.QuotesResponse](ctx, s, CapabilityLatestQuotes, "/price/"+url.PathEscape(r.Symbol)+"/quotes/latest", q)
}
func (s *service) GetLatestTrades(ctx context.Context, r dto.GetLatestTradesRequest) (*dto.LatestTradesResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	return get[dto.LatestTradesResponse](ctx, s, CapabilityLatestTrades, "/price/"+url.PathEscape(r.Symbol)+"/trades/latest", q)
}
func (s *service) GetOHLC(ctx context.Context, r dto.GetOHLCRequest) (*dto.OhlcResponse, error) {
	q := url.Values{}
	set(q, "symbol", r.Symbol)
	set(q, "type", r.Type)
	set(q, "resolution", r.Resolution)
	setInt64(q, "from", r.From)
	setInt64(q, "to", r.To)
	return get[dto.OhlcResponse](ctx, s, CapabilityOHLC, "/price/ohlc", q)
}
func (s *service) GetClosePrice(ctx context.Context, r dto.GetClosePriceRequest) (*dto.PriceSymbolCloseResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	return get[dto.PriceSymbolCloseResponse](ctx, s, CapabilityClosePrice, "/price/"+url.PathEscape(r.Symbol)+"/close", q)
}
func (s *service) GetQuoteHistory(ctx context.Context, r dto.GetQuoteHistoryRequest) (*dto.QuotesResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	setInt64(q, "from", r.From)
	setInt64(q, "to", r.To)
	setInt(q, "limit", r.Limit)
	return get[dto.QuotesResponse](ctx, s, CapabilityQuoteHistory, "/price/"+url.PathEscape(r.Symbol)+"/quotes", q)
}
func (s *service) GetSecurityDefinition(ctx context.Context, r dto.GetSecurityDefinitionRequest) (*dto.SecdefResponse, error) {
	q := url.Values{}
	set(q, "boardId", r.BoardID)
	return get[dto.SecdefResponse](ctx, s, CapabilitySecurityDefinition, "/price/"+url.PathEscape(r.Symbol)+"/secdef", q)
}
func (s *service) GetWorkingDates(ctx context.Context, _ dto.GetWorkingDatesRequest) (*dto.WorkingDatesResponse, error) {
	return get[dto.WorkingDatesResponse](ctx, s, CapabilityWorkingDates, "/market/working-dates", nil)
}

func get[T any](ctx context.Context, s *service, capability core.Capability, path string, q url.Values) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	endpoint := strings.TrimRight(s.dependencies.BaseURL, "/") + path
	if encoded := q.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	headers := map[string]string{}
	if s.dependencies.Headers != nil {
		headers = s.dependencies.Headers()
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{Method: "GET", URL: endpoint, Headers: headers})
	if err != nil {
		return nil, err
	}
	result := new(T)
	payload := response.Raw
	if len(payload) == 0 {
		payload, err = json.Marshal(response.Body)
		if err != nil {
			return nil, err
		}
	}
	if err = json.Unmarshal(payload, result); err != nil {
		return nil, sdkerrors.Decode("dnse", string(capability), "decode DNSE native market data response", response.Body, err)
	}
	return result, nil
}
func set(q url.Values, key, value string) {
	if value != "" {
		q.Set(key, value)
	}
}
func setInt(q url.Values, key string, value int) {
	if value != 0 {
		q.Set(key, strconv.Itoa(value))
	}
}
func setInt64(q url.Values, key string, value int64) {
	if value != 0 {
		q.Set(key, strconv.FormatInt(value, 10))
	}
}
