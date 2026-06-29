package vnbrokers

import (
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func NewDNSE(config dnse.Config) *dnse.Broker {
	return dnse.NewBroker(config)
}

func NewEntrade(config entrade.Config) *entrade.Broker {
	return entrade.NewBroker(config)
}

func NewSSI(config ssi.Config) *ssi.Broker {
	return ssi.NewBroker(config)
}

func NewTCBS(config tcbs.Config) *tcbs.Broker {
	return tcbs.NewBroker(config)
}

func NewFHSC(config fhsc.Config) *fhsc.Broker {
	return fhsc.NewBroker(config)
}

type BrokerFactory = core.BrokerFactory
type BrokerDescriptor = core.BrokerDescriptor
type BrokerConfig = core.BrokerConfig
type Brokers = core.Brokers

// FactoryConfig is kept for transition while callers migrate to broker-specific config values.
// Deprecated: use NewBroker with a broker-specific Config value instead.
type FactoryConfig struct {
	DNSE    *dnse.Config
	Entrade *entrade.Config
	FHSC    *fhsc.Config
	SSI     *ssi.Config
	TCBS    *tcbs.Config
}

// RegisterBroker registers a broker adapter factory.
func RegisterBroker(descriptor BrokerDescriptor) error {
	return core.RegisterBroker(descriptor)
}

// MustRegisterBroker registers a broker adapter factory and panics on error.
func MustRegisterBroker(descriptor BrokerDescriptor) {
	core.MustRegisterBroker(descriptor)
}

// RegisteredBrokers returns the sorted names of registered broker adapters.
func RegisteredBrokers() []string {
	return core.RegisteredBrokers()
}

// NewBroker builds a registered broker by name from broker-specific config.
func NewBroker(name string, config any) (core.Broker, error) {
	return core.NewBroker(name, config)
}

// NewBrokers builds multiple registered brokers keyed by caller-defined ID.
func NewBrokers(configs []BrokerConfig) (Brokers, error) {
	return core.NewBrokers(configs)
}
