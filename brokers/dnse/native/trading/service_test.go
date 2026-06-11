package trading

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestTradingServiceExposesEveryDocumentedRESTOperation(t *testing.T) {
	var service Service
	ctx := context.Background()

	compile := []func() error{
		func() error { _, err := service.GetAccounts(ctx, dto.GetAccountsRequest{}); return err },
		func() error { _, err := service.GetAccountBalances(ctx, dto.GetAccountBalancesRequest{}); return err },
		func() error {
			_, err := service.GetCorporateActionHistory(ctx, dto.GetCorporateActionHistoryRequest{})
			return err
		},
		func() error { _, err := service.GetExecutions(ctx, dto.GetExecutionsRequest{}); return err },
		func() error { _, err := service.GetLoanPackages(ctx, dto.GetLoanPackagesRequest{}); return err },
		func() error { _, err := service.GetOrderHistory(ctx, dto.GetOrderHistoryRequest{}); return err },
		func() error { _, err := service.GetOrders(ctx, dto.GetOrdersRequest{}); return err },
		func() error {
			_, err := service.GetPositionPnLConfigs(ctx, dto.GetPositionPnLConfigsRequest{})
			return err
		},
		func() error { _, err := service.GetPosition(ctx, dto.GetPositionRequest{}); return err },
		func() error { _, err := service.GetPositions(ctx, dto.GetPositionsRequest{}); return err },
		func() error { _, err := service.GetPPSE(ctx, dto.GetPPSERequest{}); return err },
		func() error { _, err := service.CancelOrder(ctx, dto.CancelOrderRequest{}); return err },
		func() error { _, err := service.ClosePosition(ctx, dto.ClosePositionRequest{}); return err },
		func() error { _, err := service.GetOrder(ctx, dto.GetOrderRequest{}); return err },
		func() error { _, err := service.PlaceOrder(ctx, dto.PlaceOrderRequest{}); return err },
		func() error {
			_, err := service.SetPositionPnLConfigs(ctx, dto.SetPositionPnLConfigsRequest{})
			return err
		},
		func() error { _, err := service.ReplaceOrder(ctx, dto.ReplaceOrderRequest{}); return err },
	}
	_ = compile
}

func TestPlaceOrderEncodesPriceAsJSONNumber(t *testing.T) {
	price := 23000.5
	var request transport.HTTPRequest
	service := NewService(Dependencies{
		BaseURL:           "https://openapi.dnse.com.vn",
		RequireCapability: func(core.Capability) error { return nil },
		Send: func(_ context.Context, _ string, got transport.HTTPRequest) (transport.HTTPResponse, error) {
			request = got
			return transport.HTTPResponse{Raw: []byte(`{}`)}, nil
		},
	})
	if _, err := service.PlaceOrder(context.Background(), dto.PlaceOrderRequest{Price: &price}); err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(request.JSON)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), `"price":"23000.5"`) {
		t.Fatalf("price encoded as string: %s", payload)
	}
	if !strings.Contains(string(payload), `"price":23000.5`) {
		t.Fatalf("price missing as number: %s", payload)
	}
}
