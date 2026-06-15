// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type PlaceStockOrderBody struct {
	ExecType  string `json:"execType"`
	Price     int64  `json:"price"`
	PriceType string `json:"priceType"`
	Quantity  int64  `json:"quantity"`
	Symbol    string `json:"symbol"`
}

type PlaceStockOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type UpdateStockOrderBody struct {
	Price    int64 `json:"price"`
	Quantity int64 `json:"quantity"`
}

type UpdateStockOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type CancelStockOrderBody struct {
	OrdersList []StockOrderID `json:"ordersList"`
}

type CancelStockOrderResponse struct {
	Data       []CancelStockOrderResult `json:"data"`
	Object     string                   `json:"object"`
	PageIndex  int64                    `json:"pageIndex"`
	PageSize   int64                    `json:"pageSize"`
	TotalCount int64                    `json:"totalCount"`
}

type CancelStockOrderDetail struct {
	Deleted     string `json:"deleted"`
	ErrorCode   string `json:"errorCode"`
	ErrorMesage string `json:"errorMesage"`
	OrderID     string `json:"orderID"`
}

type CancelStockOrderResult struct {
	Details []CancelStockOrderDetail `json:"details"`
	Object  string                   `json:"object"`
}

type StockOrderID struct {
	OrderID string `json:"orderID"`
}

type StockOrderPage struct {
	Data       []StockOrder `json:"data"`
	Object     string       `json:"object"`
	PageIndex  int64        `json:"pageIndex"`
	PageSize   int64        `json:"pageSize"`
	TotalCount int64        `json:"totalCount"`
}

type StockOrder struct {
	AccountNo    string  `json:"accountNo"`
	BRatio       float64 `json:"bRatio"`
	CancelQtty   float64 `json:"cancelQtty"`
	CodeID       string  `json:"codeID"`
	ExecQtty     float64 `json:"execQtty"`
	ExecType     string  `json:"execType"`
	ExpDate      string  `json:"expDate"`
	FeeAcr       float64 `json:"feeAcr"`
	IsAmend      string  `json:"isAmend"`
	IsCancel     string  `json:"isCancel"`
	IsDisposal   string  `json:"isDisposal"`
	IsFOOrder    string  `json:"isFOOrder"`
	LimitPrice   float64 `json:"limitPrice"`
	MatchAmount  float64 `json:"matchAmount"`
	MatchPrice   float64 `json:"matchPrice"`
	MatchType    string  `json:"matchType"`
	MMType       string  `json:"mmType"`
	Object       string  `json:"object"`
	OdTimeStamp  string  `json:"odTimeStamp"`
	OrStatus     string  `json:"orStatus"`
	OrderID      string  `json:"orderID"`
	OrderQtty    float64 `json:"orderQtty"`
	OrsOrderID   string  `json:"orsOrderID"`
	PriceType    string  `json:"priceType"`
	QuotePrice   float64 `json:"quotePrice"`
	RemainQtty   float64 `json:"remainQtty"`
	SecType      string  `json:"sectype"`
	Symbol       string  `json:"symbol"`
	TaxSellAmout float64 `json:"taxSellAmout"`
	TimeType     string  `json:"timeType"`
	TradePlace   string  `json:"tradePlace"`
	TxDate       string  `json:"txdate"`
	TxTime       string  `json:"txtime"`
	UserName     string  `json:"userName"`
	Via          string  `json:"via"`
}

type StockMatchingDetails struct {
	Data       []StockMatchingDetail `json:"data"`
	Object     string                `json:"object"`
	PageIndex  int64                 `json:"pageIndex"`
	PageSize   int64                 `json:"pageSize"`
	TotalCount int64                 `json:"totalCount"`
}

type StockMatchingDetail struct {
	OrderID    string  `json:"orderId"`
	Price      float64 `json:"price"`
	Qtty       float64 `json:"qtty"`
	QuotePrice float64 `json:"quotePrice"`
	QuoteQtty  float64 `json:"quoteQtty"`
	Side       string  `json:"side"`
	Symbol     string  `json:"symbol"`
	TimeExec   float64 `json:"timeExec"`
	TradeID    string  `json:"tradeId"`
}
