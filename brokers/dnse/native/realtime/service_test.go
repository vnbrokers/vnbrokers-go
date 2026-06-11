package realtime

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestBuildAuthMessageUsesProvidedValues(t *testing.T) {
	message := BuildAuthMessage("key", "secret", 123, "nonce")
	if message["api_key"] != "key" || message["timestamp"] != int64(123) || message["nonce"] != "nonce" {
		t.Fatalf("message = %#v", message)
	}
}

func TestSubscribeRespondsToServerPing(t *testing.T) {
	socket := &fakeWebSocket{sent: make(chan []byte, 4), received: make(chan []byte, 1), closed: make(chan struct{})}
	subscription, err := Subscribe[map[string]any](context.Background(), Dependencies{
		APIKey: "key", APISecret: "secret", StreamURL: "wss://example.test/v1/stream", Encoding: "json", PongInterval: time.Hour,
		WebSocketFactory: func(context.Context, string) (transport.WebSocketTransport, error) { return socket, nil },
	}, map[string]any{"action": "subscribe"}, func(map[string]any) bool { return false }, func(message map[string]any) (map[string]any, error) { return message, nil })
	if err != nil {
		t.Fatal(err)
	}
	defer subscription.Close()
	<-socket.sent
	<-socket.sent
	socket.received <- []byte(`{"action":"ping"}`)
	select {
	case payload := <-socket.sent:
		message, err := decodeMessage(payload, "json")
		if err != nil {
			t.Fatal(err)
		}
		if message["action"] != "pong" {
			t.Fatalf("message = %#v", message)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for pong")
	}
}

type fakeWebSocket struct {
	sent, received chan []byte
	closed         chan struct{}
	once           sync.Once
}

func (s *fakeWebSocket) Send(ctx context.Context, message transport.WebSocketMessage) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-s.closed:
		return fmt.Errorf("closed")
	case s.sent <- message:
		return nil
	}
}
func (s *fakeWebSocket) Receive(ctx context.Context) (transport.WebSocketMessage, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-s.closed:
		return nil, fmt.Errorf("closed")
	case message := <-s.received:
		return message, nil
	}
}
func (s *fakeWebSocket) Close() error { s.once.Do(func() { close(s.closed) }); return nil }

func TestDependenciesStreamURLUsesEncoding(t *testing.T) {
	dependencies := Dependencies{StreamURL: "wss://example.test/v1/stream?encoding=json", Encoding: "msgpack", PongInterval: time.Second}
	url, err := dependencies.URL()
	if err != nil {
		t.Fatal(err)
	}
	if url != "wss://example.test/v1/stream?encoding=msgpack" {
		t.Fatalf("url = %q", url)
	}
}
