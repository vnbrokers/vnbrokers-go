package ssi

import "github.com/vnbrokers/vnbrokers-go/core"

type Config struct {
	BaseURL     string
	AccessToken string
}

type Broker struct {
	core.BaseBroker
	config Config
}

func NewBroker(config Config) *Broker {
	return &Broker{
		BaseBroker: core.BaseBroker{
			BrokerName:         "ssi",
			BrokerCapabilities: []core.Capability{},
		},
		config: config,
	}
}
