package core

import "testing"

type registryTestBroker struct{ BaseBroker }

func TestBrokerRegistryBuildsRegisteredBroker(t *testing.T) {
	const name = "registry-test-build"
	if err := RegisterBroker(BrokerDescriptor{
		Name:    name,
		Aliases: []string{"registry-test-build-alias"},
		New: func(config any) (Broker, error) {
			return &registryTestBroker{BaseBroker: BaseBroker{BrokerName: config.(string)}}, nil
		},
	}); err != nil {
		t.Fatalf("register broker: %v", err)
	}

	broker, err := NewBroker(" REGISTRY-TEST-BUILD-ALIAS ", "built")
	if err != nil {
		t.Fatalf("new broker: %v", err)
	}
	if broker.Name() != "built" {
		t.Fatalf("broker name = %q", broker.Name())
	}
}

func TestBrokerRegistryRejectsDuplicateName(t *testing.T) {
	const name = "registry-test-duplicate"
	descriptor := BrokerDescriptor{
		Name: name,
		New:  func(any) (Broker, error) { return &registryTestBroker{}, nil },
	}
	if err := RegisterBroker(descriptor); err != nil {
		t.Fatalf("first register: %v", err)
	}
	if err := RegisterBroker(descriptor); err == nil {
		t.Fatal("expected duplicate registration error")
	}
}

func TestBrokerRegistryRejectsUnknownBroker(t *testing.T) {
	if _, err := NewBroker("registry-test-unknown", nil); err == nil {
		t.Fatal("expected unknown broker error")
	}
}
