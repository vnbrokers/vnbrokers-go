package dto_test

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func TestNonRequestNumericFieldsUseDecimal(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob dto files: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}

		t.Run(file, func(t *testing.T) {
			fset := token.NewFileSet()
			parsed, err := parser.ParseFile(fset, file, nil, parser.SkipObjectResolution)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}

			ast.Inspect(parsed, func(node ast.Node) bool {
				typeSpec, ok := node.(*ast.TypeSpec)
				if !ok {
					return true
				}
				if strings.HasSuffix(typeSpec.Name.Name, "Request") || strings.HasSuffix(typeSpec.Name.Name, "Body") {
					return false
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					return false
				}

				for _, field := range structType.Fields.List {
					typeName, ok := primitiveNumericType(field.Type)
					if !ok {
						continue
					}
					for _, name := range field.Names {
						pos := fset.Position(field.Pos())
						t.Errorf("%s.%s at %s uses %s, want decimal.Decimal", typeSpec.Name.Name, name.Name, pos, typeName)
					}
				}

				return false
			})
		})
	}
}

func primitiveNumericType(expr ast.Expr) (string, bool) {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return "", false
	}
	switch ident.Name {
	case "float64", "int", "int64":
		return ident.Name, true
	default:
		return "", false
	}
}

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
				if !ordersResponse.PageSize.Equal(mustDecimal("10")) || !ordersResponse.Data[0].OrderQtty.Equal(mustDecimal("100.5")) {
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
				if !purchasingPowerResponse.Price.Equal(mustDecimal("120000")) || !purchasingPowerResponse.MaxBuyQuantity.Equal(mustDecimal("458600")) {
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
				if !derivativeCashResponse.Data.Cash.Equal(mustDecimal("600089123")) || !derivativeCashResponse.Data.VM.Equal(mustDecimal("-4200000")) {
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
				if stockTickersResponse.Data[0].Symbol != "FPT" || !stockTickersResponse.Data[0].TotalVol.Equal(mustDecimal("833900")) {
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

func mustDecimal(value string) decimal.Decimal {
	parsed, err := decimal.NewFromString(value)
	if err != nil {
		panic(err)
	}
	return parsed
}
