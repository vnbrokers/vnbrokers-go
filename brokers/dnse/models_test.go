package dnse

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func TestUnmarshalRawPayloadDecodesSecurityDefinitionsFromBrunoSchema(t *testing.T) {
	payload := domain.RawPayload{
		Source: "dnse",
		Data: []any{
			map[string]any{
				"marketId":                         "STO",
				"boardId":                          "G1",
				"isin":                             "VN000000HPG4",
				"symbol":                           "HPG",
				"productGrpId":                     "STO",
				"securityGroupId":                  "ST",
				"basicPrice":                       26.8,
				"ceilingPrice":                     28.65,
				"floorPrice":                       24.95,
				"securityStatus":                   "UNSPECIFIED",
				"symbolAdminStatusCode":            "NRM",
				"symbolTradingMethodStatusCode":    "NRM",
				"symbolTradingSanctionStatusCode":  "NRM",
				"finalTradeDate":                   nil,
				"listingDate":                      "2007-11-15T00:00:00Z",
				"time":                             "2026-05-26 22:01:01.303",
			},
		},
	}

	var secdefs SecurityDefinitionList
	if err := UnmarshalRawPayload(payload, &secdefs); err != nil {
		t.Fatalf("unmarshal secdef: %v", err)
	}
	if len(secdefs) != 1 {
		t.Fatalf("secdefs len = %d", len(secdefs))
	}
	if secdefs[0].MarketID != "STO" {
		t.Fatalf("market id = %s", secdefs[0].MarketID)
	}
	if secdefs[0].ProductGrpID != "STO" {
		t.Fatalf("product group id = %s", secdefs[0].ProductGrpID)
	}
	if secdefs[0].FloorPrice == nil || !secdefs[0].FloorPrice.Equal(decimal.RequireFromString("24.95")) {
		t.Fatalf("floor price = %v", secdefs[0].FloorPrice)
	}
	floorPrice, ok := secdefs.FloorPrice("HPG")
	if !ok || !floorPrice.Equal(decimal.RequireFromString("24.95")) {
		t.Fatalf("floor price lookup = %v, %t", floorPrice, ok)
	}
}

func TestUnmarshalRawPayloadDecodesLoanPackagesFromBrunoSchema(t *testing.T) {
	payload := domain.RawPayload{
		Source: "dnse",
		Bytes: []byte(`{
			"symbolType": "",
			"marketType": "STOCK",
			"loanPackages": [{
				"id": 1775,
				"name": "GD Tien mat",
				"initialRate": 100,
				"maintenanceRate": 80,
				"liquidRate": 70,
				"tradingFee": {
					"id": 1,
					"name": "default",
					"scope": "ALL",
					"channel": "ONLINE",
					"schemaType": "PROGRESS",
					"createdDate": "2026-01-01T00:00:00Z",
					"modifiedDate": "2026-01-02T00:00:00Z",
					"fixedTradingFee": 0,
					"fixedDailyCloseTradingFee": 0,
					"progressTradingFee": [{"fromQuantity": 0, "toQuantity": 1000, "fee": 0.001}],
					"progressDailyCloseTradingFee": [{"fromQuantity": 0, "toQuantity": 1000, "fee": 0.002}]
				}
			}]
		}`),
	}

	var response LoanPackagesResponse
	if err := UnmarshalRawPayload(payload, &response); err != nil {
		t.Fatalf("unmarshal loan packages: %v", err)
	}
	if response.MarketType != "STOCK" {
		t.Fatalf("market type = %s", response.MarketType)
	}
	if len(response.LoanPackages) != 1 {
		t.Fatalf("loan package len = %d", len(response.LoanPackages))
	}
	pkg := response.LoanPackages[0]
	if pkg.ID != 1775 {
		t.Fatalf("loan package id = %d", pkg.ID)
	}
	if pkg.MaintenanceRate == nil || !pkg.MaintenanceRate.Equal(decimal.NewFromInt(80)) {
		t.Fatalf("maintenance rate = %v", pkg.MaintenanceRate)
	}
	if pkg.TradingFee == nil || len(pkg.TradingFee.ProgressTradingFee) != 1 {
		t.Fatalf("trading fee = %+v", pkg.TradingFee)
	}
}
