package marketdata

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetDerivativeTickers(ctx context.Context, _ dto.GetDerivativeTickersRequest) (*dto.GetDerivativeTickersResponse, error) {
	return do[dto.GetDerivativeTickersResponse](s, ctx, CapabilityGetDerivativeTickers, "/tartarus/v1/derivatives", url.Values{})
}
