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
		{"foreign trading", "/price/ACB/foreign-trading?boardId=G1&from=1&limit=100&order=desc&to=2", func(s Service) error {
			_, err := s.GetForeignTrading(context.Background(), dto.GetForeignTradingRequest{Symbol: "ACB", BoardID: "G1", From: 1, To: 2, Limit: 100, Order: "desc"})
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
		{"trading session", "/market/trading-session?boardId=G1&tscProdGrpId=STO", func(s Service) error {
			_, err := s.GetTradingSession(context.Background(), dto.GetTradingSessionRequest{BoardID: "G1", TSCProdGrpID: "STO"})
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

func TestMarketDataServiceDecodesTradingSessionResponse(t *testing.T) {
	service := NewService(Dependencies{
		BaseURL:           "https://openapi.dnse.com.vn",
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(_ context.Context, _ string, _ transport.HTTPRequest) (transport.HTTPResponse, error) {
			return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{
				"tradingSessions": [{
					"marketId": "STO",
					"boardId": "G1",
					"tscProdGrpId": "STO",
					"tradingSessionId": "99",
					"eventId": "AC2",
					"time": "2026-06-26 14:45:00.960"
				}]
			}`)}, nil
		},
	})

	response, err := service.GetTradingSession(context.Background(), dto.GetTradingSessionRequest{TSCProdGrpID: "STO"})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.TradingSessions) != 1 {
		t.Fatalf("trading sessions length = %d", len(response.TradingSessions))
	}
	session := response.TradingSessions[0]
	if session.MarketID != "STO" || session.BoardID != "G1" || session.TSCProdGrpID != "STO" || session.TradingSessionID != "99" || session.EventID != "AC2" {
		t.Fatalf("session = %+v", session)
	}
}

func TestMarketDataServiceDecodesForeignTradingResponse(t *testing.T) {
	service := NewService(Dependencies{
		BaseURL:           "https://openapi.dnse.com.vn",
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(_ context.Context, _ string, _ transport.HTTPRequest) (transport.HTTPResponse, error) {
			return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{
				"foreigners": [{
					"marketId": "STO",
					"boardId": "G1",
					"symbol": "ACB",
					"tradingSessionId": "99",
					"sellVolume": 1239300,
					"sellTradedAmount": 32683015000,
					"buyVolume": 1299400,
					"buyTradedAmount": 34345090000,
					"totalSellVolume": 1239352,
					"totalSellTradedAmount": 32684385850,
					"totalBuyVolume": 1299415,
					"totalBuyTradedAmount": 34345486100,
					"foreignerOrderLimitQuantity": 1743293567,
					"foreignerBuyPossibleQuantity": 2053706002,
					"time": "2026-06-11 15:33:00.368"
				}],
				"nextPageToken": "next-token"
			}`)}, nil
		},
	})

	response, err := service.GetForeignTrading(context.Background(), dto.GetForeignTradingRequest{Symbol: "ACB", From: 1, To: 2})
	if err != nil {
		t.Fatal(err)
	}
	if response.NextPageToken != "next-token" {
		t.Fatalf("NextPageToken = %q", response.NextPageToken)
	}
	if len(response.Foreigners) != 1 {
		t.Fatalf("foreigners length = %d", len(response.Foreigners))
	}
	foreign := response.Foreigners[0]
	if foreign.Symbol != "ACB" || foreign.MarketID != "STO" || foreign.Time != "2026-06-11 15:33:00.368" {
		t.Fatalf("foreign = %+v", foreign)
	}
	if foreign.BuyTradedAmount == nil || foreign.BuyTradedAmount.String() != "34345090000" {
		t.Fatalf("BuyTradedAmount = %v", foreign.BuyTradedAmount)
	}
}
