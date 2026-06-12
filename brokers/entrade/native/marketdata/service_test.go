package marketdata

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestGetDerivativesConstructsRequestAndDecodesResponse(t *testing.T) {
	var capability core.Capability
	var request transport.HTTPRequest
	service := NewService(Dependencies{
		BaseURL: "https://entrade.example/api",
		Headers: func(bool) map[string]string { return map[string]string{"Authorization": "Bearer jwt-token"} },
		RequireCapability: func(value core.Capability) error {
			capability = value
			return nil
		},
		Send: func(_ context.Context, _ string, value transport.HTTPRequest) (transport.HTTPResponse, error) {
			request = value
			return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{"data":[{"symbol":"VN30F2512","type":"VN30F1M","extra":true}],"total":1,"unknown":true}`)}, nil
		},
	})

	response, err := service.GetDerivatives(context.Background(), dto.GetDerivativesRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if capability != CapabilityDerivatives || request.Method != "GET" || request.URL != "https://entrade.example/api/derivatives" {
		t.Fatalf("capability=%q request=%#v", capability, request)
	}
	if response.Total != 1 || len(response.Data) != 1 || response.Data[0].Symbol != "VN30F2512" || response.Data[0].Type != "VN30F1M" {
		t.Fatalf("response=%#v", response)
	}
}

func TestGetDerivativesDecodeErrorPreservesRawResponse(t *testing.T) {
	raw := []byte(`{"data":`)
	service := NewService(Dependencies{
		BaseURL:           "https://entrade.example/api",
		Headers:           func(bool) map[string]string { return nil },
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(context.Context, string, transport.HTTPRequest) (transport.HTTPResponse, error) {
			return transport.HTTPResponse{StatusCode: 200, Raw: raw}, nil
		},
	})

	_, err := service.GetDerivatives(context.Background(), dto.GetDerivativesRequest{})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error type=%T error=%v", err, err)
	}
	preserved, ok := brokerErr.Raw.([]byte)
	if brokerErr.Category != sdkerrors.CategoryDecode || !ok || !bytes.Equal(preserved, raw) {
		t.Fatalf("broker error=%#v", brokerErr)
	}
}
