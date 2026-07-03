package entrade

import (
	"context"
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/brokers/entrade/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type AuthService struct {
	broker *Broker
}

func (s *AuthService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthTradingToken); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.login", false, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.authURL("/v2/auth"),
		Headers: s.broker.headers(false, true),
		JSON:    request,
	})
	if err != nil {
		return nil, err
	}
	loginResponse, err := httpx.DecodeResponse[dto.LoginResponse]("entrade", "auth.login", "decode login response", response)
	if err != nil {
		return nil, err
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

func (b *Broker) authURL(path string) string {
	return httpx.URL(b.config.AuthBaseURL, path, nil)
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
