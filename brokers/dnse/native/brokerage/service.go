package brokerage

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const CapabilityCareByAccounts core.Capability = "native.brokerage.care_by_accounts"

type Service interface {
	GetCareByAccounts(context.Context, dto.GetCareByAccountsRequest) (*dto.CareByResponse, error)
}
type Dependencies struct {
	BaseURL           string
	Headers           func() map[string]string
	RequireCapability func(core.Capability) error
	Send              func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error)
}
type service struct{ dependencies Dependencies }

func NewService(d Dependencies) Service { return &service{dependencies: d} }
func (s *service) GetCareByAccounts(ctx context.Context, r dto.GetCareByAccountsRequest) (*dto.CareByResponse, error) {
	if err := s.dependencies.RequireCapability(CapabilityCareByAccounts); err != nil {
		return nil, err
	}
	headers := s.dependencies.Headers()
	if r.Version != "" {
		headers["version"] = r.Version
	}
	response, err := s.dependencies.Send(ctx, string(CapabilityCareByAccounts), transport.HTTPRequest{
		Method:  "GET",
		URL:     httpx.URL(s.dependencies.BaseURL, "/brokers/accounts/care-by", nil),
		Headers: headers,
	})
	if err != nil {
		return nil, err
	}
	return httpx.DecodeResponse[dto.CareByResponse]("dnse", string(CapabilityCareByAccounts), "decode DNSE native brokerage response", response)
}
