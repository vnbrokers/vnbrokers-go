package tcbs

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestPlaceStockOrderBuildsTCBSRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"orderId": "OID-1",
				"message": "ok",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Trading().Orders().Place(context.Background(), "0001170730", PlaceOrderRequest{
		ExecType:  "NB",
		Price:     23500,
		PriceType: "LO",
		Quantity:  100,
		Symbol:    "ACB",
	})
	if err != nil {
		t.Fatalf("place order: %v", err)
	}
	if response.OrderID != "OID-1" {
		t.Fatalf("order id = %s", response.OrderID)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://api.tcbs.example/akhlys/v1/accounts/0001170730/orders" {
		t.Fatalf("url = %s", request.URL)
	}
	if request.Headers["Authorization"] != "Bearer jwt-123" {
		t.Fatalf("authorization = %s", request.Headers["Authorization"])
	}
	body, ok := request.JSON.(PlaceOrderRequest)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if body.Symbol != "ACB" || body.Price != 23500 || body.Quantity != 100 {
		t.Fatalf("body = %#v", body)
	}
}

func TestUpdateAndCancelStockOrdersBuildTCBSRequests(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{
				StatusCode: 200,
				Body: map[string]any{
					"orderId": "OID-1",
					"message": "updated",
				},
			},
			{
				StatusCode: 200,
				Body: map[string]any{
					"data": []any{},
				},
			},
		},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Trading().Orders().Update(context.Background(), "0001170730", "OID-1", UpdateOrderRequest{
		Price:    23600,
		Quantity: 200,
	})
	if err != nil {
		t.Fatalf("update order: %v", err)
	}
	_, err = broker.Trading().Orders().Cancel(context.Background(), "0001170730", CancelOrderRequest{
		OrdersList: []OrderIDRef{{OrderID: "OID-1"}},
	})
	if err != nil {
		t.Fatalf("cancel order: %v", err)
	}

	if got := httpTransport.requests[0].URL; got != "https://api.tcbs.example/akhlys/v1/accounts/0001170730/orders/OID-1" {
		t.Fatalf("update url = %s", got)
	}
	if httpTransport.requests[0].Method != "PUT" {
		t.Fatalf("update method = %s", httpTransport.requests[0].Method)
	}
	if got := httpTransport.requests[1].URL; got != "https://api.tcbs.example/akhlys/v1/accounts/0001170730/cancel-orders" {
		t.Fatalf("cancel url = %s", got)
	}
	if httpTransport.requests[1].Method != "PUT" {
		t.Fatalf("cancel method = %s", httpTransport.requests[1].Method)
	}
}

func TestTradingAccountsQueryBuildsTCBSRequests(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{StatusCode: 200, Body: map[string]any{"data": []any{map[string]any{"orderID": "OID-1"}}}},
			{StatusCode: 200, Body: map[string]any{"accountNo": "0001170730", "ppse": 1000000}},
			{StatusCode: 200, Body: map[string]any{"accountNo": "0001170730", "stock": []any{}}},
			{StatusCode: 200, Body: map[string]any{"data": []any{map[string]any{"accountNo": "0001170730", "cashBalance": 1200000}}}},
		},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	if _, err := broker.Trading().Accounts().Orders(context.Background(), "0001170730"); err != nil {
		t.Fatalf("orders: %v", err)
	}
	if _, err := broker.Trading().Accounts().PurchasingPowerBySymbolPrice(context.Background(), "0001170730", "ACB", "23500"); err != nil {
		t.Fatalf("ppse by symbol price: %v", err)
	}
	if _, err := broker.Trading().Accounts().StockAssets(context.Background(), "0001170730"); err != nil {
		t.Fatalf("stock assets: %v", err)
	}
	if _, err := broker.Trading().Accounts().CashBalance(context.Background(), "0001170730"); err != nil {
		t.Fatalf("cash balance: %v", err)
	}

	expectedURLs := []string{
		"https://api.tcbs.example/aion/v1/accounts/0001170730/orders",
		"https://api.tcbs.example/aion/v1/accounts/0001170730/ppse/ACB/23500",
		"https://api.tcbs.example/aion/v1/accounts/0001170730/se",
		"https://api.tcbs.example/aion/v1/accounts/0001170730/cashInvestments",
	}
	for i, expected := range expectedURLs {
		if httpTransport.requests[i].URL != expected {
			t.Fatalf("request %d url = %s", i, httpTransport.requests[i].URL)
		}
		if httpTransport.requests[i].Headers["Authorization"] != "Bearer jwt-123" {
			t.Fatalf("request %d authorization = %s", i, httpTransport.requests[i].Headers["Authorization"])
		}
	}
}

func TestPurchasingPowerBySymbolPriceDecodesTCBSMixedTypes(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"object":             "ppse",
				"rateBrkS":           0.03,
				"rateBrkB":           0.15,
				"accountNo":          "",
				"custodyID":          nil,
				"symbol":             "",
				"price":              0,
				"pp0":                100000,
				"ppse":               100000,
				"availableTrade":     0,
				"ppseref":            100000,
				"maxBuyQuantity":     0,
				"realMaxBuyQuantity": 0,
				"minBuyQuantity":     0,
				"marginRatioLoan":    "0",
				"marginPriceLoan":    "0",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Trading().Accounts().PurchasingPowerBySymbolPrice(
		context.Background(),
		"0001170730",
		"ACB",
		"23500",
	)
	if err != nil {
		t.Fatalf("ppse by symbol price: %v", err)
	}
	if response.Object != "ppse" {
		t.Fatalf("object = %s", response.Object)
	}
	if response.PP0 != 100000 {
		t.Fatalf("pp0 = %v", response.PP0)
	}
	if response.RateBrkS != 0.03 {
		t.Fatalf("rateBrkS = %v", response.RateBrkS)
	}
	if response.RateBrkB != 0.15 {
		t.Fatalf("rateBrkB = %v", response.RateBrkB)
	}
	if response.MarginRatioLoan != "0" {
		t.Fatalf("marginRatioLoan = %s", response.MarginRatioLoan)
	}
	if response.MarginPriceLoan != "0" {
		t.Fatalf("marginPriceLoan = %s", response.MarginPriceLoan)
	}
}

func TestTCBSTradingMappers(t *testing.T) {
	order := MapOrderInfo(OrderInfo{
		AccountNo:  "0001170730",
		OrderID:    "OID-1",
		ExecType:   "NB",
		OrderQtty:  100,
		OrStatus:   "Matched",
		PriceType:  "LO",
		Symbol:     "ACB",
		LimitPrice: 23500,
	})
	if order.Broker != "tcbs" || order.AccountID != "0001170730" || order.OrderID != "OID-1" {
		t.Fatalf("order = %#v", order)
	}
	if order.Side != domain.OrderSideBuy {
		t.Fatalf("side = %s", order.Side)
	}
	if order.Status != domain.OrderStatusFilled {
		t.Fatalf("status = %s", order.Status)
	}
	if !order.Quantity.Equal(decimal.NewFromInt(100)) {
		t.Fatalf("quantity = %s", order.Quantity)
	}

	positions := MapStockHoldings("0001170730", SeInfoDTO{
		Stock: []StockHoldingInfo{{
			Symbol:           "ACB",
			TotalQtty:        300,
			AvailableTrading: 200,
			CurrentPrice:     23500,
		}},
	})
	if len(positions) != 1 || positions[0].Symbol != "ACB" {
		t.Fatalf("positions = %#v", positions)
	}
	if !positions[0].Quantity.Equal(decimal.NewFromInt(300)) {
		t.Fatalf("position quantity = %s", positions[0].Quantity)
	}

	balance := MapCashBalance("0001170730", CashInvestmentResponse{
		Data: []CashInvestment{{
			AccountNo:   "0001170730",
			CashBalance: 1200000,
			PP0:         1000000,
		}},
	})
	if balance.AccountID != "0001170730" {
		t.Fatalf("balance account = %s", balance.AccountID)
	}
	if balance.CashTotal == nil || !balance.CashTotal.Equal(decimal.NewFromInt(1200000)) {
		t.Fatalf("cash total = %v", balance.CashTotal)
	}
	if balance.BuyingPower == nil || !balance.BuyingPower.Equal(decimal.NewFromInt(1000000)) {
		t.Fatalf("buying power = %v", balance.BuyingPower)
	}
}

type fakeTCBSWebSocketFactory struct {
	url     string
	headers map[string]string
	socket  *fakeWebSocketTransport
}

func (f *fakeTCBSWebSocketFactory) Connect(_ context.Context, url string, headers map[string]string) (transport.WebSocketTransport, error) {
	f.url = url
	f.headers = headers
	return f.socket, nil
}

type fakeWebSocketTransport struct {
	received []transport.WebSocketMessage
	closed   bool
}

func (f *fakeWebSocketTransport) Send(context.Context, transport.WebSocketMessage) error {
	return nil
}

func (f *fakeWebSocketTransport) Receive(context.Context) (transport.WebSocketMessage, error) {
	if len(f.received) == 0 {
		return nil, errors.New("closed")
	}
	message := f.received[0]
	f.received = f.received[1:]
	return message, nil
}

func (f *fakeWebSocketTransport) Close() error {
	f.closed = true
	return nil
}

func TestSubscribeStockMatchesConnectsAndPublishesOrderEvents(t *testing.T) {
	socket := &fakeWebSocketTransport{
		received: []transport.WebSocketMessage{
			[]byte(`{"accountNo":"0001170730","orderID":"OID-1","symbol":"ACB","side":"B","orStatus":"Matched","execQtty":100,"txtime":"09:30:00"}`),
		},
	}
	factory := &fakeTCBSWebSocketFactory{socket: socket}
	broker := NewBroker(Config{
		BaseURL:          "https://api.tcbs.example",
		AccessToken:      "jwt-123",
		WebSocketFactory: factory.Connect,
	})

	subscription, err := broker.Trading().Realtime().SubscribeStockMatches(context.Background())
	if err != nil {
		t.Fatalf("subscribe stock matches: %v", err)
	}
	defer subscription.Close()

	if factory.url != "wss://api.tcbs.example/ws/aither" {
		t.Fatalf("url = %s", factory.url)
	}
	if factory.headers["Authorization"] != "Bearer jwt-123" {
		t.Fatalf("authorization = %s", factory.headers["Authorization"])
	}

	select {
	case event := <-subscription.Events():
		if event.Broker != "tcbs" || event.AccountID != "0001170730" || event.OrderID != "OID-1" {
			t.Fatalf("event = %#v", event)
		}
		if event.Status != domain.OrderStatusFilled {
			t.Fatalf("status = %s", event.Status)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for order event")
	}

	select {
	case status := <-subscription.Status():
		if status != realtime.StatusConnecting && status != realtime.StatusConnected && status != realtime.StatusSubscribed && status != realtime.StatusClosed {
			t.Fatalf("unexpected status = %s", status)
		}
	default:
	}
}
