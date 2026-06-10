package marketdata

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) GetSecurities(ctx context.Context, request dto.GetSecuritiesRequest) (*dto.GetSecuritiesResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "market", request.Market)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetSecuritiesResponse](ctx, s, CapabilitySecurities, "/api/v2/Market/Securities", params)
}

func (s *service) GetSecuritiesDetails(ctx context.Context, request dto.GetSecuritiesDetailsRequest) (*dto.GetSecuritiesDetailsResponse, error) {
	request.PageIndex, request.PageSize = normalizePagination(request.PageIndex, request.PageSize)
	params := url.Values{}
	setOptionalString(params, "market", request.Market)
	setOptionalString(params, "symbol", request.Symbol)
	setPagination(params, request.PageIndex, request.PageSize)
	return get[dto.GetSecuritiesDetailsResponse](ctx, s, CapabilitySecuritiesDetails, "/api/v2/Market/SecuritiesDetails", params)
}
