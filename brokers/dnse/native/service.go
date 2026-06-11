package native

import (
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/brokerage"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/trading"
)

type Service interface {
	MarketData() marketdata.Service
	Trading() trading.Service
	Brokerage() brokerage.Service
}

type service struct {
	marketData marketdata.Service
	trading    trading.Service
	brokerage  brokerage.Service
}

func NewService(marketData marketdata.Service, trading trading.Service, brokerage brokerage.Service) Service {
	return &service{marketData: marketData, trading: trading, brokerage: brokerage}
}

func (s *service) MarketData() marketdata.Service { return s.marketData }
func (s *service) Trading() trading.Service       { return s.trading }
func (s *service) Brokerage() brokerage.Service   { return s.brokerage }
