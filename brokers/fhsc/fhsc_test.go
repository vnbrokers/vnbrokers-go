package fhsc

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
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

func TestBrokerExposesAuthAndNativeServices(t *testing.T) {
	broker := NewBroker(Config{})
	if broker.Auth() == nil {
		t.Fatal("Auth() is nil")
	}
	if broker.Native() == nil || broker.Native().MarketData() == nil || broker.Native().Trading() == nil {
		t.Fatal("native services are not fully wired")
	}
}

func TestAuthGetCurrentUserSignsReadRequestAndStoresUserID(t *testing.T) {
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 200,
		Raw:        []byte(`{"status":200,"data":{"user_id":123456}}`),
	}}}
	fixedNow := time.Unix(1710000000, 0).UTC()
	broker := NewBroker(Config{
		BaseURL:       "https://open-api.fhsc.example",
		APIKey:        "api-key",
		APISecret:     "secret-key",
		HTTPTransport: httpTransport,
		Now:           func() time.Time { return fixedNow },
		Nonce:         func() string { return "nonce123" },
	})

	response, err := broker.Auth().GetCurrentUser(context.Background())
	if err != nil {
		t.Fatalf("GetCurrentUser: %v", err)
	}
	if response.Data.UserID != 123456 {
		t.Fatalf("user id = %d", response.Data.UserID)
	}
	if broker.config.UserID != 123456 {
		t.Fatalf("stored user id = %d", broker.config.UserID)
	}
	request := httpTransport.requests[0]
	if request.Method != "GET" || request.URL != "https://open-api.fhsc.example/users/v1/users/me" {
		t.Fatalf("request = %s %s", request.Method, request.URL)
	}
	if request.Headers["X-FH-APIKEY"] != "api-key" {
		t.Fatalf("api key header = %q", request.Headers["X-FH-APIKEY"])
	}
	if request.Headers["X-FH-USER-ID"] != "" {
		t.Fatalf("expected empty user-id header on bootstrap call, got %q", request.Headers["X-FH-USER-ID"])
	}
	if request.Headers["X-FH-TIMESTAMP"] != fmt.Sprintf("%d", fixedNow.UnixMilli()) {
		t.Fatalf("timestamp = %q", request.Headers["X-FH-TIMESTAMP"])
	}
	if request.Headers["X-FH-NONCE"] != "nonce123" {
		t.Fatalf("nonce = %q", request.Headers["X-FH-NONCE"])
	}
	wantSig := signForTest("secret-key", fixedNow.UnixMilli(), "GET", "/users/v1/users/me", "")
	if request.Headers["X-FH-SIGNATURE"] != wantSig {
		t.Fatalf("signature = %q, want %q", request.Headers["X-FH-SIGNATURE"], wantSig)
	}
	if got := request.Headers["X-FH-BODYHASH"]; got != "" {
		t.Fatalf("unexpected body hash = %q", got)
	}
}

func TestNativeTradingGetAvailableTradeSignsQueryString(t *testing.T) {
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 200,
		Raw:        []byte(`{"status":200,"result":{}}`),
	}}}
	fixedNow := time.Unix(1710000000, 0).UTC()
	broker := NewBroker(Config{
		BaseURL:       "https://open-api.fhsc.example",
		APIKey:        "api-key",
		APISecret:     "secret-key",
		UserID:        123456,
		HTTPTransport: httpTransport,
		Now:           func() time.Time { return fixedNow },
		Nonce:         func() string { return "nonce123" },
	})

	_, err := broker.Native().Trading().GetAvailableTrade(context.Background(), dto.GetAvailableTradeRequest{SubAccountID: "0001234567", OrderSide: "BUY", Symbol: "HPG", QuotePrice: 0})
	if err != nil {
		t.Fatalf("GetAvailableTrade: %v", err)
	}
	request := httpTransport.requests[0]
	wantURL := "https://open-api.fhsc.example/trading/v2/accounts/0001234567/available-trade?orderSide=BUY&quotePrice=0&symbol=HPG"
	if request.URL != wantURL {
		t.Fatalf("url = %s", request.URL)
	}
	wantSig := signForTest("secret-key", fixedNow.UnixMilli(), "GET", "/trading/v2/accounts/0001234567/available-trade?orderSide=BUY&quotePrice=0&symbol=HPG", "")
	if request.Headers["X-FH-SIGNATURE"] != wantSig {
		t.Fatalf("signature = %q, want %q", request.Headers["X-FH-SIGNATURE"], wantSig)
	}
	if request.Headers["X-FH-USER-ID"] != "123456" {
		t.Fatalf("user id header = %q", request.Headers["X-FH-USER-ID"])
	}
}

func TestPlaceOrderAddsBodyHashAndTwoFactorHeader(t *testing.T) {
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 200,
		Raw:        []byte(`{"result":[]}`),
	}}}
	fixedNow := time.Unix(1710000000, 0).UTC()
	limitPrice := int64(25000)
	body := dto.CreateOrderRequest{SubAccount: "0881234567.4", Side: "BUY", Symbol: "HPG", Quantity: 100, TypeValue: "LIMIT", LimitPrice: &limitPrice, StockType: "STOCK"}
	broker := NewBroker(Config{
		BaseURL:        "https://open-api.fhsc.example",
		APIKey:         "api-key",
		APISecret:      "secret-key",
		UserID:         123456,
		TwoFactorToken: "2fa-token",
		HTTPTransport:  httpTransport,
		Now:            func() time.Time { return fixedNow },
		Nonce:          func() string { return "nonce123" },
	})

	_, err := broker.Native().Trading().PlaceOrder(context.Background(), dto.PlaceOrderRequest{SubAccountID: "0881234567", Body: body})
	if err != nil {
		t.Fatalf("PlaceOrder: %v", err)
	}
	request := httpTransport.requests[0]
	if request.Method != "POST" || request.URL != "https://open-api.fhsc.example/trading/oa/sub-accounts/0881234567/orders" {
		t.Fatalf("request = %s %s", request.Method, request.URL)
	}
	if !reflect.DeepEqual(request.JSON, body) {
		t.Fatalf("json body = %#v", request.JSON)
	}
	payloadHash := bodyHashForTest(body)
	if request.Headers["X-FH-BODYHASH"] != payloadHash {
		t.Fatalf("body hash = %q, want %q", request.Headers["X-FH-BODYHASH"], payloadHash)
	}
	wantSig := signForTest("secret-key", fixedNow.UnixMilli(), "POST", "/trading/oa/sub-accounts/0881234567/orders", payloadHash)
	if request.Headers["X-FH-SIGNATURE"] != wantSig {
		t.Fatalf("signature = %q, want %q", request.Headers["X-FH-SIGNATURE"], wantSig)
	}
	if request.Headers["X-FH-2FA-TOKEN"] != "2fa-token" {
		t.Fatalf("2fa header = %q", request.Headers["X-FH-2FA-TOKEN"])
	}
}

func TestVerifyTwoFactorStoresSessionToken(t *testing.T) {
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 200,
		Raw:        []byte(`{"session_token":"session-123","expires_at":"2026-06-23T00:00:00Z","expires_at_epoch":1782172800}`),
	}}}
	broker := NewBroker(Config{
		BaseURL:       "https://open-api.fhsc.example",
		APIKey:        "api-key",
		APISecret:     "secret-key",
		UserID:        123456,
		HTTPTransport: httpTransport,
		Now:           func() time.Time { return time.Unix(1710000000, 0).UTC() },
		Nonce:         func() string { return "nonce123" },
	})

	response, err := broker.Auth().VerifyTwoFactorOTP(context.Background(), dto.TwoFactorVerifyPayload{TicketID: "ticket-1", OTPCode: "123456"})
	if err != nil {
		t.Fatalf("VerifyTwoFactorOTP: %v", err)
	}
	if response.SessionToken != "session-123" {
		t.Fatalf("session token = %q", response.SessionToken)
	}
	if broker.config.TwoFactorToken != "session-123" {
		t.Fatalf("stored 2fa token = %q", broker.config.TwoFactorToken)
	}
	if !broker.HasTwoFactorSession() {
		t.Fatal("expected active 2fa session")
	}
}

func TestInferIdentityPrefersOrderSubAccountExt(t *testing.T) {
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 200,
		Raw:        []byte(`{"status":200,"data":{"user_id":123456}}`),
	}, {
		StatusCode: 200,
		Raw:        []byte(`{"status":200,"result":[{"id":"0001","type":"normal","sub_account_ext":"0001.1"},{"id":"0004","type":"normal","sub_account_ext":"0004.4"}]}`),
	}}}
	broker := NewBroker(Config{
		BaseURL:       "https://open-api.fhsc.example",
		APIKey:        "api-key",
		APISecret:     "secret-key",
		HTTPTransport: httpTransport,
		Now:           func() time.Time { return time.Unix(1710000000, 0).UTC() },
		Nonce:         func() string { return "nonce123" },
	})

	identity, err := broker.Auth().InferIdentity(context.Background())
	if err != nil {
		t.Fatalf("InferIdentity: %v", err)
	}
	if identity.UserID != 123456 {
		t.Fatalf("user id = %d", identity.UserID)
	}
	orderAccount := identity.OrderSubAccount()
	if orderAccount == nil || orderAccount.SubAccountExt != "0004.4" {
		t.Fatalf("order sub-account = %#v", orderAccount)
	}
}

func signForTest(secret string, timestampMillis int64, method string, path string, bodyHash string) string {
	payload := fmt.Sprintf("%d\n%s\n%s\n", timestampMillis, method, path)
	if bodyHash != "" {
		payload += bodyHash
	}
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func bodyHashForTest(value any) string {
	payload, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	digest := sha256.Sum256(payload)
	return hex.EncodeToString(digest[:])
}
