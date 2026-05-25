package vnbrokers

import (
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	"github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	"github.com/vnbrokers/vnbrokers-go/core"
)

func NewDNSE(config dnse.Config) *dnse.Broker {
	return dnse.NewBroker(config)
}

func NewSSI(config ssi.Config) *ssi.Broker {
	return ssi.NewBroker(config)
}

func NewTCBS(config tcbs.Config) *tcbs.Broker {
	return tcbs.NewBroker(config)
}

type FactoryConfig struct {
	DNSE *dnse.Config
	SSI  *ssi.Config
	TCBS *tcbs.Config
}

func NewBroker(name string, config FactoryConfig) (core.Broker, error) {
	switch name {
	case "dnse":
		if config.DNSE == nil {
			return nil, fmt.Errorf("dnse config is required")
		}
		return dnse.NewBroker(*config.DNSE), nil
	case "ssi":
		if config.SSI == nil {
			return nil, fmt.Errorf("ssi config is required")
		}
		return ssi.NewBroker(*config.SSI), nil
	case "tcbs":
		if config.TCBS == nil {
			return nil, fmt.Errorf("tcbs config is required")
		}
		return tcbs.NewBroker(*config.TCBS), nil
	default:
		return nil, fmt.Errorf("unsupported broker: %s", name)
	}
}
