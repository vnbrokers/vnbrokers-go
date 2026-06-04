package entrade

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
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

func TestLoginUsesAuthBaseURLAndStoresBearerToken(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"token": "jwt-token",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://entrade.example/api",
		AuthBaseURL:   "https://entrade.example/auth-api",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Auth().Login(context.Background(), "alice", "secret")
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if response.Token != "jwt-token" {
		t.Fatalf("token = %s", response.Token)
	}
	if broker.config.Token != "jwt-token" {
		t.Fatalf("stored token = %s", broker.config.Token)
	}
	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://entrade.example/auth-api/v2/auth" {
		t.Fatalf("url = %s", request.URL)
	}
	body, ok := request.JSON.(LoginRequest)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if body.Username != "alice" || body.Password != "secret" {
		t.Fatalf("body = %#v", body)
	}
}

func TestAccountRequestsUseBearerTokenAndEntradePaths(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{StatusCode: 200, Body: map[string]any{"id": 1000000036}},
			{StatusCode: 200, Body: map[string]any{"availableCash": 1000}},
			{StatusCode: 200, Body: map[string]any{"total": 1}},
			{StatusCode: 200, Body: map[string]any{"qmax": 2}},
		},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://entrade.example/api",
		Token:         "jwt-token",
		HTTPTransport: httpTransport,
	})

	_, _ = broker.Trading().Accounts().Master(context.Background(), "1000000036")
	_, _ = broker.Trading().Accounts().Balance(context.Background(), "1000000036")
	_, _ = broker.Trading().Accounts().LoanPackages(context.Background(), "1000000036")
	_, _ = broker.Trading().Accounts().BuyingPower(context.Background(), BuyingPowerRequest{
		InvestorID:            "1000000036",
		BankMarginPortfolioID: "34",
		Symbol:                "VN30F2512",
		Price:                 decimal.RequireFromString("1922.8"),
		Side:                  "NB",
	})

	expectedURLs := []string{
		"https://entrade.example/api/investors/1000000036/investor_account",
		"https://entrade.example/api/account_balances/1000000036",
		"https://entrade.example/api/investors/1000000036/derivative_margin_portfolios",
		"https://entrade.example/api/derivative/investors/1000000036/ppse?bankMarginPortfolioId=34&price=1922.8&side=NB&symbol=VN30F2512",
	}
	for i, expected := range expectedURLs {
		request := httpTransport.requests[i]
		if request.URL != expected {
			t.Fatalf("request %d url = %s", i, request.URL)
		}
		if request.Headers["Authorization"] != "Bearer jwt-token" {
			t.Fatalf("request %d auth = %s", i, request.Headers["Authorization"])
		}
	}
}

func TestOrderRequestsUseDerivativeEndpoints(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{StatusCode: 200, Body: map[string]any{"id": 1110910}},
			{StatusCode: 200, Body: map[string]any{"total": 1}},
			{StatusCode: 200, Body: map[string]any{"id": 1110910}},
			{StatusCode: 200, Body: map[string]any{"id": 1110910}},
		},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://entrade.example/api",
		Token:         "jwt-token",
		HTTPTransport: httpTransport,
	})

	_, _ = broker.Trading().Orders().Place(context.Background(), PlaceDerivativeOrderRequest{
		BankMarginPortfolioID: 34,
		InvestorID:            1000000036,
		Symbol:                "VN30F2512",
		Price:                 decimal.RequireFromString("1920.9"),
		OrderType:             "LO",
		Side:                  "NB",
		Quantity:              1,
	})
	_, _ = broker.Trading().Orders().List(context.Background(), ListOrdersRequest{
		InvestorAccountID: "1000000036",
		Start:             0,
		End:               20,
	})
	_, _ = broker.Trading().Orders().Get(context.Background(), "1110910")
	_, _ = broker.Trading().Orders().Cancel(context.Background(), "1110910")

	if got := httpTransport.requests[0].URL; got != "https://entrade.example/api/derivative/orders" {
		t.Fatalf("place url = %s", got)
	}
	body, ok := httpTransport.requests[0].JSON.(map[string]any)
	if !ok {
		t.Fatalf("place body type = %T", httpTransport.requests[0].JSON)
	}
	if body["symbol"] != "VN30F2512" || body["quantity"] != 1 {
		t.Fatalf("place body = %#v", body)
	}
	if body["price"] != 1920.9 {
		t.Fatalf("place price = %#v", body["price"])
	}
	if got := httpTransport.requests[1].URL; got != "https://entrade.example/api/derivative/orders?_end=20&_start=0&investorAccountId=1000000036" {
		t.Fatalf("list url = %s", got)
	}
	if got := httpTransport.requests[2].URL; got != "https://entrade.example/api/derivative/orders/1110910" {
		t.Fatalf("get url = %s", got)
	}
	if got := httpTransport.requests[3].Method; got != "DELETE" {
		t.Fatalf("cancel method = %s", got)
	}
}

func TestDealsAndRiskLiveUnderTrading(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{
			{StatusCode: 200, Body: map[string]any{"total": 1}},
			{StatusCode: 200, Body: map[string]any{"data": []any{}}},
			{StatusCode: 200, Body: map[string]any{"total": 1}},
			{StatusCode: 200, Body: map[string]any{"cutLossRate": 0.24}},
		},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://entrade.example/api",
		Token:         "jwt-token",
		HTTPTransport: httpTransport,
	})

	_, _ = broker.Trading().Deals().List(context.Background(), ListDealsRequest{
		InvestorAccountID: "1000000036",
		Start:             0,
		End:               20,
	})
	_, _ = broker.Trading().Deals().Close(context.Background(), "1000546", "LO")
	_, _ = broker.Trading().Risk().Config(context.Background(), "1000000036")
	_, _ = broker.Trading().Risk().UpdateConfig(context.Background(), "1000000036", RiskConfigRequest{
		CutLossRate:               decimal.RequireFromString("0.24"),
		InvestorAccountID:         1000000036,
		TrailingEnabled:           false,
		InvestorID:                1000000036,
		AutoIncreaseDealRate:      true,
		EnableAutoDealDepositNoti: true,
	})

	expected := []string{
		"https://entrade.example/api/derivative/deals?_end=20&_start=0&investorAccountId=1000000036",
		"https://entrade.example/api/derivative/deals/1000546/_close_deal",
		"https://entrade.example/api/risk_configs?investorAccountId=1000000036",
		"https://entrade.example/api/risk_configs/1000000036",
	}
	for i, expectedURL := range expected {
		if got := httpTransport.requests[i].URL; got != expectedURL {
			t.Fatalf("request %d url = %s", i, got)
		}
	}
	if got := httpTransport.requests[1].JSON.(map[string]any)["triggeredBy"]; got != "close-deal" {
		t.Fatalf("triggeredBy = %v", got)
	}
	if got := httpTransport.requests[3].Method; got != "PATCH" {
		t.Fatalf("risk update method = %s", got)
	}
	body, ok := httpTransport.requests[3].JSON.(map[string]any)
	if !ok {
		t.Fatalf("risk body type = %T", httpTransport.requests[3].JSON)
	}
	if body["cutLossRate"] != 0.24 {
		t.Fatalf("cutLossRate = %#v", body["cutLossRate"])
	}
}

func TestDerivativesListMapsSymbols(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{
						"symbol": "VN30F2512",
						"type":   "VN30F1M",
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://entrade.example/api",
		Token:         "jwt-token",
		HTTPTransport: httpTransport,
	})

	symbols, err := broker.MarketData().Derivatives().List(context.Background())
	if err != nil {
		t.Fatalf("list derivatives: %v", err)
	}
	if got := httpTransport.requests[0].URL; got != "https://entrade.example/api/derivatives" {
		t.Fatalf("url = %s", got)
	}
	if symbols[0].Symbol != "VN30F2512" || symbols[0].DisplayName != "VN30F1M" {
		t.Fatalf("symbol = %#v", symbols[0])
	}
}
