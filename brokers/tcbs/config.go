package tcbs

import (
	"context"
	"net/http"

	"github.com/coder/websocket"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	ProductionBaseURL = "https://openapi.tcbs.com.vn"
	SITBaseURL        = "https://openapisit.tcbs.com.vn"
)

type Config struct {
	BaseURL          string
	AccessToken      string
	HTTPClient       *http.Client
	HTTPTransport    transport.HTTPTransport
	WebSocketFactory TCBSWebSocketFactory
}

type TCBSWebSocketFactory func(context.Context, string, map[string]string) (transport.WebSocketTransport, error)

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = ProductionBaseURL
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	if c.WebSocketFactory == nil {
		c.WebSocketFactory = ConnectTCBSWebSocket
	}
	return c
}

func ConnectTCBSWebSocket(
	ctx context.Context,
	url string,
	headers map[string]string,
) (transport.WebSocketTransport, error) {
	httpHeaders := http.Header{}
	for key, value := range headers {
		httpHeaders.Set(key, value)
	}
	conn, _, err := websocket.Dial(ctx, url, &websocket.DialOptions{HTTPHeader: httpHeaders})
	if err != nil {
		return nil, err
	}
	return &tcbsWebSocketTransport{conn: conn}, nil
}

type tcbsWebSocketTransport struct {
	conn *websocket.Conn
}

func (t *tcbsWebSocketTransport) Send(ctx context.Context, message transport.WebSocketMessage) error {
	return t.conn.Write(ctx, websocket.MessageText, message)
}

func (t *tcbsWebSocketTransport) Receive(ctx context.Context) (transport.WebSocketMessage, error) {
	_, payload, err := t.conn.Read(ctx)
	return payload, err
}

func (t *tcbsWebSocketTransport) Close() error {
	return t.conn.Close(websocket.StatusNormalClosure, "")
}
