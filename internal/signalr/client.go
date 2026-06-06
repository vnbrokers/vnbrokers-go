package signalr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/coder/websocket"
)

type Client struct {
	BaseURL        string
	ClientProtocol string
	Hubs           []string

	Headers    http.Header
	Query      url.Values
	HTTPClient *http.Client

	connToken string
	connID    string
	ws        *websocket.Conn

	mu       sync.Mutex
	msgID    int
	handlers map[string]map[string]func(args []json.RawMessage)
	onError  func(error)
}

type hubData struct {
	Name string `json:"name"`
}

type negotiateResponse struct {
	ConnectionID            string  `json:"ConnectionId"`
	ConnectionToken         string  `json:"ConnectionToken"`
	ProtocolVersion         string  `json:"ProtocolVersion"`
	TryWebSockets           bool    `json:"TryWebSockets"`
	KeepAliveTimeout        float64 `json:"KeepAliveTimeout"`
	DisconnectTimeout       float64 `json:"DisconnectTimeout"`
	TransportConnectTimeout float64 `json:"TransportConnectTimeout"`
}

type hubInvoke struct {
	Hub    string `json:"H"`
	Method string `json:"M"`
	Args   []any  `json:"A"`
	ID     int    `json:"I"`
}

type serverMessage struct {
	C string `json:"C"`
	S int    `json:"S"`
	M []struct {
		Hub    string            `json:"H"`
		Method string            `json:"M"`
		Args   []json.RawMessage `json:"A"`
	} `json:"M"`

	ID     string          `json:"I"`
	Result json.RawMessage `json:"R"`
	Error  string          `json:"E"`
}

func NewClient(baseURL string, hubs []string) *Client {
	return &Client{
		BaseURL:        strings.TrimRight(baseURL, "/"),
		ClientProtocol: "2.1",
		Hubs:           hubs,
		Headers:        make(http.Header),
		Query:          make(url.Values),
		HTTPClient:     &http.Client{Timeout: 15 * time.Second},
		handlers:       make(map[string]map[string]func(args []json.RawMessage)),
	}
}

func (c *Client) cleanHubData() []hubData {
	out := make([]hubData, 0, len(c.Hubs))
	for _, hub := range c.Hubs {
		hub = strings.TrimSpace(hub)
		if hub != "" {
			out = append(out, hubData{Name: strings.ToLower(hub)})
		}
	}
	return out
}

func (c *Client) connectionData() (string, error) {
	data, err := json.Marshal(c.cleanHubData())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *Client) Negotiate(ctx context.Context) error {
	connectionData, err := c.connectionData()
	if err != nil {
		return err
	}

	u := c.BaseURL + "/negotiate?" + c.signalRQuery(connectionData, nil).Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header = c.Headers.Clone()

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusFound {
		return fmt.Errorf("signalr negotiate unauthorized: status=%d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("signalr negotiate failed: status=%d", resp.StatusCode)
	}

	var nr negotiateResponse
	if err := json.NewDecoder(resp.Body).Decode(&nr); err != nil {
		return err
	}
	if !nr.TryWebSockets {
		return fmt.Errorf("server does not allow websockets")
	}

	c.mu.Lock()
	c.connToken = nr.ConnectionToken
	c.connID = nr.ConnectionID
	c.mu.Unlock()
	return nil
}

func (c *Client) Connect(ctx context.Context) error {
	c.mu.Lock()
	hasToken := c.connToken != ""
	c.mu.Unlock()
	if !hasToken {
		if err := c.Negotiate(ctx); err != nil {
			return err
		}
	}

	connectionData, err := c.connectionData()
	if err != nil {
		return err
	}

	c.mu.Lock()
	connToken := c.connToken
	c.mu.Unlock()
	params := c.signalRQuery(connectionData, map[string]string{
		"transport":       "webSockets",
		"connectionToken": connToken,
		"tid":             "10",
	})

	wsURL, err := toWebSocketURL(c.BaseURL + "/connect?" + params.Encode())
	if err != nil {
		return err
	}

	ws, resp, err := websocket.Dial(ctx, wsURL, &websocket.DialOptions{HTTPHeader: c.Headers})
	if err != nil {
		if resp != nil {
			return fmt.Errorf("websocket connect failed: status=%d err=%w", resp.StatusCode, err)
		}
		return err
	}

	c.mu.Lock()
	c.ws = ws
	c.mu.Unlock()

	if err := c.Start(ctx); err != nil {
		_ = ws.Close(websocket.StatusNormalClosure, "")
		c.mu.Lock()
		if c.ws == ws {
			c.ws = nil
		}
		c.mu.Unlock()
		return err
	}

	go c.readLoop(ws)
	return nil
}

func (c *Client) Start(ctx context.Context) error {
	connectionData, err := c.connectionData()
	if err != nil {
		return err
	}

	c.mu.Lock()
	connToken := c.connToken
	c.mu.Unlock()
	params := c.signalRQuery(connectionData, map[string]string{
		"transport":       "webSockets",
		"connectionToken": connToken,
	})

	u := c.BaseURL + "/start?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header = c.Headers.Clone()

	resp, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusFound {
		return fmt.Errorf("signalr start unauthorized: status=%d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("signalr start failed: status=%d", resp.StatusCode)
	}
	return nil
}

func (c *Client) Invoke(hub string, method string, args ...any) error {
	c.mu.Lock()
	c.msgID++
	id := c.msgID
	ws := c.ws
	c.mu.Unlock()

	if ws == nil {
		return fmt.Errorf("websocket is not connected")
	}

	payload := hubInvoke{
		Hub:    hub,
		Method: method,
		Args:   args,
		ID:     id,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return ws.Write(context.Background(), websocket.MessageText, data)
}

func (c *Client) On(hub string, method string, fn func(args []json.RawMessage)) {
	hub = strings.ToLower(hub)
	method = strings.ToLower(method)

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.handlers[hub] == nil {
		c.handlers[hub] = make(map[string]func(args []json.RawMessage))
	}
	c.handlers[hub][method] = fn
}

func (c *Client) OnError(fn func(error)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.onError = fn
}

func (c *Client) SetHeader(key string, value string) {
	c.Headers.Set(key, value)
}

func (c *Client) SetQuery(key string, value string) {
	c.Query.Set(key, value)
}

func (c *Client) Close() error {
	c.mu.Lock()
	ws := c.ws
	c.ws = nil
	c.mu.Unlock()

	if ws == nil {
		return nil
	}
	return ws.Close(websocket.StatusNormalClosure, "")
}

func (c *Client) readLoop(ws *websocket.Conn) {
	for {
		_, data, err := ws.Read(context.Background())
		if err != nil {
			c.mu.Lock()
			fn := c.onError
			c.mu.Unlock()
			if fn != nil {
				fn(err)
			}
			return
		}
		c.dispatch(data)
	}
}

func (c *Client) dispatch(data []byte) {
	if string(data) == "{}" {
		return
	}

	var msg serverMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return
	}

	for _, message := range msg.M {
		hub := strings.ToLower(message.Hub)
		method := strings.ToLower(message.Method)

		c.mu.Lock()
		fn := c.handlers[hub][method]
		c.mu.Unlock()

		if fn != nil {
			fn(message.Args)
		}
	}
}

func (c *Client) signalRQuery(connectionData string, extra map[string]string) url.Values {
	query := url.Values{}
	query.Set("clientProtocol", c.ClientProtocol)
	query.Set("connectionData", connectionData)
	for key, values := range c.Query {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	for key, value := range extra {
		query.Set(key, value)
	}
	return query
}

func (c *Client) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}

func toWebSocketURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	switch u.Scheme {
	case "http":
		u.Scheme = "ws"
	case "https":
		u.Scheme = "wss"
	case "ws", "wss":
	default:
		return "", fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}

	return u.String(), nil
}
