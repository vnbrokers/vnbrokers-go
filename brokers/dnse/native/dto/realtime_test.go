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
