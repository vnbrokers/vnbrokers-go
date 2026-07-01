package marketdata

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	sdkmarketdata "github.com/vnbrokers/vnbrokers-go/marketdata"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilitySecurities        core.Capability = "native.marketdata.securities"
	CapabilitySecuritiesDetails core.Capability = "native.marketdata.securities_details"
	CapabilityIndexComponents   core.Capability = "native.marketdata.index_components"
	CapabilityIndexList         core.Capability = "native.marketdata.index_list"
	CapabilityDailyOhlc         core.Capability = "native.marketdata.daily_ohlc"
	CapabilityIntradayOhlc      core.Capability = "native.marketdata.intraday_ohlc"
	CapabilityDailyIndex        core.Capability = "native.marketdata.daily_index"
	CapabilityDailyStockPrice   core.Capability = "native.marketdata.daily_stock_price"
)

type Service interface {
	Realtime() RealtimeService
	GetSecurities(context.Context, dto.GetSecuritiesRequest) (*dto.GetSecuritiesResponse, error)
	GetSecuritiesDetails(context.Context, dto.GetSecuritiesDetailsRequest) (*dto.GetSecuritiesDetailsResponse, error)
	GetIndexComponents(context.Context, dto.GetIndexComponentsRequest) (*dto.GetIndexComponentsResponse, error)
	GetIndexList(context.Context, dto.GetIndexListRequest) (*dto.GetIndexListResponse, error)
	GetDailyOhlc(context.Context, dto.GetDailyOhlcRequest) (*dto.GetDailyOhlcResponse, error)
	GetIntradayOhlc(context.Context, dto.GetIntradayOhlcRequest) (*dto.GetIntradayOhlcResponse, error)
	GetDailyIndex(context.Context, dto.GetDailyIndexRequest) (*dto.GetDailyIndexResponse, error)
	GetDailyStockPrice(context.Context, dto.GetDailyStockPriceRequest) (*dto.GetDailyStockPriceResponse, error)
}

type RealtimeService interface {
	SubscribeTradingStatus(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.TradingStatusEvent], error)
	SubscribeQuotes(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.QuoteEvent], error)
	SubscribeTrades(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.TradeEvent], error)
	SubscribeSnapshots(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.SnapshotEvent], error)
	SubscribeForeignRooms(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.ForeignRoomEvent], error)
	SubscribeMarketIndexes(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.MarketIndexEvent], error)
	SubscribeOHLCV(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.OHLCVEvent], error)
	SubscribeOddLots(context.Context, sdkmarketdata.SubscribeSymbolRequest) (realtime.Subscription[dto.OddLotEvent], error)
	SubscribeRawChannel(context.Context, string) (realtime.Subscription[domain.RawPayload], error)
}

type Dependencies struct {
	BaseURL           string
	DataToken         func() string
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

	headers := map[string]string{"Accept": "application/json"}
	if token := s.dependencies.DataToken(); token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method:  "GET",
		URL:     httpx.URL(s.dependencies.BaseURL, path, params),
		Headers: headers,
	})
	if err != nil {
		return nil, err
	}

	result, err := httpx.DecodeResponse[T]("ssi", string(capability), "decode SSI native market data response", response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func normalizePagination(pageIndex int, pageSize int) (int, int) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageIndex, pageSize
}

func setPagination(params url.Values, pageIndex int, pageSize int) {
	params.Set("pageIndex", strconv.Itoa(pageIndex))
	params.Set("pageSize", strconv.Itoa(pageSize))
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

func setOptionalBool(params url.Values, key string, value bool) {
	if value {
		params.Set(key, "true")
	}
}
