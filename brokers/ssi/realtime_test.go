package ssi

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/marketdata"
	"github.com/vnbrokers/vnbrokers-go/trading"
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

func TestSSITradingRealtimePublishesOrderEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	var baseURL string
	var hubs []string
	broker := NewBroker(Config{
		AccessToken: "access-token",
		SignalRFactory: func(url string, names []string) SignalRClient {
			baseURL = url
			hubs = names
			return fake
		},
	})

	subscription, err := broker.Trading().Realtime().SubscribeOrders(
		context.Background(),
		trading.SubscribeOrdersRequest{},
	)
	if err != nil {
		t.Fatalf("subscribe orders: %v", err)
	}
	defer subscription.Close()

	if baseURL != "https://fc-tradehub.ssi.com.vn/v2.0/signalr" || len(hubs) != 1 || hubs[0] != "BroadcastHubV2" {
		t.Fatalf("connection = %s %+v", baseURL, hubs)
	}
	if fake.headers["Authorization"] != "Bearer access-token" || fake.query["notify_id"] != "-1" {
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

func TestSSIRealtimeRequiresAccessToken(t *testing.T) {
	broker := NewBroker(Config{})
	if _, err := broker.MarketData().Realtime().SubscribeRawChannel(context.Background(), "X:ALL"); err == nil {
		t.Fatalf("expected access token error")
	}
}

func TestSSITradingRealtimePublishesPositions(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		AccessToken:    "access-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.Trading().Realtime().SubscribePositions(
		context.Background(),
		trading.SubscribePositionsRequest{},
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

func TestSSITradingRealtimePublishesFCOEvents(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		AccessToken:    "access-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.Trading().Realtime().SubscribeFCOEvents(context.Background())
	if err != nil {
		t.Fatalf("subscribe conditional orders: %v", err)
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
		event.ProcessStatus != "INIT" || event.MatchedQuantity.String() != "0" ||
		event.IsPlaceOrder || event.Symbol != "ssi" || event.AccountID != "1231496" ||
		event.Quantity.String() != "4000" || event.Price != "MP" {
		t.Fatalf("event = %+v", event)
	}
}

func TestSSIMarketDataRealtimeSwitchesChannelAndPublishesTopPrice(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		AccessToken:    "access-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.MarketData().Realtime().SubscribeTopPrice(
		context.Background(),
		marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
	)
	if err != nil {
		t.Fatalf("subscribe top price: %v", err)
	}
	defer subscription.Close()

	if len(fake.invokes) != 1 || fake.invokes[0].hub != "FcMarketDataV2Hub" ||
		fake.invokes[0].method != "SwitchChannels" || fake.invokes[0].args[0] != "X:SSI" {
		t.Fatalf("invokes = %+v", fake.invokes)
	}

	fake.emit("FcMarketDataV2Hub", "broadcast", nestedSSIArg(t, "X", map[string]any{
		"Symbol":    "SSI",
		"BidPrice1": 21000,
		"BidVol1":   100,
		"AskPrice1": 21100,
		"AskVol1":   200,
		"Time":      "14:45:00",
	}))

	top := receiveEvent(t, subscription.Events())
	if top.Symbol != "SSI" || top.BidPrice.String() != "21000" || top.AskPrice.String() != "21100" {
		t.Fatalf("top price = %+v", top)
	}
}

func TestSSIMarketDataRealtimePublishesRawChannel(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		AccessToken:    "access-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.MarketData().Realtime().SubscribeRawChannel(context.Background(), "X:ALL")
	if err != nil {
		t.Fatalf("subscribe raw: %v", err)
	}
	defer subscription.Close()

	fake.emit("FcMarketDataV2Hub", "update", nestedSSIArg(t, "X", map[string]any{"Symbol": "SSI"}))
	payload := receiveEvent(t, subscription.Events())
	data, ok := payload.Data.(map[string]any)
	if !ok || data["Symbol"] != "SSI" {
		t.Fatalf("payload = %+v", payload)
	}
}

func nestedSSIArg(t *testing.T, dataType string, content map[string]any) json.RawMessage {
	t.Helper()
	contentJSON, err := json.Marshal(content)
	if err != nil {
		t.Fatalf("marshal content: %v", err)
	}
	envelopeJSON, err := json.Marshal(map[string]any{
		"DataType": dataType,
		"Content":  string(contentJSON),
	})
	if err != nil {
		t.Fatalf("marshal envelope: %v", err)
	}
	arg, err := json.Marshal(string(envelopeJSON))
	if err != nil {
		t.Fatalf("marshal arg: %v", err)
	}
	return arg
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
