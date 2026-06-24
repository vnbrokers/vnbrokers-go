package trading_test

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestTradingHTTPContracts(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		wantBody     any
		responseBody string
		invoke       func(context.Context, trading.Service) error
	}{
		{
			name:         "account summary",
			method:       "GET",
			url:          "/trading/accounts/0001234567/summary",
			responseBody: `{"status":200,"result":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetAccountSummary(ctx, dto.GetAccountSummaryRequest{SubAccountID: "0001234567"})
				return err
			},
		},
		{
			name:         "user assets summary",
			method:       "GET",
			url:          "/users/v3/users/123456/assets/summary?cache-control=CACHE",
			responseBody: `{"status":200,"result":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetUserAssetsSummary(ctx, dto.GetUserAssetsSummaryRequest{UserID: 123456, CacheControl: "CACHE"})
				return err
			},
		},
		{
			name:         "pnl today",
			method:       "GET",
			url:          "/trading/pnl-today/123456?sub-account-id=ALL",
			responseBody: `{"status":200,"data":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetPnLToday(ctx, dto.GetPnLTodayRequest{UserID: 123456, SubAccountID: "ALL"})
				return err
			},
		},
		{
			name:         "portfolio",
			method:       "GET",
			url:          "/trading/v2/sub-accounts/0001234567/portfolio?cache-control=NOCACHE",
			responseBody: `{"status":200,"data":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetPortfolio(ctx, dto.GetPortfolioRequest{SubAccountID: "0001234567", CacheControl: "NOCACHE"})
				return err
			},
		},
		{
			name:         "available trade",
			method:       "GET",
			url:          "/trading/v2/accounts/0001234567/available-trade?orderSide=BUY&quotePrice=0&symbol=HPG",
			responseBody: `{"status":200,"result":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetAvailableTrade(ctx, dto.GetAvailableTradeRequest{SubAccountID: "0001234567", OrderSide: "BUY", Symbol: "HPG", QuotePrice: 0})
				return err
			},
		},
		{
			name:         "order history",
			method:       "GET",
			url:          "/trading/sub-accounts/0001234567/orders?fromDate=2024-01-01&orderStatus=ALL&page=1&symbol=HPG&toDate=2024-01-31",
			responseBody: `{"status":200,"result":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetOrderHistory(ctx, dto.GetOrderHistoryRequest{SubAccountID: "0001234567", FromDate: "2024-01-01", ToDate: "2024-01-31", Page: 1, OrderStatus: "ALL", Symbol: "HPG"})
				return err
			},
		},
		{
			name:         "order book detail",
			method:       "GET",
			url:          "/trading/v1/accounts/0001234567/order-book/ORD123456?cache-control=NOCACHE",
			responseBody: `{"status":200,"data":{}}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetOrderBookDetail(ctx, dto.GetOrderBookDetailRequest{SubAccountID: "0001234567", OrderID: "ORD123456", CacheControl: "NOCACHE"})
				return err
			},
		},
		{
			name:         "order book list",
			method:       "GET",
			url:          "/trading/v1/accounts/0001234567/order-book?cache-control=NOCACHE",
			responseBody: `{"status":200,"result":[]}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.GetOrderBook(ctx, dto.GetOrderBookRequest{SubAccountID: "0001234567", CacheControl: "NOCACHE"})
				return err
			},
		},
		{
			name:   "place order",
			method: "POST",
			url:    "/trading/oa/sub-accounts/0881234567/orders",
			wantBody: dto.CreateOrderRequest{
				SubAccount: "0881234567.4",
				Side:       "BUY",
				Symbol:     "HPG",
				Quantity:   100,
				TypeValue:  "LIMIT",
				LimitPrice: int64Ptr(25000),
				StockType:  "STOCK",
			},
			responseBody: `{"result":[]}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.PlaceOrder(ctx, dto.PlaceOrderRequest{
					SubAccountID: "0881234567",
					Body: dto.CreateOrderRequest{
						SubAccount: "0881234567.4",
						Side:       "BUY",
						Symbol:     "HPG",
						Quantity:   100,
						TypeValue:  "LIMIT",
						LimitPrice: int64Ptr(25000),
						StockType:  "STOCK",
					},
				})
				return err
			},
		},
		{
			name:   "cancel order",
			method: "DELETE",
			url:    "/trading/oa/sub-accounts/0881234567/orders/ORD123456",
			wantBody: dto.CancelOrderRequest{
				SubAccount: "0881234567.4",
			},
			responseBody: `{"result":[]}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.CancelOrder(ctx, dto.CancelOrderOperationRequest{
					SubAccountID: "0881234567",
					OrderID:      "ORD123456",
					Body:         dto.CancelOrderRequest{SubAccount: "0881234567.4"},
				})
				return err
			},
		},
		{
			name:   "update order",
			method: "PUT",
			url:    "/trading/oa/sub-accounts/0881234567/orders/ORD123456",
			wantBody: dto.UpdateOrderRequest{
				Quantity: 200,
				Price:    25500,
			},
			responseBody: `{"result":[]}`,
			invoke: func(ctx context.Context, s trading.Service) error {
				_, err := s.UpdateOrder(ctx, dto.UpdateOrderOperationRequest{
					SubAccountID: "0881234567",
					OrderID:      "ORD123456",
					Body:         dto.UpdateOrderRequest{Quantity: 200, Price: 25500},
				})
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var gotCapability core.Capability
			var gotRequest transport.HTTPRequest
			service := trading.NewService(trading.Dependencies{
				BaseURL:           "https://api.example",
				Headers:           func(bool, bool) map[string]string { return map[string]string{"Authorization": "Bearer token"} },
				RequireCapability: func(capability core.Capability) error { gotCapability = capability; return nil },
				Send: func(_ context.Context, _ string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
					gotRequest = request
					return transport.HTTPResponse{StatusCode: 200, Raw: []byte(test.responseBody)}, nil
				},
			})

			if err := test.invoke(context.Background(), service); err != nil {
				t.Fatalf("invoke: %v", err)
			}
			if gotCapability == "" {
				t.Fatal("capability was not checked")
			}
			if gotRequest.Method != test.method || gotRequest.URL != "https://api.example"+test.url {
				t.Fatalf("request = %s %s", gotRequest.Method, gotRequest.URL)
			}
			if !reflect.DeepEqual(gotRequest.JSON, test.wantBody) {
				t.Fatalf("body = %#v, want %#v", gotRequest.JSON, test.wantBody)
			}
			if gotRequest.Headers["Authorization"] != "Bearer token" {
				t.Fatalf("headers = %#v", gotRequest.Headers)
			}
		})
	}
}

func TestTradingDecodeErrorPreservesRawResponse(t *testing.T) {
	raw := []byte(`{"result":`)
	service := trading.NewService(trading.Dependencies{
		BaseURL:           "https://api.example",
		Headers:           func(bool, bool) map[string]string { return map[string]string{"Authorization": "Bearer token"} },
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(_ context.Context, _ string, _ transport.HTTPRequest) (transport.HTTPResponse, error) {
			return transport.HTTPResponse{StatusCode: 200, Raw: raw}, nil
		},
	})

	_, err := service.GetAccountSummary(context.Background(), dto.GetAccountSummaryRequest{SubAccountID: "0001234567"})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error type=%T error=%v", err, err)
	}
	preserved, ok := brokerErr.Raw.([]byte)
	if brokerErr.Category != sdkerrors.CategoryDecode || !ok || !bytes.Equal(preserved, raw) {
		t.Fatalf("broker error=%#v", brokerErr)
	}
}

func TestTradingBrokerRejectedOnAPIErrorCode(t *testing.T) {
	service := trading.NewService(trading.Dependencies{
		BaseURL:           "https://api.example",
		Headers:           func(bool, bool) map[string]string { return map[string]string{"Authorization": "Bearer token"} },
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(_ context.Context, _ string, _ transport.HTTPRequest) (transport.HTTPResponse, error) {
			return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"error_code":"1001","message":"invalid sub account"}`)}, nil
		},
	})

	_, err := service.GetAvailableTrade(context.Background(), dto.GetAvailableTradeRequest{SubAccountID: "0001234567", OrderSide: "BUY", Symbol: "HPG", QuotePrice: 0})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error type=%T error=%v", err, err)
	}
	if brokerErr.Category != sdkerrors.CategoryBrokerRejected || brokerErr.Code != "1001" {
		t.Fatalf("broker error=%#v", brokerErr)
	}
}

func int64Ptr(value int64) *int64 { return &value }
