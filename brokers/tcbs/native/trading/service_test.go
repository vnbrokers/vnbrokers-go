package trading_test

import (
	"context"
	"strings"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/trading"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestTradingHTTPContracts(t *testing.T) {
	tests := []struct {
		name   string
		method string
		url    string
		body   bool
		invoke func(context.Context, trading.Service) error
	}{
		{"sub account", "GET", "/eros/v2/get-profile/by-username/105C%2F1?fields=basicInfo", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetSubAccountInformation(ctx, dto.GetSubAccountInformationRequest{CustodyCode: "105C/1", Fields: "basicInfo"})
			return err
		}},
		{"transfer", "POST", "/physis/v1/stock/transfer", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.TransferBetweenSubaccounts(ctx, dto.TransferBetweenSubaccountsRequest{})
			return err
		}},
		{"withdraw derivative margin", "POST", "/khronos/v1/cash/withdraw/update", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.WithdrawDerivativeMargin(ctx, dto.WithdrawDerivativeMarginRequest{})
			return err
		}},
		{"deposit derivative margin", "POST", "/khronos/v1/cash/deposit/update", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.DepositDerivativeMargin(ctx, dto.DepositDerivativeMarginRequest{})
			return err
		}},
		{"place stock order", "POST", "/akhlys/v1/accounts/0001%2FA/orders", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.PlaceStockOrder(ctx, dto.PlaceStockOrderRequest{AccountNo: "0001/A"})
			return err
		}},
		{"update stock order", "PUT", "/akhlys/v1/accounts/0001/orders/42%2F1", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.UpdateStockOrder(ctx, dto.UpdateStockOrderRequest{AccountNo: "0001", OrderID: "42/1"})
			return err
		}},
		{"cancel stock order", "PUT", "/akhlys/v1/accounts/0001/cancel-orders", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.CancelStockOrder(ctx, dto.CancelStockOrderRequest{AccountNo: "0001"})
			return err
		}},
		{"stock orders", "GET", "/aion/v1/accounts/0001/orders", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockOrders(ctx, dto.GetStockOrdersRequest{AccountNo: "0001"})
			return err
		}},
		{"stock order", "GET", "/aion/v1/accounts/0001/orders/42", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockOrder(ctx, dto.GetStockOrderRequest{AccountNo: "0001", OrderID: "42"})
			return err
		}},
		{"stock matching", "GET", "/aion/v1/accounts/0001/matching-details", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockMatchingDetails(ctx, dto.GetStockMatchingDetailsRequest{AccountNo: "0001"})
			return err
		}},
		{"stock purchasing power", "GET", "/aion/v1/accounts/0001/ppse", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockPurchasingPower(ctx, dto.GetStockPurchasingPowerRequest{AccountNo: "0001"})
			return err
		}},
		{"stock purchasing power symbol", "GET", "/aion/v1/accounts/0001/ppse/FPT", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockPurchasingPowerBySymbol(ctx, dto.GetStockPurchasingPowerBySymbolRequest{AccountNo: "0001", Symbol: "FPT"})
			return err
		}},
		{"stock purchasing power price", "GET", "/aion/v1/accounts/0001/ppse/FPT/52000", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockPurchasingPowerBySymbolPrice(ctx, dto.GetStockPurchasingPowerBySymbolPriceRequest{AccountNo: "0001", Symbol: "FPT", Price: "52000"})
			return err
		}},
		{"margin quota", "GET", "/aion/v1/customers/105C/accounts", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetMarginQuota(ctx, dto.GetMarginQuotaRequest{CustodyID: "105C"})
			return err
		}},
		{"margin account", "GET", "/hydros/v1/account/0001/risk", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetMarginAccountInformation(ctx, dto.GetMarginAccountInformationRequest{AccountNo: "0001"})
			return err
		}},
		{"loan packages", "GET", "/campaign-management/v1/margin/subscription/0001/addons/detail", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetSupplementaryLoanPackages(ctx, dto.GetSupplementaryLoanPackagesRequest{AccountNo: "0001"})
			return err
		}},
		{"loans", "GET", "/khaos/v1/loan/0001", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetLoans(ctx, dto.GetLoansRequest{AccountNo: "0001"})
			return err
		}},
		{"stock assets", "GET", "/aion/v1/accounts/0001/se", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetStockAssets(ctx, dto.GetStockAssetsRequest{AccountNo: "0001"})
			return err
		}},
		{"cash investments", "GET", "/aion/v1/accounts/0001/cashInvestments", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetCashInvestments(ctx, dto.GetCashInvestmentsRequest{AccountNo: "0001"})
			return err
		}},
		{"cash statements", "GET", "/erebos/v2/digital/trans-hist-cashStatements?accountNo=0001&fromDate=2025-01-01&pageIndex=1&pageSize=25&toDate=2025-01-15&transactionCode=1153", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetCashStatements(ctx, dto.GetCashStatementsRequest{AccountNo: "0001", FromDate: "2025-01-01", ToDate: "2025-01-15", PageSize: "25", PageIndex: "1", TransactionCode: "1153"})
			return err
		}},
		{"cash statements without transaction code", "GET", "/erebos/v2/digital/trans-hist-cashStatements?accountNo=0001&fromDate=2025-01-01&pageIndex=1&pageSize=25&toDate=2025-01-15", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetCashStatements(ctx, dto.GetCashStatementsRequest{AccountNo: "0001", FromDate: "2025-01-01", ToDate: "2025-01-15", PageSize: "25", PageIndex: "1"})
			return err
		}},
		{"margin information", "GET", "/erebos/v2/digital/margin-info?acctno=0001&custodycd=105C&fromdate=2025-01-01&page=1&size=25&toDate=2025-01-15", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetMarginInformation(ctx, dto.GetMarginInformationRequest{AccountNo: "0001", FromDate: "2025-01-01", ToDate: "2025-01-15", Page: "1", Size: "25", CustodyCode: "105C"})
			return err
		}},
		{"derivative cash", "GET", "/khronos/v1/account/status?accountId=105C&getType=1&subAccountId=105CA", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetDerivativeCash(ctx, dto.GetDerivativeCashRequest{AccountID: "105C", SubAccountID: "105CA", GetType: "1"})
			return err
		}},
		{"closed derivative positions", "GET", "/khronos/v1/account/portfolio/position/close?accountId=105C&pageNo=1&pageSize=20&subAccountId=105CA&symbol=VN30F", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetClosedDerivativePositions(ctx, dto.GetClosedDerivativePositionsRequest{AccountID: "105C", SubAccountID: "105CA", Symbol: "VN30F", PageNo: 1, PageSize: 20})
			return err
		}},
		{"open derivative positions", "GET", "/khronos/v1/account/portfolio/status?accountId=105C&subAccountId=105CA", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetOpenDerivativePositions(ctx, dto.GetOpenDerivativePositionsRequest{AccountID: "105C", SubAccountID: "105CA"})
			return err
		}},
		{"derivative orders", "GET", "/khronos/v1/order/in-day?accountId=105C&orderType=N&pageNo=1&pageSize=20&refId=ref&status=0&symbol=VN30F", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetDerivativeOrders(ctx, dto.GetDerivativeOrdersRequest{PageNo: 1, PageSize: 20, AccountID: "105C", Symbol: "VN30F", RefID: "ref", OrderType: "N", Status: "0"})
			return err
		}},
		{"derivative conditional orders", "GET", "/khronos/v1/order/condition/detail?PageSize=25&Symbol=VN30F&accountId=105C&orderStatus=0&orderType=ALL&pageNo=1&subAccountID=105CA", false, func(ctx context.Context, s trading.Service) error {
			_, err := s.GetDerivativeConditionalOrders(ctx, dto.GetDerivativeConditionalOrdersRequest{PageNo: "1", PageSize: "25", AccountID: "105C", SubAccountID: "105CA", OrderStatus: "0", OrderType: "ALL", Symbol: "VN30F"})
			return err
		}},
		{"place derivative order", "POST", "/khronos/v1/order/place", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.PlaceDerivativeOrder(ctx, dto.PlaceDerivativeOrderRequest{})
			return err
		}},
		{"place derivative conditional order", "POST", "/khronos/v1/order/condition/place", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.PlaceDerivativeConditionalOrder(ctx, dto.PlaceDerivativeConditionalOrderRequest{})
			return err
		}},
		{"update derivative order", "POST", "/khronos/v1/order/change", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.UpdateDerivativeOrder(ctx, dto.UpdateDerivativeOrderRequest{})
			return err
		}},
		{"update derivative conditional order", "POST", "/khronos/v2/order/condition/change", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.UpdateDerivativeConditionalOrder(ctx, dto.UpdateDerivativeConditionalOrderRequest{})
			return err
		}},
		{"cancel derivative order", "POST", "/khronos/v1/order/cancel", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.CancelDerivativeOrder(ctx, dto.CancelDerivativeOrderRequest{})
			return err
		}},
		{"cancel derivative conditional order", "POST", "/khronos/v1/order/condition/cancel", true, func(ctx context.Context, s trading.Service) error {
			_, err := s.CancelDerivativeConditionalOrder(ctx, dto.CancelDerivativeConditionalOrderRequest{})
			return err
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var gotCapability core.Capability
			var gotRequest transport.HTTPRequest
			service := trading.NewService(trading.Dependencies{
				BaseURL:           "https://api.example",
				Headers:           func(bool, bool) map[string]string { return map[string]string{"Authorization": "Bearer token"} },
				RequireCapability: func(capability core.Capability) error { gotCapability = capability; return nil },
				Send: func(_ context.Context, operation string, request transport.HTTPRequest) (transport.HTTPResponse, error) {
					gotRequest = request
					raw := []byte(`{}`)
					if strings.Contains(operation, "margin_quota") || strings.Contains(operation, "margin_account_information") {
						raw = []byte(`[]`)
					}
					return transport.HTTPResponse{StatusCode: 200, Raw: raw}, nil
				},
			})

			if err := test.invoke(context.Background(), service); err != nil {
				t.Fatalf("invoke: %v", err)
			}
			if gotCapability == "" {
				t.Fatal("capability was not checked")
			}
			if gotRequest.Method != test.method || gotRequest.URL != "https://api.example"+test.url {
				t.Fatalf("request = %s %s", gotRequest.Method, gotRequest.URL)
			}
			if (gotRequest.JSON != nil) != test.body {
				t.Fatalf("body presence = %t, want %t", gotRequest.JSON != nil, test.body)
			}
			if gotRequest.Headers["Authorization"] != "Bearer token" {
				t.Fatalf("headers = %#v", gotRequest.Headers)
			}
		})
	}
}
