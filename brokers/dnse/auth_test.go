package dnse

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestSignerAddsDNSEHeaders(t *testing.T) {
	signer := HMACSigner{
		APIKey:    "key",
		APISecret: "secret",
		Now: func() time.Time {
			return time.Date(2026, 5, 25, 1, 2, 3, 0, time.UTC)
		},
		Nonce: func() string { return "nonce" },
	}

	request, err := signer.Sign(context.Background(), transport.HTTPRequest{
		Method:  "GET",
		URL:     "https://api.dnse.example/accounts",
		Headers: map[string]string{},
	})
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	if request.Headers["X-API-Key"] != "key" {
		t.Fatalf("missing api key header")
	}
	if request.Headers["X-Aux-Date"] != "Mon, 25 May 2026 01:02:03 +0000" {
		t.Fatalf("bad date header: %s", request.Headers["X-Aux-Date"])
	}
	if !strings.Contains(request.Headers["X-Signature"], `nonce="nonce"`) {
		t.Fatalf("signature header missing nonce: %s", request.Headers["X-Signature"])
	}
}

func TestSignerOmitsNonceFromSigningStringWhenNonceIsEmpty(t *testing.T) {
	signer := HMACSigner{
		APIKey:    "key",
		APISecret: "secret",
		Now: func() time.Time {
			return time.Date(2026, 5, 25, 1, 2, 3, 0, time.UTC)
		},
		Nonce: func() string { return "" },
	}

	request, err := signer.Sign(context.Background(), transport.HTTPRequest{
		Method:  "GET",
		URL:     "https://api.dnse.example/accounts",
		Headers: map[string]string{},
	})
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	if !strings.Contains(request.Headers["X-Signature"], `nonce=""`) {
		t.Fatalf("signature header should include empty nonce field: %s", request.Headers["X-Signature"])
	}
}

func TestNewRESTNonceReturnsUUIDLikeHex(t *testing.T) {
	nonce, err := newRESTNonce()
	if err != nil {
		t.Fatalf("new nonce: %v", err)
	}
	if len(nonce) != 32 {
		t.Fatalf("nonce length = %d", len(nonce))
	}
	if nonce[12] != '4' {
		t.Fatalf("nonce version nibble = %q", nonce[12])
	}
	if !strings.Contains("89ab", string(nonce[16])) {
		t.Fatalf("nonce variant nibble = %q", nonce[16])
	}
}
