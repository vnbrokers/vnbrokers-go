package ssi

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/trading"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeHTTPTransport struct {
	requests  []transport.HTTPRequest
	responses []transport.HTTPResponse
}

type tokenRefreshTransport struct{}

func (tokenRefreshTransport) Send(_ context.Context, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	token := "data-token"
	if strings.Contains(request.URL, "/Trading/") {
		token = "trading-token"
	}
	return transport.HTTPResponse{
		StatusCode: 200,
		Body: map[string]any{
			"message": "Success",
			"status":  200,
			"data": map[string]any{
				"accessToken": token,
			},
		},
	}, nil
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
	if token.AccessToken != "data-token" || broker.dataToken() != "data-token" {
		t.Fatalf("token = %+v stored = %q", token, broker.dataToken())
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
	if token.AccessToken != "trading-token" || broker.tradingToken() != "trading-token" {
		t.Fatalf("token = %+v stored = %q", token, broker.tradingToken())
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

func TestGetAccessTokenRejectsMissingTokenWithoutReplacingExistingToken(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data":    map[string]any{},
			},
		}},
	}
	broker := NewBroker(Config{
		DataToken:     "existing-data-token",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Auth().GetAccessToken(context.Background())
	assertTokenDecodeError(t, err, "auth.get_data_access_token", "data")
	if broker.dataToken() != "existing-data-token" {
		t.Fatalf("data access token = %q", broker.dataToken())
	}
}

func TestGetTradingTokenRejectsEmptyTokenWithoutReplacingExistingToken(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"message": "Success",
				"status":  200,
				"data": map[string]any{
					"accessToken": "",
				},
			},
		}},
	}
	broker := NewBroker(Config{
		TradingToken:  "existing-trading-token",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Auth().GetTradingToken(context.Background(), TradingTokenRequest{})
	assertTokenDecodeError(t, err, "auth.get_trading_token", "trading")
	if broker.tradingToken() != "existing-trading-token" {
		t.Fatalf("trading access token = %q", broker.tradingToken())
	}
}

func assertTokenDecodeError(t *testing.T, err error, operation string, tokenType string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected decode error")
	}
	var brokerError *sdkerrors.BrokerError
	if !errors.As(err, &brokerError) {
		t.Fatalf("error type = %T", err)
	}
	if brokerError.Category != sdkerrors.CategoryDecode {
		t.Fatalf("category = %s", brokerError.Category)
	}
	if brokerError.Operation != operation {
		t.Fatalf("operation = %s", brokerError.Operation)
	}
	if brokerError.Cause == nil || !strings.Contains(brokerError.Cause.Error(), tokenType) ||
		!strings.Contains(brokerError.Cause.Error(), "missing accessToken") {
		t.Fatalf("cause = %v", brokerError.Cause)
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
		DataToken:    "configured-data-token",
		TradingToken: "configured-trading-token",
	})

	if broker.dataToken() != "configured-data-token" {
		t.Fatalf("data access token = %q", broker.dataToken())
	}
	if broker.tradingToken() != "configured-trading-token" {
		t.Fatalf("trading access token = %q", broker.tradingToken())
	}
}

func TestSSIServiceTokensConcurrentRefreshAndRESTRealtimeReads(t *testing.T) {
	broker := NewBroker(Config{
		DataToken:     "initial-data-token",
		TradingToken:  "initial-trading-token",
		HTTPTransport: tokenRefreshTransport{},
		SignalRFactory: func(string, []string) SignalRClient {
			return newFakeSignalRClient()
		},
	})

	const iterations = 200
	start := make(chan struct{})
	errs := make(chan error, 4)
	var workers sync.WaitGroup
	workers.Add(4)

	go func() {
		defer workers.Done()
		<-start
		for range iterations {
			if _, err := broker.Auth().GetAccessToken(context.Background()); err != nil {
				errs <- err
				return
			}
		}
	}()
	go func() {
		defer workers.Done()
		<-start
		for range iterations {
			if _, err := broker.Auth().GetTradingToken(context.Background(), TradingTokenRequest{}); err != nil {
				errs <- err
				return
			}
		}
	}()
	go func() {
		defer workers.Done()
		<-start
		for range iterations {
			header := broker.withTradingAuthorization(nil)["Authorization"]
			if header != "Bearer initial-trading-token" && header != "Bearer trading-token" {
				errs <- errors.New("unexpected REST authorization: " + header)
				return
			}
		}
	}()
	go func() {
		defer workers.Done()
		<-start
		for range iterations {
			tradingSubscription, err := broker.Native().Trading().Realtime().SubscribeOrders(context.Background(), trading.SubscribeOrdersRequest{})
			if err != nil {
				errs <- err
				return
			}
			_ = tradingSubscription.Close()

			dataSubscription, err := broker.Native().MarketData().Realtime().SubscribeRawChannel(context.Background(), "X:ALL")
			if err != nil {
				errs <- err
				return
			}
			_ = dataSubscription.Close()
		}
	}()

	close(start)
	workers.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}
}
