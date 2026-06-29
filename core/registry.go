package core

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// BrokerFactory builds a broker from broker-specific configuration.
type BrokerFactory func(config any) (Broker, error)

// BrokerDescriptor describes a broker adapter that can be created by name.
type BrokerDescriptor struct {
	Name    string
	Aliases []string
	New     BrokerFactory
}

var brokerRegistry = struct {
	sync.RWMutex
	factories map[string]BrokerFactory
	primary   map[string]struct{}
}{
	factories: make(map[string]BrokerFactory),
	primary:   make(map[string]struct{}),
}

// RegisterBroker registers a broker adapter factory under its name and aliases.
func RegisterBroker(descriptor BrokerDescriptor) error {
	name := normalizeBrokerName(descriptor.Name)
	if name == "" {
		return fmt.Errorf("broker name is required")
	}
	if descriptor.New == nil {
		return fmt.Errorf("broker factory is required for %q", name)
	}

	names := append([]string{name}, descriptor.Aliases...)
	normalized := make([]string, 0, len(names))
	seen := make(map[string]struct{}, len(names))
	for _, value := range names {
		candidate := normalizeBrokerName(value)
		if candidate == "" {
			continue
		}
		if _, exists := seen[candidate]; exists {
			continue
		}
		seen[candidate] = struct{}{}
		normalized = append(normalized, candidate)
	}

	brokerRegistry.Lock()
	defer brokerRegistry.Unlock()
	for _, candidate := range normalized {
		if _, exists := brokerRegistry.factories[candidate]; exists {
			return fmt.Errorf("broker %q is already registered", candidate)
		}
	}
	for _, candidate := range normalized {
		brokerRegistry.factories[candidate] = descriptor.New
	}
	brokerRegistry.primary[name] = struct{}{}
	return nil
}

// MustRegisterBroker registers a broker adapter factory and panics on error.
func MustRegisterBroker(descriptor BrokerDescriptor) {
	if err := RegisterBroker(descriptor); err != nil {
		panic(err)
	}
}

// RegisteredBrokers returns the sorted primary names of registered brokers.
func RegisteredBrokers() []string {
	brokerRegistry.RLock()
	defer brokerRegistry.RUnlock()

	out := make([]string, 0, len(brokerRegistry.primary))
	for name := range brokerRegistry.primary {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}

// NewBroker builds a registered broker by name.
func NewBroker(name string, config any) (Broker, error) {
	name = normalizeBrokerName(name)
	if name == "" {
		return nil, fmt.Errorf("broker name is required")
	}

	brokerRegistry.RLock()
	factory := brokerRegistry.factories[name]
	brokerRegistry.RUnlock()
	if factory == nil {
		return nil, fmt.Errorf("unsupported broker: %s", name)
	}
	return factory(config)
}

func normalizeBrokerName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
