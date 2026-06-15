package dto

type SubscribeStockOrdersRequest struct{}

type StockOrderEvent struct {
	Object      string `json:"object"`
	AccountNo   string `json:"accountNo"`
	OrderID     string `json:"orderId"`
	ExecType    string `json:"execType"`
	OrderQtty   string `json:"orderQtty"`
	Symbol      string `json:"symbol"`
	PriceType   string `json:"priceType"`
	TxTime      string `json:"txTime"`
	TxDate      string `json:"txDate"`
	ExpDate     string `json:"expDate"`
	TimeType    string `json:"timeType"`
	OrStatus    string `json:"orStatus"`
	LimitPrice  string `json:"limitPrice"`
	RemainQtty  string `json:"remainQtty"`
	Via         string `json:"via"`
	QuotePrice  string `json:"quotePrice"`
	TradePlace  string `json:"tradePlace"`
	MatchType   string `json:"matchType"`
	IsDisposal  string `json:"isDisposal"`
	IsCancel    string `json:"isCancel"`
	IsAmend     string `json:"isAmend"`
	UserName    string `json:"userName"`
	ORSOrderID  string `json:"orsOrderId"`
	SecType     string `json:"secType"`
	IsFOOrder   string `json:"isFOOrder"`
	ODTimestamp string `json:"odTimeStamp"`
}

type SubscribeDerivativeOrdersRequest struct{}

type DerivativeOrderEvent struct {
	SubAccount   string  `json:"subAccount"`
	OrderNo      string  `json:"orderNo"`
	PKOrderNo    string  `json:"pkOrderNo"`
	OrderTime    string  `json:"orderTime"`
	AccountCode  string  `json:"accountCode"`
	Side         string  `json:"side"`
	Symbol       string  `json:"symbol"`
	Volume       int64   `json:"volume"`
	ShowPrice    string  `json:"showPrice"`
	MatchVolume  int64   `json:"matchVolume"`
	MatchPriceBQ float64 `json:"matchPriceBQ"`
	Status       string  `json:"status"`
	OrderStatus  string  `json:"orderStatus"`
	Channel      string  `json:"channel"`
	Group        string  `json:"group"`
	CancelTime   string  `json:"cancelTime"`
	IsCancel     string  `json:"isCancel"`
	IsAmend      string  `json:"isAmend"`
	Info         string  `json:"info"`
	MaxPrice     int64   `json:"maxPrice"`
	MatchValue   float64 `json:"matchValue"`
	Quote        string  `json:"quote"`
	AutoType     string  `json:"autoType"`
	Product      string  `json:"product"`
	OrderType    string  `json:"orderType"`
	Source       string  `json:"source"`
	TraderCode   string  `json:"traderCode"`
}

type SubscribeDerivativeOpenPositionsRequest struct{}

type DerivativeOpenPositionEvent struct {
	SubAccount    string  `json:"subAccount"`
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	Account       string  `json:"account"`
	LastPrice     float64 `json:"lastPrice"`
	AvgRemain     float64 `json:"avgRemain"`
	IMValue       float64 `json:"imValue"`
	Net           float64 `json:"net"`
	StopLoss      float64 `json:"stoploss"`
	TakeProfit    float64 `json:"takeprofit"`
	PCRemain      float64 `json:"pcRemain"`
	VMRemain      float64 `json:"vmRemain"`
	DueDate       string  `json:"duedate"`
	NetOffVolume  int64   `json:"netoffvol"`
	TriggerType   string  `json:"triggerType"`
	CallbackPoint float64 `json:"callBackPoint"`
	TrailingPrice float64 `json:"trailingPrice"`
	TotalVMValue  float64 `json:"totalVmValue"`
}

type SubscribeDerivativeConditionalOrdersRequest struct{}

type DerivativeConditionalOrderEvent struct {
	SubAccount       string  `json:"subAccount"`
	OrderNo          string  `json:"orderNo"`
	GroupOrder       string  `json:"groupOrder"`
	PKOrderNo        string  `json:"pkOrderNo"`
	AccountCode      string  `json:"accountCode"`
	Side             string  `json:"side"`
	Symbol           string  `json:"symbol"`
	Volume           int64   `json:"volume"`
	ShowPrice        string  `json:"showPrice"`
	Condition        string  `json:"condition"`
	Result           string  `json:"result"`
	ActiveTime       string  `json:"activeTime"`
	SendTime         string  `json:"sendTime"`
	CancelTime       string  `json:"cancelTime"`
	Group            string  `json:"group"`
	Channel          string  `json:"channel"`
	SOPrice          float64 `json:"soPrice"`
	OrderType        string  `json:"orderType"`
	FromTime         string  `json:"fromTime"`
	ExpirationTime   string  `json:"expTime"`
	Status           string  `json:"status"`
	Details          string  `json:"details"`
	Message          string  `json:"message"`
	Notes            string  `json:"notes"`
	CallbackPoint    string  `json:"callBackPoint"`
	TrailingPrice    string  `json:"trailingPrice"`
	TriggerCondition string  `json:"triggerCondition"`
}
