package ssi

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"math/big"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

func TestSignerSignsPayloadFromBase64XMLPrivateKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	privateKey := testBase64XMLPrivateKey(key)

	signer := Signer{PrivateKey: privateKey}
	request, err := signer.Sign(context.Background(), transport.HTTPRequest{
		Method:  "POST",
		URL:     "https://ssi.example/api/v2/Trading/NewOrder",
		Headers: map[string]string{},
		JSON: map[string]any{
			"account":   "0901351",
			"requestID": "12345678",
		},
	})
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	signatureHex := request.Headers["X-Signature"]
	if signatureHex == "" {
		t.Fatalf("missing signature")
	}
	payload, ok := request.JSON.(signedJSON)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	digest := sha256.Sum256(payload.Bytes)
	if err := rsa.VerifyPSS(&key.PublicKey, crypto.SHA256, digest[:], payload.Signature, nil); err != nil {
		t.Fatalf("verify signature: %v", err)
	}
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

func unsignedBytes(value *big.Int) []byte {
	bytes := value.Bytes()
	if len(bytes) == 0 {
		return []byte{0}
	}
	return bytes
}
