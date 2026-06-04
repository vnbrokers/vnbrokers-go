package ssi

import (
	"context"
	"testing"

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

func TestGetOTPSendsSSIConsumerCredentials(t *testing.T) {
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
		BaseURL:        "https://ssi.example",
		ConsumerID:     "consumer",
		ConsumerSecret: "secret",
		HTTPTransport:  httpTransport,
	})

	payload, err := broker.Auth().GetOTP(context.Background())
	if err != nil {
		t.Fatalf("get otp: %v", err)
	}
	if payload.Source != "ssi" {
		t.Fatalf("payload source = %s", payload.Source)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if got := request.URL; got != "https://ssi.example/api/v2/Trading/GetOTP" {
		t.Fatalf("url = %s", got)
	}
	body, ok := request.JSON.(map[string]any)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if body["consumerID"] != "consumer" || body["consumerSecret"] != "secret" {
		t.Fatalf("body = %+v", body)
	}
	if request.Headers["Authorization"] != "" {
		t.Fatalf("auth header should be omitted for OTP")
	}
}

func TestGetAccessTokenStoresTokenForTradingRequests(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{
				StatusCode: 200,
				Body: map[string]any{
					"message": "Success",
					"status":  200,
					"data": map[string]any{
						"accessToken": "access-token",
					},
				},
			},
			{
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
			},
		},
	}
	broker := NewBroker(Config{
		BaseURL:        "https://ssi.example",
		ConsumerID:     "consumer",
		ConsumerSecret: "secret",
		HTTPTransport:  httpTransport,
	})

	token, err := broker.Auth().GetAccessToken(context.Background(), AccessTokenRequest{
		TwoFactorType: 1,
		Code:          "123456",
		IsSave:        true,
	})
	if err != nil {
		t.Fatalf("get access token: %v", err)
	}
	if token.AccessToken != "access-token" {
		t.Fatalf("access token = %s", token.AccessToken)
	}

	_, err = broker.Trading().Accounts().StockBalance(context.Background(), "0901351")
	if err != nil {
		t.Fatalf("stock balance: %v", err)
	}
	request := httpTransport.requests[1]
	if request.Headers["Authorization"] != "Bearer access-token" {
		t.Fatalf("authorization = %s", request.Headers["Authorization"])
	}
}
