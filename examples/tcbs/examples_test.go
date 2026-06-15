package tcbs_examples_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"unicode"
)

func TestNativeExampleCoverage(t *testing.T) {
	root := repositoryRoot(t)
	examples := map[string]string{
		"GetToken": "auth/get_token",

		"GetSubAccountInformation":             "native/trading/get_sub_account_information",
		"TransferBetweenSubaccounts":           "native/trading/transfer_between_subaccounts",
		"WithdrawDerivativeMargin":             "native/trading/withdraw_derivative_margin",
		"DepositDerivativeMargin":              "native/trading/deposit_derivative_margin",
		"PlaceStockOrder":                      "native/trading/place_stock_order",
		"UpdateStockOrder":                     "native/trading/update_stock_order",
		"CancelStockOrder":                     "native/trading/cancel_stock_order",
		"GetStockOrders":                       "native/trading/get_stock_orders",
		"GetStockOrder":                        "native/trading/get_stock_order",
		"GetStockMatchingDetails":              "native/trading/get_stock_matching_details",
		"GetStockPurchasingPower":              "native/trading/get_stock_purchasing_power",
		"GetStockPurchasingPowerBySymbol":      "native/trading/get_stock_purchasing_power_by_symbol",
		"GetStockPurchasingPowerBySymbolPrice": "native/trading/get_stock_purchasing_power_by_symbol_price",
		"GetMarginQuota":                       "native/trading/get_margin_quota",
		"GetMarginAccountInformation":          "native/trading/get_margin_account_information",
		"GetSupplementaryLoanPackages":         "native/trading/get_supplementary_loan_packages",
		"GetLoans":                             "native/trading/get_loans",
		"GetStockAssets":                       "native/trading/get_stock_assets",
		"GetCashInvestments":                   "native/trading/get_cash_investments",
		"GetCashStatements":                    "native/trading/get_cash_statements",
		"GetMarginInformation":                 "native/trading/get_margin_information",
		"GetDerivativeCash":                    "native/trading/get_derivative_cash",
		"GetClosedDerivativePositions":         "native/trading/get_closed_derivative_positions",
		"GetOpenDerivativePositions":           "native/trading/get_open_derivative_positions",
		"GetDerivativeOrders":                  "native/trading/get_derivative_orders",
		"GetDerivativeConditionalOrders":       "native/trading/get_derivative_conditional_orders",
		"PlaceDerivativeOrder":                 "native/trading/place_derivative_order",
		"PlaceDerivativeConditionalOrder":      "native/trading/place_derivative_conditional_order",
		"UpdateDerivativeOrder":                "native/trading/update_derivative_order",
		"UpdateDerivativeConditionalOrder":     "native/trading/update_derivative_conditional_order",
		"CancelDerivativeOrder":                "native/trading/cancel_derivative_order",
		"CancelDerivativeConditionalOrder":     "native/trading/cancel_derivative_conditional_order",

		"SubscribeStockOrders":                     "native/trading/realtime/subscribe_stock_orders",
		"SubscribeDerivativeOrders":                "native/trading/realtime/subscribe_derivative_orders",
		"SubscribeDerivativeOpenPositions":         "native/trading/realtime/subscribe_derivative_open_positions",
		"SubscribeDerivativeConditionalOrders":     "native/trading/realtime/subscribe_derivative_conditional_orders",
		"GetDerivativeTickers":                     "native/marketdata/get_derivative_tickers",
		"GetStockTickers":                          "native/marketdata/get_stock_tickers",
		"GetStockForeignRooms":                     "native/marketdata/get_stock_foreign_rooms",
		"GetStockPutThroughs":                      "native/marketdata/get_stock_put_throughs",
		"GetStockTradeHistory":                     "native/marketdata/get_stock_trade_history",
		"GetStockSupplyDemand15Minutes":            "native/marketdata/get_stock_supply_demand_15_minutes",
		"GetStockSupplyDemandDaily":                "native/marketdata/get_stock_supply_demand_daily",
		"GetStockSupplyDemandMonthly":              "native/marketdata/get_stock_supply_demand_monthly",
		"SubscribeStockPrices":                     "native/marketdata/realtime/subscribe_stock_prices",
		"SubscribeStockTradeHistory":               "native/marketdata/realtime/subscribe_stock_trade_history",
		"SubscribeStockSupplyDemandOneMinute":      "native/marketdata/realtime/subscribe_stock_supply_demand_one_minute",
		"SubscribeStockSupplyDemandFifteenMinutes": "native/marketdata/realtime/subscribe_stock_supply_demand_fifteen_minutes",
		"SubscribeDerivativeBidPrices":             "native/marketdata/realtime/subscribe_derivative_bid_prices",
		"SubscribeDerivativeOfferPrices":           "native/marketdata/realtime/subscribe_derivative_offer_prices",
		"SubscribeDerivativeForeignTrading":        "native/marketdata/realtime/subscribe_derivative_foreign_trading",
		"SubscribeDerivativeBasePrices":            "native/marketdata/realtime/subscribe_derivative_base_prices",
		"SubscribeDerivativeMatchedPrices":         "native/marketdata/realtime/subscribe_derivative_matched_prices",
		"SubscribeDerivativeTickerMatches":         "native/marketdata/realtime/subscribe_derivative_ticker_matches",
		"SubscribeDerivativeIndexes":               "native/marketdata/realtime/subscribe_derivative_indexes",
	}

	wantMethods := publicOperations(t, root)
	gotMethods := make([]string, 0, len(examples))
	for method, relativePath := range examples {
		gotMethods = append(gotMethods, method)
		mainPath := filepath.Join(root, "examples", "tcbs", relativePath, "main.go")
		source, err := os.ReadFile(mainPath)
		if err != nil {
			t.Errorf("%s example: %v", method, err)
			continue
		}
		if !strings.Contains(string(source), "."+method+"(") {
			t.Errorf("%s example does not call %s", method, method)
		}
		if filepath.Base(relativePath) != snakeCase(method) {
			t.Errorf("%s example directory = %q, want %q", method, filepath.Base(relativePath), snakeCase(method))
		}
	}
	sort.Strings(gotMethods)
	if strings.Join(gotMethods, "\n") != strings.Join(wantMethods, "\n") {
		t.Errorf("example methods do not match public operations\ngot:\n%s\nwant:\n%s", strings.Join(gotMethods, "\n"), strings.Join(wantMethods, "\n"))
	}
}

func publicOperations(t *testing.T, root string) []string {
	t.Helper()
	operations := []string{"GetToken"}
	files := []struct {
		path          string
		interfaceName string
	}{
		{"brokers/tcbs/native/trading/service.go", "Service"},
		{"brokers/tcbs/native/trading/realtime.go", "RealtimeService"},
		{"brokers/tcbs/native/marketdata/service.go", "Service"},
		{"brokers/tcbs/native/marketdata/realtime.go", "RealtimeService"},
	}
	for _, source := range files {
		parsed, err := parser.ParseFile(token.NewFileSet(), filepath.Join(root, source.path), nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range parsed.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec, ok := specification.(*ast.TypeSpec)
				if !ok || typeSpec.Name.Name != source.interfaceName {
					continue
				}
				service, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok {
					t.Fatalf("%s is not an interface", source.interfaceName)
				}
				for _, method := range service.Methods.List {
					if len(method.Names) == 1 && method.Names[0].IsExported() && method.Names[0].Name != "Realtime" {
						operations = append(operations, method.Names[0].Name)
					}
				}
			}
		}
	}
	sort.Strings(operations)
	return operations
}

func repositoryRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func snakeCase(value string) string {
	var output []rune
	characters := []rune(value)
	for index, current := range characters {
		startsWord := unicode.IsUpper(current)
		startsNumber := unicode.IsDigit(current) && index > 0 && !unicode.IsDigit(characters[index-1])
		if index > 0 && (startsWord || startsNumber) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(current))
	}
	return string(output)
}
