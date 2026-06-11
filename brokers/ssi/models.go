package ssi

import (
	"encoding/json"

	"github.com/vnbrokers/vnbrokers-go/domain"
)

func UnmarshalRawPayload(payload domain.RawPayload, out any) error {
	if len(payload.Bytes) > 0 {
		return json.Unmarshal(payload.Bytes, out)
	}
	bytes, err := json.Marshal(payload.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}

type TradingResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type TradingTokenRequest struct {
	TwoFactorType int
	Code          string
	IsSave        bool
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}
