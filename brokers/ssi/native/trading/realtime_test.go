package trading_test

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	nativetrading "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
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

func requireNoEvent[T any](t *testing.T, events <-chan T) {
	t.Helper()
	select {
	case event := <-events:
		t.Fatalf("unexpected event: %+v", event)
	case <-time.After(20 * time.Millisecond):
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
	if _, err := svc.SubscribeOrderEvents(context.Background()); err == nil {
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

	subscription, err := svc.SubscribeOrderEvents(context.Background())
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

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "orderError", map[string]any{
		"errorCode": "IGNORED",
	}))
	requireNoEvent(t, subscription.Events())

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "orderEvent", map[string]any{
		"orderID":             "26061100279501",
		"notifyID":            420825,
		"data":                nil,
		"instrumentID":        "AAA",
		"lastAction":          "NEW_ORDER",
		"uniqueID":            "60736013",
		"buySell":             "B",
		"orderType":           "LO",
		"ipAddress":           "171.225.132.124, 10.232.81.28, 10.236.229.4",
		"price":               7070.0,
		"prefix":              "dgd",
		"quantity":            1,
		"brokerId":            "",
		"marketID":            "VN",
		"origOrderId":         "26061100279501",
		"brokerIdUpdate":      nil,
		"account":             "Q712496",
		"cancelQty":           0,
		"osQty":               1,
		"filledQty":           0,
		"avgPrice":            0.0,
		"channel":             "IM",
		"inputTime":           "1781159807172",
		"modifiedTime":        "1781159807194",
		"isForceSell":         "F",
		"isShortSell":         nil,
		"orderStatus":         "QU",
		"rejectReason":        "",
		"origRequestID":       "60736013",
		"stopOrder":           false,
		"stopPrice":           0.0,
		"stopType":            "",
		"stopStep":            0.0,
		"profitPrice":         0.0,
		"modifiable":          true,
		"note":                "",
		"approveComment":      "",
		"orderApproval":       false,
		"taxRate":             0.0,
		"feeRate":             0.0035,
		"source":              "LFO",
		"lastOrderUpdateTime": "1781159807194",
		"exchangeReplyTime":   "1781159807194",
		"isCloseout":          false,
		"isOrderMM":           false,
	}))

	event := receiveEvent(t, subscription.Events())
	if event.OrderID != "26061100279501" || event.NotifyID != 420825 ||
		event.LastAction != "NEW_ORDER" || event.OSQuantity != 1 ||
		event.FeeRate != 0.0035 || event.BrokerIDUpdate != nil {
		t.Fatalf("order = %+v", event)
	}
}

func TestTradingRealtimePublishesOrderErrors(t *testing.T) {
	fake := newFakeSignalRClient()
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(string, []string) nativetrading.SignalRClient {
			return fake
		},
	})

	subscription, err := svc.SubscribeOrderErrors(context.Background())
	if err != nil {
		t.Fatalf("subscribe order errors: %v", err)
	}
	defer subscription.Close()

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "orderError", map[string]any{
		"message":      "This channel has been block; disallow to place order ",
		"notifyID":     0,
		"errorCode":    "ORD015",
		"uniqueID":     "6163422",
		"orderID":      "T20230504w3806163422",
		"instrumentID": "SSI",
		"price":        19600,
		"quantity":     200,
		"modifiable":   false,
	}))

	event := receiveEvent(t, subscription.Events())
	if event.ErrorCode != "ORD015" || event.Message == "" ||
		event.OrderID != "T20230504w3806163422" || event.Quantity != 200 {
		t.Fatalf("order error = %+v", event)
	}
}

func TestTradingRealtimePublishesOrderMatchEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(string, []string) nativetrading.SignalRClient {
			return fake
		},
	})

	subscription, err := svc.SubscribeOrderMatchEvents(context.Background())
	if err != nil {
		t.Fatalf("subscribe order match events: %v", err)
	}
	defer subscription.Close()

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "orderMatchEvent", map[string]any{
		"orderID":      "16201867",
		"notifyID":     101180,
		"instrumentID": "BVS",
		"uniqueID":     "24194396",
		"buySell":      "B",
		"matchPrice":   1000,
		"matchQty":     100,
		"matchTime":    "1656665019000",
	}))

	event := receiveEvent(t, subscription.Events())
	if event.OrderID != "16201867" || event.MatchPrice != 1000 ||
		event.MatchQuantity != 100 || event.MatchTime != "1656665019000" {
		t.Fatalf("order match = %+v", event)
	}
}

func TestTradingRealtimePublishesClientPortfolioEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	svc := nativetrading.NewRealtimeService(nativetrading.RealtimeDependencies{
		TradingToken:      func() string { return "trading-token" },
		TradingStreamURL:  "https://fc-tradehub.ssi.com.vn/v2.0/signalr",
		RequireCapability: func(core.Capability) error { return nil },
		NewSignalRClient: func(string, []string) nativetrading.SignalRClient {
			return fake
		},
	})

	subscription, err := svc.SubscribeClientPortfolioEvents(context.Background())
	if err != nil {
		t.Fatalf("subscribe positions: %v", err)
	}
	defer subscription.Close()

	fake.emit("BroadcastHubV2", "Broadcast", broadcastSSIArg(t, "clientPortfolioEvent", map[string]any{
		"account":  "0901358",
		"notifyID": 27,
		"data":     nil,
		"clientPortfoliosOpen": []map[string]any{
			{
				"martketID":    "VNFE",
				"instrumentID": "VN30F2106",
				"longQty":      9,
				"shortQty":     0,
				"net":          9,
				"bidAvgPrice":  1402.4000244140625,
				"askAvgPrice":  0,
				"tradePrice":   0,
				"marketPrice":  873,
				"floatingPL":   -476460000,
				"tradingPL":    0,
			},
		},
		"uniqueID":              nil,
		"clientPortfoliosClose": nil,
		"connectionID":          "",
		"ipAddress":             nil,
		"prefix":                nil,
	}))

	event := receiveEvent(t, subscription.Events())
	if event.Account != "0901358" || event.NotifyID != 27 || len(event.ClientPortfoliosOpen) != 1 ||
		event.ClientPortfoliosClose != nil || event.UniqueID != nil || event.IPAddress != nil {
		t.Fatalf("portfolio event = %+v", event)
	}
	position := event.ClientPortfoliosOpen[0]
	if position.MarketID != "VNFE" || position.InstrumentID != "VN30F2106" ||
		position.LongQuantity != 9 || position.Net != 9 || position.FloatingPL != -476460000 {
		t.Fatalf("portfolio = %+v", position)
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
		"notifyID":        1249008,
		"processStatus":   "INIT",
		"matchedQuantity": 0,
		"isPlaceOrder":    false,
		"instrumentID":    "ssi",
		"quantity":        4000,
		"price":           "MP",
		"account":         "1231496",
		"updatedTime":     "1751966448000",
		"status":          "200",
		"message":         "Success",
		"username":        "123149",
	}))

	event := receiveEvent(t, subscription.Events())
	if event.FCOID != "7528ac20-a340-4233-8bb7-a2379fc3c638" || event.NotifyID != 1249008 ||
		event.ProcessStatus != "INIT" || event.MatchedQuantity != 0 ||
		event.IsPlaceOrder || event.Symbol != "ssi" || event.AccountID != "1231496" ||
		event.Quantity != 4000 || event.Price != "MP" || event.Status != "200" || event.Username != "123149" {
		t.Fatalf("event = %+v", event)
	}
}
