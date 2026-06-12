package marketdata

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
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
		Method: "GET", URL: strings.TrimRight(s.dependencies.BaseURL, "/") + "/derivatives",
		Headers: s.dependencies.Headers(false),
	})
	if err != nil {
		return nil, err
	}
	var out dto.GetDerivativesResponse
	if len(response.Raw) > 0 {
		err = json.Unmarshal(response.Raw, &out)
	} else {
		var payload []byte
		payload, err = json.Marshal(response.Body)
		if err == nil {
			err = json.Unmarshal(payload, &out)
		}
	}
	if err != nil {
		return nil, sdkerrors.Decode("entrade", "native.marketdata.get_derivatives", "decode response", responsePayload(response), err)
	}
	return &out, nil
}

func responsePayload(response transport.HTTPResponse) any {
	if len(response.Raw) > 0 {
		return response.Raw
	}
	return response.Body
}
