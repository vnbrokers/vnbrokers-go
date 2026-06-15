package trading_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeSocket struct {
	incoming chan transport.WebSocketMessage
	sent     chan string
	closed   chan struct{}
}

func newFakeSocket(messages ...string) *fakeSocket {
	socket := &fakeSocket{incoming: make(chan transport.WebSocketMessage, len(messages)), sent: make(chan string, 16), closed: make(chan struct{})}
	for _, message := range messages {
		socket.incoming <- transport.WebSocketMessage(message)
	}
	return socket
}

func (s *fakeSocket) Send(_ context.Context, message transport.WebSocketMessage) error {
	s.sent <- string(message)
	return nil
}

func (s *fakeSocket) Receive(ctx context.Context) (transport.WebSocketMessage, error) {
	select {
	case message := <-s.incoming:
		return message, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closed:
		return nil, errors.New("closed")
	}
}

func (s *fakeSocket) Close() error {
	select {
	case <-s.closed:
	default:
		close(s.closed)
	}
	return nil
}

func TestTradingRealtimeStockOrdersProtocol(t *testing.T) {
	socket := newFakeSocket(
		`authenticate|{"success":true,"error":null}`,
		`pingTimeout|7`,
		`message_proto|STOCK_ORDER|{"orderId":"42","orderQtty":"100","symbol":"FPT"}`,
	)
	var connectedURL string
	service := trading.NewRealtimeService(trading.RealtimeDependencies{
		BaseURL: "https://api.example", AccessToken: func() string { return "jwt" }, PingInterval: time.Hour,
		RequireCapability: func(core.Capability) error { return nil },
		WebSocketFactory: func(_ context.Context, url string, _ map[string]string) (transport.WebSocketTransport, error) {
			connectedURL = url
			return socket, nil
		},
	})

	subscription, err := service.SubscribeStockOrders(context.Background(), dto.SubscribeStockOrdersRequest{})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()

	if connectedURL != "wss://api.example/ws/aither" {
		t.Fatalf("url = %q", connectedURL)
	}
	if got := receiveString(t, socket.sent); got != "authenticate|eyJqd3QiOiJqd3QifQ==" {
		t.Fatalf("auth frame = %q", got)
	}
	if got := receiveString(t, socket.sent); got != "subscribe|eyJ0b3BpYyI6IlNUT0NLX09SREVSIn0=" {
		t.Fatalf("subscribe frame = %q", got)
	}
	event := receiveEvent(t, subscription.Events())
	if event.OrderID != "42" || event.Symbol != "FPT" {
		t.Fatalf("event = %#v", event)
	}
}

func TestTradingRealtimeDerivativeOrdersProtocol(t *testing.T) {
	socket := newFakeSocket(
		`authenticate|{"success":true,"error":null}`,
		`pingTimeout|45`,
		`message_proto|DE_ORDER|{"orderNo":"41","volume":1,"symbol":"VN30F1M"}`,
	)
	service := trading.NewRealtimeService(trading.RealtimeDependencies{
		BaseURL: "https://api.example", AccessToken: func() string { return "jwt" }, PingInterval: time.Hour,
		RequireCapability: func(core.Capability) error { return nil },
		WebSocketFactory: func(context.Context, string, map[string]string) (transport.WebSocketTransport, error) {
			return socket, nil
		},
	})

	subscription, err := service.SubscribeDerivativeOrders(context.Background(), dto.SubscribeDerivativeOrdersRequest{})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()
	_ = receiveString(t, socket.sent)
	if got := receiveString(t, socket.sent); got != "subscribe|eyJ0b3BpYyI6IkRFX09SREVSIn0=" {
		t.Fatalf("subscribe frame = %q", got)
	}
	event := receiveEvent(t, subscription.Events())
	if event.OrderNo != "41" || event.Volume != 1 {
		t.Fatalf("event = %#v", event)
	}
}

func receiveString(t *testing.T, values <-chan string) string {
	t.Helper()
	select {
	case value := <-values:
		return value
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for frame")
		return ""
	}
}

func receiveEvent[T any](t *testing.T, values <-chan T) T {
	t.Helper()
	select {
	case value := <-values:
		return value
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for event")
		var zero T
		return zero
	}
}
