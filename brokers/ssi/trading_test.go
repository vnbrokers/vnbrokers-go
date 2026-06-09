package ssi

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestPlaceWithRequestBuildsSignedSSIOrderRequest(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": map[string]any{
					"requestID": "12345678",
					"requestData": map[string]any{
						"account":      "0901351",
						"instrumentID": "SSI",
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://ssi.example",
		DataToken:     "data-token",
		TradingToken:  "trading-token",
		PrivateKey:    testBase64XMLPrivateKey(key),
		ChannelID:     "TA",
		DeviceID:      "device",
		UserAgent:     "FCTrading",
		RequestID:     func() string { return "12345678" },
		HTTPTransport: httpTransport,
	})
	price := decimal.NewFromInt(21000)

	response, err := broker.Trading().Orders().PlaceWithRequest(context.Background(), PlaceOrderRequest{
		PlaceOrderRequest: domain.PlaceOrderRequest{
			AccountID: "0901351",
			Symbol:    "SSI",
			Side:      domain.OrderSideBuy,
			Type:      domain.OrderTypeLimit,
			Quantity:  decimal.NewFromInt(100),
			Price:     &price,
		},
		Code: "123456",
	})
	if err != nil {
		t.Fatalf("place order: %v", err)
	}
	if response.Data.RequestID != "12345678" {
		t.Fatalf("response request id = %s", response.Data.RequestID)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if got := request.URL; got != "https://ssi.example/api/v2/Trading/NewOrder" {
		t.Fatalf("url = %s", got)
	}
	if request.Headers["Authorization"] != "Bearer trading-token" {
		t.Fatalf("authorization = %s", request.Headers["Authorization"])
	}
	if request.Headers["X-Signature"] == "" {
		t.Fatalf("missing signature")
	}
	body, ok := request.JSON.(signedJSON)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	fields, ok := body.Value.(map[string]any)
	if !ok {
		t.Fatalf("signed value type = %T", body.Value)
	}
	if fields["requestID"] != "12345678" {
		t.Fatalf("requestID = %v", fields["requestID"])
	}
	if fields["channelID"] != "TA" || fields["deviceId"] != "device" || fields["userAgent"] != "FCTrading" {
		t.Fatalf("device fields = %+v", fields)
	}
	if fields["marketID"] != "VN" || fields["buySell"] != "B" || fields["orderType"] != "LO" {
		t.Fatalf("order fields = %+v", fields)
	}
}

func TestStockBalanceMapsCommonDomainBalance(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": []any{
					map[string]any{
						"account":         "0901351",
						"cashbal":         1000000,
						"withdrawable":    900000,
						"purchasingpower": 800000,
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://ssi.example",
		TradingToken:  "trading-token",
		HTTPTransport: httpTransport,
	})

	balance, err := broker.Trading().Accounts().Balance(context.Background(), "0901351")
	if err != nil {
		t.Fatalf("balance: %v", err)
	}
	if balance.AccountID != "0901351" {
		t.Fatalf("account id = %s", balance.AccountID)
	}
	if balance.CashAvailable == nil || !balance.CashAvailable.Equal(decimal.NewFromInt(900000)) {
		t.Fatalf("cash available = %v", balance.CashAvailable)
	}
	if balance.BuyingPower == nil || !balance.BuyingPower.Equal(decimal.NewFromInt(800000)) {
		t.Fatalf("buying power = %v", balance.BuyingPower)
	}
}

func TestOrdersDecodesSSIOrderBookObjectResponse(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": map[string]any{
					"account": "",
					"orders": []any{
						map[string]any{
							"uniqueID":     "59753588",
							"orderID":      "26060500251342",
							"buySell":      "B",
							"price":        6970.0,
							"quantity":     1,
							"filledQty":    0,
							"orderStatus":  "RS",
							"marketID":     "VN",
							"inputTime":    "1780633829138",
							"modifiedTime": "1780633829138",
							"instrumentID": "AAA",
							"orderType":    "LO",
							"cancelQty":    0,
							"avgPrice":     0.0,
							"isForcesell":  "F",
							"isShortsell":  nil,
							"rejectReason": "",
							"note":         "",
						},
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://ssi.example",
		TradingToken:  "trading-token",
		HTTPTransport: httpTransport,
	})

	orders, err := broker.Trading().Accounts().Orders(context.Background(), "0901351")
	if err != nil {
		t.Fatalf("orders: %v", err)
	}
	if len(orders) != 1 {
		t.Fatalf("orders len = %d", len(orders))
	}
	order := orders[0]
	if order.OrderID != "26060500251342" || order.Symbol != "AAA" {
		t.Fatalf("order = %+v", order)
	}
	if order.Status != domain.OrderStatusAccepted {
		t.Fatalf("status = %s", order.Status)
	}
	if !order.Quantity.Equal(decimal.NewFromInt(1)) {
		t.Fatalf("quantity = %s", order.Quantity)
	}
}

func TestSSIJSONStatusErrorReturnsBrokerRejected(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Invalid token",
				"status":  401,
				"data":    nil,
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://ssi.example",
		TradingToken:  "bad-token",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Trading().Accounts().StockBalance(context.Background(), "0901351")
	brokerErr, ok := err.(*errors.BrokerError)
	if !ok {
		t.Fatalf("error type = %T, %v", err, err)
	}
	if brokerErr.Category != errors.CategoryBrokerRejected || brokerErr.Code != "401" {
		t.Fatalf("broker error = %+v", brokerErr)
	}
}
