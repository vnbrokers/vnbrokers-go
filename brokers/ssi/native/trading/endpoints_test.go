package trading_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/xml"
	"math/big"
	"net/url"
	"testing"

	ssi "github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestNativeTradingGETEndpointsBuildRequests(t *testing.T) {
	tests := []struct {
		name string
		call func(context.Context, *ssi.Broker) error
		url  string
	}{
		{
			name: "cash in advance amount",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().CashInAdvanceAmount(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/cash/cashInAdvanceAmount?account=0901351",
		},
		{
			name: "unsettle sold transaction",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().UnsettleSoldTransaction(ctx, "0901351", "10/03/2023")
				return err
			},
			url: "https://ssi.example/api/v2/cash/unsettleSoldTransaction?account=0901351&settleDate=10%2F03%2F2023",
		},
		{
			name: "cash transfer histories",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().TransferHistories(ctx, "0901351", "10/01/2023", "10/02/2023")
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferHistories?account=0901351&fromDate=10%2F01%2F2023&toDate=10%2F02%2F2023",
		},
		{
			name: "stock transferable",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().Transferable(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/stock/transferable?account=0901351",
		},
		{
			name: "rights dividends",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().Dividend(ctx, "0901351")
				return err
			},
			url: "https://ssi.example/api/v2/ors/dividend?account=0901351",
		},
		{
			name: "conditional order list",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().FcoList(ctx, url.Values{
					"account":   {"0901351"},
					"pageIndex": {"1"},
					"pageSize":  {"50"},
				})
				return err
			},
			url: "https://ssi.example/api/v2/fco/list?account=0901351&pageIndex=1&pageSize=50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := ssi.NewBroker(ssi.Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "GET" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] != "" {
				t.Fatalf("GET should not be signed")
			}
		})
	}
}

func TestNativeTradingPOSTEndpointsBuildSignedRequests(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	tests := []struct {
		name string
		call func(context.Context, *ssi.Broker) error
		url  string
	}{
		{
			name: "cash transfer internal",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().TransferInternal(ctx, "0901351", "0901356", "50000", "test", "123456")
				return err
			},
			url: "https://ssi.example/api/v2/cash/transferInternal",
		},
		{
			name: "stock transfer",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().StockTransfer(ctx, map[string]any{
					"account":            "0901351",
					"beneficiaryAccount": "0901356",
					"exchangeID":         "HOSE",
					"instrumentID":       "SSI",
					"quantity":           100,
					"code":               "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/stock/transfer",
		},
		{
			name: "rights create",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().CreateRight(ctx, map[string]any{
					"account":       "0901351",
					"instrumentID":  "SSI",
					"entitlementID": "913312",
					"quantity":      100,
					"amount":        1000,
					"code":          "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/ors/create",
		},
		{
			name: "conditional new order",
			call: func(ctx context.Context, b *ssi.Broker) error {
				_, err := b.Native().Trading().FcoNewOrder(ctx, map[string]any{
					"instrumentID": "SSI",
					"side":         "B",
					"type":         "stop",
					"price":        "21000",
					"quantity":     100,
					"account":      "0901351",
					"stopPrice":    21100,
					"operator":     "greater_or_equal",
					"code":         "123456",
				})
				return err
			},
			url: "https://ssi.example/api/v2/fco/neworder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpTransport := &fakeHTTPTransport{
				responses: []transport.HTTPResponse{{
					StatusCode: 200,
					Body: map[string]any{
						"message": "Success",
						"status":  200,
						"data":    nil,
					},
				}},
			}
			broker := ssi.NewBroker(ssi.Config{
				BaseURL:       "https://ssi.example",
				TradingToken:  "trading-token",
				PrivateKey:    testBase64XMLPrivateKey(key),
				HTTPTransport: httpTransport,
			})
			if err := tt.call(context.Background(), broker); err != nil {
				t.Fatalf("call: %v", err)
			}
			request := httpTransport.requests[0]
			if request.Method != "POST" {
				t.Fatalf("method = %s", request.Method)
			}
			if request.URL != tt.url {
				t.Fatalf("url = %s", request.URL)
			}
			if request.Headers["Authorization"] != "Bearer trading-token" {
				t.Fatalf("authorization = %s", request.Headers["Authorization"])
			}
			if request.Headers["X-Signature"] == "" {
				t.Fatalf("missing signature")
			}
		})
	}
}

func unsignedBytes(value *big.Int) []byte {
	bytes := value.Bytes()
	if len(bytes) == 0 {
		return []byte{0}
	}
	return bytes
}

func testBase64XMLPrivateKey(key *rsa.PrivateKey) string {
	type rsaKeyValue struct {
		XMLName  xml.Name `xml:"RSAKeyValue"`
		Modulus  string   `xml:"Modulus"`
		Exponent string   `xml:"Exponent"`
		P        string   `xml:"P"`
		Q        string   `xml:"Q"`
		DP       string   `xml:"DP"`
		DQ       string   `xml:"DQ"`
		InverseQ string   `xml:"InverseQ"`
		D        string   `xml:"D"`
	}
	key.Precompute()
	value := rsaKeyValue{
		Modulus:  base64.StdEncoding.EncodeToString(unsignedBytes(key.N)),
		Exponent: base64.StdEncoding.EncodeToString(unsignedBytes(big.NewInt(int64(key.E)))),
		P:        base64.StdEncoding.EncodeToString(unsignedBytes(key.Primes[0])),
		Q:        base64.StdEncoding.EncodeToString(unsignedBytes(key.Primes[1])),
		DP:       base64.StdEncoding.EncodeToString(unsignedBytes(key.Precomputed.Dp)),
		DQ:       base64.StdEncoding.EncodeToString(unsignedBytes(key.Precomputed.Dq)),
		InverseQ: base64.StdEncoding.EncodeToString(unsignedBytes(key.Precomputed.Qinv)),
		D:        base64.StdEncoding.EncodeToString(unsignedBytes(key.D)),
	}
	bytes, err := xml.Marshal(value)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
