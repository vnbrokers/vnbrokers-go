package dto_test

import (
	"encoding/json"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func TestPrimitiveOpenAPIContractsDecode(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		decode  func([]byte) error
		check   func(t *testing.T)
	}{
		{
			name:    "account",
			payload: `{"basicInfo":{"code105C":"105C001"},"bankAccounts":[{"bankAccountNumber":"123"}]}`,
			decode: func(data []byte) error {
				return json.Unmarshal(data, &accountResponse)
			},
			check: func(t *testing.T) {
				if accountResponse.BasicInfo.Code105C != "105C001" || len(accountResponse.BankAccounts) != 1 {
					t.Fatalf("account response = %#v", accountResponse)
				}
			},
		},
		{
			name:    "stock orders",
			payload: `{"pageSize":10,"pageIndex":1,"totalCount":1,"data":[{"orderID":"42","orderQtty":100.5}]}`,
			decode: func(data []byte) error {
				return json.Unmarshal(data, &ordersResponse)
			},
			check: func(t *testing.T) {
				if ordersResponse.PageSize != 10 || ordersResponse.Data[0].OrderQtty != 100.5 {
					t.Fatalf("orders response = %#v", ordersResponse)
				}
			},
		},
		{
			name:    "purchasing power",
			payload: `{"accountNo":"0001","price":120000,"ppse":50000000001,"maxBuyQuantity":458600}`,
			decode: func(data []byte) error {
				return json.Unmarshal(data, &purchasingPowerResponse)
			},
			check: func(t *testing.T) {
				if purchasingPowerResponse.Price != 120000 || purchasingPowerResponse.MaxBuyQuantity != 458600 {
					t.Fatalf("purchasing response = %#v", purchasingPowerResponse)
				}
			},
		},
		{
			name:    "derivative cash",
			payload: `{"cmd":"Web.Portfolio.AccountStatus","rc":"1","data":{"cash":600089123,"vm":-4200000}}`,
			decode: func(data []byte) error {
				return json.Unmarshal(data, &derivativeCashResponse)
			},
			check: func(t *testing.T) {
				if derivativeCashResponse.Data.Cash != 600089123 || derivativeCashResponse.Data.VM != -4200000 {
					t.Fatalf("derivative cash response = %#v", derivativeCashResponse)
				}
			},
		},
		{
			name:    "stock tickers",
			payload: `{"data":[{"symbol":"FPT","ceilPrice":106100,"totalVol":833900}],"tradingDate":"29/10/2024"}`,
			decode: func(data []byte) error {
				return json.Unmarshal(data, &stockTickersResponse)
			},
			check: func(t *testing.T) {
				if stockTickersResponse.Data[0].Symbol != "FPT" || stockTickersResponse.Data[0].TotalVol != 833900 {
					t.Fatalf("stock tickers response = %#v", stockTickersResponse)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.decode([]byte(test.payload)); err != nil {
				t.Fatalf("decode: %v", err)
			}
			test.check(t)
		})
	}
}

var (
	accountResponse         dto.GetSubAccountInformationResponse
	ordersResponse          dto.GetStockOrdersResponse
	purchasingPowerResponse dto.GetStockPurchasingPowerResponse
	derivativeCashResponse  dto.GetDerivativeCashResponse
	stockTickersResponse    dto.GetStockTickersResponse
)
