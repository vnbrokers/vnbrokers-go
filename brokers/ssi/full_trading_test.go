package ssi

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"strings"
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
		{
			name: "cash in advance amount",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Cash().CashInAdvanceAmount(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/cash/cashInAdvanceAmount?account=0901351",
		},
		{
			name: "unsettled sold transactions",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Cash().UnsettledSoldTransactions(ctx, CashUnsettledSoldTransactionsRequest{
					AccountID:  "0901351",
					SettleDate: "10/03/2023",
				})
				return err
			},
			url: "https://ssi.example/api/v2/cash/unsettleSoldTransaction?account=0901351&settleDate=10%2F03%2F2023",
		},
		{
			name: "cash transfer histories",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Cash().TransferHistories(ctx, CashTransferHistoriesRequest{
					AccountID: "0901351",
					FromDate:  "10/01/2023",
					ToDate:    "10/02/2023",
				})
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferHistories?account=0901351&fromDate=10%2F01%2F2023&toDate=10%2F02%2F2023",
		},
		{
			name: "stock transferable",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().StockTransfers().Transferable(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/stock/transferable?account=0901351",
		},
		{
			name: "rights dividends",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Rights().Dividends(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/ors/dividend?account=0901351",
		},
		{
			name: "conditional order list",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().ConditionalOrders().List(ctx, ConditionalOrderListRequest{
					AccountID: "0901351",
					PageIndex: 1,
					PageSize:  50,
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
				AccessToken:   "access-token",
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
			if request.Headers["Authorization"] != "Bearer access-token" {
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
		{
			name: "cash transfer internal",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Cash().TransferInternalWithRequest(ctx, CashTransferInternalRequest{
					AccountID:          "0901351",
					BeneficiaryAccount: "0901356",
					Amount:             decimal.NewFromInt(50000),
					Remark:             "test",
					Code:               "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferInternal",
		},
		{
			name: "stock transfer",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().StockTransfers().TransferWithRequest(ctx, StockTransferRequest{
					AccountID:          "0901351",
					BeneficiaryAccount: "0901356",
					ExchangeID:         "HOSE",
					Symbol:             "SSI",
					Quantity:           decimal.NewFromInt(100),
					Code:               "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/stock/transfer",
		},
		{
			name: "rights create",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().Rights().CreateWithRequest(ctx, RightsCreateRequest{
					AccountID:     "0901351",
					Symbol:        "SSI",
					EntitlementID: "913312",
					Quantity:      decimal.NewFromInt(100),
					Amount:        decimal.NewFromInt(1000),
					Code:          "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/ors/create",
		},
		{
			name: "conditional new order",
			call: func(ctx context.Context, b *Broker) error {
				_, err := b.Trading().ConditionalOrders().NewWithRequest(ctx, ConditionalOrderNewRequest{
					AccountID: "0901351",
					Symbol:    "SSI",
					Side:      "B",
					Type:      "stop",
					Price:     "21000",
					Quantity:  decimal.NewFromInt(100),
					StopPrice: decimal.NewFromInt(21100),
					Operator:  "greater_or_equal",
					Code:      "123456",
					DeviceID:  "device",
					UserAgent: "FCTrading",
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
				AccessToken:   "access-token",
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
			if request.Headers["Authorization"] != "Bearer access-token" {
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

func TestSSIRawServicesAreExposed(t *testing.T) {
	broker := NewBroker(Config{})
	services := []string{
		"cash", "stock transfers", "rights", "conditional orders",
	}
	if broker.Trading().Cash() == nil ||
		broker.Trading().StockTransfers() == nil ||
		broker.Trading().Rights() == nil ||
		broker.Trading().ConditionalOrders() == nil {
		t.Fatalf("missing services: %s", strings.Join(services, ", "))
	}
}
