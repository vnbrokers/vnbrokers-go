package dnse

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

func TestListPositionsBuildsDNSERequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"positions": []any{
					map[string]any{
						"accountNo":    "0001179019",
						"symbol":       "VN30F2506",
						"openQuantity": 3,
					},
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:           "https://api.dnse.example",
		APIKey:            "key",
		MarketType:        "STOCK",
		PositionsPageSize: 50,
		HTTPTransport:     httpTransport,
	})

	positions, err := broker.Trading().Positions().List(context.Background(), "0001179019")
	if err != nil {
		t.Fatalf("list positions: %v", err)
	}
	if got := httpTransport.requests[0].URL; got != "https://api.dnse.example/accounts/0001179019/positions?marketType=STOCK&pageSize=50" {
		t.Fatalf("url = %s", got)
	}
	if positions[0].Symbol != "VN30F2506" {
		t.Fatalf("symbol = %s", positions[0].Symbol)
	}
}
