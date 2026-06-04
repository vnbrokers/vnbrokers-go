package ssi

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/big"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Signer struct {
	PrivateKey string
}

type signedJSON struct {
	Value     any
	Bytes     []byte
	Signature []byte
}

func (s signedJSON) MarshalJSON() ([]byte, error) {
	return s.Bytes, nil
}

func (s Signer) Sign(_ context.Context, request transport.HTTPRequest) (transport.HTTPRequest, error) {
	if s.PrivateKey == "" {
		return request, fmt.Errorf("ssi private key is required for signed request")
	}
	payload, err := json.Marshal(request.JSON)
	if err != nil {
		return request, err
	}
	privateKey, err := parseRSAPrivateKey(s.PrivateKey)
	if err != nil {
		return request, err
	}
	digest := sha256.Sum256(payload)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest[:])
	if err != nil {
		return request, err
	}
	headers := cloneHeaders(request.Headers)
	headers["X-Signature"] = hex.EncodeToString(signature)
	request.Headers = headers
	request.JSON = signedJSON{
		Value:     request.JSON,
		Bytes:     payload,
		Signature: signature,
	}
	return request, nil
}

type rsaKeyValue struct {
	Modulus  string `xml:"Modulus"`
	Exponent string `xml:"Exponent"`
	P        string `xml:"P"`
	Q        string `xml:"Q"`
	D        string `xml:"D"`
}

func parseRSAPrivateKey(privateKeyBase64XML string) (*rsa.PrivateKey, error) {
	xmlBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64XML)
	if err != nil {
		return nil, fmt.Errorf("decode SSI private key XML: %w", err)
	}
	var value rsaKeyValue
	if err := xml.Unmarshal(xmlBytes, &value); err != nil {
		return nil, fmt.Errorf("parse SSI private key XML: %w", err)
	}
	n, err := base64BigInt(value.Modulus)
	if err != nil {
		return nil, fmt.Errorf("parse RSA modulus: %w", err)
	}
	e, err := base64BigInt(value.Exponent)
	if err != nil {
		return nil, fmt.Errorf("parse RSA exponent: %w", err)
	}
	d, err := base64BigInt(value.D)
	if err != nil {
		return nil, fmt.Errorf("parse RSA private exponent: %w", err)
	}
	p, err := base64BigInt(value.P)
	if err != nil {
		return nil, fmt.Errorf("parse RSA prime P: %w", err)
	}
	q, err := base64BigInt(value.Q)
	if err != nil {
		return nil, fmt.Errorf("parse RSA prime Q: %w", err)
	}
	key := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: n,
			E: int(e.Int64()),
		},
		D:      d,
		Primes: []*big.Int{p, q},
	}
	if err := key.Validate(); err != nil {
		return nil, fmt.Errorf("validate SSI RSA private key: %w", err)
	}
	key.Precompute()
	return key, nil
}

func base64BigInt(value string) (*big.Int, error) {
	bytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(bytes), nil
}
