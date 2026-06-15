package tcbs_test

import (
	"context"
	"errors"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type authTransport struct {
	request transport.HTTPRequest
}

type responseTransport struct {
	response transport.HTTPResponse
}

func (f responseTransport) Send(context.Context, transport.HTTPRequest) (transport.HTTPResponse, error) {
	return f.response, nil
}

func (f *authTransport) Send(_ context.Context, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	f.request = request
	return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"token":"jwt-token"}`)}, nil
}

func TestNativeHTTPErrorDecodesRawBrokerResponse(t *testing.T) {
	broker := tcbs.NewBroker(tcbs.Config{HTTPTransport: responseTransport{response: transport.HTTPResponse{
		StatusCode: 400,
		Raw:        []byte(`{"code":"INVALID","message":"bad request"}`),
	}}})

	_, err := broker.Native().MarketData().GetStockTickers(context.Background(), nativedto.GetStockTickersRequest{})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error = %T %v", err, err)
	}
	if brokerErr.Code != "INVALID" || brokerErr.Message != "bad request" {
		t.Fatalf("broker error = %#v", brokerErr)
	}
	if brokerErr.Operation != string(tcbs.CapabilityNativeMarketDataGetStockTickers) {
		t.Fatalf("operation = %q", brokerErr.Operation)
	}
}

func TestBrokerExposesTypedAuthAndNativeServices(t *testing.T) {
	transport := &authTransport{}
	broker := tcbs.NewBroker(tcbs.Config{HTTPTransport: transport})

	response, err := broker.Auth().GetToken(context.Background(), nativedto.GetTokenRequest{
		APIKey: "api-key",
		OTP:    "123456",
	})
	if err != nil {
		t.Fatalf("GetToken() error = %v", err)
	}
	if response.Token != "jwt-token" {
		t.Fatalf("token = %q", response.Token)
	}
	if broker.Native() == nil || broker.Native().Trading() == nil || broker.Native().MarketData() == nil {
		t.Fatal("native services are not initialized")
	}
	if broker.Native().Trading().Realtime() == nil || broker.Native().MarketData().Realtime() == nil {
		t.Fatal("native realtime services are not initialized")
	}
	if got := transport.request.URL; got != tcbs.ProductionBaseURL+"/gaia/v1/oauth2/openapi/token" {
		t.Fatalf("auth URL = %q", got)
	}
	if !broker.Supports(tcbs.CapabilityNativeTradingGetSubAccountInformation) {
		t.Fatal("missing native trading capability")
	}
	if !broker.Supports(tcbs.CapabilityNativeMarketDataGetStockTickers) {
		t.Fatal("missing native market-data capability")
	}
}

func TestCapabilitiesCoverAuthAndEveryNativeMethod(t *testing.T) {
	if len(tcbs.Capabilities) != 56 {
		t.Fatalf("capability count = %d, want 56", len(tcbs.Capabilities))
	}
	seen := make(map[string]struct{}, len(tcbs.Capabilities))
	for _, capability := range tcbs.Capabilities {
		value := string(capability)
		if _, exists := seen[value]; exists {
			t.Fatalf("duplicate capability %q", value)
		}
		seen[value] = struct{}{}
	}
}
