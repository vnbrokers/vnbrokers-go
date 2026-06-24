// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

// Schema inferred from finhay.sh. The client reads data.user_id from this payload.
type GetCurrentUserResponse struct {
	Status    *int64          `json:"status,omitempty"`
	ErrorCode *string         `json:"error_code,omitempty"`
	Message   *string         `json:"message,omitempty"`
	Data      UserProfileLite `json:"data,omitempty"`
}

// Schema inferred from finhay.sh. The client accepts either result[] or data[] for the sub-account array.
type GetSubAccountsResponse struct {
	Status    *int64       `json:"status,omitempty"`
	ErrorCode *string      `json:"error_code,omitempty"`
	Message   *string      `json:"message,omitempty"`
	Result    []SubAccount `json:"result,omitempty"`
	Data      []SubAccount `json:"data,omitempty"`
}

// Schema inferred from finhay.sh infer flow.
type SubAccount struct {
	// Sub-account type; client uppercases this to build env vars such as SUB_ACCOUNT_NORMAL
	TypeValue string `json:"type,omitempty"`
	ID        string `json:"id,omitempty"`
	// Extended sub-account ID. Values ending with .4 are treated as the order-execution account by finhay.sh.
	SubAccountExt string `json:"sub_account_ext,omitempty"`
}

type TwoFactorRequestPayload struct {
	// Enum: EMAIL
	Channel string `json:"channel"`
}

// Schema inferred from finhay.sh; ticket_id and masked_destination are consumed by the client.
type TwoFactorRequestResponse struct {
	TicketID          string  `json:"ticket_id"`
	MaskedDestination string  `json:"masked_destination,omitempty"`
	Message           *string `json:"message,omitempty"`
	ErrorCode         *string `json:"error_code,omitempty"`
}

type TwoFactorRevokePayload struct {
	SessionToken string `json:"session_token"`
}

// Schema is not explicitly documented in the repository; represented as a permissive success envelope inferred from CLI usage.
type TwoFactorRevokeResponse struct {
	Status    *int64  `json:"status,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	Message   *string `json:"message,omitempty"`
}

type TwoFactorVerifyPayload struct {
	TicketID string `json:"ticket_id"`
	OTPCode  string `json:"otp_code"`
}

// Schema inferred from finhay.sh; session_token, expires_at, and expires_at_epoch are consumed by the client.
type TwoFactorVerifyResponse struct {
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at,omitempty"`
	// Unix epoch seconds
	ExpiresAtEpoch int64   `json:"expires_at_epoch,omitempty"`
	Message        *string `json:"message,omitempty"`
	ErrorCode      *string `json:"error_code,omitempty"`
}

// Minimal user profile object inferred from finhay.sh infer flow.
type UserProfileLite struct {
	UserID int64 `json:"user_id"`
}
