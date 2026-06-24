// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type APISuccessEnvelope struct {
	Status    *int64  `json:"status,omitempty"`
	ErrorCode *string `json:"error_code,omitempty"`
	Message   *string `json:"message,omitempty"`
}

type StandardError struct {
	// Application-level error code
	ErrorCode *string `json:"error_code,omitempty"`
	// Human-readable status or error message
	Message *string `json:"message,omitempty"`
	Data    *any    `json:"data,omitempty"`
	Result  *any    `json:"result,omitempty"`
	// HTTP-like status field used by some endpoints
	Status *int64 `json:"status,omitempty"`
}
