package ssi

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	nativedto "github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestSSINativeMarketDataServiceIsExposed(t *testing.T) {
	broker := NewBroker(Config{})
	if broker.Native() == nil {
		t.Fatal("expected native service")
	}
	if broker.Native().MarketData() == nil {
		t.Fatal("expected native market data service")
	}
}

func TestNativeMarketDataMethodsRequireOwnCapabilities(t *testing.T) {
	tests := []struct {
		name       string
		capability core.Capability
		call       func(context.Context, *Broker) error
	}{
		{"securities", CapabilityNativeMarketDataSecurities, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetSecurities(ctx, nativedto.GetSecuritiesRequest{})
			return err
		}},
		{"securities details", CapabilityNativeMarketDataSecuritiesDetails, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetSecuritiesDetails(ctx, nativedto.GetSecuritiesDetailsRequest{})
			return err
		}},
		{"index components", CapabilityNativeMarketDataIndexComponents, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetIndexComponents(ctx, nativedto.GetIndexComponentsRequest{})
			return err
		}},
		{"index list", CapabilityNativeMarketDataIndexList, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetIndexList(ctx, nativedto.GetIndexListRequest{})
			return err
		}},
		{"daily ohlc", CapabilityNativeMarketDataDailyOhlc, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetDailyOhlc(ctx, nativedto.GetDailyOhlcRequest{})
			return err
		}},
		{"intraday ohlc", CapabilityNativeMarketDataIntradayOhlc, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetIntradayOhlc(ctx, nativedto.GetIntradayOhlcRequest{})
			return err
		}},
		{"daily index", CapabilityNativeMarketDataDailyIndex, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetDailyIndex(ctx, nativedto.GetDailyIndexRequest{})
			return err
		}},
		{"daily stock price", CapabilityNativeMarketDataDailyStockPrice, func(ctx context.Context, b *Broker) error {
			_, err := b.Native().MarketData().GetDailyStockPrice(ctx, nativedto.GetDailyStockPriceRequest{})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{}
			broker := NewBroker(Config{HTTPTransport: httpTransport})
			broker.BrokerCapabilities = nil

			err := tt.call(context.Background(), broker)
			if err == nil {
				t.Fatal("expected unsupported capability error")
			}
			var brokerErr *sdkerrors.BrokerError
			if !errors.As(err, &brokerErr) || brokerErr.Category != sdkerrors.CategoryCapabilityUnsupported || brokerErr.Code != string(tt.capability) {
				t.Fatalf("error = %#v", err)
			}
			if len(httpTransport.requests) != 0 {
				t.Fatalf("requests = %d, want 0", len(httpTransport.requests))
			}
		})
	}
}

func TestNativeMarketDataMapsBrokerErrorFixture(t *testing.T) {
	raw, err := os.ReadFile("native/marketdata/testdata/broker_error.json")
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
	broker := NewBroker(Config{DataBaseURL: "https://data.ssi.example", HTTPTransport: httpTransport})

	_, err = broker.Native().MarketData().GetSecurities(context.Background(), nativedto.GetSecuritiesRequest{})
	var brokerErr *sdkerrors.BrokerError
	if !errors.As(err, &brokerErr) {
		t.Fatalf("error = %#v", err)
	}
	if brokerErr.Category != sdkerrors.CategoryBrokerRejected || brokerErr.Code != "400" || brokerErr.Operation != string(CapabilityNativeMarketDataSecurities) {
		t.Fatalf("broker error = %+v", brokerErr)
	}
}
