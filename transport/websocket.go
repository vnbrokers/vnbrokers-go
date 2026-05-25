package transport

import (
	"context"

	"github.com/coder/websocket"
)

type WebSocketMessage []byte

type WebSocketTransport interface {
	Send(context.Context, WebSocketMessage) error
	Receive(context.Context) (WebSocketMessage, error)
	Close() error
}

type WebSocketFactory func(context.Context, string) (WebSocketTransport, error)

type CoderWebSocketTransport struct {
	conn *websocket.Conn
}

func ConnectWebSocket(ctx context.Context, url string) (WebSocketTransport, error) {
	conn, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return &CoderWebSocketTransport{conn: conn}, nil
}

func (t *CoderWebSocketTransport) Send(ctx context.Context, message WebSocketMessage) error {
	return t.conn.Write(ctx, websocket.MessageBinary, message)
}

func (t *CoderWebSocketTransport) Receive(ctx context.Context) (WebSocketMessage, error) {
	_, payload, err := t.conn.Read(ctx)
	return payload, err
}

func (t *CoderWebSocketTransport) Close() error {
	return t.conn.Close(websocket.StatusNormalClosure, "")
}
