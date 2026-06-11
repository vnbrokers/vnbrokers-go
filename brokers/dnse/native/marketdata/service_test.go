package marketdata

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestMarketDataServiceBuildsTypedRequests(t *testing.T) {
	tests := []struct {
		name string
		want string
		call func(Service) error
	}{
		{"trade history", "/price/ACB/trades?boardId=G1&from=1&limit=100&to=2", func(s Service) error {
			_, err := s.GetTradeHistory(context.Background(), dto.GetTradeHistoryRequest{Symbol: "ACB", BoardID: "G1", From: 1, To: 2, Limit: 100})
			return err
		}},
		{"instrument details", "/instruments?indexName=VN30&limit=10&marketId=STO&page=1&securityGroupId=ST&symbol=ACB", func(s Service) error {
			_, err := s.GetInstrumentDetails(context.Background(), dto.GetInstrumentDetailsRequest{Symbol: "ACB", MarketID: "STO", SecurityGroupID: "ST", IndexName: "VN30", Limit: 10, Page: 1})
			return err
		}},
		{"instruments", "/instruments?limit=1000&symbol=ACB", func(s Service) error {
			_, err := s.GetInstruments(context.Background(), dto.GetInstrumentsRequest{Symbol: "ACB", Limit: 1000})
			return err
		}},
		{"latest quotes", "/price/ACB/quotes/latest?boardId=G1", func(s Service) error {
			_, err := s.GetLatestQuotes(context.Background(), dto.GetLatestQuotesRequest{Symbol: "ACB", BoardID: "G1"})
			return err
		}},
		{"latest trades", "/price/ACB/trades/latest?boardId=G1", func(s Service) error {
			_, err := s.GetLatestTrades(context.Background(), dto.GetLatestTradesRequest{Symbol: "ACB", BoardID: "G1"})
			return err
		}},
		{"ohlc", "/price/ohlc?from=1&resolution=15&symbol=ACB&to=2&type=STOCK", func(s Service) error {
			_, err := s.GetOHLC(context.Background(), dto.GetOHLCRequest{Symbol: "ACB", Type: "STOCK", Resolution: "15", From: 1, To: 2})
			return err
		}},
		{"close price", "/price/ACB/close?boardId=G1", func(s Service) error {
			_, err := s.GetClosePrice(context.Background(), dto.GetClosePriceRequest{Symbol: "ACB", BoardID: "G1"})
			return err
		}},
		{"quote history", "/price/ACB/quotes?boardId=G1&from=1&limit=100&to=2", func(s Service) error {
			_, err := s.GetQuoteHistory(context.Background(), dto.GetQuoteHistoryRequest{Symbol: "ACB", BoardID: "G1", From: 1, To: 2, Limit: 100})
			return err
		}},
		{"security definition", "/price/ACB/secdef?boardId=G1", func(s Service) error {
			_, err := s.GetSecurityDefinition(context.Background(), dto.GetSecurityDefinitionRequest{Symbol: "ACB", BoardID: "G1"})
			return err
		}},
		{"working dates", "/market/working-dates", func(s Service) error {
			_, err := s.GetWorkingDates(context.Background(), dto.GetWorkingDatesRequest{})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got transport.HTTPRequest
			service := NewService(Dependencies{
				BaseURL:           "https://openapi.dnse.com.vn",
				RequireCapability: func(core.Capability) error { return nil },
				Send: func(_ context.Context, _ string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
					got = request
					return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`null`)}, nil
				},
			})
			if err := tt.call(service); err != nil {
				t.Fatal(err)
			}
			if got.URL != "https://openapi.dnse.com.vn"+tt.want {
				t.Fatalf("URL = %q", got.URL)
			}
		})
	}
}
