package ssi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type AuthService struct {
	broker *Broker
}

func (s *AuthService) GetOTP(ctx context.Context) (domain.RawPayload, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthSendOTP); err != nil {
		return domain.RawPayload{}, err
	}
	response, err := s.broker.send(ctx, "auth.get_otp", false, false, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/api/v2/Trading/GetOTP"),
		Headers: s.broker.headers(false, true),
		JSON: map[string]any{
			"consumerID":     s.broker.config.ConsumerID,
			"consumerSecret": s.broker.config.TradingConsumerSecret,
		},
	})
	if err != nil {
		return domain.RawPayload{}, err
	}
	return rawPayload(response.Body, response.Raw), nil
}

func (s *AuthService) GetAccessToken(
	ctx context.Context,
) (AccessTokenResponse, error) {
	response, err := s.broker.send(ctx, "auth.get_data_access_token", false, false, transport.HTTPRequest{
		Method:  "POST",
		URL:     strings.TrimRight(s.broker.config.DataBaseURL, "/") + "/api/v2/Market/AccessToken",
		Headers: s.broker.headers(false, true),
		JSON: map[string]any{
			"consumerID":     s.broker.config.ConsumerID,
			"consumerSecret": s.broker.config.DataConsumerSecret,
		},
	})
	if err != nil {
		return AccessTokenResponse{}, err
	}
	var envelope TradingResponse[AccessTokenResponse]
	if err := decode(response, &envelope); err != nil {
		return AccessTokenResponse{}, sdkerrors.Decode("ssi", "auth.get_data_access_token", "decode SSI data access token response", response.Body, err)
	}
	s.broker.dataAccessToken = envelope.Data.AccessToken
	return envelope.Data, nil
}

func (s *AuthService) GetTradingToken(
	ctx context.Context,
	request TradingTokenRequest,
) (AccessTokenResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthTradingToken); err != nil {
		return AccessTokenResponse{}, err
	}
	response, err := s.broker.send(ctx, "auth.get_trading_token", false, false, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/api/v2/Trading/AccessToken"),
		Headers: s.broker.headers(false, true),
		JSON: map[string]any{
			"consumerID":     s.broker.config.ConsumerID,
			"consumerSecret": s.broker.config.TradingConsumerSecret,
			"twoFactorType":  request.TwoFactorType,
			"code":           request.Code,
			"isSave":         request.IsSave,
		},
	})
	if err != nil {
		return AccessTokenResponse{}, err
	}
	var envelope TradingResponse[AccessTokenResponse]
	if err := decode(response, &envelope); err != nil {
		return AccessTokenResponse{}, sdkerrors.Decode("ssi", "auth.get_trading_token", "decode SSI trading access token response", response.Body, err)
	}
	s.broker.tradingAccessToken = envelope.Data.AccessToken
	return envelope.Data, nil
}

func (b *Broker) send(
	ctx context.Context,
	operation string,
	authenticated bool,
	sign bool,
	request transport.HTTPRequest,
) (transport.HTTPResponse, error) {
	if authenticated {
		request.Headers = b.withAuthorization(request.Headers)
	}
	if sign {
		signer := Signer{PrivateKey: b.config.PrivateKey}
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
		message := stringify(body["message"])
		if message == "" {
			message = fmt.Sprintf("SSI request failed with status %d", response.StatusCode)
		}
		return transport.HTTPResponse{}, sdkerrors.BrokerRejected("ssi", operation, stringify(firstNonNil(body["status"], body["code"])), message, response.Body)
	}
	if rejected, code, message := rejectedStatus(response.Body); rejected {
		if message == "" {
			message = "SSI request rejected"
		}
		return transport.HTTPResponse{}, sdkerrors.BrokerRejected("ssi", operation, code, message, response.Body)
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

func (b *Broker) query(path string, values url.Values) string {
	if len(values) == 0 {
		return b.url(path)
	}
	return b.url(path + "?" + values.Encode())
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

func rawPayload(data any, raw []byte) domain.RawPayload {
	return domain.RawPayload{Source: "ssi", Data: data, Bytes: raw}
}

func cloneHeaders(headers map[string]string) map[string]string {
	out := make(map[string]string, len(headers))
	for key, value := range headers {
		out[key] = value
	}
	return out
}

func rejectedStatus(body any) (bool, string, string) {
	payload, _ := body.(map[string]any)
	status := stringify(payload["status"])
	if status == "" || status == "200" || strings.EqualFold(status, "Success") {
		return false, "", ""
	}
	return true, status, stringify(payload["message"])
}

func stringify(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d", int64(v))
		}
		return fmt.Sprintf("%v", v)
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case json.Number:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func firstNonNil(values ...any) any {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}
