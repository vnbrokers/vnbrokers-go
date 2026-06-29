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

// NamedBrokerConfig is implemented by broker-specific config values that know
// which registered broker adapter should build them.
type NamedBrokerConfig interface {
	BrokerName() string
}

// BrokerConfig describes one broker instance to build.
type BrokerConfig struct {
	// ID is the caller-defined instance key, for example "dnse-main" or "tcbs-paper".
	ID string
	// Config is the broker-specific config value accepted by the adapter factory.
	Config NamedBrokerConfig
}

// Brokers is a keyed set of broker instances.
type Brokers map[string]Broker

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

// NewBrokers builds multiple registered broker instances keyed by caller-defined ID.
func NewBrokers(configs []BrokerConfig) (Brokers, error) {
	brokers := make(Brokers, len(configs))
	for _, config := range configs {
		id := strings.TrimSpace(config.ID)
		if id == "" {
			return nil, fmt.Errorf("broker instance id is required")
		}
		if _, exists := brokers[id]; exists {
			return nil, fmt.Errorf("duplicate broker instance id: %s", id)
		}

		if config.Config == nil {
			return nil, fmt.Errorf("broker config is required for %q", id)
		}

		name := normalizeBrokerName(config.Config.BrokerName())
		broker, err := NewBroker(name, config.Config)
		if err != nil {
			return nil, fmt.Errorf("build broker %q (%s): %w", id, name, err)
		}
		brokers[id] = broker
	}
	return brokers, nil
}

// Get returns a broker by instance ID.
func (b Brokers) Get(id string) (Broker, bool) {
	broker, ok := b[strings.TrimSpace(id)]
	return broker, ok
}

func normalizeBrokerName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
