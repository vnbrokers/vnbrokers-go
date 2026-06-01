package dnse

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeHTTPTransport struct {
	requests  []transport.HTTPRequest
	responses []transport.HTTPResponse
}

func (f *fakeHTTPTransport) Send(_ context.Context, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	f.requests = append(f.requests, request)
	response := f.responses[0]
	f.responses = f.responses[1:]
	return response, nil
}

func TestListPositionsBuildsDNSERequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"positions": []any{
					map[string]any{
						"accountNo":    "0001179019",
						"symbol":       "VN30F2506",
						"openQuantity": 3,
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:           "https://api.dnse.example",
		APIKey:            "key",
		MarketType:        "STOCK",
		PositionsPageSize: 50,
		HTTPTransport:     httpTransport,
	})

	positions, err := broker.Trading().Positions().List(context.Background(), "0001179019")
	if err != nil {
		t.Fatalf("list positions: %v", err)
	}
	if got := httpTransport.requests[0].URL; got != "https://api.dnse.example/accounts/0001179019/positions?marketType=STOCK&pageSize=50" {
		t.Fatalf("url = %s", got)
	}
	if positions[0].Symbol != "VN30F2506" {
		t.Fatalf("symbol = %s", positions[0].Symbol)
	}
}

func TestReplaceOrderBuildsDNSERequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"id":       1626,
				"quantity": 3,
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.dnse.example",
		TradingToken:  "trading-token",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Trading().Orders().Replace(context.Background(), sdktrading.ReplaceOrderRequest{
		AccountID:     "0001179019",
		OrderID:       "1626",
		Price:         decimal.RequireFromString("1851.5"),
		Quantity:      3,
		MarketType:    sdktrading.MarketTypeDerivative,
		OrderCategory: sdktrading.OrderCategoryNormal,
	})
	if err != nil {
		t.Fatalf("replace order: %v", err)
	}
	request := httpTransport.requests[0]
	if request.Method != "PUT" {
		t.Fatalf("method = %s", request.Method)
	}
	if got := request.URL; got != "https://api.dnse.example/accounts/0001179019/orders/1626?marketType=DERIVATIVE&orderCategory=NORMAL" {
		t.Fatalf("url = %s", got)
	}
	body, ok := request.JSON.(map[string]any)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if body["price"] != 1851.5 {
		t.Fatalf("price = %v", body["price"])
	}
	if body["quantity"] != 3 {
		t.Fatalf("quantity = %v", body["quantity"])
	}
	if request.Headers["trading-token"] != "trading-token" {
		t.Fatalf("missing trading token header")
	}
}

func TestCancelOrderCanReturnRawPayload(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"id":          1626,
				"orderStatus": "PendingCancel",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.dnse.example",
		TradingToken:  "trading-token",
		HTTPTransport: httpTransport,
	})

	payload, err := broker.Trading().Orders().CancelWithRequest(context.Background(), sdktrading.CancelOrderRequest{
		AccountID: "0001179019",
		OrderID:   "1626",
	})
	if err != nil {
		t.Fatalf("cancel order: %v", err)
	}
	if got := httpTransport.requests[0].URL; got != "https://api.dnse.example/accounts/0001179019/orders/1626?marketType=DERIVATIVE&orderCategory=NORMAL" {
		t.Fatalf("url = %s", got)
	}
	data, ok := payload.Data.(map[string]any)
	if !ok {
		t.Fatalf("payload data type = %T", payload.Data)
	}
	if data["orderStatus"] != "PendingCancel" {
		t.Fatalf("orderStatus = %v", data["orderStatus"])
	}
}
