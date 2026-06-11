package marketdata_test

import (
	"context"
	"testing"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	ssi "github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestSecuritiesBuildsSSIDataRequestAndDecodesResponse(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{
						"Market":      "HOSE",
						"Symbol":      "AAA",
						"StockName":   "CTCP NHUA&MT XANH AN PHAT",
						"StockEnName": "An Phat Bioplastics Joint Stock Company",
					},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetSecurities(context.Background(), nativedto.GetSecuritiesRequest{
		Market:    "HOSE",
		PageIndex: 2,
		PageSize:  20,
	})
	if err != nil {
		t.Fatalf("securities: %v", err)
	}
	if response.Status != "Success" || response.TotalRecord != 1 || len(response.Data) != 1 {
		t.Fatalf("response = %+v", response)
	}
	if response.Data[0].Symbol != "AAA" || response.Data[0].Market != "HOSE" {
		t.Fatalf("security = %+v", response.Data[0])
	}

	request := httpTransport.requests[0]
	if request.Method != "GET" {
		t.Fatalf("method = %s", request.Method)
	}
	wantURL := "https://data.ssi.example/api/v2/Market/Securities?market=HOSE&pageIndex=2&pageSize=20"
	if request.URL != wantURL {
		t.Fatalf("url = %s, want %s", request.URL, wantURL)
	}
	if request.Headers["Authorization"] != "Bearer data-token" {
		t.Fatalf("authorization = %s", request.Headers["Authorization"])
	}
	if request.Headers["Accept"] != "application/json" {
		t.Fatalf("accept = %s", request.Headers["Accept"])
	}
}

func TestSecuritiesUsesPaginationDefaultsAndOmitsEmptyMarket(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data":        []any{},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 0,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example/",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	if _, err := broker.Native().MarketData().GetSecurities(context.Background(), nativedto.GetSecuritiesRequest{}); err != nil {
		t.Fatalf("securities: %v", err)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/Securities?pageIndex=1&pageSize=10"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestSecuritiesDetailsBuildsSSIDataRequestAndDecodesResponse(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{
						"RType":      "y",
						"ReportDate": "19/01/2023",
						"TotalNoSym": "1",
						"RepeatedInfo": []any{
							map[string]any{
								"Symbol":        "SSI",
								"SymbolName":    "CTCP CHUNG KHOAN SSI",
								"SymbolEngName": "SSI Securities Corporation",
							},
						},
					},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetSecuritiesDetails(context.Background(), nativedto.GetSecuritiesDetailsRequest{
		Market:    "HOSE",
		Symbol:    "SSI",
		PageIndex: 2,
		PageSize:  20,
	})
	if err != nil {
		t.Fatalf("securities details: %v", err)
	}
	if response.TotalRecord != 1 || len(response.Data) != 1 || response.Data[0].RepeatedInfo[0].Symbol != "SSI" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/SecuritiesDetails?market=HOSE&pageIndex=2&pageSize=20&symbol=SSI"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestIndexComponentsBuildsSSIDataRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{
						"IndexCode": "VN30",
						"IndexComponent": []any{
							map[string]any{"Isin": "ACB", "StockSymbol": "ACB"},
						},
					},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetIndexComponents(context.Background(), nativedto.GetIndexComponentsRequest{
		IndexCode: "VN30",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("index components: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].IndexComponent[0].StockSymbol != "ACB" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/IndexComponents?indexCode=VN30&pageIndex=1&pageSize=10"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestIndexListBuildsSSIDataRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{"IndexCode": "VN30", "IndexName": "VN30", "Exchange": "HOSE"},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetIndexList(context.Background(), nativedto.GetIndexListRequest{
		Exchange:  "HOSE",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("index list: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].IndexCode != "VN30" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/IndexList?exchange=HOSE&pageIndex=1&pageSize=10"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestDailyOhlcUsesDefaultsAndOmitsEmptyOptionals(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{"Symbol": "SSI", "Open": "28600", "Close": "28100"},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetDailyOhlc(context.Background(), nativedto.GetDailyOhlcRequest{
		Symbol:   "SSI",
		FromDate: "10/08/2023",
		ToDate:   "13/08/2023",
	})
	if err != nil {
		t.Fatalf("daily ohlc: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].Close != "28100" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/DailyOhlc?fromDate=10%2F08%2F2023&pageIndex=1&pageSize=10&symbol=SSI&toDate=13%2F08%2F2023"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestIntradayOhlcIncludesResolutionAndAscending(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{"Symbol": "SSI", "Time": "14:45:04", "Volume": "529200"},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetIntradayOhlc(context.Background(), nativedto.GetIntradayOhlcRequest{
		Symbol:     "SSI",
		FromDate:   "14/08/2023",
		ToDate:     "14/08/2023",
		PageIndex:  1,
		PageSize:   10,
		Ascending:  true,
		Resolution: 1,
	})
	if err != nil {
		t.Fatalf("intraday ohlc: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].Time != "14:45:04" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/IntradayOhlc?ascending=true&fromDate=14%2F08%2F2023&pageIndex=1&pageSize=10&resolution=1&symbol=SSI&toDate=14%2F08%2F2023"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestDailyIndexBuildsSSIDataRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{"IndexId": "HNX30", "IndexValue": "510.56", "TradingSession": "C"},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetDailyIndex(context.Background(), nativedto.GetDailyIndexRequest{
		IndexID:   "HNX30",
		FromDate:  "14/08/2023",
		ToDate:    "14/08/2023",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("daily index: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].IndexID != "HNX30" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/DailyIndex?fromDate=14%2F08%2F2023&indexId=HNX30&pageIndex=1&pageSize=10&toDate=14%2F08%2F2023"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}

func TestDailyStockPriceBuildsSSIDataRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"data": []any{
					map[string]any{"Symbol": "HUB", "ClosePrice": "20000", "TotalTradedValue": "380230000"},
				},
				"message":     "Success",
				"status":      "Success",
				"totalRecord": 1,
			},
		}},
	}
	broker := ssi.NewBroker(ssi.Config{
		DataBaseURL:   "https://data.ssi.example",
		DataToken:     "data-token",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Native().MarketData().GetDailyStockPrice(context.Background(), nativedto.GetDailyStockPriceRequest{
		Symbol:    "SSI",
		Market:    "HOSE",
		FromDate:  "19/07/2023",
		ToDate:    "19/07/2023",
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("daily stock price: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].ClosePrice != "20000" {
		t.Fatalf("response = %+v", response)
	}

	wantURL := "https://data.ssi.example/api/v2/Market/DailyStockPrice?fromDate=19%2F07%2F2023&market=HOSE&pageIndex=1&pageSize=10&symbol=SSI&toDate=19%2F07%2F2023"
	if got := httpTransport.requests[0].URL; got != wantURL {
		t.Fatalf("url = %s, want %s", got, wantURL)
	}
}
