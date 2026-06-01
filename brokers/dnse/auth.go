package dnse

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	"github.com/vnbrokers/vnbrokers-go/errors"
	sdktrading "github.com/vnbrokers/vnbrokers-go/trading"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Signer struct {
	APIKey    string
	APISecret string
	Now       func() time.Time
	Nonce     func() string
}

type HMACSigner = Signer

func (s Signer) Sign(_ context.Context, request transport.HTTPRequest) (transport.HTTPRequest, error) {
	now := time.Now
	if s.Now != nil {
		now = s.Now
	}
	nonce, err := newRESTNonce()
	if err != nil {
		return request, err
	}
	if s.Nonce != nil {
		nonce = s.Nonce()
	}
	parsed, err := url.Parse(request.URL)
	if err != nil {
		return request, err
	}
	date := now().UTC().Format("Mon, 02 Jan 2006 15:04:05 +0000")
	signingString := fmt.Sprintf(
		"(request-target): %s %s\nx-aux-date: %s",
		strings.ToLower(request.Method),
		parsed.EscapedPath(),
		date,
	)
	if nonce != "" {
		signingString += "\nnonce: " + nonce
	}
	mac := hmac.New(sha256.New, []byte(s.APISecret))
	_, _ = mac.Write([]byte(signingString))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	headers := cloneHeaders(request.Headers)
	headers["X-API-Key"] = s.APIKey
	headers["X-Aux-Date"] = date
	headers["X-Signature"] = fmt.Sprintf(
		`Signature keyId="%s",algorithm="hmac-sha256",headers="(request-target) x-aux-date",signature="%s",nonce="%s"`,
		s.APIKey,
		signature,
		nonce,
	)
	request.Headers = headers
	return request, nil
}

func newRESTNonce() (string, error) {
	var uuid [16]byte
	if _, err := rand.Read(uuid[:]); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return hex.EncodeToString(uuid[:]), nil
}

type AuthService struct {
	broker *Broker
}

func (s *AuthService) SendEmailOTP(ctx context.Context) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthSendOTP); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "auth.send_email_otp", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/registration/send-email-otp"),
		Headers: s.broker.apiHeaders(),
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *AuthService) GetTradingToken(
	ctx context.Context,
	otpType string,
	passcode string,
) (domain.RawPayload, error) {
	return s.GetTradingTokenWithRequest(ctx, sdktrading.TradingTokenRequest{
		OTPType:  sdktrading.OTPType(otpType),
		Passcode: passcode,
	})
}

func (s *AuthService) GetTradingTokenWithRequest(
	ctx context.Context,
	request sdktrading.TradingTokenRequest,
) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthTradingToken); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "auth.get_trading_token", transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/registration/trading-token"),
		Headers: withContentType(s.broker.apiHeaders()),
		JSON: map[string]any{
			"otpType":  string(request.OTPType),
			"passcode": request.Passcode,
		},
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (b *Broker) signer() *Signer {
	if b.config.APIKey == "" || b.config.APISecret == "" {
		return nil
	}
	return &Signer{APIKey: b.config.APIKey, APISecret: b.config.APISecret}
}

func (b *Broker) send(
	ctx context.Context,
	operation string,
	request transport.HTTPRequest,
) (transport.HTTPResponse, error) {
	signer := b.signer()
	if signer != nil {
		signed, err := signer.Sign(ctx, request)
		if err != nil {
			return transport.HTTPResponse{}, err
		}
		request = signed
	}
	response, err := b.config.HTTPTransport.Send(ctx, request)
	if err != nil {
		return transport.HTTPResponse{}, err
	}
	if response.StatusCode >= 400 {
		body, _ := response.Body.(map[string]any)
		code := stringify(body["code"])
		message := stringify(body["message"])
		if message == "" {
			message = fmt.Sprintf("DNSE request failed with status %d", response.StatusCode)
		}
		return transport.HTTPResponse{}, errors.BrokerRejected("dnse", operation, code, message, response.Body)
	}
	return response, nil
}

func (b *Broker) url(path string) string {
	return strings.TrimRight(b.config.BaseURL, "/") + path
}

func (b *Broker) apiHeaders() map[string]string {
	headers := map[string]string{}
	if b.config.APIKey != "" {
		headers["X-API-Key"] = b.config.APIKey
	}
	return headers
}

func (b *Broker) tradingHeaders(includeContentType bool) map[string]string {
	headers := b.apiHeaders()
	if b.config.TradingToken != "" {
		headers["trading-token"] = b.config.TradingToken
	}
	if includeContentType {
		headers["Content-Type"] = "application/json"
	}
	return headers
}

func cloneHeaders(headers map[string]string) map[string]string {
	out := make(map[string]string, len(headers))
	for key, value := range headers {
		out[key] = value
	}
	return out
}

func withContentType(headers map[string]string) map[string]string {
	out := cloneHeaders(headers)
	out["Content-Type"] = "application/json"
	return out
}
