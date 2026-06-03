package ssi

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func TestUnmarshalRawPayloadDecodesMarketDataSecuritiesSchema(t *testing.T) {
	payload := domain.RawPayload{
		Source: "ssi",
		Data: map[string]any{
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
	}

	var response DataResponse[Security]
	if err := UnmarshalRawPayload(payload, &response); err != nil {
		t.Fatalf("unmarshal securities: %v", err)
	}
	if response.Status != "Success" {
		t.Fatalf("status = %s", response.Status)
	}
	if response.TotalRecord != 1 {
		t.Fatalf("total record = %d", response.TotalRecord)
	}
	if len(response.Data) != 1 || response.Data[0].Symbol != "AAA" {
		t.Fatalf("data = %+v", response.Data)
	}
}

func TestUnmarshalRawPayloadDecodesTradingStockBalancesSchema(t *testing.T) {
	payload := domain.RawPayload{
		Source: "ssi",
		Bytes: []byte(`{
			"message": "Success",
			"status": 200,
			"data": [{
				"account": "0901358",
				"cashbal": 100000000,
				"cashonhold": 0,
				"secureamount": 0,
				"withdrawable": 100000000,
				"receivingcasht1": 0,
				"receivingcasht2": 0,
				"matchedbuyvolume": 0,
				"matchedsellvolume": 0,
				"unmatchedbuyvolume": 0,
				"unmatchedsellvolume": 0,
				"paidcasht1": 0,
				"paidcasht2": 0,
				"cia": 0,
				"debt": 0,
				"purchasingpower": 100000000,
				"totalasset": 100000000
			}]
		}`),
	}

	var response TradingResponse[[]StockAccountBalance]
	if err := UnmarshalRawPayload(payload, &response); err != nil {
		t.Fatalf("unmarshal stock balances: %v", err)
	}
	if response.Status != 200 {
		t.Fatalf("status = %d", response.Status)
	}
	if len(response.Data) != 1 {
		t.Fatalf("balance len = %d", len(response.Data))
	}
	if response.Data[0].Account != "0901358" {
		t.Fatalf("account = %s", response.Data[0].Account)
	}
	if response.Data[0].PurchasingPower == nil || !response.Data[0].PurchasingPower.Equal(decimal.NewFromInt(100000000)) {
		t.Fatalf("purchasing power = %v", response.Data[0].PurchasingPower)
	}
}
