package dnse

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestMarketDataSubscribeMessageSupportsMultipleSymbols(t *testing.T) {
	message := BuildMarketDataSubscribeMessage("top_price.G1.json", []string{"ACB", "HPG"})
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)
	symbols := channel["symbols"].([]any)

	if len(symbols) != 2 {
		t.Fatalf("symbols len = %d", len(symbols))
	}
}

func TestMarketDataSubscribeMessageSupportsAllSymbols(t *testing.T) {
	message := BuildMarketDataSubscribeMessage("top_price.G1.json", nil)
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)

	if _, ok := channel["symbols"]; ok {
		t.Fatalf("all-symbol subscribe should omit symbols")
	}
}

func TestTradingSubscribeOrdersMessageUsesEncoding(t *testing.T) {
	message := BuildStreamSubscribeOrdersMessage("STOCK", "msgpack")
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)

	if channel["name"] != "order.STOCK.msgpack" {
		t.Fatalf("channel = %s", channel["name"])
	}
}

func TestConfigDefaultsStreamPongInterval(t *testing.T) {
	config := Config{}.withDefaults()

	if config.StreamPongInterval != 30*time.Second {
		t.Fatalf("stream pong interval = %s", config.StreamPongInterval)
	}
}

func TestStreamRespondsPongToServerPing(t *testing.T) {
	socket := newFakeWebSocket()
	broker := NewBroker(Config{
		APIKey:             "key",
		APISecret:          "secret",
		StreamPongInterval: time.Hour,
		WebSocketFactory: func(context.Context, string) (transport.WebSocketTransport, error) {
			return socket, nil
		},
	})

	subscription, err := startRealtimeSubscription[map[string]any](
		context.Background(),
		broker,
		map[string]any{"action": "subscribe", "channels": []any{}},
		func(map[string]any) bool { return false },
		func(*realtime.QueueSubscription[map[string]any], map[string]any) {},
	)
	if err != nil {
		t.Fatalf("start subscription: %v", err)
	}
	defer subscription.Close()

	_ = receiveSentMessage(t, socket)
	_ = receiveSentMessage(t, socket)
	socket.receiveJSON(t, map[string]any{"action": "ping"})
	message := receiveSentMessage(t, socket)

	if message["action"] != "pong" {
		t.Fatalf("action = %v", message["action"])
	}
}

func TestStreamSendsProactivePong(t *testing.T) {
	socket := newFakeWebSocket()
	broker := NewBroker(Config{
		APIKey:             "key",
		APISecret:          "secret",
		StreamPongInterval: 5 * time.Millisecond,
		WebSocketFactory: func(context.Context, string) (transport.WebSocketTransport, error) {
			return socket, nil
		},
	})

	subscription, err := startRealtimeSubscription[map[string]any](
		context.Background(),
		broker,
		map[string]any{"action": "subscribe", "channels": []any{}},
		func(map[string]any) bool { return false },
		func(*realtime.QueueSubscription[map[string]any], map[string]any) {},
	)
	if err != nil {
		t.Fatalf("start subscription: %v", err)
	}
	defer subscription.Close()

	_ = receiveSentMessage(t, socket)
	_ = receiveSentMessage(t, socket)
	message := receiveSentMessage(t, socket)

	if message["action"] != "pong" {
		t.Fatalf("action = %v", message["action"])
	}
}

func TestDecodeCandleUsesResolutionAndTime(t *testing.T) {
	candle := decodeCandle(map[string]any{
		"data": map[string]any{
			"symbol":     "ACB",
			"resolution": "15",
			"time":       "1773657310",
			"open":       10,
			"high":       11,
			"low":        9,
			"close":      10.5,
			"volume":     1000,
		},
	})

	if candle.Resolution != "15" {
		t.Fatalf("resolution = %q", candle.Resolution)
	}
	if candle.Time != "1773657310" {
		t.Fatalf("time = %q", candle.Time)
	}
}

type fakeWebSocket struct {
	sent     chan transport.WebSocketMessage
	received chan transport.WebSocketMessage
	closed   chan struct{}
	once     sync.Once
}

func newFakeWebSocket() *fakeWebSocket {
	return &fakeWebSocket{
		sent:     make(chan transport.WebSocketMessage, 8),
		received: make(chan transport.WebSocketMessage, 8),
		closed:   make(chan struct{}),
	}
}

func (s *fakeWebSocket) Send(ctx context.Context, message transport.WebSocketMessage) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closed:
		return fmt.Errorf("websocket closed")
	case s.sent <- message:
		return nil
	}
}

func (s *fakeWebSocket) Receive(ctx context.Context) (transport.WebSocketMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closed:
		return nil, fmt.Errorf("websocket closed")
	case message := <-s.received:
		return message, nil
	}
}

func (s *fakeWebSocket) Close() error {
	s.once.Do(func() {
		close(s.closed)
	})
	return nil
}

func (s *fakeWebSocket) receiveJSON(t *testing.T, value map[string]any) {
	t.Helper()
	message, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal message: %v", err)
	}
	select {
	case s.received <- message:
	case <-time.After(time.Second):
		t.Fatalf("timed out sending fake websocket message")
	}
}

func receiveSentMessage(t *testing.T, socket *fakeWebSocket) map[string]any {
	t.Helper()
	select {
	case raw := <-socket.sent:
		message := map[string]any{}
		if err := json.Unmarshal(raw, &message); err != nil {
			t.Fatalf("decode sent message: %v", err)
		}
		return message
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for websocket send")
		return nil
	}
}
