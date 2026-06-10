package ssi

import (
	"context"
	"encoding/json"
	"strings"
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
		DataToken:    "data-token",
		TradingToken: "trading-token",
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

func TestSSITradingRealtimeRequiresTradingAccessToken(t *testing.T) {
	broker := NewBroker(Config{DataToken: "data-token"})
	if _, err := broker.Trading().Realtime().SubscribeOrders(context.Background(), trading.SubscribeOrdersRequest{}); err == nil {
		t.Fatalf("expected trading access token error")
	}
}

func TestSSIMarketDataRealtimeRequiresDataAccessToken(t *testing.T) {
	broker := NewBroker(Config{TradingToken: "trading-token"})
	if _, err := broker.Native().MarketData().Realtime().SubscribeRawChannel(context.Background(), "X:ALL"); err == nil {
		t.Fatalf("expected data access token error")
	}
}

func TestSSITradingRealtimePublishesPositions(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		TradingToken:   "trading-token",
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
		TradingToken:   "trading-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.Trading().Realtime().SubscribeFCOEvents(context.Background())
	if err != nil {
		t.Fatalf("subscribe conditional orders: %v", err)
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

func TestBuildSSIChannel(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		symbols []string
		want    string
	}{
		{name: "all", prefix: "x-trade", want: "X-TRADE:ALL"},
		{name: "single", prefix: "f", symbols: []string{" ssi "}, want: "F:SSI"},
		{name: "multiple", prefix: "mi", symbols: []string{"vn30", " HNXIndex "}, want: "MI:VN30-HNXINDEX"},
		{name: "blank values", prefix: "ol", symbols: []string{" ", ""}, want: "OL:ALL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildSSIChannel(tt.prefix, tt.symbols); got != tt.want {
				t.Fatalf("buildSSIChannel() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSSIMarketDataRealtimeTypedChannels(t *testing.T) {
	tests := []struct {
		name      string
		channel   string
		subscribe func(*Broker) error
	}{
		{
			name:    "trading status",
			channel: "F:SSI-PAN",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeTradingStatus(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", " pan "}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "quotes",
			channel: "X-QUOTE:SSI-PAN",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeQuotes(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "pan"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "trades",
			channel: "X-TRADE:SSI-PAN",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeTrades(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "pan"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "snapshots",
			channel: "X:SSI-PAN",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeSnapshots(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "pan"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "foreign rooms",
			channel: "R:SSI-PAN",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeForeignRooms(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "pan"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "market indexes",
			channel: "MI:VN30-HNXINDEX",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeMarketIndexes(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"vn30", "HNXIndex"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "ohlcv",
			channel: "B:SSI-VN30",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeOHLCV(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "vn30"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
		{
			name:    "odd lots",
			channel: "OL:SSI-VND",
			subscribe: func(broker *Broker) error {
				subscription, err := broker.Native().MarketData().Realtime().SubscribeOddLots(
					context.Background(), marketdata.SubscribeSymbolRequest{Symbols: []string{"ssi", "vnd"}},
				)
				if err == nil {
					defer subscription.Close()
				}
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := newFakeSignalRClient()
			broker := NewBroker(Config{
				DataToken:      "data-token",
				SignalRFactory: func(string, []string) SignalRClient { return fake },
			})

			if err := tt.subscribe(broker); err != nil {
				t.Fatalf("subscribe: %v", err)
			}
			if len(fake.invokes) != 1 || fake.invokes[0].hub != ssiMarketDataHub ||
				fake.invokes[0].method != "SwitchChannels" || fake.invokes[0].args[0] != tt.channel {
				t.Fatalf("invokes = %+v", fake.invokes)
			}
			if fake.headers["Authorization"] != "Bearer data-token" {
				t.Fatalf("authorization = %s", fake.headers["Authorization"])
			}
		})
	}
}

func TestSSIMarketDataRealtimePublishesTypedEvents(t *testing.T) {
	newBroker := func() (*Broker, *fakeSignalRClient) {
		fake := newFakeSignalRClient()
		broker := NewBroker(Config{
			DataToken:      "data-token",
			SignalRFactory: func(string, []string) SignalRClient { return fake },
		})
		return broker, fake
	}

	t.Run("trading status", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeTradingStatus(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "broadcast", nestedSSIArg(t, "F", map[string]any{
			"RType": "F", "Symbol": "SSI", "TradingStatus": "N",
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "F" || event.Symbol != "SSI" || event.TradingStatus != "N" {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("quote", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeQuotes(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "update", nestedSSIArg(t, "X-QUOTE", map[string]any{
			"RType": "X-QUOTE", "Symbol": "SSI", "BidPrice1": 21000, "AskPrice10": 22000,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "X-QUOTE" || event.Symbol != "SSI" || event.BidPrice1 != 21000 || event.AskPrice10 != 22000 {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("trade", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeTrades(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "broadcast", nestedSSIArg(t, "X-TRADE", map[string]any{
			"RType": "X-TRADE", "Symbol": "SSI", "LastPrice": 21500, "Side": "BU",
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "X-TRADE" || event.Symbol != "SSI" || event.LastPrice != 21500 || event.Side != "BU" {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("snapshot", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeSnapshots(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "update", nestedSSIArg(t, "X", map[string]any{
			"RType": "X", "Symbol": "SSI", "Close": 21600, "BidPrice10": 20500,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "X" || event.Symbol != "SSI" || event.Close != 21600 || event.BidPrice10 != 20500 {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("foreign room", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeForeignRooms(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "broadcast", nestedSSIArg(t, "R", map[string]any{
			"RType": "R", "Symbol": "SSI", "CurrentRoom": 806887173,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "R" || event.Symbol != "SSI" || event.CurrentRoom != 806887173 {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("market index", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeMarketIndexes(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "VN30"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "update", nestedSSIArg(t, "MI", map[string]any{
			"RType": "MI", "IndexId": "VN30", "IndexValue": 1238.76,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "MI" || event.IndexID != "VN30" || event.IndexValue != 1238.76 {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("ohlcv", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeOHLCV(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "broadcast", nestedSSIArg(t, "B", map[string]any{
			"RType": "B", "Symbol": "SSI", "TradingTime": "14:28:33", "Volume": 5000,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "B" || event.Symbol != "SSI" || event.TradingTime != "14:28:33" || event.Volume != 5000 {
			t.Fatalf("event = %+v", event)
		}
	})

	t.Run("odd lot", func(t *testing.T) {
		broker, fake := newBroker()
		subscription, err := broker.Native().MarketData().Realtime().SubscribeOddLots(
			context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
		)
		if err != nil {
			t.Fatalf("subscribe: %v", err)
		}
		defer subscription.Close()

		fake.emit(ssiMarketDataHub, "update", nestedSSIArg(t, "OL", map[string]any{
			"RType": "OL", "Symbol": "SSI", "StockNo": 2027, "LastPrice": 22750,
		}))
		event := receiveEvent(t, subscription.Events())
		if event.RType != "OL" || event.Symbol != "SSI" || event.StockNo != 2027 || event.LastPrice != 22750 {
			t.Fatalf("event = %+v", event)
		}
	})
}

func TestSSIMarketDataRealtimeReportsTypedDecodeErrors(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		DataToken:      "data-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})
	subscription, err := broker.Native().MarketData().Realtime().SubscribeQuotes(
		context.Background(), marketdata.SubscribeSymbolRequest{Symbol: "SSI"},
	)
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer subscription.Close()

	fake.emit(ssiMarketDataHub, "broadcast", nestedSSIArg(t, "X-QUOTE", map[string]any{
		"Symbol": "SSI", "BidPrice1": "invalid",
	}))

	select {
	case err := <-subscription.Errors():
		if !strings.Contains(err.Error(), "ssi realtime decode FcMarketDataV2Hub.broadcast") {
			t.Fatalf("error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for decode error")
	}

	select {
	case event := <-subscription.Events():
		t.Fatalf("unexpected event = %+v", event)
	case <-time.After(25 * time.Millisecond):
	}
}

func TestSSIMarketDataRealtimePublishesRawChannel(t *testing.T) {
	fake := newFakeSignalRClient()
	broker := NewBroker(Config{
		DataToken:      "data-token",
		SignalRFactory: func(string, []string) SignalRClient { return fake },
	})

	subscription, err := broker.Native().MarketData().Realtime().SubscribeRawChannel(context.Background(), "X:ALL")
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
