package ssi

import (
	"testing"

	"github.com/vnbrokers/vnbrokers-go/domain"
)

func TestUnmarshalRawPayloadDecodesFromBytes(t *testing.T) {
	payload := domain.RawPayload{
		Source: "ssi",
		Bytes:  []byte(`{"message":"Success","status":200,"data":{"accessToken":"abc123"}}`),
	}

	var response TradingResponse[AccessTokenResponse]
	if err := UnmarshalRawPayload(payload, &response); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if response.Status != 200 {
		t.Fatalf("status = %d", response.Status)
	}
	if response.Data.AccessToken != "abc123" {
		t.Fatalf("access token = %s", response.Data.AccessToken)
	}
}

func TestUnmarshalRawPayloadDecodesFromData(t *testing.T) {
	payload := domain.RawPayload{
		Source: "ssi",
		Data: map[string]any{
			"message": "Success",
			"status":  200,
			"data": map[string]any{
				"accessToken": "def456",
			},
		},
	}

	var response TradingResponse[AccessTokenResponse]
	if err := UnmarshalRawPayload(payload, &response); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if response.Status != 200 {
		t.Fatalf("status = %d", response.Status)
	}
	if response.Data.AccessToken != "def456" {
		t.Fatalf("access token = %s", response.Data.AccessToken)
	}
}
