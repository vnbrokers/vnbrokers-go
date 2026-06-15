package dto

type GetTokenRequest struct {
	APIKey string `json:"apiKey"`
	OTP    string `json:"otp"`
}

type GetTokenResponse struct {
	Token string `json:"token"`
}

type TokenError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
