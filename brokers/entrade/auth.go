package entrade

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type AuthService struct {
	broker *Broker
}

func (s *AuthService) Login(ctx context.Context, username string, password string) (LoginResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthTradingToken); err != nil {
		return LoginResponse{}, err
	}
	response, err := s.broker.send(ctx, "auth.login", false, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.authURL("/v2/auth"),
		Headers: s.broker.headers(false, true),
		JSON: LoginRequest{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return LoginResponse{}, err
	}
	var loginResponse LoginResponse
	if err := decode(response, &loginResponse); err != nil {
		return LoginResponse{}, sdkerrors.Decode("entrade", "auth.login", "decode login response", response.Body, err)
	}
	s.broker.config.Token = loginResponse.Token
	return loginResponse, nil
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
			message = fmt.Sprintf("Entrade request failed with status %d", response.StatusCode)
		}
		return transport.HTTPResponse{}, sdkerrors.BrokerRejected("entrade", operation, code, message, response.Body)
	}
	return response, nil
}

func (b *Broker) sendRaw(
	ctx context.Context,
	operation string,
	method string,
	path string,
	body any,
) (domain.RawPayload, error) {
	response, err := b.send(ctx, operation, true, transport.HTTPRequest{
		Method:  method,
		URL:     b.url(path),
		Headers: b.headers(true, body != nil),
		JSON:    body,
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (b *Broker) url(path string) string {
	return strings.TrimRight(b.config.BaseURL, "/") + path
}

func (b *Broker) authURL(path string) string {
	return strings.TrimRight(b.config.AuthBaseURL, "/") + path
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
	out := make(map[string]string, len(headers)+1)
	for key, value := range headers {
		out[key] = value
	}
	if b.config.Token != "" {
		out["Authorization"] = "Bearer " + b.config.Token
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
