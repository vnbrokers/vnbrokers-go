package ssi

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"net/url"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestSSIAdditionalGETEndpointsBuildRequests(t *testing.T) {
	tests := []struct {
		name string
		call func(context.Context, *Broker) error
		url  string
	}{
		{
			name: "audit order book",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Accounts().AuditOrders(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/Trading/auditOrderBook?account=0901351",
		},
		{
			name: "account asset",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Accounts().AccountAsset(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/Trading/ppmmraccount?account=0901351",
		},
		{
			name: "max sell quantity",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Accounts().MaxSellQuantity(ctx, MaxSellQuantityRequest{
					AccountID: "0901351",
					Symbol:    "SSI",
					Price:     decimal.NewFromInt(21000),
				})
				return err
			},
			url: "https://ssi.example/api/v2/Trading/maxSellQty?account=0901351&instrumentID=SSI&price=21000",
		},
		{
			name: "rate limit",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Accounts().RateLimit(ctx)
				return err
			},
			url: "https://ssi.example/api/v2/Trading/rateLimit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := NewBroker(Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "GET" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] != "" {
				t.Fatalf("GET should not be signed")
			}
		})
	}
}

func TestSSIAdditionalPOSTEndpointsBuildSignedRequests(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	tests := []struct {
		name string
		call func(context.Context, *Broker) error
		url  string
	}{
		{
			name: "derivative new order",
			call: func(ctx context.Context, b *Broker) error {
				price := decimal.NewFromInt(1200)
				_, err := b.Trading().Orders().DerivativePlaceWithRequest(ctx, PlaceOrderRequest{
					PlaceOrderRequest: domain.PlaceOrderRequest{
						AccountID: "0901358",
						Symbol:    "VN30F2306",
						Side:      domain.OrderSideBuy,
						Type:      domain.OrderTypeLimit,
						Quantity:  decimal.NewFromInt(1),
						Price:     &price,
					},
					MarketID: "VNFE",
					Code:     "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/Trading/derNewOrder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := NewBroker(Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				PrivateKey:    testBase64XMLPrivateKey(key),
				RequestID:     func() string { return "12345678" },
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "POST" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] == "" {
				t.Fatalf("missing signature")
			}
			if _, ok := request.JSON.(signedJSON); !ok {
				t.Fatalf("json body type = %T", request.JSON)
			}
		})
	}
}

func TestSSINativeTradingGETEndpointsBuildRequests(t *testing.T) {
	tests := []struct {
		name string
		call func(context.Context, *Broker) error
		url  string
	}{
		{
			name: "cash in advance amount",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().CashInAdvanceAmount(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/cash/cashInAdvanceAmount?account=0901351",
		},
		{
			name: "unsettle sold transaction",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().UnsettleSoldTransaction(ctx, "0901351", "10/03/2023")
				return err
			},
			url: "https://ssi.example/api/v2/cash/unsettleSoldTransaction?account=0901351&settleDate=10%2F03%2F2023",
		},
		{
			name: "cash transfer histories",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().TransferHistories(ctx, "0901351", "10/01/2023", "10/02/2023")
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferHistories?account=0901351&fromDate=10%2F01%2F2023&toDate=10%2F02%2F2023",
		},
		{
			name: "stock transferable",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().Transferable(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/stock/transferable?account=0901351",
		},
		{
			name: "rights dividends",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().Dividend(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/ors/dividend?account=0901351",
		},
		{
			name: "conditional order list",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().FcoList(ctx, url.Values{
					"account":   {"0901351"},
					"pageIndex": {"1"},
					"pageSize":  {"50"},
				})
				return err
			},
			url: "https://ssi.example/api/v2/fco/list?account=0901351&pageIndex=1&pageSize=50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := NewBroker(Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "GET" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] != "" {
				t.Fatalf("GET should not be signed")
			}
		})
	}
}

func TestSSINativeTradingPOSTEndpointsBuildSignedRequests(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	tests := []struct {
		name string
		call func(context.Context, *Broker) error
		url  string
	}{
		{
			name: "cash transfer internal",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().TransferInternal(ctx, "0901351", "0901356", "50000", "test", "123456")
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferInternal",
		},
		{
			name: "stock transfer",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().StockTransfer(ctx, map[string]any{
					"account":            "0901351",
					"beneficiaryAccount": "0901356",
					"exchangeID":         "HOSE",
					"instrumentID":       "SSI",
					"quantity":           100,
					"code":               "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/stock/transfer",
		},
		{
			name: "rights create",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().CreateRight(ctx, map[string]any{
					"account":       "0901351",
					"instrumentID":  "SSI",
					"entitlementID": "913312",
					"quantity":      100,
					"amount":        1000,
					"code":          "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/ors/create",
		},
		{
			name: "conditional new order",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Native().Trading().FcoNewOrder(ctx, map[string]any{
					"instrumentID": "SSI",
					"side":         "B",
					"type":         "stop",
					"price":        "21000",
					"quantity":     100,
					"account":      "0901351",
					"stopPrice":    21100,
					"operator":     "greater_or_equal",
					"code":         "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/fco/neworder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := NewBroker(Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				PrivateKey:    testBase64XMLPrivateKey(key),
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "POST" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] == "" {
				t.Fatalf("missing signature")
			}
		})
	}
}
