package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetSubAccountInformation(ctx context.Context, request dto.GetSubAccountInformationRequest) (*dto.GetSubAccountInformationResponse, error) {
	query := url.Values{}
	set(query, "fields", request.Fields)
	return do[dto.GetSubAccountInformationResponse](s, ctx, CapabilityGetSubAccountInformation, "GET", "/eros/v2/get-profile/by-username/"+escaped(request.CustodyCode), query, nil)
}
