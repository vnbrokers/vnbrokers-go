package fhsc

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/internal/httpx"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityGetCurrentUser  core.Capability = "native.helper.get_current_user"
	CapabilityGetSubAccounts  core.Capability = "native.helper.get_sub_accounts"
	CapabilityVerifyTwoFactor core.Capability = "native.auth.verify_two_factor"
	CapabilityRevokeTwoFactor core.Capability = "native.auth.revoke_two_factor"
)

type AuthService struct {
	broker *Broker
}

type Identity struct {
	UserID      int64
	SubAccounts []dto.SubAccount
}

func (i Identity) OrderSubAccount() *dto.SubAccount {
	for index := range i.SubAccounts {
		candidate := i.SubAccounts[index]
		if strings.HasSuffix(candidate.SubAccountExt, ".4") {
			return &candidate
		}
	}
	if len(i.SubAccounts) == 0 {
		return nil
	}
	candidate := i.SubAccounts[0]
	return &candidate
}

func (i Identity) SubAccountsByType() map[string]dto.SubAccount {
	out := make(map[string]dto.SubAccount, len(i.SubAccounts))
	for _, account := range i.SubAccounts {
		if account.TypeValue == "" {
			continue
		}
		out[strings.ToUpper(account.TypeValue)] = account
	}
	return out
}

func (s *AuthService) GetCurrentUser(ctx context.Context) (*dto.GetCurrentUserResponse, error) {
	if err := s.broker.RequireCapability(CapabilityGetCurrentUser); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.get_current_user", true, transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url("/users/v1/users/me"),
		Headers: s.broker.headers(true, false),
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.GetCurrentUserResponse]("fhsc", "auth.get_current_user", "decode current user response", response)
	if err != nil {
		return nil, err
	}
	if out.Data.UserID != 0 {
		s.broker.config.UserID = out.Data.UserID
	}
	return out, nil
}

func (s *AuthService) GetSubAccounts(ctx context.Context, userID int64) (*dto.GetSubAccountsResponse, error) {
	if err := s.broker.RequireCapability(CapabilityGetSubAccounts); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.get_sub_accounts", true, transport.HTTPRequest{
		Method:  "GET",
		URL:     s.broker.url(fmt.Sprintf("/users/v1/users/%d/sub-accounts", userID)),
		Headers: s.broker.headers(true, false),
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.GetSubAccountsResponse]("fhsc", "auth.get_sub_accounts", "decode sub accounts response", response)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *AuthService) RequestTwoFactorOTP(ctx context.Context, request dto.TwoFactorRequestPayload) (*dto.TwoFactorRequestResponse, error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingAuthSendOTP); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.request_two_factor", true, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/auth/v1/openapi/2fa/request"),
		Headers: s.broker.headers(true, true),
		JSON:    request,
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.TwoFactorRequestResponse]("fhsc", "auth.request_two_factor", "decode request 2fa response", response)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *AuthService) VerifyTwoFactorOTP(ctx context.Context, request dto.TwoFactorVerifyPayload) (*dto.TwoFactorVerifyResponse, error) {
	if err := s.broker.RequireCapability(CapabilityVerifyTwoFactor); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.verify_two_factor", true, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/auth/v1/openapi/2fa/verify"),
		Headers: s.broker.headers(true, true),
		JSON:    request,
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.TwoFactorVerifyResponse]("fhsc", "auth.verify_two_factor", "decode verify 2fa response", response)
	if err != nil {
		return nil, err
	}
	if out.SessionToken != "" {
		s.broker.config.TwoFactorToken = out.SessionToken
	}
	return out, nil
}

func (s *AuthService) RevokeTwoFactorSession(ctx context.Context, request dto.TwoFactorRevokePayload) (*dto.TwoFactorRevokeResponse, error) {
	if err := s.broker.RequireCapability(CapabilityRevokeTwoFactor); err != nil {
		return nil, err
	}
	response, err := s.broker.send(ctx, "auth.revoke_two_factor", true, transport.HTTPRequest{
		Method:  "POST",
		URL:     s.broker.url("/auth/v1/openapi/2fa/revoke"),
		Headers: s.broker.headers(true, true),
		JSON:    request,
	})
	if err != nil {
		return nil, err
	}
	out, err := httpx.DecodeResponse[dto.TwoFactorRevokeResponse]("fhsc", "auth.revoke_two_factor", "decode revoke 2fa response", response)
	if err != nil {
		return nil, err
	}
	if request.SessionToken != "" && request.SessionToken == s.broker.config.TwoFactorToken {
		s.broker.config.TwoFactorToken = ""
	}
	return out, nil
}

func (s *AuthService) InferIdentity(ctx context.Context) (*Identity, error) {
	currentUser, err := s.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	subAccounts, err := s.GetSubAccounts(ctx, currentUser.Data.UserID)
	if err != nil {
		return nil, err
	}
	accounts := subAccounts.Result
	if len(accounts) == 0 {
		accounts = subAccounts.Data
	}
	return &Identity{UserID: currentUser.Data.UserID, SubAccounts: accounts}, nil
}

func (s *AuthService) ResolveOrderSubAccount(ctx context.Context) (*dto.SubAccount, error) {
	identity, err := s.InferIdentity(ctx)
	if err != nil {
		return nil, err
	}
	account := identity.OrderSubAccount()
	if account == nil {
		return nil, sdkerrors.BrokerRejected("fhsc", "auth.resolve_order_sub_account", "NO_SUB_ACCOUNT", "no FHSC sub-account found", nil)
	}
	return account, nil
}

func (b *Broker) HasTwoFactorSession() bool {
	return strings.TrimSpace(b.config.TwoFactorToken) != ""
}

func (b *Broker) SetTwoFactorToken(token string) {
	b.config.TwoFactorToken = strings.TrimSpace(token)
}

func (b *Broker) send(ctx context.Context, operation string, authenticated bool, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	if authenticated {
		signed, err := b.signRequest(request)
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
		body := expectObject(response.Body)
		raw := httpx.RawPayload(response)
		if len(response.Raw) > 0 {
			var decoded map[string]any
			if json.Unmarshal(response.Raw, &decoded) == nil {
				body = decoded
			}
		}
		code := stringify(body["error_code"])
		message := stringify(body["message"])
		if message == "" {
			message = fmt.Sprintf("FHSC request failed with status %d", response.StatusCode)
		}
		return transport.HTTPResponse{}, sdkerrors.BrokerRejected("fhsc", operation, code, message, raw)
	}
	return response, nil
}

func (b *Broker) signRequest(request transport.HTTPRequest) (transport.HTTPRequest, error) {
	headers := cloneHeaders(request.Headers)
	if headers == nil {
		headers = map[string]string{}
	}
	parsed, err := url.Parse(request.URL)
	if err != nil {
		return request, err
	}
	signPath := parsed.EscapedPath()
	if parsed.RawQuery != "" {
		signPath += "?" + parsed.RawQuery
	}
	timestamp := fmt.Sprintf("%d", b.config.Now().UTC().UnixMilli())
	nonce := b.config.Nonce()
	bodyHash := ""
	if request.JSON != nil {
		payload, err := json.Marshal(request.JSON)
		if err != nil {
			return request, err
		}
		digest := sha256.Sum256(payload)
		bodyHash = hex.EncodeToString(digest[:])
		headers["X-FH-BODYHASH"] = bodyHash
	}
	signingPayload := timestamp + "\n" + strings.ToUpper(request.Method) + "\n" + signPath + "\n"
	if bodyHash != "" {
		signingPayload += bodyHash
	}
	mac := hmac.New(sha256.New, []byte(b.config.APISecret))
	_, _ = mac.Write([]byte(signingPayload))
	headers["X-FH-APIKEY"] = b.config.APIKey
	if b.config.UserID != 0 {
		headers["X-FH-USER-ID"] = fmt.Sprintf("%d", b.config.UserID)
	}
	headers["X-FH-TIMESTAMP"] = timestamp
	headers["X-FH-NONCE"] = nonce
	headers["X-FH-SIGNATURE"] = hex.EncodeToString(mac.Sum(nil))
	if b.config.TwoFactorToken != "" {
		headers["X-FH-2FA-TOKEN"] = b.config.TwoFactorToken
	}
	if b.config.OpenAPISkillVersion != "" {
		headers["X-FH-OPENAPI-SKILL-VERSION"] = b.config.OpenAPISkillVersion
	}
	if b.config.OpenAPIOperatingSystem != "" {
		headers["X-FH-OPENAPI-OS"] = b.config.OpenAPIOperatingSystem
	}
	if b.config.OpenAPIAgent != "" {
		headers["X-FH-OPENAPI-AGENT"] = b.config.OpenAPIAgent
	}
	if b.config.UserAgent != "" {
		headers["User-Agent"] = b.config.UserAgent
	}
	request.Headers = headers
	return request, nil
}

func (b *Broker) url(path string) string {
	return httpx.URL(b.config.BaseURL, path, nil)
}

func (b *Broker) headers(authenticated bool, includeContentType bool) map[string]string {
	headers := map[string]string{"Accept": "application/json"}
	if includeContentType {
		headers["Content-Type"] = "application/json"
	}
	return headers
}
