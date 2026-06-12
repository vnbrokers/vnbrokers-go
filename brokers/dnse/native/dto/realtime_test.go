package dto

import (
	"encoding/json"
	"testing"
)

func TestForeignEventDecodesWireFieldNames(t *testing.T) {
	var event ForeignEvent
	if err := json.Unmarshal([]byte(`{"marketId":"STO","boardId":"G1","symbol":"FPT","sellVolume":10,"currentRoom":123.5,"additiveField":true}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.MarketID != "STO" || event.BoardID != "G1" || event.Symbol != "FPT" {
		t.Fatalf("event = %+v", event)
	}
	if event.CurrentRoom == nil || event.CurrentRoom.String() != "123.5" {
		t.Fatalf("current room = %v", event.CurrentRoom)
	}
}

func TestMarketIndexEventAllowsNullableCounts(t *testing.T) {
	var event MarketIndexEvent
	if err := json.Unmarshal([]byte(`{"T":"mi","indexName":"VNINDEX","fluctuationLowerLimitIssueCount":null,"fluctuationUpperLimitIssueCount":7}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.FluctuationLowerLimitIssueCount != nil {
		t.Fatalf("lower-limit count = %v", event.FluctuationLowerLimitIssueCount)
	}
	if event.FluctuationUpperLimitIssueCount == nil || *event.FluctuationUpperLimitIssueCount != 7 {
		t.Fatalf("upper-limit count = %v", event.FluctuationUpperLimitIssueCount)
	}
}

func TestMarketIndexEventDecodesCurrentWireFields(t *testing.T) {
	var event MarketIndexEvent
	if err := json.Unmarshal([]byte(`{
		"T":"mi",
		"blkTrdAccTrdVal":41.69273115,
		"blkTrdAccTrdVol":1738007,
		"changedRatio":0.91,
		"changedValue":17.74,
		"contauctAccTrdVal":2369.54154935,
		"contauctAccTrdVol":85122929,
		"currencyCode":"VND",
		"grossTradeAmount":2411.2342805,
		"highestValueIndexes":1969.19,
		"indexName":"VN30",
		"indexTypeCode":"101",
		"lowestValueIndexes":1947.28,
		"marketId":"STO",
		"marketIndexClass":"HSX",
		"multicastReceiveTime":{"Nanos":109572169,"Seconds":1781234600},
		"priorValueIndexes":1947.28,
		"totalVolumeTraded":86860936,
		"tradingSessionId":"40",
		"transactTime":{"Nanos":63000000,"Seconds":1781234600},
		"valueIndexes":1965.02
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.IndexName != "VN30" || event.MarketID != "STO" || event.CurrencyCode != "VND" {
		t.Fatalf("event = %+v", event)
	}
	if event.ValueIndexes == nil || event.ValueIndexes.String() != "1965.02" {
		t.Fatalf("value indexes = %v", event.ValueIndexes)
	}
	if event.GrossTradeAmount == nil || event.GrossTradeAmount.String() != "2411.2342805" {
		t.Fatalf("gross trade amount = %v", event.GrossTradeAmount)
	}
	if event.TotalVolumeTraded != 86860936 || event.TransactTime.Seconds != 1781234600 {
		t.Fatalf("volume/time = %d/%+v", event.TotalVolumeTraded, event.TransactTime)
	}
}

func TestEstimatedMarketIndexEventDecodesWireFieldNames(t *testing.T) {
	var event EstimatedMarketIndexEvent
	if err := json.Unmarshal([]byte(`{
		"indexName":"VN30",
		"valueIndexes":1948.57,
		"changedValue":-37.71,
		"changedRatio":-1.9,
		"fluctuationUpIssueCount":5,
		"fluctuationDownIssueCount":25,
		"fluctuationSteadinessIssueCount":0,
		"grossTradeAmount":6391.86,
		"totalVolumeTraded":184907600,
		"time":"2026-06-08 13:56:29.371"
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.IndexName != "VN30" || event.Time != "2026-06-08 13:56:29.371" {
		t.Fatalf("event = %+v", event)
	}
	if event.ValueIndexes == nil || event.ValueIndexes.String() != "1948.57" {
		t.Fatalf("value indexes = %v", event.ValueIndexes)
	}
	if event.ChangedValue == nil || event.ChangedValue.String() != "-37.71" {
		t.Fatalf("changed value = %v", event.ChangedValue)
	}
	if event.ChangedRatio == nil || event.ChangedRatio.String() != "-1.9" {
		t.Fatalf("changed ratio = %v", event.ChangedRatio)
	}
	if event.GrossTradeAmount == nil || event.GrossTradeAmount.String() != "6391.86" {
		t.Fatalf("gross trade amount = %v", event.GrossTradeAmount)
	}
	if event.FluctuationUpIssueCount != 5 || event.FluctuationDownIssueCount != 25 || event.FluctuationSteadinessIssueCount != 0 {
		t.Fatalf("fluctuation counts = %d/%d/%d", event.FluctuationUpIssueCount, event.FluctuationDownIssueCount, event.FluctuationSteadinessIssueCount)
	}
	if event.TotalVolumeTraded != 184907600 {
		t.Fatalf("total volume traded = %d", event.TotalVolumeTraded)
	}
}

func TestBrokerOrderEventDecodesWireFieldNames(t *testing.T) {
	var event BrokerOrderEvent
	if err := json.Unmarshal([]byte(`{
		"id":596,
		"side":"NS",
		"accountNo":"0001179019",
		"symbol":"41I1G5000",
		"orderType":"LO",
		"price":1920.0,
		"quantity":5,
		"fillQuantity":2,
		"canceledQuantity":0,
		"leaveQuantity":3,
		"orderStatus":"PartiallyFilled",
		"loanPackageId":2278,
		"marketType":"DERIVATIVE",
		"transDate":"2026-04-06T00:00:00Z",
		"createdDate":"2026-04-13T04:24:05.274Z",
		"modifiedDate":"2026-04-13T04:28:27.749Z"
	}`), &event); err != nil {
		t.Fatal(err)
	}
	if event.ID != 596 || event.AccountNo != "0001179019" || event.Symbol != "41I1G5000" {
		t.Fatalf("event = %+v", event)
	}
	if event.Price == nil || event.Price.String() != "1920" {
		t.Fatalf("price = %v", event.Price)
	}
	if event.Quantity == nil || event.Quantity.String() != "5" || event.FillQuantity == nil || event.FillQuantity.String() != "2" {
		t.Fatalf("quantities = %v/%v", event.Quantity, event.FillQuantity)
	}
	if event.MarketType != "DERIVATIVE" || event.OrderStatus != "PartiallyFilled" || event.LoanPackageID != 2278 {
		t.Fatalf("event = %+v", event)
	}
}
