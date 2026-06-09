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
		BaseURL:               "https://ssi.example",
		ConsumerID:            "consumer",
		TradingConsumerSecret: "trading-secret",
		HTTPTransport:         httpTransport,
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
	if len(body) != 2 || body["consumerID"] != "consumer" || body["consumerSecret"] != "trading-secret" {
		t.Fatalf("body = %+v", body)
	}
	if request.Headers["Authorization"] != "" {
		t.Fatalf("auth header should be omitted for OTP")
	}
}

func TestGetAccessTokenUsesDataCredentialsWithoutTwoFactor(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": map[string]any{
					"accessToken": "data-token",
				},
			},
		}},
	}
	broker := NewBroker(Config{
		DataBaseURL:        "https://data.ssi.example",
		ConsumerID:         "consumer",
		DataConsumerSecret: "data-secret",
		HTTPTransport:      httpTransport,
	})

	token, err := broker.Auth().GetAccessToken(context.Background())
	if err != nil {
		t.Fatalf("get data access token: %v", err)
	}
	if token.AccessToken != "data-token" || broker.dataAccessToken != "data-token" {
		t.Fatalf("token = %+v stored = %q", token, broker.dataAccessToken)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://data.ssi.example/api/v2/Market/AccessToken" {
		t.Fatalf("url = %s", request.URL)
	}
	body, ok := request.JSON.(map[string]any)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if len(body) != 2 || body["consumerID"] != "consumer" || body["consumerSecret"] != "data-secret" {
		t.Fatalf("body = %+v", body)
	}
	if request.Headers["Authorization"] != "" {
		t.Fatalf("auth header should be omitted for data token")
	}
}

func TestGetTradingTokenUsesTradingCredentialsAndTwoFactor(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": map[string]any{
					"accessToken": "trading-token",
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:               "https://trade.ssi.example",
		ConsumerID:            "consumer",
		TradingConsumerSecret: "trading-secret",
		HTTPTransport:         httpTransport,
	})

	token, err := broker.Auth().GetTradingToken(context.Background(), TradingTokenRequest{
		TwoFactorType: 1,
		Code:          "123456",
		IsSave:        true,
	})
	if err != nil {
		t.Fatalf("get trading token: %v", err)
	}
	if token.AccessToken != "trading-token" || broker.tradingAccessToken != "trading-token" {
		t.Fatalf("token = %+v stored = %q", token, broker.tradingAccessToken)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://trade.ssi.example/api/v2/Trading/AccessToken" {
		t.Fatalf("url = %s", request.URL)
	}
	body, ok := request.JSON.(map[string]any)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if len(body) != 5 || body["consumerID"] != "consumer" || body["consumerSecret"] != "trading-secret" ||
		body["twoFactorType"] != 1 || body["code"] != "123456" || body["isSave"] != true {
		t.Fatalf("body = %+v", body)
	}
	if request.Headers["Authorization"] != "" {
		t.Fatalf("auth header should be omitted for trading token")
	}
}

func TestSSIConfigDefaultsDataBaseURL(t *testing.T) {
	config := Config{}.withDefaults()
	if config.DataBaseURL != "https://fc-data.ssi.com.vn" {
		t.Fatalf("data base URL = %s", config.DataBaseURL)
	}
}

func TestNewBrokerInitializesServiceSpecificAccessTokens(t *testing.T) {
	broker := NewBroker(Config{
		DataAccessToken:    "configured-data-token",
		TradingAccessToken: "configured-trading-token",
	})

	if broker.dataAccessToken != "configured-data-token" {
		t.Fatalf("data access token = %q", broker.dataAccessToken)
	}
	if broker.tradingAccessToken != "configured-trading-token" {
		t.Fatalf("trading access token = %q", broker.tradingAccessToken)
	}
}
