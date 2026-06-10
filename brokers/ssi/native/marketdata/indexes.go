package marketdata

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) GetIndexComponents(ctx context.Context, request dto.GetIndexComponentsRequest) (*dto.GetIndexComponentsResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "indexCode", request.IndexCode)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetIndexComponentsResponse](ctx, s, CapabilityIndexComponents, "/api/v2/Market/IndexComponents", params)
}

func (s *service) GetIndexList(ctx context.Context, request dto.GetIndexListRequest) (*dto.GetIndexListResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "exchange", request.Exchange)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetIndexListResponse](ctx, s, CapabilityIndexList, "/api/v2/Market/IndexList", params)
}

func (s *service) GetDailyIndex(ctx context.Context, request dto.GetDailyIndexRequest) (*dto.GetDailyIndexResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "indexId", request.IndexID)
	setOptionalString(params, "fromDate", request.FromDate)
	setOptionalString(params, "toDate", request.ToDate)
	setOptionalBool(params, "ascending", request.Ascending)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetDailyIndexResponse](ctx, s, CapabilityDailyIndex, "/api/v2/Market/DailyIndex", params)
}
