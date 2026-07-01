package marketdata

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const CapabilityDerivatives core.Capability = "native.marketdata.derivatives"

type Service interface {
	GetDerivatives(context.Context, dto.GetDerivativesRequest) (*dto.GetDerivativesResponse, error)
}

type Dependencies struct {
	BaseURL           string
	Headers           func(bool) map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}

type service struct{ dependencies Dependencies }

func NewService(dependencies Dependencies) Service { return &service{dependencies: dependencies} }

func (s *service) GetDerivatives(ctx context.Context, _ dto.GetDerivativesRequest) (*dto.GetDerivativesResponse, error) {
	if err := s.dependencies.RequireCapability(CapabilityDerivatives); err != nil {
		return nil, err
	}
	response, err := s.dependencies.Send(ctx, "native.marketdata.get_derivatives", transport.HTTPRequest{
		Method: "GET", URL: httpx.URL(s.dependencies.BaseURL, "/derivatives", nil),
		Headers: s.dependencies.Headers(false),
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.GetDerivativesResponse](
		"entrade",
		"native.marketdata.get_derivatives",
		"decode response",
		response,
	)
	if err != nil {
		return nil, err
	}
	return out, nil
}
