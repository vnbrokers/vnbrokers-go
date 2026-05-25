package core

import sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"

type Broker interface {
	Name() string
	Capabilities() []Capability
	Supports(Capability) bool
	RequireCapability(Capability) error
}

type BaseBroker struct {
	BrokerName         string
	BrokerCapabilities []Capability
}

func (b BaseBroker) Name() string {
	return b.BrokerName
}

func (b BaseBroker) Capabilities() []Capability {
	out := make([]Capability, len(b.BrokerCapabilities))
	copy(out, b.BrokerCapabilities)
	return out
}

func (b BaseBroker) Supports(capability Capability) bool {
	for _, supported := range b.BrokerCapabilities {
		if supported == capability {
			return true
		}
	}
	return false
}

func (b BaseBroker) RequireCapability(capability Capability) error {
	if b.Supports(capability) {
		return nil
	}
	return sdkerrors.UnsupportedCapability(b.BrokerName, string(capability))
}
