// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type DerivativeCash struct {
	Assets           float64 `json:"assets"`
	AvaiCash         float64 `json:"avaiCash"`
	AvaiColla        float64 `json:"avaiColla"`
	Cash             float64 `json:"cash"`
	CashOut          float64 `json:"cashOut"`
	CashWithdraw     float64 `json:"cashWithdraw"`
	Cashavaiwithdraw float64 `json:"cashavaiwithdraw"`
	Collateral       float64 `json:"collateral"`
	Color            string  `json:"color"`
	Debt             string  `json:"debt"`
	DM               float64 `json:"dm"`
	FeeCTCK          float64 `json:"feeCTCK"`
	FeeHNX           float64 `json:"feeHNX"`
	FeeMan           float64 `json:"feeMan"`
	FeePos           float64 `json:"feePos"`
	IM               float64 `json:"im"`
	Info             string  `json:"info"`
	Limit            float64 `json:"limit"`
	MR               float64 `json:"mr"`
	NAV              float64 `json:"nav"`
	Net              string  `json:"net"`
	Others           float64 `json:"others"`
	Package          string  `json:"package"`
	Product          string  `json:"product"`
	Status           string  `json:"status"`
	Stock            float64 `json:"stock"`
	Tax              float64 `json:"tax"`
	Tienbosung       float64 `json:"tienbosung"`
	Tyle             string  `json:"tyle"`
	Type             string  `json:"type"`
	UnrelizeVM       float64 `json:"unrelizeVM"`
	VM               float64 `json:"vm"`
	VmEod            string  `json:"vm_eod"`
	Vmunpay          float64 `json:"vmunpay"`
	W1               float64 `json:"w1"`
	W2               float64 `json:"w2"`
}

type ClosedDerivativePosition struct {
	ClosePC       string `json:"closePC"`
	ClosePosition string `json:"closePosition"`
	ClosePrice    string `json:"closePrice"`
	CloseVM       string `json:"closeVM"`
	Fee           string `json:"fee"`
	OpenPrice     string `json:"openPrice"`
	Side          string `json:"side"`
	Symbol        string `json:"symbol"`
	Tax           string `json:"tax"`
	Time          string `json:"time"`
	Unrealize     string `json:"unrealize"`
}

type OpenDerivativePosition struct {
	Account    string  `json:"account"`
	AvgRemain  float64 `json:"avg_remain"`
	Deliver    int64   `json:"deliver"`
	Duedate    string  `json:"duedate"`
	IM         string  `json:"im"`
	ImValue    float64 `json:"imValue"`
	LastPrice  float64 `json:"lastPrice"`
	MrValue    float64 `json:"mrValue"`
	Net        int64   `json:"net"`
	Netoffvol  int64   `json:"netoffvol"`
	PcRemain   float64 `json:"pc_remain"`
	Receive    int64   `json:"receive"`
	Side       string  `json:"side"`
	Stoploss   string  `json:"stoploss"`
	Symbol     string  `json:"symbol"`
	Takeprofit string  `json:"takeprofit"`
	VmValue    float64 `json:"vmValue"`
	VmRemain   float64 `json:"vm_remain"`
	Wapb       float64 `json:"wapb"`
	Wasp       float64 `json:"wasp"`
}
