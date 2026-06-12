package native

import (
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/trading"
)

type Service interface {
	Trading() trading.Service
	MarketData() marketdata.Service
}

type service struct {
	trading    trading.Service
	marketData marketdata.Service
}

func NewService(tradingService trading.Service, marketDataService marketdata.Service) Service {
	return &service{trading: tradingService, marketData: marketDataService}
}

func (s *service) Trading() trading.Service       { return s.trading }
func (s *service) MarketData() marketdata.Service { return s.marketData }
