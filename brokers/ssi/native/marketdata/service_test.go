package marketdata_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	ssi "github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestNativeMarketDataServiceIsExposed(t *testing.T) {
	broker := ssi.NewBroker(ssi.Config{})
	if broker.Native() == nil {
		t.Fatal("expected native service")
	}
	if broker.Native().MarketData() == nil {
		t.Fatal("expected native market data service")
	}
	if broker.Native().MarketData().Realtime() == nil {
		t.Fatal("expected native market data realtime service")
	}
}

func TestNativeMarketDataMethodsRequireOwnCapabilities(t *testing.T) {
	tests := []struct {
		name       string
		capability string
		call       func(context.Context, *ssi.Broker) error
	}{
		{"securities", string(ssi.CapabilityNativeMarketDataSecurities), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetSecurities(ctx, nativedto.GetSecuritiesRequest{})
			return err
		}},
		{"securities details", string(ssi.CapabilityNativeMarketDataSecuritiesDetails), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetSecuritiesDetails(ctx, nativedto.GetSecuritiesDetailsRequest{})
			return err
		}},
		{"index components", string(ssi.CapabilityNativeMarketDataIndexComponents), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetIndexComponents(ctx, nativedto.GetIndexComponentsRequest{})
			return err
		}},
		{"index list", string(ssi.CapabilityNativeMarketDataIndexList), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetIndexList(ctx, nativedto.GetIndexListRequest{})
			return err
		}},
		{"daily ohlc", string(ssi.CapabilityNativeMarketDataDailyOhlc), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetDailyOhlc(ctx, nativedto.GetDailyOhlcRequest{})
			return err
		}},
		{"intraday ohlc", string(ssi.CapabilityNativeMarketDataIntradayOhlc), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetIntradayOhlc(ctx, nativedto.GetIntradayOhlcRequest{})
			return err
		}},
		{"daily index", string(ssi.CapabilityNativeMarketDataDailyIndex), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetDailyIndex(ctx, nativedto.GetDailyIndexRequest{})
			return err
		}},
		{"daily stock price", string(ssi.CapabilityNativeMarketDataDailyStockPrice), func(ctx context.Context, b *ssi.Broker) error {
			_, err := b.Native().MarketData().GetDailyStockPrice(ctx, nativedto.GetDailyStockPriceRequest{})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{}
			broker := ssi.NewBroker(ssi.Config{HTTPTransport: httpTransport})
			broker.BrokerCapabilities = nil

			err := tt.call(context.Background(), broker)
			if err == nil {
				t.Fatal("expected unsupported capability error")
			}
			var brokerErr *sdkerrors.BrokerError
			if !errors.As(err, &brokerErr) || brokerErr.Category != sdkerrors.CategoryCapabilityUnsupported || brokerErr.Code != tt.capability {
				t.Fatalf("error = %#v", err)
			}
			if len(httpTransport.requests) != 0 {
				t.Fatalf("requests = %d, want 0", len(httpTransport.requests))
			}
		})
	}
}

func TestNativeMarketDataMapsBrokerErrorFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/broker_error.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}
	httpTransport := &fakeHTTPTransport{responses: []transport.HTTPResponse{{
		StatusCode: 400,
		Body:       body,
		Raw:        raw,
	}}}
	broker := ssi.NewBroker(ssi.Config{DataBaseURL: "https://data.ssi.example", HTTPTransport: httpTransport})

	_, err = broker.Native().MarketData().GetSecurities(context.Background(), nativedto.GetSecuritiesRequest{})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error = %#v", err)
	}
	if brokerErr.Category != sdkerrors.CategoryBrokerRejected || brokerErr.Code != "400" || brokerErr.Operation != string(ssi.CapabilityNativeMarketDataSecurities) {
		t.Fatalf("broker error = %+v", brokerErr)
	}
}
