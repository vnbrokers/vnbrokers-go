package marketdata

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetStockTickers(ctx context.Context, request dto.GetStockTickersRequest) (*dto.GetStockTickersResponse, error) {
	query := url.Values{}
	set(query, "tickers", request.Tickers)
	set(query, "index", strconv.FormatFloat(request.Index, 'f', -1, 64))
	return do[dto.GetStockTickersResponse](s, ctx, CapabilityGetStockTickers, "/tartarus/v1/tickerCommons", query)
}

func (s *service) GetStockForeignRooms(ctx context.Context, request dto.GetStockForeignRoomsRequest) (*dto.GetStockForeignRoomsResponse, error) {
	query := url.Values{}
	set(query, "index", strconv.FormatFloat(request.Index, 'f', -1, 64))
	return do[dto.GetStockForeignRoomsResponse](s, ctx, CapabilityGetStockForeignRooms, "/tartarus/v1/tickerSnaps", query)
}

func (s *service) GetStockPutThroughs(ctx context.Context, request dto.GetStockPutThroughsRequest) (*dto.GetStockPutThroughsResponse, error) {
	query := url.Values{}
	set(query, "floor", strconv.FormatFloat(request.Floor, 'f', -1, 64))
	return do[dto.GetStockPutThroughsResponse](s, ctx, CapabilityGetStockPutThroughs, "/tartarus/v1/putThroughSnaps", query)
}

func (s *service) GetStockTradeHistory(ctx context.Context, request dto.GetStockTradeHistoryRequest) (*dto.GetStockTradeHistoryResponse, error) {
	query := url.Values{}
	set(query, "page", strconv.FormatFloat(request.Page, 'f', -1, 64))
	set(query, "size", strconv.FormatFloat(request.Size, 'f', -1, 64))
	set(query, "headIndex", strconv.FormatFloat(request.HeadIndex, 'f', -1, 64))
	return do[dto.GetStockTradeHistoryResponse](s, ctx, CapabilityGetStockTradeHistory, "/nyx/v1/intraday/"+escaped(request.Ticker)+"/his/paging", query)
}

func (s *service) GetStockSupplyDemand15Minutes(ctx context.Context, request dto.GetStockSupplyDemand15MinutesRequest) (*dto.GetStockSupplyDemand15MinutesResponse, error) {
	query := url.Values{}
	set(query, "timeWindow", request.TimeWindow)
	set(query, "tWindow", request.TWindow)
	set(query, "type", request.Type)
	return do[dto.GetStockSupplyDemand15MinutesResponse](s, ctx, CapabilityGetStockSupplyDemand15Minutes, "/nyx/v1/intraday/"+escaped(request.Ticker)+"/bsa-ext", query)
}

func (s *service) GetStockSupplyDemandDaily(ctx context.Context, request dto.GetStockSupplyDemandDailyRequest) (*dto.GetStockSupplyDemandDailyResponse, error) {
	query := url.Values{}
	set(query, "type", request.Type)
	return do[dto.GetStockSupplyDemandDailyResponse](s, ctx, CapabilityGetStockSupplyDemandDaily, "/nyx/v1/intraday/"+escaped(request.Ticker)+"/bsa", query)
}

func (s *service) GetStockSupplyDemandMonthly(ctx context.Context, request dto.GetStockSupplyDemandMonthlyRequest) (*dto.GetStockSupplyDemandMonthlyResponse, error) {
	query := url.Values{}
	set(query, "timeWindow", request.TimeWindow)
	set(query, "type", request.Type)
	return do[dto.GetStockSupplyDemandMonthlyResponse](s, ctx, CapabilityGetStockSupplyDemandMonthly, "/nyx/v1/intraday/"+escaped(request.Ticker)+"/bsa-month", query)
}
