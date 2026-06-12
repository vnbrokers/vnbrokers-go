package dto

type GetDerivativesRequest struct{}

type Derivative struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
}

type GetDerivativesResponse struct {
	Data  []Derivative `json:"data"`
	Total int          `json:"total,omitempty"`
}
