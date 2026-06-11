package dto

// Authentication requests.

type AccessTokenRequest struct {
	ConsumerID     string `json:"consumerID,omitempty"`
	ConsumerSecret string `json:"consumerSecret,omitempty"`
	Code           string `json:"code,omitempty"`
	TwoFactorType  int    `json:"twoFactorType,omitempty"`
	IsSave         bool   `json:"isSave,omitempty"`
}

type GetOTPRequest struct {
	ConsumerID     string `json:"consumerID,omitempty"`
	ConsumerSecret string `json:"consumerSecret,omitempty"`
}

// Query requests.

type AccountRequest struct {
	Account string `json:"account,omitempty"`
}

type AuditOrderBookRequest = AccountRequest
type OrderBookRequest = AccountRequest
type PpmmrAccountRequest = AccountRequest
type CashAccountBalanceRequest = AccountRequest
type StockPositionRequest = AccountRequest
type DerivativeAccountBalanceRequest = AccountRequest
type CashInAdvanceAmountRequest = AccountRequest
type TransferableRequest = AccountRequest
type DividendRequest = AccountRequest
type ExercisableQuantityRequest = AccountRequest

type OrderHistoryRequest struct {
	Account   string `json:"account,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
}

type DerivativePositionRequest struct {
	Account      string `json:"account,omitempty"`
	QuerySummary bool   `json:"querySummary,omitempty"`
}

type MaxBuyQtyRequest struct {
	Account      string  `json:"account,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

type MaxSellQtyRequest struct {
	Account      string  `json:"account,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

type RateLimitRequest struct{}

type UnsettleSoldTransactionRequest struct {
	Account    string `json:"account,omitempty"`
	SettleDate string `json:"settleDate,omitempty"`
}

type CashTransferHistoriesRequest struct {
	Account  string `json:"account,omitempty"`
	FromDate string `json:"fromDate,omitempty"`
	ToDate   string `json:"toDate,omitempty"`
}

type CashInAdvanceHistoriesRequest = CashTransferHistoriesRequest

type EstCashInAdvanceFeeRequest struct {
	Account       string  `json:"account,omitempty"`
	CIAAmount     float64 `json:"ciaAmount,omitempty"`
	ReceiveAmount float64 `json:"receiveAmount,omitempty"`
}

type StockTransferHistoriesRequest struct {
	Account   string `json:"account,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type RightsHistoriesRequest = StockTransferHistoriesRequest

type FCOOrderBookRequest struct {
	FCOID     string `json:"fcoId,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type FCOStatusHistoryRequest = FCOOrderBookRequest

type FCOListRequest struct {
	FCOID         string `json:"fcoId,omitempty"`
	Account       string `json:"account,omitempty"`
	ProcessStatus string `json:"processStatus,omitempty"`
	Type          string `json:"type,omitempty"`
	InstrumentID  string `json:"instrumentID,omitempty"`
	Side          string `json:"side,omitempty"`
	FromDate      string `json:"fromDate,omitempty"`
	ToDate        string `json:"toDate,omitempty"`
	PageIndex     int    `json:"pageIndex,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
}

// ──────────────────────────────── Cash ────────────────────────────────

type CashInAdvanceAmount struct {
	Account    string              `json:"account,omitempty"`
	CIAAmounts []CashInAdvanceItem `json:"ciaAmounts,omitempty"`
}

type CashInAdvanceItem struct {
	DueDate      string  `json:"dueDate,omitempty"`
	SellValue    float64 `json:"sellValue,omitempty"`
	NetSellValue float64 `json:"netSellValue,omitempty"`
	Advance      float64 `json:"advance,omitempty"`
	CashAdvance  float64 `json:"cashAdvance,omitempty"`
}

type UnsettledSoldTransactions struct {
	Account                   string                     `json:"account,omitempty"`
	UnsettledSoldTransactions []UnsettledSoldTransaction `json:"unsettledSoldTransactions,omitempty"`
}

type UnsettledSoldTransaction struct {
	TradeDate    string  `json:"tradeDate,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	NetSellValue float64 `json:"netSellValue,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Price        float64 `json:"price,omitempty"`
}

type CashTransferHistories struct {
	TransferHistories []CashTransferHistory `json:"transferHistories,omitempty"`
}

type CashTransferHistory struct {
	TransactionID      string  `json:"transactionID,omitempty"`
	Date               string  `json:"date,omitempty"`
	Account            string  `json:"account,omitempty"`
	BeneficiaryAccount string  `json:"beneficiaryAccount,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	BankName           string  `json:"bankName,omitempty"`
	BankBranchName     string  `json:"bankBranchName,omitempty"`
	Beneficiary        string  `json:"beneficiary,omitempty"`
	Remark             string  `json:"remark,omitempty"`
	Type               string  `json:"type,omitempty"`
	Status             string  `json:"status,omitempty"`
}

type CashInAdvanceHistories struct {
	Account      string                 `json:"account,omitempty"`
	CIAHistories []CashInAdvanceHistory `json:"ciaHistories,omitempty"`
}

type CashInAdvanceHistory struct {
	TransactionID string                       `json:"transactionID,omitempty"`
	Date          string                       `json:"date,omitempty"`
	DateTime      string                       `json:"dateTime,omitempty"`
	TotalAmount   float64                      `json:"totalAmount,omitempty"`
	Details       []CashInAdvanceHistoryDetail `json:"details,omitempty"`
	Status        string                       `json:"status,omitempty"`
}

type CashInAdvanceHistoryDetail struct {
	Type       string  `json:"type,omitempty"`
	Value      float64 `json:"value,omitempty"`
	SettleDate *string `json:"settleDate,omitempty"`
}

type EstimateCashInAdvanceFee struct {
	Account       string  `json:"account,omitempty"`
	CIAAmount     float64 `json:"ciaAmount,omitempty"`
	ReceiveAmount float64 `json:"receiveAmount,omitempty"`
	Fee           float64 `json:"fee,omitempty"`
}

type Transaction struct {
	Account       string `json:"account,omitempty"`
	TransactionID string `json:"transactionID,omitempty"`
}

type VSDCashDWRequest struct {
	Account string  `json:"account,omitempty"`
	Amount  float64 `json:"amount,omitempty"`
	Type    string  `json:"type,omitempty"`
	Remark  string  `json:"remark,omitempty"`
	Code    string  `json:"code,omitempty"`
}

type TransferInternalRequest struct {
	Account            string  `json:"account,omitempty"`
	BeneficiaryAccount string  `json:"beneficiaryAccount,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	Remark             string  `json:"remark,omitempty"`
	Code               string  `json:"code,omitempty"`
}

type CreateCashInAdvanceRequest struct {
	Account       string  `json:"account,omitempty"`
	CIAAmount     float64 `json:"ciaAmount,omitempty"`
	ReceiveAmount float64 `json:"receiveAmount,omitempty"`
	Code          int     `json:"code,omitempty"`
}

// ──────────────────────────────── Core Trading ────────────────────────────────

type CashAccountBalance struct {
	Account             string  `json:"account,omitempty"`
	CashBalance         float64 `json:"cashBal,omitempty"`
	CashOnHold          float64 `json:"cashOnHold,omitempty"`
	SecureAmount        float64 `json:"secureAmount,omitempty"`
	Withdrawable        float64 `json:"withdrawable,omitempty"`
	ReceivingCashT1     float64 `json:"receivingCashT1,omitempty"`
	ReceivingCashT2     float64 `json:"receivingCashT2,omitempty"`
	MatchedBuyVolume    float64 `json:"matchedBuyVolume,omitempty"`
	MatchedSellVolume   float64 `json:"matchedSellVolume,omitempty"`
	UnMatchedBuyVolume  float64 `json:"unMatchedBuyVolume,omitempty"`
	UnMatchedSellVolume float64 `json:"unMatchedSellVolume,omitempty"`
	PaidCashT1          float64 `json:"paidCashT1,omitempty"`
	PaidCashT2          float64 `json:"paidCashT2,omitempty"`
	CIA                 float64 `json:"cia,omitempty"`
	Debt                float64 `json:"debt,omitempty"`
	PurchasingPower     float64 `json:"purchasingPower,omitempty"`
	TotalAssets         float64 `json:"totalAssets,omitempty"`
}

type DerivativeInternalAssets struct {
	Cash             float64 `json:"cash,omitempty"`
	ValidNonCash     float64 `json:"validNonCash,omitempty"`
	TotalValue       float64 `json:"totalValue,omitempty"`
	MaxValidNonCash  float64 `json:"maxValidNonCash,omitempty"`
	CashWithdrawable float64 `json:"cashWithdrawable,omitempty"`
	EE               float64 `json:"ee,omitempty"`
}

type DerivativeExchangeAssets struct {
	Cash             float64 `json:"cash,omitempty"`
	ValidNonCash     float64 `json:"validNonCash,omitempty"`
	TotalValue       float64 `json:"totalValue,omitempty"`
	MaxValidNonCash  float64 `json:"maxValidNonCash,omitempty"`
	CashWithdrawable float64 `json:"cashWithdrawable,omitempty"`
	EE               float64 `json:"ee,omitempty"`
}

type DerivativeInternalMargin struct {
	InitialMargin          float64 `json:"initialMargin,omitempty"`
	DeliveryMargin         float64 `json:"deliveryMargin,omitempty"`
	MarginReq              float64 `json:"marginReq,omitempty"`
	AccountRatio           float64 `json:"accountRatio,omitempty"`
	UsedLimitWarningLevel1 float64 `json:"usedLimitWarningLevel1,omitempty"`
	UsedLimitWarningLevel2 float64 `json:"usedLimitWarningLevel2,omitempty"`
	UsedLimitWarningLevel3 float64 `json:"usedLimitWarningLevel3,omitempty"`
	MarginCall             float64 `json:"marginCall,omitempty"`
}

type DerivativeExchangeMargin struct {
	MarginReq              float64 `json:"marginReq,omitempty"`
	AccountRatio           float64 `json:"accountRatio,omitempty"`
	UsedLimitWarningLevel1 float64 `json:"usedLimitWarningLevel1,omitempty"`
	UsedLimitWarningLevel2 float64 `json:"usedLimitWarningLevel2,omitempty"`
	UsedLimitWarningLevel3 float64 `json:"usedLimitWarningLevel3,omitempty"`
	MarginCall             float64 `json:"marginCall,omitempty"`
}

type DerivativeAccountBalance struct {
	Account               string                    `json:"account,omitempty"`
	AccountBalance        float64                   `json:"accountBalance,omitempty"`
	Fee                   float64                   `json:"fee,omitempty"`
	Commission            float64                   `json:"commission,omitempty"`
	Interest              float64                   `json:"interest,omitempty"`
	Loan                  float64                   `json:"loan,omitempty"`
	DeliveryAmount        float64                   `json:"deliveryAmount,omitempty"`
	FloatingPL            float64                   `json:"floatingPL,omitempty"`
	TotalPL               float64                   `json:"totalPL,omitempty"`
	Marginable            float64                   `json:"marginable,omitempty"`
	Depositable           float64                   `json:"depositable,omitempty"`
	RCCall                float64                   `json:"rcCall,omitempty"`
	Withdrawable          float64                   `json:"withdrawable,omitempty"`
	NonCashDrawableRCCall float64                   `json:"nonCashDrawableRCCall,omitempty"`
	InternalAssets        *DerivativeInternalAssets `json:"internalAssets,omitempty"`
	ExchangeAssets        *DerivativeExchangeAssets `json:"exchangeAssets,omitempty"`
	InternalMargin        *DerivativeInternalMargin `json:"internalMargin,omitempty"`
	ExchangeMargin        *DerivativeExchangeMargin `json:"exchangeMargin,omitempty"`
	NAV                   float64                   `json:"nav,omitempty"`
	OrigMarginRatio       float64                   `json:"origMarginRatio,omitempty"`
}

type MaxBuyQuantity struct {
	Account         string  `json:"account,omitempty"`
	MaxBuyQty       int     `json:"maxBuyQty,omitempty"`
	MarginRatio     string  `json:"marginRatio,omitempty"`
	PurchasingPower float64 `json:"purchasingPower,omitempty"`
	OrigMarginRatio string  `json:"origMarginRatio,omitempty"`
}

type MaxSellQuantity struct {
	Account    string `json:"account,omitempty"`
	MaxSellQty int    `json:"maxSellQty,omitempty"`
}

type Order struct {
	UniqueID       *string `json:"uniqueID,omitempty"`
	OrderID        string  `json:"orderID,omitempty"`
	BuySell        string  `json:"buySell,omitempty"`
	Price          float64 `json:"price,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
	FilledQty      int     `json:"filledQty,omitempty"`
	OrderStatus    string  `json:"orderStatus,omitempty"`
	MarketID       string  `json:"marketID,omitempty"`
	InputTime      string  `json:"inputTime,omitempty"`
	ModifiedTime   string  `json:"modifiedTime,omitempty"`
	InstrumentID   string  `json:"instrumentID,omitempty"`
	OrderType      string  `json:"orderType,omitempty"`
	CancelQty      int     `json:"cancelQty,omitempty"`
	AveragePrice   float64 `json:"avgPrice,omitempty"`
	IsForceSell    *string `json:"isForcesell,omitempty"`
	IsShortSell    *string `json:"isShortsell,omitempty"`
	RejectReason   string  `json:"rejectReason,omitempty"`
	LastErrorEvent any     `json:"lastErrorEvent,omitempty"`
}

type OrderBook struct {
	Account string  `json:"account,omitempty"`
	Orders  []Order `json:"orders,omitempty"`
}

type OrderHistory struct {
	OrderHistories []Order `json:"orderHistories,omitempty"`
	Account        string  `json:"account,omitempty"`
}

type OrderRequest struct {
	Account      string  `json:"account,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	Market       string  `json:"market,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	OrderID      string  `json:"orderID,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	StopOrder    bool    `json:"stopOrder,omitempty"`
	StopPrice    float64 `json:"stopPrice,omitempty"`
	StopType     string  `json:"stopType,omitempty"`
	StopStep     float64 `json:"stopStep,omitempty"`
	LossStep     float64 `json:"lossStep,omitempty"`
	ProfitStep   float64 `json:"profitStep,omitempty"`
	ForceSell    bool    `json:"forceSell,omitempty"`
	Modifiable   bool    `json:"modifiable,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type StockPortfolio struct {
	Account          string          `json:"account,omitempty"`
	TotalMarketValue float64         `json:"totalMarketValue,omitempty"`
	StockPositions   []StockPosition `json:"stockPositions,omitempty"`
}

type StockPosition struct {
	MarketID     string  `json:"marketID,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	OnHand       int     `json:"onHand,omitempty"`
	Block        int     `json:"block,omitempty"`
	Bonus        int     `json:"bonus,omitempty"`
	BuyT0        int     `json:"buyT0,omitempty"`
	BuyT1        int     `json:"buyT1,omitempty"`
	BuyT2        int     `json:"buyT2,omitempty"`
	SellT0       int     `json:"sellT0,omitempty"`
	SellT1       int     `json:"sellT1,omitempty"`
	SellT2       int     `json:"sellT2,omitempty"`
	AveragePrice float64 `json:"avgPrice,omitempty"`
	Mortgage     int     `json:"mortgage,omitempty"`
	SellableQty  int     `json:"sellableQty,omitempty"`
	HoldForTrade int     `json:"holdForTrade,omitempty"`
	MarketPrice  float64 `json:"marketPrice,omitempty"`
	ExchangeID   string  `json:"exchangeID,omitempty"`
	SellingQty   string  `json:"sellingQty,omitempty"`
	BuyingQty    string  `json:"buyingQty,omitempty"`
}

type DerivativePosition struct {
	MarketID     string  `json:"marketID,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	LongQty      int     `json:"longQty,omitempty"`
	ShortQty     int     `json:"shortQty,omitempty"`
	Net          int     `json:"net,omitempty"`
	BidAvgPrice  float64 `json:"bidAvgPrice,omitempty"`
	AskAvgPrice  float64 `json:"askAvgPrice,omitempty"`
	TradePrice   float64 `json:"tradePrice,omitempty"`
	MarketPrice  float64 `json:"marketPrice,omitempty"`
	FloatingPL   float64 `json:"floatingPL,omitempty"`
	TradingPL    float64 `json:"tradingPL,omitempty"`
}

type DerivativePositions struct {
	Account        string               `json:"account,omitempty"`
	OpenPositions  []DerivativePosition `json:"openPosition,omitempty"`
	ClosePositions []DerivativePosition `json:"closePosition,omitempty"`
}

type AccountAsset struct {
	Account           string  `json:"account,omitempty"`
	CollateralAsset   float64 `json:"collateralAsset,omitempty"`
	CallLMW           float64 `json:"callLMW,omitempty"`
	Liability         float64 `json:"liability,omitempty"`
	EEOrigin          float64 `json:"eeOrigin,omitempty"`
	ForceLMV          float64 `json:"forceLMV,omitempty"`
	Equity            float64 `json:"equity,omitempty"`
	EE                float64 `json:"ee,omitempty"`
	CallMargin        float64 `json:"callMargin,omitempty"`
	CashBalance       float64 `json:"cashBalance,omitempty"`
	PurchasingPower   float64 `json:"purchasingPower,omitempty"`
	CallForceSell     float64 `json:"callForceSell,omitempty"`
	LMV               float64 `json:"lmv,omitempty"`
	MarginCall        float64 `json:"marginCall,omitempty"`
	Withdrawal        float64 `json:"withdrawal,omitempty"`
	CollateralA       float64 `json:"collateralA,omitempty"`
	Action            string  `json:"action,omitempty"`
	MarginRatio       float64 `json:"marginRatio,omitempty"`
	Debt              float64 `json:"debt,omitempty"`
	AccruedInterest   float64 `json:"accruedInterest,omitempty"`
	HoldRight         float64 `json:"holdRight,omitempty"`
	PreLoan           float64 `json:"preLoan,omitempty"`
	Fees              float64 `json:"fees,omitempty"`
	BuyUnmatch        float64 `json:"buyUnmatch,omitempty"`
	AP                float64 `json:"ap,omitempty"`
	APT1              float64 `json:"apT1,omitempty"`
	SellUnmatch       float64 `json:"sellUnmatch,omitempty"`
	CIA               float64 `json:"cia,omitempty"`
	AR                float64 `json:"ar,omitempty"`
	ART1              float64 `json:"arT1,omitempty"`
	PPCredit          float64 `json:"ppCredit,omitempty"`
	CreditLimit       float64 `json:"creditLimit,omitempty"`
	TotalAsset        float64 `json:"totalAsset,omitempty"`
	TotalAssets       float64 `json:"totalAssets,omitempty"`
	MarginCallLMVSold float64 `json:"marginCallLMVSold,omitempty"`
	LMVNonMarginable  float64 `json:"lmvNonMarginable,omitempty"`
	EECredit          float64 `json:"eeCredit,omitempty"`
	TotalEquity       float64 `json:"totalEquity,omitempty"`
	EE90              float64 `json:"eE90,omitempty"`
	EE80              float64 `json:"eE80,omitempty"`
	EE70              float64 `json:"eE70,omitempty"`
	EE60              float64 `json:"eE60,omitempty"`
	EE50              float64 `json:"eE50,omitempty"`
	Dividend          float64 `json:"dividend,omitempty"`
	MaintenanceRatio  float64 `json:"maintenanceRatio,omitempty"`
	WarningRatio      float64 `json:"warningRatio,omitempty"`
}

type APILimit struct {
	Endpoint string `json:"endpoint,omitempty"`
	Period   string `json:"period,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

// ──────────────────────────────── Order Input Types ────────────────────────────────

type NewOrderRequest struct {
	InstrumentID string  `json:"instrumentID,omitempty"`
	Market       string  `json:"market,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Account      string  `json:"account,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	StopOrder    bool    `json:"stopOrder,omitempty"`
	StopPrice    float64 `json:"stopPrice,omitempty"`
	StopType     string  `json:"stopType,omitempty"`
	StopStep     float64 `json:"stopStep,omitempty"`
	LossStep     float64 `json:"lossStep,omitempty"`
	ProfitStep   float64 `json:"profitStep,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	ForceSell    bool    `json:"forceSell,omitempty"`
	Modifiable   bool    `json:"modifiable,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type CancelOrderRequest struct {
	OrderID      string  `json:"orderID,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	Market       string  `json:"market,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Account      string  `json:"account,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type ModifyOrderRequest struct {
	OrderID      string  `json:"orderID,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	Market       string  `json:"market,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Account      string  `json:"account,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type DerNewOrderRequest struct {
	InstrumentID string  `json:"instrumentID,omitempty"`
	Market       string  `json:"market,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Account      string  `json:"account,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	StopOrder    bool    `json:"stopOrder,omitempty"`
	StopPrice    float64 `json:"stopPrice,omitempty"`
	StopType     string  `json:"stopType,omitempty"`
	StopStep     float64 `json:"stopStep,omitempty"`
	LossStep     float64 `json:"lossStep,omitempty"`
	ProfitStep   float64 `json:"profitStep,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	ForceSell    bool    `json:"forceSell,omitempty"`
	Modifiable   bool    `json:"modifiable,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type DerModifyOrderRequest struct {
	OrderID      string  `json:"orderID,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	Market       string  `json:"market,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Account      string  `json:"account,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

type DerCancelOrderRequest struct {
	OrderID      string  `json:"orderID,omitempty"`
	Account      string  `json:"account,omitempty"`
	MarketID     string  `json:"marketID,omitempty"`
	Market       string  `json:"market,omitempty"`
	InstrumentID string  `json:"instrumentID,omitempty"`
	BuySell      string  `json:"buySell,omitempty"`
	OrderType    string  `json:"orderType,omitempty"`
	ChannelID    string  `json:"channelID,omitempty"`
	Price        float64 `json:"price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	RequestID    string  `json:"requestID,omitempty"`
	Code         string  `json:"code,omitempty"`
	DeviceID     string  `json:"deviceId,omitempty"`
	UserAgent    string  `json:"userAgent,omitempty"`
	Note         string  `json:"note,omitempty"`
	BrokerID     string  `json:"brokerID,omitempty"`
}

// ──────────────────────────────── Stock Transfer ────────────────────────────────

type TransferableStockAccount struct {
	Account            string              `json:"account,omitempty"`
	TransferableStocks []TransferableStock `json:"transferableStocks,omitempty"`
}

type TransferableStock struct {
	InstrumentID   string `json:"instrumentID,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
	InstrumentType string `json:"instrumentType,omitempty"`
}

type StockTransferHistoryAccount struct {
	Account                string                 `json:"account,omitempty"`
	StockTransferHistories []StockTransferHistory `json:"stockTransferHistories,omitempty"`
}

type StockTransferHistory struct {
	TransactionID      string `json:"transactionID,omitempty"`
	BeneficiaryAccount string `json:"beneficiaryAccount,omitempty"`
	InstrumentID       string `json:"instrumentID,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	DateTime           string `json:"dateTime,omitempty"`
	Status             string `json:"status,omitempty"`
	Remark             string `json:"remark,omitempty"`
	AuditRemark        string `json:"auditRemark,omitempty"`
}

type StockTransferRequest struct {
	Account            string `json:"account,omitempty"`
	BeneficiaryAccount string `json:"beneficiaryAccount,omitempty"`
	ExchangeID         string `json:"exchangeID,omitempty"`
	InstrumentID       string `json:"instrumentID,omitempty"`
	Quantity           int    `json:"quantity,omitempty"`
	Code               string `json:"code,omitempty"`
}

// ──────────────────────────────── Rights ────────────────────────────────

type Dividends struct {
	Account   string     `json:"account,omitempty"`
	Dividends []Dividend `json:"dividends,omitempty"`
}

type Dividend struct {
	StockDividend          string  `json:"stockDividend,omitempty"`
	InstrumentID           string  `json:"instrumentID,omitempty"`
	Quantity               int     `json:"quantity,omitempty"`
	ExecutedRate           string  `json:"executedRate,omitempty"`
	CloseDate              string  `json:"closeDate,omitempty"`
	PaidDate               string  `json:"paidDate,omitempty"`
	Amount                 float64 `json:"amount,omitempty"`
	Status                 string  `json:"status,omitempty"`
	ReceivedQuantity       int     `json:"receivedQuantity,omitempty"`
	IssueInstrument        string  `json:"issueInstrument,omitempty"`
	DistributedFlag        string  `json:"distributedFlag,omitempty"`
	PayableDate            *string `json:"payableDate,omitempty"`
	SubscriptionPrice      float64 `json:"subscriptionPrice,omitempty"`
	SubscriptionAmount     float64 `json:"subscriptionAmount,omitempty"`
	SubscriptionQuantity   int     `json:"subscriptionQuantity,omitempty"`
	SubscriptionPeriodFrom *string `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo   *string `json:"subscriptionPeriodTo,omitempty"`
	EntitlementID          *string `json:"entitlementID,omitempty"`
	ExchangeID             string  `json:"exchangeID,omitempty"`
}

type ExercisableQuantities struct {
	Account               string                `json:"account,omitempty"`
	ExercisableQuantities []ExercisableQuantity `json:"exercisableQuantities,omitempty"`
}

type ExercisableQuantity struct {
	EntitlementID               string  `json:"entitlementID,omitempty"`
	InstrumentID                string  `json:"instrumentID,omitempty"`
	SubscriptionPrice           float64 `json:"subscriptionPrice,omitempty"`
	ExecutedRateFrom            float64 `json:"executedRateFrom,omitempty"`
	SubscriptionPeriodFrom      string  `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo        string  `json:"subscriptionPeriodTo,omitempty"`
	ExerciseableQuantity        int     `json:"exerciseableQuantity,omitempty"`
	ExerciseableReceiveQuantity int     `json:"exerciseableReceiveQuantity,omitempty"`
	ExercisedReceiveQuantity    int     `json:"exercisedReceiveQuantity,omitempty"`
	ExecutedRateTo              float64 `json:"executedRateTo,omitempty"`
	ExercisedQuantity           int     `json:"exercisedQuantity,omitempty"`
	PayableDate                 *string `json:"payableDate,omitempty"`
}

type RightsHistories struct {
	Account                          string          `json:"account,omitempty"`
	OnlineRightSubscriptionHistories []RightsHistory `json:"onlineRightSubscriptionHistories,omitempty"`
}

type RightsHistory struct {
	TransactionID             string  `json:"transactionID,omitempty"`
	DateTime                  string  `json:"dateTime,omitempty"`
	InstrumentID              string  `json:"instrumentID,omitempty"`
	RatioFrom                 float64 `json:"ratioFrom,omitempty"`
	SubscriptionPrice         float64 `json:"subscriptionPrice,omitempty"`
	SubscriptionPeriodFrom    string  `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo      string  `json:"subscriptionPeriodTo,omitempty"`
	ExercisedReceivedQuantity int     `json:"exercisedReceivedQty,omitempty"`
	Amount                    float64 `json:"amount,omitempty"`
	Status                    string  `json:"status,omitempty"`
	RatioTo                   float64 `json:"ratioTo,omitempty"`
	UnderlyingInstrumentID    string  `json:"underlyingInstrumentID,omitempty"`
}

type RightsCreateRequest struct {
	Account       string  `json:"account,omitempty"`
	InstrumentID  string  `json:"instrumentID,omitempty"`
	EntitlementID string  `json:"entitlementID,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Code          string  `json:"code,omitempty"`
}

// ──────────────────────────────── Conditional Orders (FCO) ────────────────────────────────

type ConditionalOrderResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	FCOID   string `json:"fcoid,omitempty"`
	DCOID   string `json:"dcoid,omitempty"`
}

type ConditionalOrderPage[T any] struct {
	Message       string  `json:"message,omitempty"`
	Status        int     `json:"status,omitempty"`
	Data          []T     `json:"data,omitempty"`
	FCOList       []T     `json:"fcoList,omitempty"`
	PageSize      int     `json:"pageSize,omitempty"`
	PageNumber    int     `json:"pageNumber,omitempty"`
	ItemsCount    int     `json:"itemsCount,omitempty"`
	PageCount     float64 `json:"pageCount,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	Page          int     `json:"page,omitempty"`
	Size          int     `json:"size,omitempty"`
}

type ConditionalTriggeredOrder struct {
	FCOID           string  `json:"fcoId,omitempty"`
	Account         string  `json:"account,omitempty"`
	AccountNo       string  `json:"accountNo,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	Price           string  `json:"price,omitempty"`
	InstrumentID    string  `json:"instrumentId,omitempty"`
	Symbol          string  `json:"symbol,omitempty"`
	Side            string  `json:"side,omitempty"`
	OrderType       string  `json:"orderType,omitempty"`
	RunningMode     bool    `json:"runningMode,omitempty"`
	MainOrder       bool    `json:"mainOrder,omitempty"`
	AttachedOrder   bool    `json:"attachedOrder,omitempty"`
	IsMainOrder     bool    `json:"isMainOrder,omitempty"`
	IsAttachedOrder bool    `json:"isAttachedOrder,omitempty"`
	CreatedTime     string  `json:"createdTime,omitempty"`
	UpdatedTime     string  `json:"updatedTime,omitempty"`
	UniqueID        *string `json:"uniqueId,omitempty"`
	OrderID         string  `json:"orderId,omitempty"`
	MatchedQuantity int     `json:"matchedQuantity,omitempty"`
	FilledQuantity  int     `json:"filledQuantity,omitempty"`
	OSQuantity      int     `json:"osQuantity,omitempty"`
	AveragePrice    float64 `json:"avgPrice,omitempty"`
	Status          string  `json:"status,omitempty"`
	Detail          string  `json:"detail,omitempty"`
}

type ConditionalOrderStatus struct {
	State  string `json:"state,omitempty"`
	Time   string `json:"time,omitempty"`
	Code   int    `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type ConditionalOrderParameters struct {
	StopPrice      float64 `json:"stopPrice,omitempty"`
	Side           string  `json:"side,omitempty"`
	ActivePrice    float64 `json:"activePrice,omitempty"`
	TrailingAmount float64 `json:"trailingAmount,omitempty"`
	TpActivePrice  float64 `json:"tpActivePrice,omitempty"`
	SlActivePrice  float64 `json:"slActivePrice,omitempty"`
	TpPrice        float64 `json:"tpPrice,omitempty"`
	SlPrice        float64 `json:"slPrice,omitempty"`
	TpSlip         float64 `json:"tpSlip,omitempty"`
	SlSlip         float64 `json:"slSlip,omitempty"`
	Operator       string  `json:"operator,omitempty"`
}

type ConditionalOrder struct {
	FCOID           string                      `json:"fcoId,omitempty"`
	DCOID           string                      `json:"dcoId,omitempty"`
	Username        string                      `json:"username,omitempty"`
	UserID          string                      `json:"userid,omitempty"`
	Account         string                      `json:"account,omitempty"`
	AccountNo       string                      `json:"accountNo,omitempty"`
	Quantity        int                         `json:"quantity,omitempty"`
	Price           string                      `json:"price,omitempty"`
	PriceSlip       float64                     `json:"priceSlip,omitempty"`
	InstrumentID    string                      `json:"instrumentID,omitempty"`
	Side            string                      `json:"side,omitempty"`
	Type            string                      `json:"type,omitempty"`
	ProcessStatus   string                      `json:"processStatus,omitempty"`
	OrderType       *string                     `json:"orderType,omitempty"`
	RunningMode     bool                        `json:"runningMode,omitempty"`
	MainOrder       bool                        `json:"mainOrder,omitempty"`
	AttachedOrder   bool                        `json:"attachedOrder,omitempty"`
	CreatedDate     string                      `json:"createdDate,omitempty"`
	CreatedTime     *string                     `json:"createdTime,omitempty"`
	UpdatedTime     *string                     `json:"updatedTime,omitempty"`
	MatchedQuantity int                         `json:"matchedQuantity,omitempty"`
	PlaceOrder      bool                        `json:"placeOrder,omitempty"`
	IsPlaceOrder    bool                        `json:"isPlaceOrder,omitempty"`
	Status          *string                     `json:"status,omitempty"`
	StatusDetail    string                      `json:"statusDetail,omitempty"`
	Detail          *string                     `json:"detail,omitempty"`
	Params          *ConditionalOrderParameters `json:"params,omitempty"`
	OrderID         string                      `json:"orderID,omitempty"`
	OrderStatus     string                      `json:"orderStatus,omitempty"`
}

type FCONewOrderRequest struct {
	InstrumentID   string  `json:"instrumentID,omitempty"`
	Side           string  `json:"side,omitempty"`
	Type           string  `json:"type,omitempty"`
	Price          string  `json:"price,omitempty"`
	PriceSlip      float64 `json:"priceSlip,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
	Account        string  `json:"account,omitempty"`
	FromDate       string  `json:"fromDate,omitempty"`
	ToDate         string  `json:"toDate,omitempty"`
	StopPrice      float64 `json:"stopPrice,omitempty"`
	ActivePrice    float64 `json:"activePrice,omitempty"`
	TrailingAmount float64 `json:"trailingAmount,omitempty"`
	TpActivePrice  float64 `json:"tpActivePrice,omitempty"`
	SlActivePrice  float64 `json:"slActivePrice,omitempty"`
	TpPrice        string  `json:"tpPrice,omitempty"`
	SlPrice        string  `json:"slPrice,omitempty"`
	TpSlip         float64 `json:"tpSlip,omitempty"`
	SlSlip         float64 `json:"slSlip,omitempty"`
	Operator       string  `json:"operator,omitempty"`
	Code           string  `json:"code,omitempty"`
	UserAgent      string  `json:"userAgent,omitempty"`
	DeviceID       string  `json:"deviceId,omitempty"`
}

type FCOCancelOrderRequest struct {
	FCOID     string `json:"fcoId,omitempty"`
	Code      string `json:"code,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	DeviceID  string `json:"deviceId,omitempty"`
}

// ──────────────────────────────── Wrapper types ────────────────────────────────

type TradingResponse[T any] struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type DataResponse[T any] struct {
	Data        []T    `json:"data,omitempty"`
	Message     string `json:"message,omitempty"`
	Status      string `json:"status,omitempty"`
	TotalRecord int    `json:"totalRecord,omitempty"`
}

// Endpoint responses.

type AccessToken struct {
	AccessToken string `json:"accessToken,omitempty"`
}

type AccessTokenResponse struct {
	Message string       `json:"message,omitempty"`
	Status  int          `json:"status,omitempty"`
	Data    *AccessToken `json:"data,omitempty"`
}

type GetOTPResponse struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type AuditOrderBookResponse struct{ TradingResponse[OrderBook] }
type OrderBookResponse struct{ TradingResponse[OrderBook] }
type OrderHistoryResponse struct{ TradingResponse[OrderHistory] }
type PpmmrAccountResponse struct{ TradingResponse[AccountAsset] }
type CashAccountBalanceResponse struct {
	TradingResponse[CashAccountBalance]
}
type StockPositionResponse struct {
	TradingResponse[StockPortfolio]
}
type DerivativeAccountBalanceResponse struct {
	TradingResponse[DerivativeAccountBalance]
}
type DerivativePositionResponse struct {
	TradingResponse[DerivativePositions]
}
type MaxBuyQtyResponse struct {
	TradingResponse[MaxBuyQuantity]
}
type MaxSellQtyResponse struct {
	TradingResponse[MaxSellQuantity]
}
type RateLimitResponse struct{ TradingResponse[[]APILimit] }

type OrderSubmission[T any] struct {
	RequestID   string `json:"requestID,omitempty"`
	RequestData T      `json:"requestData,omitempty"`
}

type NewOrderResponse struct {
	TradingResponse[OrderSubmission[NewOrderRequest]]
}
type ModifyOrderResponse struct {
	TradingResponse[OrderSubmission[ModifyOrderRequest]]
}
type CancelOrderResponse struct {
	TradingResponse[OrderSubmission[CancelOrderRequest]]
}
type DerNewOrderResponse struct {
	TradingResponse[OrderSubmission[DerNewOrderRequest]]
}
type DerModifyOrderResponse struct {
	TradingResponse[OrderSubmission[DerModifyOrderRequest]]
}
type DerCancelOrderResponse struct {
	TradingResponse[OrderSubmission[DerCancelOrderRequest]]
}

type CashInAdvanceAmountResponse struct {
	TradingResponse[CashInAdvanceAmount]
}
type UnsettleSoldTransactionResponse struct {
	TradingResponse[UnsettledSoldTransactions]
}
type CashTransferHistoriesResponse struct {
	TradingResponse[CashTransferHistories]
}
type CashInAdvanceHistoriesResponse struct {
	TradingResponse[CashInAdvanceHistories]
}
type EstCashInAdvanceFeeResponse struct {
	TradingResponse[EstimateCashInAdvanceFee]
}
type VSDCashDWResponse struct{ TradingResponse[Transaction] }
type TransferInternalResponse struct{ TradingResponse[Transaction] }
type CreateCashInAdvanceResponse struct{ TradingResponse[Transaction] }

type TransferableResponse struct {
	TradingResponse[TransferableStockAccount]
}
type StockTransferHistoriesResponse struct {
	TradingResponse[StockTransferHistoryAccount]
}
type StockTransferResponse struct{ TradingResponse[Transaction] }

type DividendResponse struct{ TradingResponse[Dividends] }
type ExercisableQuantityResponse struct {
	TradingResponse[ExercisableQuantities]
}
type RightsHistoriesResponse struct {
	TradingResponse[RightsHistories]
}
type RightsCreateResponse struct{ TradingResponse[Transaction] }

type FCONewOrderResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	FCOID   string `json:"fcoid,omitempty"`
}

type FCOCancelOrderResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	DCOID   string `json:"dcoid,omitempty"`
}

type FCOOrderBookResponse struct {
	ConditionalOrderPage[ConditionalTriggeredOrder]
}
type FCOStatusHistoryResponse struct {
	ConditionalOrderPage[ConditionalOrderStatus]
}
type FCOListResponse struct {
	ConditionalOrderPage[ConditionalOrder]
}
