package dnseintegration

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func TestDNSERestServicesAccept200Responses(t *testing.T) {
	statuses := []int{}
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Skipf("local HTTP server is not permitted in this environment: %v", err)
	}
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statuses = append(statuses, http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := responseFor(r.Method, r.URL.Path)
		if response == nil {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(map[string]any{"message": "not found"})
			return
		}
		_ = json.NewEncoder(w).Encode(response)
	}))
	server.Listener = listener
	server.Start()
	defer server.Close()

	price := decimal.NewFromInt(23000)
	orderPrice, _ := price.Float64()
	broker := dnse.NewBroker(dnse.Config{
		BaseURL:      server.URL,
		APIKey:       "key",
		APISecret:    "secret",
		TradingToken: "token",
		MarketType:   "STOCK",
	})

	if _, err := broker.Auth().SendEmailOTP(context.Background()); err != nil {
		t.Fatalf("send otp: %v", err)
	}
	if _, err := broker.Auth().GetTradingToken(context.Background(), nativedto.GetTradingTokenRequest{OTPType: "EMAIL", Passcode: "123456"}); err != nil {
		t.Fatalf("trading token: %v", err)
	}
	if _, err := broker.Native().Trading().GetAccounts(context.Background(), nativedto.GetAccountsRequest{}); err != nil {
		t.Fatalf("accounts: %v", err)
	}
	if _, err := broker.Native().Trading().GetAccountBalances(context.Background(), nativedto.GetAccountBalancesRequest{AccountNo: "0001179019"}); err != nil {
		t.Fatalf("balance: %v", err)
	}
	if _, err := broker.Native().Trading().GetOrders(context.Background(), nativedto.GetOrdersRequest{AccountNo: "0001179019", MarketType: "STOCK", OrderCategory: "NORMAL"}); err != nil {
		t.Fatalf("orders: %v", err)
	}
	if _, err := broker.Native().Trading().GetOrderHistory(context.Background(), nativedto.GetOrderHistoryRequest{AccountNo: "0001179019", MarketType: "STOCK", From: "2026-05-01", To: "2026-05-23"}); err != nil {
		t.Fatalf("order history: %v", err)
	}
	if _, err := broker.Native().Trading().GetExecutions(context.Background(), nativedto.GetExecutionsRequest{AccountNo: "0001179019", OrderID: "42", MarketType: "STOCK"}); err != nil {
		t.Fatalf("executions: %v", err)
	}
	if _, err := broker.Native().Trading().GetPPSE(context.Background(), nativedto.GetPPSERequest{AccountNo: "0001179019", Symbol: "ACB", MarketType: "STOCK", Price: &price}); err != nil {
		t.Fatalf("ppse: %v", err)
	}
	if _, err := broker.Native().Trading().GetLoanPackages(context.Background(), nativedto.GetLoanPackagesRequest{AccountNo: "0001179019", Symbol: "ACB", MarketType: "STOCK"}); err != nil {
		t.Fatalf("loan packages: %v", err)
	}
	if _, err := broker.Native().Trading().GetPositions(context.Background(), nativedto.GetPositionsRequest{AccountNo: "0001179019", MarketType: "STOCK"}); err != nil {
		t.Fatalf("positions: %v", err)
	}
	if _, err := broker.Native().Trading().GetPosition(context.Background(), nativedto.GetPositionRequest{PositionID: "177796763592657", MarketType: "STOCK"}); err != nil {
		t.Fatalf("position: %v", err)
	}
	if _, err := broker.Native().Trading().ClosePosition(context.Background(), nativedto.ClosePositionRequest{PositionID: "177796763592657", MarketType: "STOCK"}); err != nil {
		t.Fatalf("close position: %v", err)
	}
	if _, err := broker.Native().Trading().PlaceOrder(context.Background(), nativedto.PlaceOrderRequest{
		AccountNo: "0001179019", Symbol: "ACB", Side: "NB", OrderType: "LO",
		Quantity: 1, Price: &orderPrice, MarketType: "STOCK", OrderCategory: "NORMAL",
	}); err != nil {
		t.Fatalf("place order: %v", err)
	}
	if _, err := broker.Native().Trading().CancelOrder(context.Background(), nativedto.CancelOrderRequest{AccountNo: "0001179019", OrderID: "42", MarketType: "STOCK", OrderCategory: "NORMAL"}); err != nil {
		t.Fatalf("cancel order: %v", err)
	}
	if _, err := broker.Native().Trading().GetOrder(context.Background(), nativedto.GetOrderRequest{AccountNo: "0001179019", OrderID: "42", MarketType: "STOCK", OrderCategory: "NORMAL"}); err != nil {
		t.Fatalf("get order: %v", err)
	}
	if _, err := broker.Native().Trading().ReplaceOrder(context.Background(), nativedto.ReplaceOrderRequest{AccountNo: "0001179019", OrderID: "42", Price: &orderPrice, Quantity: 1, MarketType: "STOCK", OrderCategory: "NORMAL"}); err != nil {
		t.Fatalf("update order: %v", err)
	}
	if _, err := broker.Native().MarketData().GetInstruments(context.Background(), nativedto.GetInstrumentsRequest{Symbol: "ACB", Limit: 10}); err != nil {
		t.Fatalf("symbols: %v", err)
	}
	if _, err := broker.Native().MarketData().GetSecurityDefinition(context.Background(), nativedto.GetSecurityDefinitionRequest{Symbol: "HPG", BoardID: "G1"}); err != nil {
		t.Fatalf("secdef: %v", err)
	}
	if _, err := broker.Native().MarketData().GetWorkingDates(context.Background(), nativedto.GetWorkingDatesRequest{}); err != nil {
		t.Fatalf("working dates: %v", err)
	}
	if _, err := broker.Native().MarketData().GetLatestTrades(context.Background(), nativedto.GetLatestTradesRequest{Symbol: "ACB", BoardID: "G1"}); err != nil {
		t.Fatalf("quote: %v", err)
	}
	if _, err := broker.Native().MarketData().GetTradeHistory(context.Background(), nativedto.GetTradeHistoryRequest{Symbol: "ACB", From: 1773182637, To: 1773183637, BoardID: "G1", Limit: 100}); err != nil {
		t.Fatalf("price trades: %v", err)
	}
	if _, err := broker.Native().MarketData().GetClosePrice(context.Background(), nativedto.GetClosePriceRequest{Symbol: "HPG", BoardID: "G1"}); err != nil {
		t.Fatalf("close price: %v", err)
	}
	if _, err := broker.Native().MarketData().GetOHLC(context.Background(), nativedto.GetOHLCRequest{Symbol: "ACB", Resolution: "15", From: 1773657310, To: 1773830110, Type: "STOCK"}); err != nil {
		t.Fatalf("candles: %v", err)
	}
	if _, err := broker.Native().Brokerage().GetCareByAccounts(context.Background(), nativedto.GetCareByAccountsRequest{}); err != nil {
		t.Fatalf("brokerage care by: %v", err)
	}

	if len(statuses) != 24 {
		t.Fatalf("request count = %d", len(statuses))
	}
}

func responseFor(method string, path string) any {
	switch method + " " + path {
	case "POST /registration/send-email-otp":
		return map[string]any{}
	case "POST /registration/trading-token":
		return map[string]any{"tradingToken": "token"}
	case "GET /accounts":
		return map[string]any{"name": "Nguyen Van A", "accounts": []any{map[string]any{"id": "0001179019"}}}
	case "GET /accounts/0001179019/balances":
		return map[string]any{"stock": map[string]any{"availableCash": 1000, "totalCash": 2000}}
	case "GET /accounts/0001179019/orders":
		return map[string]any{"orders": []any{orderPayload()}}
	case "GET /accounts/0001179019/orders/history":
		return map[string]any{"data": []any{orderPayload()}}
	case "GET /accounts/0001179019/executions/42":
		return map[string]any{"reports": []any{map[string]any{"execId": "E1"}}}
	case "GET /accounts/0001179019/ppse":
		return map[string]any{"qmaxBuy": 1}
	case "GET /accounts/0001179019/loan-packages":
		return map[string]any{"loanPackages": []any{map[string]any{"id": 1775}}}
	case "GET /accounts/0001179019/positions":
		return map[string]any{"positions": []any{positionPayload()}}
	case "GET /accounts/positions/177796763592657":
		return positionPayload()
	case "POST /accounts/positions/177796763592657/close":
		return map[string]any{"id": 1}
	case "POST /accounts/orders":
		return orderPayload()
	case "DELETE /accounts/0001179019/orders/42":
		return map[string]any{}
	case "GET /accounts/0001179019/orders/42":
		return orderPayload()
	case "PUT /accounts/0001179019/orders/42":
		return orderPayload()
	case "GET /instruments":
		return map[string]any{"data": []any{map[string]any{"symbol": "ACB", "marketId": "STO"}}}
	case "GET /price/HPG/secdef":
		return []any{map[string]any{"symbol": "HPG", "boardId": "G1"}}
	case "GET /market/working-dates":
		return map[string]any{"workingDates": []any{"2026-05-25"}}
	case "GET /price/ACB/trades/latest":
		return map[string]any{"trades": []any{map[string]any{"matchPrice": 23000, "time": "1773183637"}}}
	case "GET /price/ACB/trades":
		return map[string]any{"trades": []any{map[string]any{"price": 23000}}, "nextPageToken": "next"}
	case "GET /price/HPG/close":
		return map[string]any{"prices": []any{map[string]any{"symbol": "HPG", "closePrice": 26.8}}}
	case "GET /price/ohlc":
		return map[string]any{
			"t": []any{1773657310},
			"o": []any{23000},
			"h": []any{23200},
			"l": []any{22900},
			"c": []any{23100},
			"v": []any{100000},
		}
	case "GET /brokers/accounts/care-by":
		return map[string]any{"careBy": []any{map[string]any{"accountNo": "0001179019"}}}
	default:
		return nil
	}
}

func orderPayload() map[string]any {
	return map[string]any{
		"id":          42,
		"accountNo":   "0001179019",
		"symbol":      "ACB",
		"side":        "NB",
		"orderType":   "LO",
		"orderStatus": "FILLED",
		"quantity":    1,
		"price":       23000,
	}
}

func positionPayload() map[string]any {
	return map[string]any{
		"accountNo":    "0001179019",
		"symbol":       "ACB",
		"openQuantity": 1,
		"costPrice":    23000,
		"marketPrice":  23100,
	}
}
