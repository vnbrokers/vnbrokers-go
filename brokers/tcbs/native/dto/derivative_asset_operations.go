package dto

type GetDerivativeCashRequest struct {
	AccountID    string
	SubAccountID string
	GetType      string
}

type GetDerivativeCashResponse struct {
	Cmd  string         `json:"cmd"`
	RC   string         `json:"rc"`
	RS   string         `json:"rs"`
	OID  string         `json:"oID"`
	Data DerivativeCash `json:"data"`
}

type GetClosedDerivativePositionsRequest struct {
	AccountID    string
	SubAccountID string
	Symbol       string
	PageNo       int64
	PageSize     int64
}

type GetClosedDerivativePositionsResponse struct {
	Cmd  string                     `json:"cmd"`
	RC   string                     `json:"rc"`
	RS   string                     `json:"rs"`
	OID  string                     `json:"oID"`
	Data []ClosedDerivativePosition `json:"data"`
}

type GetOpenDerivativePositionsRequest struct {
	AccountID    string
	SubAccountID string
}

type GetOpenDerivativePositionsResponse struct {
	Cmd  string                 `json:"cmd"`
	RC   string                 `json:"rc"`
	RS   string                 `json:"rs"`
	OID  string                 `json:"oID"`
	Data OpenDerivativePosition `json:"data"` // ??? TODO: Is an array or object?

}
