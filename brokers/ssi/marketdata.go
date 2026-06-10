package ssi

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func (s *MarketDataService) Securities(ctx context.Context, request dto.SecuritiesRequest) (dto.SecuritiesResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "market", request.Market)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.SecuritiesResponse](ctx, s, "marketdata.securities", "/api/v2/Market/Securities", params)
}

func (s *MarketDataService) SecuritiesDetails(ctx context.Context, request dto.SecuritiesDetailsRequest) (dto.SecuritiesDetailsResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "market", request.Market)
	setOptionalString(params, "symbol", request.Symbol)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.SecuritiesDetailsResponse](ctx, s, "marketdata.securities_details", "/api/v2/Market/SecuritiesDetails", params)
}

func (s *MarketDataService) IndexComponents(ctx context.Context, request dto.IndexComponentsRequest) (dto.IndexComponentsResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "indexCode", request.IndexCode)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.IndexComponentsResponse](ctx, s, "marketdata.index_components", "/api/v2/Market/IndexComponents", params)
}

func (s *MarketDataService) IndexList(ctx context.Context, request dto.IndexListRequest) (dto.IndexListResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "exchange", request.Exchange)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.IndexListResponse](ctx, s, "marketdata.index_list", "/api/v2/Market/IndexList", params)
}

func (s *MarketDataService) DailyOhlc(ctx context.Context, request dto.DailyOhlcRequest) (dto.DailyOhlcResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.DailyOhlcResponse](ctx, s, "marketdata.daily_ohlc", "/api/v2/Market/DailyOhlc", params)
}

func (s *MarketDataService) IntradayOhlc(ctx context.Context, request dto.IntradayOhlcRequest) (dto.IntradayOhlcResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setOptionalInt(params, "resolution", request.Resolution)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.IntradayOhlcResponse](ctx, s, "marketdata.intraday_ohlc", "/api/v2/Market/IntradayOhlc", params)
}

func (s *MarketDataService) DailyIndex(ctx context.Context, request dto.DailyIndexRequest) (dto.DailyIndexResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "indexId", request.IndexID)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.DailyIndexResponse](ctx, s, "marketdata.daily_index", "/api/v2/Market/DailyIndex", params)
}

func (s *MarketDataService) DailyStockPrice(ctx context.Context, request dto.DailyStockPriceRequest) (dto.DailyStockPriceResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalString(params, "market", request.Market)
	setPagination(params, request.PageIndex, request.PageSize)
	return getMarketData[dto.DailyStockPriceResponse](ctx, s, "marketdata.daily_stock_price", "/api/v2/Market/DailyStockPrice", params)
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

func setOptionalBool(params url.Values, key string, value bool) {
	if value {
		params.Set(key, "true")
	}
}

func getMarketData[T any](
	ctx context.Context,
	s *MarketDataService,
	operation string,
	path string,
	params url.Values,
) (T, error) {
	var zero T
	if err := s.broker.RequireCapability(core.CapabilityMarketDataSymbolsList); err != nil {
		return zero, err
	}

	endpoint := strings.TrimRight(s.broker.config.DataBaseURL, "/") + path
	if encoded := params.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	headers := map[string]string{"Accept": "application/json"}
	if token := s.broker.dataToken(); token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	response, err := s.broker.send(ctx, operation, false, false, transport.HTTPRequest{
		Method:  "GET",
		URL:     endpoint,
		Headers: headers,
	})
	if err != nil {
		return zero, err
	}

	var result T
	if err := decode(response, &result); err != nil {
		return zero, sdkerrors.Decode("ssi", operation, "decode SSI market data response", response.Body, err)
	}
	return result, nil
}
