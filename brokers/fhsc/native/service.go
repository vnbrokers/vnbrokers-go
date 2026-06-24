package native

import (
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/trading"
)

type Service interface {
	MarketData() marketdata.Service
	Trading() trading.Service
}

type service struct {
	marketData marketdata.Service
	trading    trading.Service
}

func NewService(marketData marketdata.Service, trading trading.Service) Service {
	return &service{marketData: marketData, trading: trading}
}

func (s *service) MarketData() marketdata.Service {
	return s.marketData
}

func (s *service) Trading() trading.Service {
	return s.trading
}
