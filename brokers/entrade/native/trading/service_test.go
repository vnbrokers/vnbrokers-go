package trading

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeDependencies struct {
	capability core.Capability
	request    transport.HTTPRequest
	response   transport.HTTPResponse
	err        error
	sent       bool
}

func (f *fakeDependencies) dependencies() Dependencies {
	return Dependencies{
		BaseURL: "https://entrade.example/api",
		Headers: func(body bool) map[string]string {
			headers := map[string]string{"Authorization": "Bearer jwt-token"}
			if body {
				headers["Content-Type"] = "application/json"
			}
			return headers
		},
		RequireCapability: func(capability core.Capability) error {
			f.capability = capability
			return f.err
		},
		Send: func(_ context.Context, _ string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
			f.sent = true
			f.request = request
			return f.response, nil
		},
	}
}

func TestAccountEndpointsConstructRequests(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name       string
		capability core.Capability
		method     string
		url        string
		call       func(Service) error
	}{
		{"investor account", CapabilityInvestorAccount, "GET", "https://entrade.example/api/investors/1000000036/investor_account", func(s Service) error {
			_, err := s.GetInvestorAccount(ctx, dto.GetInvestorAccountRequest{InvestorID: "1000000036"})
			return err
		}},
		{"account balance", CapabilityAccountBalance, "GET", "https://entrade.example/api/account_balances/1000000036", func(s Service) error {
			_, err := s.GetAccountBalance(ctx, dto.GetAccountBalanceRequest{InvestorID: "1000000036"})
			return err
		}},
		{"margin portfolios", CapabilityDerivativeMarginPortfolios, "GET", "https://entrade.example/api/investors/1000000036/derivative_margin_portfolios", func(s Service) error {
			_, err := s.GetDerivativeMarginPortfolios(ctx, dto.GetDerivativeMarginPortfoliosRequest{InvestorID: "1000000036"})
			return err
		}},
		{"ppse", CapabilityPPSE, "GET", "https://entrade.example/api/derivative/investors/1000000036/ppse?bankMarginPortfolioId=34&price=1922.8&side=NB&symbol=VN30F2512", func(s Service) error {
			_, err := s.GetPPSE(ctx, dto.GetPPSERequest{InvestorID: "1000000036", BankMarginPortfolioID: "34", Symbol: "VN30F2512", Price: decimal.RequireFromString("1922.8"), Side: "NB"})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fake := &fakeDependencies{response: transport.HTTPResponse{StatusCode: 200, Body: map[string]any{}}}
			if err := tc.call(NewService(fake.dependencies())); err != nil {
				t.Fatal(err)
			}
			if fake.capability != tc.capability || fake.request.Method != tc.method || fake.request.URL != tc.url {
				t.Fatalf("capability=%q method=%q url=%q", fake.capability, fake.request.Method, fake.request.URL)
			}
			if fake.request.Headers["Authorization"] != "Bearer jwt-token" {
				t.Fatalf("headers=%v", fake.request.Headers)
			}
		})
	}
}

func TestOrderDealAndRiskEndpointsConstructRequests(t *testing.T) {
	ctx := context.Background()
	fake := &fakeDependencies{response: transport.HTTPResponse{StatusCode: 200, Body: map[string]any{}}}
	service := NewService(fake.dependencies())

	_, _ = service.PlaceDerivativeOrder(ctx, dto.PlaceDerivativeOrderRequest{BankMarginPortfolioID: 34, InvestorID: 1000000036, Symbol: "VN30F2512", Price: decimal.RequireFromString("1920.9"), OrderType: "LO", Side: "NB", Quantity: 1})
	if fake.capability != CapabilityPlaceDerivativeOrder || fake.request.Method != "POST" || fake.request.URL != "https://entrade.example/api/derivative/orders" {
		t.Fatalf("place request=%#v capability=%q", fake.request, fake.capability)
	}
	body := fake.request.JSON.(dto.PlaceDerivativeOrderRequest)
	if !body.Price.Equal(decimal.RequireFromString("1920.9")) || body.Quantity != 1 {
		t.Fatalf("place body=%#v", body)
	}

	_, _ = service.GetDerivativeOrders(ctx, dto.GetDerivativeOrdersRequest{InvestorAccountID: "1000000036"})
	if fake.request.URL != "https://entrade.example/api/derivative/orders?_end=20&_start=0&investorAccountId=1000000036" {
		t.Fatalf("orders url=%s", fake.request.URL)
	}

	_, _ = service.GetDerivativeOrder(ctx, dto.GetDerivativeOrderRequest{OrderID: "1110910"})
	assertRequest(t, fake, CapabilityDerivativeOrder, "GET", "https://entrade.example/api/derivative/orders/1110910")
	_, _ = service.CancelDerivativeOrder(ctx, dto.CancelDerivativeOrderRequest{OrderID: "1110910"})
	assertRequest(t, fake, CapabilityCancelDerivativeOrder, "DELETE", "https://entrade.example/api/derivative/orders/1110910")
	_, _ = service.GetDerivativeDeals(ctx, dto.GetDerivativeDealsRequest{InvestorAccountID: "1000000036"})
	assertRequest(t, fake, CapabilityDerivativeDeals, "GET", "https://entrade.example/api/derivative/deals?_end=20&_start=0&investorAccountId=1000000036")
	_, _ = service.CloseDerivativeDeal(ctx, dto.CloseDerivativeDealRequest{DealID: "1000546", OrderType: "LO"})
	assertRequest(t, fake, CapabilityCloseDerivativeDeal, "POST", "https://entrade.example/api/derivative/deals/1000546/_close_deal")
	closeBody := fake.request.JSON.(map[string]any)
	if closeBody["orderType"] != "LO" || closeBody["triggeredBy"] != "close-deal" {
		t.Fatalf("close body=%#v", closeBody)
	}
	_, _ = service.GetRiskConfig(ctx, dto.GetRiskConfigRequest{InvestorAccountID: "1000000036"})
	assertRequest(t, fake, CapabilityRiskConfig, "GET", "https://entrade.example/api/risk_configs?investorAccountId=1000000036")
	_, _ = service.UpdateRiskConfig(ctx, dto.UpdateRiskConfigRequest{PathInvestorAccountID: "1000000036", CutLossRate: decimal.RequireFromString("0.24"), InvestorAccountID: 1000000036, InvestorID: 1000000036, AutoIncreaseDealRate: true, EnableAutoDealDepositNoti: true})
	assertRequest(t, fake, CapabilityUpdateRiskConfig, "PATCH", "https://entrade.example/api/risk_configs/1000000036")
}

func TestTradingDecodesRawAndBodyResponses(t *testing.T) {
	fake := &fakeDependencies{response: transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"id":1110910,"symbol":"VN30F2512","price":"1920.9","unknown":true}`)}}
	service := NewService(fake.dependencies())
	order, err := service.GetDerivativeOrder(context.Background(), dto.GetDerivativeOrderRequest{OrderID: "1110910"})
	if err != nil {
		t.Fatal(err)
	}
	if order.ID != 1110910 || order.Symbol != "VN30F2512" || !order.Price.Equal(decimal.RequireFromString("1920.9")) {
		t.Fatalf("order=%#v", order)
	}

	fake.response = transport.HTTPResponse{StatusCode: 200, Body: map[string]any{"availableCash": 1000.5}}
	balance, err := service.GetAccountBalance(context.Background(), dto.GetAccountBalanceRequest{InvestorID: "1000000036"})
	if err != nil || !balance.AvailableCash.Equal(decimal.RequireFromString("1000.5")) {
		t.Fatalf("balance=%#v err=%v", balance, err)
	}
}

func TestTradingDecodesCloseDealAndRiskConfigWrappers(t *testing.T) {
	fake := &fakeDependencies{response: transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"data":[{"id":1000546,"symbol":"VN30F2512"}]}`)}}
	service := NewService(fake.dependencies())
	closed, err := service.CloseDerivativeDeal(context.Background(), dto.CloseDerivativeDealRequest{DealID: "1000546", OrderType: "LO"})
	if err != nil || len(closed.Data) != 1 || closed.Data[0].ID != 1000546 {
		t.Fatalf("closed=%#v err=%v", closed, err)
	}

	fake.response = transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"data":[{"investorAccountId":1000000036,"cutLossRate":"0.24"}],"total":1}`)}
	configs, err := service.GetRiskConfig(context.Background(), dto.GetRiskConfigRequest{InvestorAccountID: "1000000036"})
	if err != nil || configs.Total != 1 || len(configs.Data) != 1 || !configs.Data[0].CutLossRate.Equal(decimal.RequireFromString("0.24")) {
		t.Fatalf("configs=%#v err=%v", configs, err)
	}
}

func TestTradingCapabilityFailureSkipsTransport(t *testing.T) {
	want := errors.New("unsupported")
	fake := &fakeDependencies{err: want}
	_, err := NewService(fake.dependencies()).GetInvestorAccount(context.Background(), dto.GetInvestorAccountRequest{})
	if !errors.Is(err, want) || fake.sent {
		t.Fatalf("err=%v sent=%v", err, fake.sent)
	}
}

func TestTradingDecodeErrorPreservesRawResponse(t *testing.T) {
	raw := []byte(`{"id":`)
	fake := &fakeDependencies{response: transport.HTTPResponse{StatusCode: 200, Raw: raw}}
	_, err := NewService(fake.dependencies()).GetDerivativeOrder(context.Background(), dto.GetDerivativeOrderRequest{OrderID: "1"})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error type=%T error=%v", err, err)
	}
	preserved, ok := brokerErr.Raw.([]byte)
	if brokerErr.Category != sdkerrors.CategoryDecode || !ok || !bytes.Equal(preserved, raw) {
		t.Fatalf("broker error=%#v", brokerErr)
	}
}

func assertRequest(t *testing.T, fake *fakeDependencies, capability core.Capability, method string, url string) {
	t.Helper()
	if fake.capability != capability || fake.request.Method != method || fake.request.URL != url {
		t.Fatalf("capability=%q method=%q url=%q", fake.capability, fake.request.Method, fake.request.URL)
	}
}
