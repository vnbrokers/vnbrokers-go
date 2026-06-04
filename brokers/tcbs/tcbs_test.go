package tcbs

import (
	"context"
	"errors"
	"testing"

	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeHTTPTransport struct {
	requests  []transport.HTTPRequest
	responses []transport.HTTPResponse
}

func (f *fakeHTTPTransport) Send(_ context.Context, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	f.requests = append(f.requests, request)
	response := f.responses[0]
	f.responses = f.responses[1:]
	return response, nil
}

func TestGetTokenBuildsTCBSRequestAndStoresToken(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"token": "jwt-123",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Auth().GetToken(context.Background(), "api-key", "111111")
	if err != nil {
		t.Fatalf("get token: %v", err)
	}
	if response.Token != "jwt-123" {
		t.Fatalf("token = %s", response.Token)
	}
	if broker.accessToken != "jwt-123" {
		t.Fatalf("stored token = %s", broker.accessToken)
	}

	request := httpTransport.requests[0]
	if request.Method != "POST" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://api.tcbs.example/gaia/v1/oauth2/openapi/token" {
		t.Fatalf("url = %s", request.URL)
	}
	body, ok := request.JSON.(TokenRequest)
	if !ok {
		t.Fatalf("json body type = %T", request.JSON)
	}
	if body.APIKey != "api-key" || body.OTP != "111111" {
		t.Fatalf("body = %#v", body)
	}
	if request.Headers["Authorization"] != "" {
		t.Fatalf("token request should not include authorization header")
	}
}

func TestGetSubAccountInfoBuildsBearerRequest(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"basicInfo": map[string]any{
					"code105C": "105C001",
					"status":   "ACTIVE",
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	response, err := broker.Account().GetSubAccountInfo(
		context.Background(),
		"105C001",
		"basicInfo,personalInfo",
	)
	if err != nil {
		t.Fatalf("get sub account info: %v", err)
	}
	if response.BasicInfo == nil || response.BasicInfo.Code105C != "105C001" {
		t.Fatalf("response = %#v", response)
	}

	request := httpTransport.requests[0]
	if request.Method != "GET" {
		t.Fatalf("method = %s", request.Method)
	}
	if request.URL != "https://api.tcbs.example/eros/v2/get-profile/by-username/105C001?fields=basicInfo%2CpersonalInfo" {
		t.Fatalf("url = %s", request.URL)
	}
	if request.Headers["Authorization"] != "Bearer jwt-123" {
		t.Fatalf("authorization = %s", request.Headers["Authorization"])
	}
	if request.Headers["Accept"] != "application/json" {
		t.Fatalf("accept = %s", request.Headers["Accept"])
	}
}

func TestGetBasicInfoMapsTCBSAccountToDomainAccount(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 200,
			Body: map[string]any{
				"basicInfo": map[string]any{
					"code105C":   "105C001",
					"status":     "ACTIVE",
					"tcbsId":     "0001738764",
					"depository": true,
				},
				"personalInfo": map[string]any{
					"fullName": "Nguyen Van Nam",
				},
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		AccessToken:   "jwt-123",
		HTTPTransport: httpTransport,
	})

	account, err := broker.Account().GetBasicInfo(context.Background(), "105C001")
	if err != nil {
		t.Fatalf("get basic info: %v", err)
	}

	if account.Broker != "tcbs" {
		t.Fatalf("broker = %s", account.Broker)
	}
	if account.AccountID != "105C001" {
		t.Fatalf("account id = %s", account.AccountID)
	}
	if account.DisplayName != "Nguyen Van Nam" {
		t.Fatalf("display name = %s", account.DisplayName)
	}
	if account.Raw.Source != "tcbs" {
		t.Fatalf("raw source = %s", account.Raw.Source)
	}
	if _, ok := account.Raw.Data.(AccountInformationResponse); !ok {
		t.Fatalf("raw data type = %T", account.Raw.Data)
	}
}

func TestMapAccountInformationFallsBackToCustodyCode(t *testing.T) {
	account := MapAccountInformation("105C001", AccountInformationResponse{})

	if account.Broker != "tcbs" {
		t.Fatalf("broker = %s", account.Broker)
	}
	if account.AccountID != "105C001" {
		t.Fatalf("account id = %s", account.AccountID)
	}
	if account.DisplayName != "" {
		t.Fatalf("display name = %s", account.DisplayName)
	}
	if account.Raw.Source != "tcbs" {
		t.Fatalf("raw source = %s", account.Raw.Source)
	}
	if _, ok := account.Raw.Data.(AccountInformationResponse); !ok {
		t.Fatalf("raw data type = %T", account.Raw.Data)
	}
}

func TestTCBSHTTPErrorMapsBrokerRejectedCodeAndMessage(t *testing.T) {
	httpTransport := &fakeHTTPTransport{
		responses: []transport.HTTPResponse{{
			StatusCode: 400,
			Body: map[string]any{
				"code":    "203074",
				"message": "The API Key is invalid",
			},
		}},
	}
	broker := NewBroker(Config{
		BaseURL:       "https://api.tcbs.example",
		HTTPTransport: httpTransport,
	})

	_, err := broker.Auth().GetToken(context.Background(), "bad-key", "111111")
	if err == nil {
		t.Fatal("expected error")
	}
	var brokerError *sdkerrors.BrokerError
	if !errors.As(err, &brokerError) {
		t.Fatalf("error type = %T", err)
	}
	if brokerError.Broker != "tcbs" {
		t.Fatalf("broker = %s", brokerError.Broker)
	}
	if brokerError.Operation != "auth.get_token" {
		t.Fatalf("operation = %s", brokerError.Operation)
	}
	if brokerError.Code != "203074" {
		t.Fatalf("code = %s", brokerError.Code)
	}
	if brokerError.Message != "The API Key is invalid" {
		t.Fatalf("message = %s", brokerError.Message)
	}
}
