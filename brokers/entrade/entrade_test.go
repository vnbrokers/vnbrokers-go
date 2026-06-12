package entrade

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeHTTPTransport struct {
	requests  []transport.HTTPRequest
	responses []transport.HTTPResponse
}

func TestHTTPErrorMapsBrokerRejectedCodeMessageAndPayload(t *testing.T) {
	body := map[string]any{"code": "ORDER_REJECTED", "message": "invalid order"}
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{StatusCode: 400, Body: body}}}
	broker := NewBroker(Config{BaseURL: "https://entrade.example/api", Token: "jwt-token", HTTPTransport: httpTransport})

	_, err := broker.Native().Trading().GetDerivativeOrder(context.Background(), nativedto.GetDerivativeOrderRequest{OrderID: "1110910"})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error type=%T error=%v", err, err)
	}
	if brokerErr.Category != sdkerrors.CategoryBrokerRejected || brokerErr.Code != "ORDER_REJECTED" || brokerErr.Message != "invalid order" || !reflect.DeepEqual(brokerErr.Raw, body) {
		t.Fatalf("broker error=%#v", brokerErr)
	}
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

	response, err := broker.Auth().Login(context.Background(), nativedto.LoginRequest{
		Username: "alice",
		Password: "secret",
	})
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
	body, ok := request.JSON.(nativedto.LoginRequest)
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

	_, _ = broker.Native().Trading().GetInvestorAccount(context.Background(), nativedto.GetInvestorAccountRequest{InvestorID: "1000000036"})
	_, _ = broker.Native().Trading().GetAccountBalance(context.Background(), nativedto.GetAccountBalanceRequest{InvestorID: "1000000036"})
	_, _ = broker.Native().Trading().GetDerivativeMarginPortfolios(context.Background(), nativedto.GetDerivativeMarginPortfoliosRequest{InvestorID: "1000000036"})
	_, _ = broker.Native().Trading().GetPPSE(context.Background(), nativedto.GetPPSERequest{
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

	_, _ = broker.Native().Trading().PlaceDerivativeOrder(context.Background(), nativedto.PlaceDerivativeOrderRequest{
		BankMarginPortfolioID: 34,
		InvestorID:            1000000036,
		Symbol:                "VN30F2512",
		Price:                 decimal.RequireFromString("1920.9"),
		OrderType:             "LO",
		Side:                  "NB",
		Quantity:              1,
	})
	_, _ = broker.Native().Trading().GetDerivativeOrders(context.Background(), nativedto.GetDerivativeOrdersRequest{
		InvestorAccountID: "1000000036",
		Start:             0,
		End:               20,
	})
	_, _ = broker.Native().Trading().GetDerivativeOrder(context.Background(), nativedto.GetDerivativeOrderRequest{OrderID: "1110910"})
	_, _ = broker.Native().Trading().CancelDerivativeOrder(context.Background(), nativedto.CancelDerivativeOrderRequest{OrderID: "1110910"})

	if got := httpTransport.requests[0].URL; got != "https://entrade.example/api/derivative/orders" {
		t.Fatalf("place url = %s", got)
	}
	body, ok := httpTransport.requests[0].JSON.(nativedto.PlaceDerivativeOrderRequest)
	if !ok {
		t.Fatalf("place body type = %T", httpTransport.requests[0].JSON)
	}
	if body.Symbol != "VN30F2512" || body.Quantity != 1 {
		t.Fatalf("place body = %#v", body)
	}
	if !body.Price.Equal(decimal.RequireFromString("1920.9")) {
		t.Fatalf("place price = %#v", body.Price)
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

	_, _ = broker.Native().Trading().GetDerivativeDeals(context.Background(), nativedto.GetDerivativeDealsRequest{
		InvestorAccountID: "1000000036",
		Start:             0,
		End:               20,
	})
	_, _ = broker.Native().Trading().CloseDerivativeDeal(context.Background(), nativedto.CloseDerivativeDealRequest{DealID: "1000546", OrderType: "LO"})
	_, _ = broker.Native().Trading().GetRiskConfig(context.Background(), nativedto.GetRiskConfigRequest{InvestorAccountID: "1000000036"})
	_, _ = broker.Native().Trading().UpdateRiskConfig(context.Background(), nativedto.UpdateRiskConfigRequest{
		PathInvestorAccountID:     "1000000036",
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
	body, ok := httpTransport.requests[3].JSON.(nativedto.UpdateRiskConfigRequest)
	if !ok {
		t.Fatalf("risk body type = %T", httpTransport.requests[3].JSON)
	}
	if !body.CutLossRate.Equal(decimal.RequireFromString("0.24")) {
		t.Fatalf("cutLossRate = %#v", body.CutLossRate)
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

	response, err := broker.Native().MarketData().GetDerivatives(context.Background(), nativedto.GetDerivativesRequest{})
	if err != nil {
		t.Fatalf("list derivatives: %v", err)
	}
	if got := httpTransport.requests[0].URL; got != "https://entrade.example/api/derivatives" {
		t.Fatalf("url = %s", got)
	}
	if response.Data[0].Symbol != "VN30F2512" || response.Data[0].Type != "VN30F1M" {
		t.Fatalf("derivative = %#v", response.Data[0])
	}
}
