package entrade

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type MarketDataService struct {
	derivatives *MarketDataDerivativesService
}

func (s *MarketDataService) Derivatives() *MarketDataDerivativesService {
	return s.derivatives
}

type MarketDataDerivativesService struct {
	broker *Broker
}

func (s *MarketDataDerivativesService) List(ctx context.Context) ([]domain.Symbol, error) {
	if err := s.broker.RequireCapability(core.CapabilityMarketDataSymbolsList); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "marketdata.derivatives.list", true, transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/derivatives"),
		Headers: s.broker.headers(true, false),
	})
	if err != nil {
		return nil, err
	}
	return MapDerivatives(expectObject(response.Body)), nil
}
