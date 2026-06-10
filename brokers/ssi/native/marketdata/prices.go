package marketdata

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) GetDailyOhlc(ctx context.Context, request dto.GetDailyOhlcRequest) (*dto.GetDailyOhlcResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetDailyOhlcResponse](ctx, s, CapabilityDailyOhlc, "/api/v2/Market/DailyOhlc", params)
}

func (s *service) GetIntradayOhlc(ctx context.Context, request dto.GetIntradayOhlcRequest) (*dto.GetIntradayOhlcResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setOptionalInt(params, "resolution", request.Resolution)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetIntradayOhlcResponse](ctx, s, CapabilityIntradayOhlc, "/api/v2/Market/IntradayOhlc", params)
}

func (s *service) GetDailyStockPrice(ctx context.Context, request dto.GetDailyStockPriceRequest) (*dto.GetDailyStockPriceResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "symbol", request.Symbol)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalString(params, "market", request.Market)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetDailyStockPriceResponse](ctx, s, CapabilityDailyStockPrice, "/api/v2/Market/DailyStockPrice", params)
}
