package tcbs

import (
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/core"
)

func init() {
	core.MustRegisterBroker(core.BrokerDescriptor{
		Name: "tcbs",
		New:  newBrokerFromConfig,
	})
}

func newBrokerFromConfig(config any) (core.Broker, error) {
	switch c := config.(type) {
	case Config:
		return NewBroker(c), nil
	case *Config:
		if c == nil {
			return nil, fmt.Errorf("tcbs config is required")
		}
		return NewBroker(*c), nil
	default:
		return nil, fmt.Errorf("tcbs config must be tcbs.Config or *tcbs.Config")
	}
}
