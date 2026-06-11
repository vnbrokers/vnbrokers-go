package trading_test

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
)

type fakeSignalRClient struct {
	mu       sync.Mutex
	handlers map[string]map[string]func([]json.RawMessage)
	onError  func(error)
	headers  map[string]string
	query    map[string]string
	invokes  []fakeSignalRInvoke
	closed   bool
}

type fakeSignalRInvoke struct {
	hub    string
	method string
	args   []any
}

func newFakeSignalRClient() *fakeSignalRClient {
	return &fakeSignalRClient{
		handlers: map[string]map[string]func([]json.RawMessage){},
		headers:  map[string]string{},
		query:    map[string]string{},
	}
}

func (c *fakeSignalRClient) Connect(context.Context) error { return nil }
func (c *fakeSignalRClient) Invoke(hub string, method string, args ...any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.invokes = append(c.invokes, fakeSignalRInvoke{hub: hub, method: method, args: args})
	return nil
}
func (c *fakeSignalRClient) On(hub string, method string, fn func([]json.RawMessage)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.handlers[hub] == nil {
		c.handlers[hub] = map[string]func([]json.RawMessage){}
	}
	c.handlers[hub][method] = fn
}
func (c *fakeSignalRClient) OnError(fn func(error))             { c.onError = fn }
func (c *fakeSignalRClient) SetHeader(key string, value string) { c.headers[key] = value }
func (c *fakeSignalRClient) SetQuery(key string, value string)  { c.query[key] = value }
func (c *fakeSignalRClient) Close() error                       { c.closed = true; return nil }
func (c *fakeSignalRClient) emit(hub string, method string, args ...json.RawMessage) {
	c.mu.Lock()
	fn := c.handlers[hub][method]
	c.mu.Unlock()
	if fn != nil {
		fn(args)
	}
}

func broadcastSSIArg(t *testing.T, eventType string, data map[string]any) json.RawMessage {
	t.Helper()
	payload, err := json.Marshal(map[string]any{
		"type": eventType,
		"data": data,
	})
	if err != nil {
		t.Fatalf("marshal broadcast payload: %v", err)
	}
	arg, err := json.Marshal(string(payload))
	if err != nil {
		t.Fatalf("marshal broadcast arg: %v", err)
	}
	return arg
}

func receiveEvent[T any](t *testing.T, events <-chan T) T {
	t.Helper()
	select {
	case event := <-events:
		return event
	case <-time.After(time.Second):
		var zero T
		t.Fatalf("timed out waiting for event")
		return zero
	}
}

func TestTradingRealtimeRequiresTradingAccessToken(t *testing.T) {
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(baseURL string, hubs []string) nativetrading.SignalRClient {
			return newFakeSignalRClient()
		},
	})
	if _, err := svc.SubscribeOrders(context.Background(), sdktrading.SubscribeOrdersRequest{}); err == nil {
		t.Fatalf("expected trading access token error")
	}
}

func TestTradingRealtimePublishesOrderEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	var baseURL string
	var hubs []string
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(url string, names []string) nativetrading.SignalRClient {
			baseURL = url
			hubs = names
			return fake
		},
	})

	subscription, err := svc.SubscribeOrders(
		context.Background(),
		sdktrading.SubscribeOrdersRequest{},
	)
	if err != nil {
		t.Fatalf("subscribe orders: %v", err)
	}
	defer subscription.Close()

	if baseURL != "https://fc-tradehub.ssi.com.vn/v2.0/signalr" || len(hubs) != 1 || hubs[0] != "BroadcastHubV2" {
		t.Fatalf("connection = %s %+v", baseURL, hubs)
	}
	if fake.headers["Authorization"] != "Bearer trading-token" || fake.query["notify_id"] != "-1" {
		t.Fatalf("auth = %+v query = %+v", fake.headers, fake.query)
	}

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "orderMatchEvent", map[string]any{
		"Account":      "0901351",
		"OrderID":      "26060500251341",
		"InstrumentID": "SSI",
		"OrderStatus":  "PF",
		"FilledQty":    50,
		"ModifiedTime": "1780633829138",
	}))

	event := receiveEvent(t, subscription.Events())
	if event.Broker != "ssi" || event.AccountID != "0901351" || event.OrderID != "26060500251341" ||
		event.Symbol != "SSI" || event.Status != domain.OrderStatusPartiallyFilled || event.FilledQuantity != "50" {
		t.Fatalf("event = %+v", event)
	}
}

func TestTradingRealtimePublishesPositions(t *testing.T) {
	fake := newFakeSignalRClient()
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(string, []string) nativetrading.SignalRClient {
			return fake
		},
	})

	subscription, err := svc.SubscribePositions(
		context.Background(),
		sdktrading.SubscribePositionsRequest{},
	)
	if err != nil {
		t.Fatalf("subscribe positions: %v", err)
	}
	defer subscription.Close()

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "clientPortfolioEvent", map[string]any{
		"Account":      "0901351",
		"InstrumentID": "SSI",
		"OnHand":       100,
		"SellableQty":  80,
		"AvgPrice":     21000,
		"MarketPrice":  22000,
	}))

	position := receiveEvent(t, subscription.Events())
	if position.AccountID != "0901351" || position.Symbol != "SSI" ||
		position.Quantity.String() != "100" || position.AvailableQuantity.String() != "80" {
		t.Fatalf("position = %+v", position)
	}
}

func TestTradingRealtimePublishesFCOEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(string, []string) nativetrading.SignalRClient {
			return fake
		},
	})

	subscription, err := svc.SubscribeFCOEvents(context.Background())
	if err != nil {
		t.Fatalf("subscribe fco events: %v", err)
	}
	defer subscription.Close()

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "fcoEvent", map[string]any{
		"fcoId":           "7528ac20-a340-4233-8bb7-a2379fc3c638",
		"notifyID": 6840196,
		"processStatus":   "INIT",
		"matchedQuantity": 0,
		"isPlaceOrder":    false,
		"instrumentID":    "ssi",
		"quantity":        4000,
		"price":           "MP",
		"account":         "",
		"updatedTime":     "1751966448000",
		"status":          "200",
		"message":         "Success",
		"username":        "123149",
	}))

	event := receiveEvent(t, subscription.Events())
	if event.FCOID != "7528ac20-a340-4233-8bb7-a2379fc3c638" || event.NotifyID != 6840196 ||
		event.ProcessStatus != "INIT" || event.MatchedQuantity.String() != "0" ||
		event.IsPlaceOrder || event.Symbol != "ssi" || event.AccountID != "" ||
		event.Quantity.String() != "4000" || event.Price != "MP" {
		t.Fatalf("event = %+v", event)
	}
}
