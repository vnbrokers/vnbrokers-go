package marketdata_test

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/marketdata"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestMarketDataHTTPContracts(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		invoke func(context.Context, marketdata.Service) error
	}{
		{"derivative tickers", "/tartarus/v1/derivatives", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetDerivativeTickers(ctx, dto.GetDerivativeTickersRequest{})
			return err
		}},
		{"stock tickers", "/tartarus/v1/tickerCommons?index=1&tickers=FPT%2CMWG", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockTickers(ctx, dto.GetStockTickersRequest{Tickers: "FPT,MWG", Index: 1})
			return err
		}},
		{"foreign rooms", "/tartarus/v1/tickerSnaps?index=1", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockForeignRooms(ctx, dto.GetStockForeignRoomsRequest{Index: 1})
			return err
		}},
		{"put throughs", "/tartarus/v1/putThroughSnaps?floor=1", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockPutThroughs(ctx, dto.GetStockPutThroughsRequest{Floor: 1})
			return err
		}},
		{"trade history", "/nyx/v1/intraday/TCB%2F1/his/paging?headIndex=-1&page=0&size=20", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockTradeHistory(ctx, dto.GetStockTradeHistoryRequest{Ticker: "TCB/1", Page: 0, Size: 20, HeadIndex: -1})
			return err
		}},
		{"supply demand 15 minutes", "/nyx/v1/intraday/TCB/bsa-ext?tWindow=15&timeWindow=15&type=all", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockSupplyDemand15Minutes(ctx, dto.GetStockSupplyDemand15MinutesRequest{Ticker: "TCB", TimeWindow: "15", TWindow: "15", Type: "all"})
			return err
		}},
		{"supply demand daily", "/nyx/v1/intraday/TCB/bsa?type=all", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockSupplyDemandDaily(ctx, dto.GetStockSupplyDemandDailyRequest{Ticker: "TCB", Type: "all"})
			return err
		}},
		{"supply demand monthly", "/nyx/v1/intraday/TCB/bsa-month?timeWindow=30&type=all", func(ctx context.Context, s marketdata.Service) error {
			_, err := s.GetStockSupplyDemandMonthly(ctx, dto.GetStockSupplyDemandMonthlyRequest{Ticker: "TCB", TimeWindow: "30", Type: "all"})
			return err
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var capability core.Capability
			var request transport.HTTPRequest
			service := marketdata.NewService(marketdata.Dependencies{
				BaseURL:           "https://api.example",
				Headers:           func(bool, bool) map[string]string { return map[string]string{"Authorization": "Bearer token"} },
				RequireCapability: func(value core.Capability) error { capability = value; return nil },
				Send: func(_ context.Context, _ string, value transport.HTTPRequest) (transport.HTTPResponse, error) {
					request = value
					return transport.HTTPResponse{StatusCode: 200, Raw: []byte(`{}`)}, nil
				},
			})

			if err := test.invoke(context.Background(), service); err != nil {
				t.Fatalf("invoke: %v", err)
			}
			if capability == "" {
				t.Fatal("capability was not checked")
			}
			if request.Method != "GET" || request.URL != "https://api.example"+test.url {
				t.Fatalf("request = %s %s", request.Method, request.URL)
			}
			if request.JSON != nil {
				t.Fatalf("unexpected body = %#v", request.JSON)
			}
			if request.Headers["Authorization"] != "Bearer token" {
				t.Fatalf("headers = %#v", request.Headers)
			}
		})
	}
}
