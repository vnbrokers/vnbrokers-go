package tcbs

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type AuthService struct {
	broker *Broker
}

func (s *AuthService) GetToken(ctx context.Context, apiKey, otp string) (TokenResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthTradingToken); err != nil {
		return TokenResponse{}, err
	}
	response, err := s.broker.send(ctx, "auth.get_token", false, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/gaia/v1/oauth2/openapi/token"),
		Headers: s.broker.headers(false, true),
		JSON: TokenRequest{
			APIKey: apiKey,
			OTP:    otp,
		},
	})
	if err != nil {
		return TokenResponse{}, err
	}
	var tokenResponse TokenResponse
	if err := decode(response, &tokenResponse); err != nil {
		return TokenResponse{}, sdkerrors.Decode("tcbs", "auth.get_token", "decode token response", response.Body, err)
	}
	s.broker.accessToken = tokenResponse.Token
	return tokenResponse, nil
}

func (b *Broker) send(
	ctx context.Context,
	operation string,
	authenticated bool,
	request transport.HTTPRequest,
) (transport.HTTPResponse, error) {
	if authenticated {
		request.Headers = b.withAuthorization(request.Headers)
	}
	response, err := b.config.HTTPTransport.Send(ctx, request)
	if err != nil {
		return transport.HTTPResponse{}, err
	}
	if response.StatusCode >= 400 {
		body := expectObject(response.Body)
		code := stringify(body["code"])
		message := stringify(body["message"])
		if message == "" {
			message = fmt.Sprintf("TCBS request failed with status %d", response.StatusCode)
		}
		return transport.HTTPResponse{}, sdkerrors.BrokerRejected("tcbs", operation, code, message, response.Body)
	}
	return response, nil
}

func (b *Broker) url(path string) string {
	return strings.TrimRight(b.config.BaseURL, "/") + path
}

func (b *Broker) headers(authenticated bool, includeContentType bool) map[string]string {
	headers := map[string]string{
		"Accept": "application/json",
	}
	if includeContentType {
		headers["Content-Type"] = "application/json"
	}
	if authenticated {
		headers = b.withAuthorization(headers)
	}
	return headers
}

func (b *Broker) withAuthorization(headers map[string]string) map[string]string {
	out := cloneHeaders(headers)
	if b.accessToken != "" {
		out["Authorization"] = "Bearer " + b.accessToken
	}
	return out
}

func decode(response transport.HTTPResponse, out any) error {
	if len(response.Raw) > 0 {
		return json.Unmarshal(response.Raw, out)
	}
	payload, err := json.Marshal(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, out)
}
