package dto

type GetDerivativeOrdersRequest struct {
	PageNo    float64
	PageSize  float64
	AccountID string
	Symbol    string
	RefID     string
	OrderType string
	Status    string
}

type GetDerivativeOrdersResponse struct {
	Cmd  string            `json:"cmd"`
	RC   string            `json:"rc"`
	RS   string            `json:"rs"`
	OID  string            `json:"oID"`
	Data []DerivativeOrder `json:"data"`
}

type GetDerivativeConditionalOrdersRequest struct {
	PageNo       string
	PageSize     string
	AccountID    string
	SubAccountID string
	OrderStatus  string
	OrderType    string
	Symbol       string
}

type GetDerivativeConditionalOrdersResponse struct {
	Cmd  string                       `json:"cmd"`
	RC   string                       `json:"rc"`
	RS   string                       `json:"rs"`
	OID  string                       `json:"oID"`
	Data []DerivativeConditionalOrder `json:"data"`
}

type PlaceDerivativeOrderRequest struct{ PlaceDerivativeOrderBody }
type PlaceDerivativeOrderResponse struct {
	Cmd  string                       `json:"cmd"`
	RC   string                       `json:"rc"`
	RS   string                       `json:"rs"`
	OID  string                       `json:"oID"`
	Data []PlaceDerivativeOrderResult `json:"data"`
}

type PlaceDerivativeConditionalOrderRequest struct {
	PlaceDerivativeConditionalOrderBody
}
type PlaceDerivativeConditionalOrderResponse struct {
	Cmd  string                                  `json:"cmd"`
	RC   string                                  `json:"rc"`
	RS   string                                  `json:"rs"`
	OID  string                                  `json:"oID"`
	Data []PlaceDerivativeConditionalOrderResult `json:"data"`
}

type UpdateDerivativeOrderRequest struct{ UpdateDerivativeOrderBody }
type UpdateDerivativeOrderResponse struct {
	Cmd  string                        `json:"cmd"`
	RC   string                        `json:"rc"`
	RS   string                        `json:"rs"`
	OID  string                        `json:"oID"`
	Data []UpdateDerivativeOrderResult `json:"data"`
}

type UpdateDerivativeConditionalOrderRequest struct {
	UpdateDerivativeConditionalOrderBody
}
type UpdateDerivativeConditionalOrderResponse struct {
	Cmd  string                                   `json:"cmd"`
	RC   string                                   `json:"rc"`
	RS   string                                   `json:"rs"`
	OID  string                                   `json:"oID"`
	Data []UpdateDerivativeConditionalOrderResult `json:"data"`
}

type CancelDerivativeOrderRequest struct{ CancelDerivativeOrderBody }
type CancelDerivativeOrderResponse struct {
	Cmd  string                        `json:"cmd"`
	RC   string                        `json:"rc"`
	RS   string                        `json:"rs"`
	OID  string                        `json:"oID"`
	Data []CancelDerivativeOrderResult `json:"data"`
}

type CancelDerivativeConditionalOrderRequest struct {
	CancelDerivativeConditionalOrderBody
}
type CancelDerivativeConditionalOrderResponse struct {
	Cmd  string                                   `json:"cmd"`
	RC   string                                   `json:"rc"`
	RS   string                                   `json:"rs"`
	OID  string                                   `json:"oID"`
	Data []CancelDerivativeConditionalOrderResult `json:"data"`
}
