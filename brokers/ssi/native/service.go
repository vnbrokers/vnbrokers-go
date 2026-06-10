package native

import "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/marketdata"

type Service interface {
	MarketData() marketdata.Service
}

type service struct {
	marketData marketdata.Service
}

func NewService(marketData marketdata.Service) Service {
	return &service{marketData: marketData}
}

func (s *service) MarketData() marketdata.Service {
	return s.marketData
}
