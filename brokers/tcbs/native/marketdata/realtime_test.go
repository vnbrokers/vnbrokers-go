package marketdata_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeMarketSocket struct {
	incoming chan transport.WebSocketMessage
	sent     chan string
	closed   chan struct{}
}

func newFakeMarketSocket(messages ...string) *fakeMarketSocket {
	socket := &fakeMarketSocket{incoming: make(chan transport.WebSocketMessage, len(messages)), sent: make(chan string, 16), closed: make(chan struct{})}
	for _, message := range messages {
		socket.incoming <- transport.WebSocketMessage(message)
	}
	return socket
}

func (s *fakeMarketSocket) Send(_ context.Context, message transport.WebSocketMessage) error {
	s.sent <- string(message)
	return nil
}
func (s *fakeMarketSocket) Receive(ctx context.Context) (transport.WebSocketMessage, error) {
	select {
	case message := <-s.incoming:
		return message, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closed:
		return nil, errors.New("closed")
	}
}
func (s *fakeMarketSocket) Close() error {
	select {
	case <-s.closed:
	default:
		close(s.closed)
	}
	return nil
}

func TestMarketDataRealtimeStockHistoryProtocol(t *testing.T) {
	socket := newFakeMarketSocket(
		`d|0|{"success":true,"error":null}`,
		`d|33|15`,
		`C001|TCB|{"symbol":"TCB","closePrice":38600,"closeVol":1000}`,
	)
	service := marketdata.NewRealtimeService(marketdata.RealtimeDependencies{
		BaseURL: "https://api.example", AccessToken: func() string { return "jwt" }, PingInterval: time.Hour,
		RequireCapability: func(core.Capability) error { return nil },
		WebSocketFactory: func(context.Context, string, map[string]string) (transport.WebSocketTransport, error) {
			return socket, nil
		},
	})

	subscription, err := service.SubscribeStockTradeHistory(context.Background(), dto.SubscribeStockTradeHistoryRequest{Tickers: []string{"TCB", "FPT"}})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()
	if got := marketFrame(t, socket.sent); got != "d|a|||and0" {
		t.Fatalf("auth frame = %q", got)
	}
	if got := marketFrame(t, socket.sent); got != "d|st|C001|TCB,FPT" {
		t.Fatalf("subscribe frame = %q", got)
	}
	event := marketEvent(t, subscription.Events())
	if event.Symbol != "TCB" || event.ClosePrice != 38600 {
		t.Fatalf("event = %#v", event)
	}
}

func TestMarketDataRealtimeDerivativeBasePriceProtocol(t *testing.T) {
	socket := newFakeMarketSocket(
		`d|0|{"success":true,"error":null}`,
		`d|33|15`,
		`s|4|{"symbol":"VN30F1M","ceilPrice":1400,"floorPrice":1200,"refPrice":1300}`,
	)
	service := marketdata.NewRealtimeService(marketdata.RealtimeDependencies{
		BaseURL: "https://api.example", AccessToken: func() string { return "jwt" }, PingInterval: time.Hour,
		RequireCapability: func(core.Capability) error { return nil },
		WebSocketFactory: func(context.Context, string, map[string]string) (transport.WebSocketTransport, error) {
			return socket, nil
		},
	})

	subscription, err := service.SubscribeDerivativeBasePrices(context.Background(), dto.SubscribeDerivativeBasePricesRequest{Symbols: []string{"VN30F1M"}})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()
	_ = marketFrame(t, socket.sent)
	if got := marketFrame(t, socket.sent); got != "d|s|tk|bp|VN30F1M" {
		t.Fatalf("subscribe frame = %q", got)
	}
	event := marketEvent(t, subscription.Events())
	if event.Symbol != "VN30F1M" || event.RefPrice != 1300 {
		t.Fatalf("event = %#v", event)
	}
}

func TestMarketDataRealtimeStockPricesRemainRaw(t *testing.T) {
	socket := newFakeMarketSocket(`raw-price-frame`)
	var connectedURL string
	service := marketdata.NewRealtimeService(marketdata.RealtimeDependencies{
		BaseURL: "https://api.example", AccessToken: func() string { return "jwt" }, PingInterval: time.Hour,
		RequireCapability: func(core.Capability) error { return nil },
		WebSocketFactory: func(_ context.Context, url string, _ map[string]string) (transport.WebSocketTransport, error) {
			connectedURL = url
			return socket, nil
		},
	})

	subscription, err := service.SubscribeStockPrices(context.Background(), dto.SubscribeStockPricesRequest{})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()
	if connectedURL != "wss://api.example/ws/thesis/v1/stream/normal" {
		t.Fatalf("url = %q", connectedURL)
	}
	select {
	case sent := <-socket.sent:
		t.Fatalf("unexpected undocumented frame %q", sent)
	case <-time.After(20 * time.Millisecond):
	}
	event := marketEvent(t, subscription.Events())
	if string(event) != "raw-price-frame" {
		t.Fatalf("raw event = %q", event)
	}
}

func marketFrame(t *testing.T, values <-chan string) string {
	t.Helper()
	select {
	case value := <-values:
		return value
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for frame")
		return ""
	}
}

func marketEvent[T any](t *testing.T, values <-chan T) T {
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
