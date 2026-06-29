package ssi

import (
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/core"
)

func init() {
	core.MustRegisterBroker(core.BrokerDescriptor{
		Name: "ssi",
		New:  newBrokerFromConfig,
	})
}

func newBrokerFromConfig(config any) (core.Broker, error) {
	switch c := config.(type) {
	case Config:
		return NewBroker(c), nil
	case *Config:
		if c == nil {
			return nil, fmt.Errorf("ssi config is required")
		}
		return NewBroker(*c), nil
	default:
		return nil, fmt.Errorf("ssi config must be ssi.Config or *ssi.Config")
	}
}
