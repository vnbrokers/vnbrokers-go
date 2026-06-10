package dto

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSecuritiesRequestMarshal(t *testing.T) {
	request := SecuritiesRequest{
		Market:    "hose",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal SecuritiesRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"market":    "hose",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestSecuritiesResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"Market": "HOSE",
				"Symbol": "AAA",
				"StockName": "CTCP NHUA&MT XANH AN PHAT",
				"StockEnName": "An Phat Bioplastics Joint Stock Company"
			},
			{
				"Market": "HOSE",
				"Symbol": "AAM",
				"StockName": "CTCP THUY SAN MEKONG",
				"StockEnName": "Mekong Fisheries Joint Stock Company"
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 2
	}`)

	var response SecuritiesResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal SecuritiesResponse: %v", err)
	}

	if response.Message != "Success" || response.Status != "Success" || response.TotalRecord != 2 {
		t.Fatalf("response metadata = message %q, status %q, total %d", response.Message, response.Status, response.TotalRecord)
	}
	want := []Security{
		{
			Market:      "HOSE",
			Symbol:      "AAA",
			StockName:   "CTCP NHUA&MT XANH AN PHAT",
			StockEnName: "An Phat Bioplastics Joint Stock Company",
		},
		{
			Market:      "HOSE",
			Symbol:      "AAM",
			StockName:   "CTCP THUY SAN MEKONG",
			StockEnName: "Mekong Fisheries Joint Stock Company",
		},
	}
	if !reflect.DeepEqual(response.Data, want) {
		t.Fatalf("response data = %#v, want %#v", response.Data, want)
	}
}

func TestSecuritiesDetailsRequestMarshal(t *testing.T) {
	request := SecuritiesDetailsRequest{
		Market:    "HOSE",
		Symbol:    "SSI",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal SecuritiesDetailsRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"market":    "HOSE",
		"symbol":    "SSI",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestSecuritiesDetailsResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"RType": "y",
				"ReportDate": "19/01/2023",
				"TotalNoSym": "1",
				"RepeatedInfo": [
					{
						"Isin": null,
						"Symbol": "SSI",
						"SymbolName": "CTCP CHUNG KHOAN SSI",
						"SymbolEngName": "SSI Securities Corporation",
						"SecType": "S",
						"MarketId": "HOSE",
						"Exchange": "HOSE",
						"Issuer": null,
						"LotSize": "100",
						"IssueDate": "",
						"MaturityDate": "",
						"FirstTradingDate": "",
						"LastTradingDate": "",
						"ContractMultiplier": "0",
						"SettlMethod": "",
						"Underlying": null,
						"PutOrCall": null,
						"ExercisePrice": "0",
						"ExerciseStyle": "",
						"ExcerciseRatio": "0",
						"ListedShare": "1501130137",
						"TickPrice1": null,
						"TickIncrement1": null,
						"TickPrice2": null,
						"TickIncrement2": null,
						"TickPrice3": null,
						"TickIncrement3": null,
						"TickPrice4": null,
						"TickIncrement4": null
					}
				]
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response SecuritiesDetailsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal SecuritiesDetailsResponse: %v", err)
	}
	if response.Status != "Success" || response.TotalRecord != 1 || len(response.Data) != 1 {
		t.Fatalf("response = %+v", response)
	}
	group := response.Data[0]
	if group.RType != "y" || group.ReportDate != "19/01/2023" || group.TotalNoSym != "1" {
		t.Fatalf("group = %+v", group)
	}
	if len(group.RepeatedInfo) != 1 {
		t.Fatalf("repeated info len = %d", len(group.RepeatedInfo))
	}
	item := group.RepeatedInfo[0]
	if item.Symbol != "SSI" || item.SymbolName != "CTCP CHUNG KHOAN SSI" || item.SymbolEngName != "SSI Securities Corporation" {
		t.Fatalf("item = %+v", item)
	}
	if item.MarketID != "HOSE" || item.Exchange != "HOSE" || item.ListedShare != "1501130137" {
		t.Fatalf("item = %+v", item)
	}
}

func TestIndexComponentsRequestMarshal(t *testing.T) {
	request := IndexComponentsRequest{
		IndexCode: "VN30",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal IndexComponentsRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"indexCode": float64(0),
	}
	delete(want, "indexCode")
	want["indexCode"] = "VN30"
	want["pageIndex"] = float64(1)
	want["pageSize"] = float64(10)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestIndexComponentsResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"IndexCode": "VN30",
				"IndexName": "VN30",
				"Exchange": "HOSE",
				"TotalSymbolNo": "30",
				"IndexComponent": [
					{"Isin": "ACB", "StockSymbol": "ACB"},
					{"Isin": "BCM", "StockSymbol": "BCM"}
				]
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response IndexComponentsResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal IndexComponentsResponse: %v", err)
	}
	if response.TotalRecord != 1 || len(response.Data) != 1 {
		t.Fatalf("response = %+v", response)
	}
	group := response.Data[0]
	if group.IndexCode != "VN30" || group.TotalSymbolNo != "30" || len(group.IndexComponent) != 2 {
		t.Fatalf("group = %+v", group)
	}
	if group.IndexComponent[0].Isin != "ACB" || group.IndexComponent[1].StockSymbol != "BCM" {
		t.Fatalf("components = %+v", group.IndexComponent)
	}
}

func TestIndexListRequestMarshal(t *testing.T) {
	request := IndexListRequest{
		Exchange:  "HOSE",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal IndexListRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"exchange":  "HOSE",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestIndexListResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{"IndexCode": "VN100", "IndexName": "VN100", "Exchange": "HOSE"},
			{"IndexCode": "VN30", "IndexName": "VN30", "Exchange": "HOSE"}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 2
	}`)

	var response IndexListResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal IndexListResponse: %v", err)
	}
	if response.TotalRecord != 2 || len(response.Data) != 2 {
		t.Fatalf("response = %+v", response)
	}
	if response.Data[0].IndexCode != "VN100" || response.Data[1].IndexName != "VN30" {
		t.Fatalf("data = %+v", response.Data)
	}
}

func TestDailyOhlcRequestMarshal(t *testing.T) {
	request := DailyOhlcRequest{
		Symbol:    "SSI",
		FromDate:  "10/08/2023",
		ToDate:    "13/08/2023",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal DailyOhlcRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"symbol":    "SSI",
		"fromDate":  "10/08/2023",
		"toDate":    "13/08/2023",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
		"ascending": false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestDailyOhlcResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"Symbol": "SSI",
				"Market": "HOSE",
				"TradingDate": "10/08/2023",
				"Time": null,
				"Open": "28600",
				"High": "28850",
				"Low": "28100",
				"Close": "28100",
				"Volume": "23382100",
				"Value": "663258204999.9850"
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response DailyOhlcResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal DailyOhlcResponse: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].Open != "28600" || response.Data[0].Value != "663258204999.9850" {
		t.Fatalf("response = %+v", response)
	}
}

func TestIntradayOhlcRequestMarshal(t *testing.T) {
	request := IntradayOhlcRequest{
		Symbol:     "SSI",
		FromDate:   "14/08/2023",
		ToDate:     "14/08/2023",
		PageIndex:  1,
		PageSize:   10,
		Resolution: 1,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal IntradayOhlcRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"symbol":     "SSI",
		"fromDate":   "14/08/2023",
		"toDate":     "14/08/2023",
		"pageIndex":  float64(1),
		"pageSize":   float64(10),
		"ascending":  false,
		"resolution": float64(1),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestIntradayOhlcResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"Symbol": "SSI",
				"Market": "HOSE",
				"TradingDate": "14/08/2023",
				"Time": "14:45:04",
				"Open": "29150",
				"High": "29150",
				"Low": "29150",
				"Close": "29150",
				"Volume": "529200",
				"Value": "29150"
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response IntradayOhlcResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal IntradayOhlcResponse: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].Time != "14:45:04" || response.Data[0].Volume != "529200" {
		t.Fatalf("response = %+v", response)
	}
}

func TestDailyIndexRequestMarshal(t *testing.T) {
	request := DailyIndexRequest{
		IndexID:   "HNX30",
		FromDate:  "14/08/2023",
		ToDate:    "14/08/2023",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal DailyIndexRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"indexId":   "HNX30",
		"fromDate":  "14/08/2023",
		"toDate":    "14/08/2023",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
		"ascending": false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestDailyIndexResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"IndexId": "HNX30",
				"IndexValue": "510.56",
				"TradingDate": "14/08/2023",
				"Time": null,
				"Change": "19.09",
				"RatioChange": "3.89",
				"TotalTrade": "0",
				"TotalMatchVol": "84693600",
				"TotalMatchVal": "1836008470000",
				"TypeIndex": null,
				"IndexName": "HNX30",
				"Advances": "21",
				"NoChanges": "4",
				"Declines": "5",
				"Ceilings": "2",
				"Floors": "0",
				"TotalDealVol": "2504000",
				"TotalDealVal": "60256000000",
				"TotalVol": "87197600",
				"TotalVal": "1896264470000",
				"TradingSession": "C"
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response DailyIndexResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal DailyIndexResponse: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].IndexID != "HNX30" || response.Data[0].IndexValue != "510.56" {
		t.Fatalf("response = %+v", response)
	}
	if response.Data[0].Advances != "21" || response.Data[0].TradingSession != "C" {
		t.Fatalf("record = %+v", response.Data[0])
	}
}

func TestDailyStockPriceRequestMarshal(t *testing.T) {
	request := DailyStockPriceRequest{
		Symbol:    "SSI",
		Market:    "HOSE",
		FromDate:  "19/07/2023",
		ToDate:    "19/07/2023",
		PageIndex: 1,
		PageSize:  10,
	}

	data, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal DailyStockPriceRequest: %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal request JSON: %v", err)
	}
	want := map[string]any{
		"symbol":    "SSI",
		"market":    "HOSE",
		"fromDate":  "19/07/2023",
		"toDate":    "19/07/2023",
		"pageIndex": float64(1),
		"pageSize":  float64(10),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("request JSON = %#v, want %#v", got, want)
	}
}

func TestDailyStockPriceResponseUnmarshal(t *testing.T) {
	data := []byte(`{
		"data": [
			{
				"TradingDate": "19/07/2023",
				"PriceChange": "-150",
				"PerPriceChange": "-0.70",
				"CeilingPrice": "21550",
				"FloorPrice": "18750",
				"RefPrice": "20150",
				"OpenPrice": "20950",
				"HighestPrice": "20950",
				"LowestPrice": "20000",
				"ClosePrice": "20000",
				"AveragePrice": "20118",
				"ClosePriceAdjusted": "17392",
				"TotalMatchVol": "18900",
				"TotalMatchVal": "380230000",
				"TotalDealVal": "0",
				"TotalDealVol": "0",
				"ForeignBuyVolTotal": "0",
				"ForeignCurrentRoom": "0",
				"ForeignSellVolTotal": "0",
				"ForeignBuyValTotal": "0",
				"ForeignSellValTotal": "0",
				"TotalBuyTrade": "0",
				"TotalBuyTradeVol": "0",
				"TotalSellTrade": "0",
				"TotalSellTradeVol": "0",
				"NetBuySellVol": "0",
				"NetBuySellVal": "0",
				"TotalTradedVol": "18900",
				"TotalTradedValue": "380230000",
				"Symbol": "HUB",
				"Time": null
			}
		],
		"message": "Success",
		"status": "Success",
		"totalRecord": 1
	}`)

	var response DailyStockPriceResponse
	if err := json.Unmarshal(data, &response); err != nil {
		t.Fatalf("unmarshal DailyStockPriceResponse: %v", err)
	}
	if len(response.Data) != 1 || response.Data[0].Symbol != "HUB" || response.Data[0].ClosePrice != "20000" {
		t.Fatalf("response = %+v", response)
	}
	if response.Data[0].TotalTradedValue != "380230000" || response.Data[0].PerPriceChange != "-0.70" {
		t.Fatalf("record = %+v", response.Data[0])
	}
}
