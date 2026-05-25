package dnse

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type BrokerageService struct {
	broker *Broker
}

func (s *BrokerageService) CareByAccounts(ctx context.Context, version string) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityBrokerageCareBy); err != nil {
		return domain.RawPayload{}, err
	}
	if version == "" {
		version = "2026-05-07"
	}
	headers := s.broker.apiHeaders()
	headers["version"] = version
	response, err := s.broker.send(ctx, "brokerage.care_by_accounts", transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/brokers/accounts/care-by"),
		Headers: headers,
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}
