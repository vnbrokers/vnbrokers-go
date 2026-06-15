package dto

type TransferBetweenSubaccountsRequest struct {
	TransferBetweenSubaccountsBody
}

type TransferBetweenSubaccountsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type WithdrawDerivativeMarginRequest struct {
	WithdrawDerivativeMarginBody
}

type WithdrawDerivativeMarginResponse struct {
	Cmd  string       `json:"cmd"`
	RC   string       `json:"rc"`
	RS   string       `json:"rs"`
	OID  string       `json:"oID"`
	Data []RawMessage `json:"data"`
}

type DepositDerivativeMarginRequest struct {
	DepositDerivativeMarginBody
}

type DepositDerivativeMarginResponse struct {
	Cmd  string                        `json:"cmd"`
	RC   string                        `json:"rc"`
	RS   string                        `json:"rs"`
	OID  string                        `json:"oID"`
	Data []DerivativeMarginTransaction `json:"data"`
}
