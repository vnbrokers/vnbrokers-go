package dto

type OrderEvent struct {
	OrderID             string  `json:"orderID,omitempty"`
	NotifyID            int64   `json:"notifyID,omitempty"`
	Data                any     `json:"data,omitempty"`
	InstrumentID        string  `json:"instrumentID,omitempty"`
	LastAction          string  `json:"lastAction,omitempty"`
	UniqueID            string  `json:"uniqueID,omitempty"`
	BuySell             string  `json:"buySell,omitempty"`
	OrderType           string  `json:"orderType,omitempty"`
	IPAddress           string  `json:"ipAddress,omitempty"`
	Price               float64 `json:"price,omitempty"`
	Prefix              string  `json:"prefix,omitempty"`
	Quantity            int64   `json:"quantity,omitempty"`
	BrokerID            string  `json:"brokerId,omitempty"`
	MarketID            string  `json:"marketID,omitempty"`
	OrigOrderID         string  `json:"origOrderId,omitempty"`
	BrokerIDUpdate      *string `json:"brokerIdUpdate,omitempty"`
	Account             string  `json:"account,omitempty"`
	CancelQuantity      int64   `json:"cancelQty,omitempty"`
	OSQuantity          int64   `json:"osQty,omitempty"`
	FilledQuantity      int64   `json:"filledQty,omitempty"`
	AveragePrice        float64 `json:"avgPrice,omitempty"`
	Channel             string  `json:"channel,omitempty"`
	InputTime           string  `json:"inputTime,omitempty"`
	ModifiedTime        string  `json:"modifiedTime,omitempty"`
	IsForceSell         string  `json:"isForceSell,omitempty"`
	IsShortSell         *string `json:"isShortSell,omitempty"`
	OrderStatus         string  `json:"orderStatus,omitempty"`
	RejectReason        string  `json:"rejectReason,omitempty"`
	OrigRequestID       string  `json:"origRequestID,omitempty"`
	StopOrder           bool    `json:"stopOrder,omitempty"`
	StopPrice           float64 `json:"stopPrice,omitempty"`
	StopType            string  `json:"stopType,omitempty"`
	StopStep            float64 `json:"stopStep,omitempty"`
	ProfitPrice         float64 `json:"profitPrice,omitempty"`
	Modifiable          bool    `json:"modifiable,omitempty"`
	Note                string  `json:"note,omitempty"`
	ApproveComment      string  `json:"approveComment,omitempty"`
	OrderApproval       bool    `json:"orderApproval,omitempty"`
	TaxRate             float64 `json:"taxRate,omitempty"`
	FeeRate             float64 `json:"feeRate,omitempty"`
	Source              string  `json:"source,omitempty"`
	LastOrderUpdateTime string  `json:"lastOrderUpdateTime,omitempty"`
	ExchangeReplyTime   string  `json:"exchangeReplyTime,omitempty"`
	IsCloseout          bool    `json:"isCloseout,omitempty"`
	IsOrderMM           bool    `json:"isOrderMM,omitempty"`
}

type OrderError struct {
	Message       string  `json:"message,omitempty"`
	NotifyID      int64   `json:"notifyID,omitempty"`
	Data          any     `json:"data,omitempty"`
	ErrorCode     string  `json:"errorCode,omitempty"`
	UniqueID      string  `json:"uniqueID,omitempty"`
	ConnectionID  string  `json:"connectionID,omitempty"`
	IPAddress     string  `json:"ipAddress,omitempty"`
	Prefix        string  `json:"prefix,omitempty"`
	OrderID       string  `json:"orderID,omitempty"`
	InstrumentID  string  `json:"instrumentID,omitempty"`
	BuySell       string  `json:"buySell,omitempty"`
	OrderType     string  `json:"orderType,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Quantity      int64   `json:"quantity,omitempty"`
	MarketID      string  `json:"marketID,omitempty"`
	OrigOrderID   string  `json:"origOrderId,omitempty"`
	Account       string  `json:"account,omitempty"`
	Channel       string  `json:"channel,omitempty"`
	InputTime     string  `json:"inputTime,omitempty"`
	ModifiedTime  string  `json:"modifiedTime,omitempty"`
	IsForceSell   string  `json:"isForceSell,omitempty"`
	IsShortSell   *string `json:"isShortSell,omitempty"`
	OrigRequestID string  `json:"origRequestID,omitempty"`
	StopOrder     bool    `json:"stopOrder,omitempty"`
	StopPrice     float64 `json:"stopPrice,omitempty"`
	StopType      string  `json:"stopType,omitempty"`
	StopStep      float64 `json:"stopStep,omitempty"`
	ProfitPrice   float64 `json:"profitPrice,omitempty"`
	Modifiable    bool    `json:"modifiable,omitempty"`
	Note          string  `json:"note,omitempty"`
}

type OrderMatchEvent struct {
	OrderID       string  `json:"orderID,omitempty"`
	NotifyID      int64   `json:"notifyID,omitempty"`
	InstrumentID  string  `json:"instrumentID,omitempty"`
	UniqueID      string  `json:"uniqueID,omitempty"`
	BuySell       string  `json:"buySell,omitempty"`
	MatchPrice    float64 `json:"matchPrice,omitempty"`
	IPAddress     string  `json:"ipAddress,omitempty"`
	MatchQuantity int64   `json:"matchQty,omitempty"`
	Prefix        string  `json:"prefix,omitempty"`
	Account       string  `json:"account,omitempty"`
	MatchTime     string  `json:"matchTime,omitempty"`
}

type ClientPortfolioEvent struct {
	Account               string            `json:"account,omitempty"`
	NotifyID              int64             `json:"notifyID,omitempty"`
	Data                  any               `json:"data,omitempty"`
	ClientPortfoliosOpen  []ClientPortfolio `json:"clientPortfoliosOpen"`
	UniqueID              *string           `json:"uniqueID,omitempty"`
	ClientPortfoliosClose []ClientPortfolio `json:"clientPortfoliosClose"`
	ConnectionID          string            `json:"connectionID,omitempty"`
	IPAddress             *string           `json:"ipAddress,omitempty"`
	Prefix                *string           `json:"prefix,omitempty"`
}

type ClientPortfolio struct {
	MarketID        string  `json:"martketID,omitempty"`
	InstrumentID    string  `json:"instrumentID,omitempty"`
	LongQuantity    int64   `json:"longQty,omitempty"`
	ShortQuantity   int64   `json:"shortQty,omitempty"`
	Net             int64   `json:"net,omitempty"`
	BidAveragePrice float64 `json:"bidAvgPrice,omitempty"`
	AskAveragePrice float64 `json:"askAvgPrice,omitempty"`
	TradePrice      float64 `json:"tradePrice,omitempty"`
	MarketPrice     float64 `json:"marketPrice,omitempty"`
	FloatingPL      float64 `json:"floatingPL,omitempty"`
	TradingPL       float64 `json:"tradingPL,omitempty"`
}

type FCOEvent struct {
	FCOID           string `json:"fcoId,omitempty"`
	NotifyID        int64  `json:"notifyID,omitempty"`
	Data            any    `json:"data,omitempty"`
	ProcessStatus   string `json:"processStatus,omitempty"`
	LastAction      string `json:"lastAction,omitempty"`
	UniqueID        string `json:"uniqueID,omitempty"`
	MatchedQuantity int64  `json:"matchedQuantity,omitempty"`
	IsPlaceOrder    bool   `json:"isPlaceOrder,omitempty"`
	IPAddress       string `json:"ipAddress,omitempty"`
	Symbol          string `json:"instrumentID,omitempty"`
	Prefix          string `json:"prefix,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	BrokerID        string `json:"brokerId,omitempty"`
	Price           string `json:"price,omitempty"`
	AccountID       string `json:"account,omitempty"`
	BrokerIDUpdate  string `json:"brokerIdUpdate,omitempty"`
	UpdatedTime     string `json:"updatedTime,omitempty"`
	Status          string `json:"status,omitempty"`
	Message         string `json:"message,omitempty"`
	Username        string `json:"username,omitempty"`
}
