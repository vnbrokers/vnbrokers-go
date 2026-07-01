package marketdata

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityGetDerivativeTickers          core.Capability = "native.marketdata.get_derivative_tickers"
	CapabilityGetStockTickers               core.Capability = "native.marketdata.get_stock_tickers"
	CapabilityGetStockForeignRooms          core.Capability = "native.marketdata.get_stock_foreign_rooms"
	CapabilityGetStockPutThroughs           core.Capability = "native.marketdata.get_stock_put_throughs"
	CapabilityGetStockTradeHistory          core.Capability = "native.marketdata.get_stock_trade_history"
	CapabilityGetStockSupplyDemand15Minutes core.Capability = "native.marketdata.get_stock_supply_demand_15_minutes"
	CapabilityGetStockSupplyDemandDaily     core.Capability = "native.marketdata.get_stock_supply_demand_daily"
	CapabilityGetStockSupplyDemandMonthly   core.Capability = "native.marketdata.get_stock_supply_demand_monthly"
)

type Service interface {
	Realtime() RealtimeService
	GetDerivativeTickers(context.Context, dto.GetDerivativeTickersRequest) (*dto.GetDerivativeTickersResponse, error)
	GetStockTickers(context.Context, dto.GetStockTickersRequest) (*dto.GetStockTickersResponse, error)
	GetStockForeignRooms(context.Context, dto.GetStockForeignRoomsRequest) (*dto.GetStockForeignRoomsResponse, error)
	GetStockPutThroughs(context.Context, dto.GetStockPutThroughsRequest) (*dto.GetStockPutThroughsResponse, error)
	GetStockTradeHistory(context.Context, dto.GetStockTradeHistoryRequest) (*dto.GetStockTradeHistoryResponse, error)
	GetStockSupplyDemand15Minutes(context.Context, dto.GetStockSupplyDemand15MinutesRequest) (*dto.GetStockSupplyDemand15MinutesResponse, error)
	GetStockSupplyDemandDaily(context.Context, dto.GetStockSupplyDemandDailyRequest) (*dto.GetStockSupplyDemandDailyResponse, error)
	GetStockSupplyDemandMonthly(context.Context, dto.GetStockSupplyDemandMonthlyRequest) (*dto.GetStockSupplyDemandMonthlyResponse, error)
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

func do[T any](s *service, ctx context.Context, capability core.Capability, path string, query url.Values) (*T, error) {
	if err := s.dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	response, err := s.dependencies.Send(ctx, string(capability), transport.HTTPRequest{
		Method: "GET", URL: httpx.URL(s.dependencies.BaseURL, path, query), Headers: s.dependencies.Headers(true, false),
	})
	if err != nil {
		return nil, err
	}
	result, err := httpx.DecodeResponse[T]("tcbs", string(capability), "decode TCBS native market-data response", response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func escaped(value string) string { return url.PathEscape(value) }

func set(query url.Values, key, value string) { query.Set(key, value) }
