package dto

// ──────────────────────────────── Cash ────────────────────────────────

type CashInAdvanceAmountData struct {
	Account    string              `json:"account,omitempty"`
	CIAAmounts []CashInAdvanceItem `json:"ciaAmounts,omitempty"`
}

type CashInAdvanceItem struct {
	DueDate      string `json:"dueDate,omitempty"`
	SellValue    string `json:"sellValue,omitempty"`
	NetSellValue string `json:"netSellValue,omitempty"`
	Advance      string `json:"advance,omitempty"`
	CashAdvance  string `json:"cashAdvance,omitempty"`
}

type UnsettledSoldTransactionsData struct {
	Account                   string                     `json:"account,omitempty"`
	UnsettledSoldTransactions []UnsettledSoldTransaction `json:"unsettledSoldTransactions,omitempty"`
}

type UnsettledSoldTransaction struct {
	TradeDate    string `json:"tradeDate,omitempty"`
	InstrumentID string `json:"instrumentID,omitempty"`
	NetSellValue string `json:"netSellValue,omitempty"`
	Quantity     string `json:"quantity,omitempty"`
	Price        string `json:"price,omitempty"`
}

type TransferHistoriesData struct {
	TransferHistories []CashTransferHistory `json:"transferHistories,omitempty"`
}

type CashTransferHistory struct {
	TransactionID      string `json:"transactionID,omitempty"`
	Date               string `json:"date,omitempty"`
	Account            string `json:"account,omitempty"`
	BeneficiaryAccount string `json:"beneficiaryAccount,omitempty"`
	BankName           string `json:"bankName,omitempty"`
	BankBranchName     string `json:"bankBranchName,omitempty"`
	Beneficiary        string `json:"beneficiary,omitempty"`
	Type               string `json:"type,omitempty"`
	SenderAccount      string `json:"senderAccount,omitempty"`
	ReceiverAccount    string `json:"receiverAccount,omitempty"`
	Amount             string `json:"amount,omitempty"`
	DateTime           string `json:"dateTime,omitempty"`
	Status             string `json:"status,omitempty"`
	Remark             string `json:"remark,omitempty"`
}

type CashInAdvanceHistoriesData struct {
	Account      string                 `json:"account,omitempty"`
	CIAHistories []CashInAdvanceHistory `json:"ciaHistories,omitempty"`
}

type CashInAdvanceHistory struct {
	TransactionID string                       `json:"transactionID,omitempty"`
	DateTime      string                       `json:"dateTime,omitempty"`
	TotalAmount   string                       `json:"totalAmount,omitempty"`
	Details       []CashInAdvanceHistoryDetail `json:"details,omitempty"`
	Status        string                       `json:"status,omitempty"`
}

type CashInAdvanceHistoryDetail struct {
	Type       string `json:"type,omitempty"`
	Value      string `json:"value,omitempty"`
	SettleDate string `json:"settleDate,omitempty"`
}

type EstimateCashInAdvanceFeeData struct {
	Account       string `json:"account,omitempty"`
	CIAAmount     string `json:"ciaAmount,omitempty"`
	ReceiveAmount string `json:"receiveAmount,omitempty"`
	Fee           string `json:"fee,omitempty"`
}

type TransactionResponse struct {
	Account       string `json:"account,omitempty"`
	TransactionID string `json:"transactionID,omitempty"`
}

// ──────────────────────────────── Core Trading ────────────────────────────────

type StockAccountBalance struct {
	Account             string `json:"account,omitempty"`
	CashBalance         string `json:"cashbal,omitempty"`
	CashOnHold          string `json:"cashonhold,omitempty"`
	SecureAmount        string `json:"secureamount,omitempty"`
	Withdrawable        string `json:"withdrawable,omitempty"`
	ReceivingCashT1     string `json:"receivingcasht1,omitempty"`
	ReceivingCashT2     string `json:"receivingcasht2,omitempty"`
	MatchedBuyVolume    string `json:"matchedbuyvolume,omitempty"`
	MatchedSellVolume   string `json:"matchedsellvolume,omitempty"`
	UnmatchedBuyVolume  string `json:"unmatchedbuyvolume,omitempty"`
	UnmatchedSellVolume string `json:"unmatchedsellvolume,omitempty"`
	PaidCashT1          string `json:"paidcasht1,omitempty"`
	PaidCashT2          string `json:"paidcasht2,omitempty"`
	CIA                 string `json:"cia,omitempty"`
	Debt                string `json:"debt,omitempty"`
	PurchasingPower     string `json:"purchasingpower,omitempty"`
	TotalAsset          string `json:"totalasset,omitempty"`
}

type DerivativeAccountBalance struct {
	Account               string `json:"account,omitempty"`
	AccountBalance        string `json:"accountbalance,omitempty"`
	Fee                   string `json:"fee,omitempty"`
	Commission            string `json:"commission,omitempty"`
	Interest              string `json:"interest,omitempty"`
	Loan                  string `json:"loan,omitempty"`
	DeliveryAmount        string `json:"deliveryamount,omitempty"`
	FloatingPL            string `json:"floatingpl,omitempty"`
	TotalPL               string `json:"totalpl,omitempty"`
	Marginable            string `json:"marginable,omitempty"`
	Depositable           string `json:"depositable,omitempty"`
	RCCall                string `json:"rccall,omitempty"`
	Withdrawable          string `json:"withdrawable,omitempty"`
	NonCashDrawableRCCall string `json:"noncashdrawablerccall,omitempty"`
	InternalAssets        any    `json:"internalassets,omitempty"`
	ExchangeAssets        any    `json:"exchangeassets,omitempty"`
	InternalMargin        any    `json:"internalmargin,omitempty"`
	ExchangeMargin        any    `json:"exchangemargin,omitempty"`
	NAV                   string `json:"nav,omitempty"`
	OrigMarginRatio       string `json:"origMarginRatio,omitempty"`
}

type MaxBuyQuantityData struct {
	Account         string `json:"account,omitempty"`
	MaxBuyQty       string `json:"maxbuyqty,omitempty"`
	MarginRatio     string `json:"marginRatio,omitempty"`
	PurchasingPower string `json:"purchasingPower,omitempty"`
	OrigMarginRatio string `json:"origMarginRatio,omitempty"`
}

type MaxSellQuantityData struct {
	Account    string `json:"account,omitempty"`
	MaxSellQty string `json:"maxSellQty,omitempty"`
}

type OrderData struct {
	UniqueID     string `json:"uniqueID,omitempty"`
	OrderID      string `json:"orderID,omitempty"`
	BuySell      string `json:"buySell,omitempty"`
	Price        string `json:"price,omitempty"`
	Quantity     string `json:"quantity,omitempty"`
	FilledQty    string `json:"filledQty,omitempty"`
	OrderStatus  string `json:"orderStatus,omitempty"`
	MarketID     string `json:"marketID,omitempty"`
	InputTime    string `json:"inputTime,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
	InstrumentID string `json:"instrumentID,omitempty"`
	OrderType    string `json:"orderType,omitempty"`
	CancelQty    string `json:"cancelQty,omitempty"`
	AveragePrice string `json:"avgPrice,omitempty"`
	IsForceSell  string `json:"isForcesell,omitempty"`
	IsShortSell  string `json:"isShortsell,omitempty"`
	RejectReason string `json:"rejectReason,omitempty"`
}

type OrderBookData struct {
	Account string      `json:"account,omitempty"`
	Orders  []OrderData `json:"orders,omitempty"`
}

type OrderHistoryData struct {
	OrderHistories []OrderData `json:"orderHistories,omitempty"`
}

type PlaceOrderResponse struct {
	Account      string `json:"account,omitempty"`
	InstrumentID string `json:"instrumentID,omitempty"`
	MarketID     string `json:"marketID,omitempty"`
	Market       string `json:"market,omitempty"`
	BuySell      string `json:"buySell,omitempty"`
	OrderType    string `json:"orderType,omitempty"`
	Price        string `json:"price,omitempty"`
	Quantity     string `json:"quantity,omitempty"`
	RequestID    string `json:"requestID,omitempty"`
	OrderID      string `json:"orderID,omitempty"`
	ChannelID    string `json:"channelID,omitempty"`
	Code         string `json:"code,omitempty"`
	DeviceID     string `json:"deviceId,omitempty"`
	UserAgent    string `json:"userAgent,omitempty"`
}

type StockPortfolioData struct {
	TotalMarketValue string              `json:"totalMarketValue,omitempty"`
	StockPositions   []StockPositionData `json:"stockPositions,omitempty"`
}

type StockPositionData struct {
	MarketID     string `json:"marketID,omitempty"`
	InstrumentID string `json:"instrumentID,omitempty"`
	OnHand       string `json:"onHand,omitempty"`
	Block        string `json:"block,omitempty"`
	Bonus        string `json:"bonus,omitempty"`
	BuyT0        string `json:"buyT0,omitempty"`
	BuyT1        string `json:"buyT1,omitempty"`
	BuyT2        string `json:"buyT2,omitempty"`
	SellT0       string `json:"sellT0,omitempty"`
	SellT1       string `json:"sellT1,omitempty"`
	SellT2       string `json:"sellT2,omitempty"`
	AveragePrice string `json:"avgPrice,omitempty"`
	Mortgage     string `json:"mortgage,omitempty"`
	SellableQty  string `json:"sellableQty,omitempty"`
	HoldForTrade string `json:"holdForTrade,omitempty"`
	MarketPrice  string `json:"marketPrice,omitempty"`
	ExchangeID   string `json:"exchangeID,omitempty"`
	SellingQty   string `json:"sellingQty,omitempty"`
	BuyingQty    string `json:"buyingQty,omitempty"`
}

type DerivativePositionData struct {
	MarketID     string `json:"marketID,omitempty"`
	InstrumentID string `json:"instrumentID,omitempty"`
	LongQty      string `json:"longQty,omitempty"`
	ShortQty     string `json:"shortQty,omitempty"`
	Net          string `json:"net,omitempty"`
	BidAvgPrice  string `json:"bidAvgPrice,omitempty"`
	AskAvgPrice  string `json:"askAvgPrice,omitempty"`
	TradePrice   string `json:"tradePrice,omitempty"`
	MarketPrice  string `json:"marketPrice,omitempty"`
	FloatingPL   string `json:"floatingPL,omitempty"`
	TradingPL    string `json:"tradingPL,omitempty"`
}

type DerivativePositionsData struct {
	Account        string                   `json:"account,omitempty"`
	OpenPositions  []DerivativePositionData `json:"openPositions,omitempty"`
	ClosePositions []DerivativePositionData `json:"closePositions,omitempty"`
}

type AccountAssetData struct {
	Account           string `json:"account,omitempty"`
	CollateralAsset   string `json:"collateralAsset,omitempty"`
	CallLMW           string `json:"callLMW,omitempty"`
	Liability         string `json:"liability,omitempty"`
	EEOrigin          string `json:"eeOrigin,omitempty"`
	ForceLMV          string `json:"forceLMV,omitempty"`
	Equity            string `json:"equity,omitempty"`
	EE                string `json:"ee,omitempty"`
	CallMargin        string `json:"callMargin,omitempty"`
	CashBalance       string `json:"cashBalance,omitempty"`
	PurchasingPower   string `json:"purchasingPower,omitempty"`
	CallForceSell     string `json:"callForceSell,omitempty"`
	LMV               string `json:"lmv,omitempty"`
	MarginCall        string `json:"marginCall,omitempty"`
	Withdrawal        string `json:"withdrawal,omitempty"`
	CollateralA       string `json:"collateralA,omitempty"`
	Action            string `json:"action,omitempty"`
	MarginRatio       string `json:"marginRatio,omitempty"`
	Debt              string `json:"debt,omitempty"`
	AccruedInterest   string `json:"accruedInterest,omitempty"`
	HoldRight         string `json:"holdRight,omitempty"`
	PreLoan           string `json:"preLoan,omitempty"`
	Fees              string `json:"fees,omitempty"`
	BuyUnmatch        string `json:"buyUnmatch,omitempty"`
	AP                string `json:"ap,omitempty"`
	APT1              string `json:"apT1,omitempty"`
	SellUnmatch       string `json:"sellUnmatch,omitempty"`
	CIA               string `json:"cia,omitempty"`
	AR                string `json:"ar,omitempty"`
	ART1              string `json:"arT1,omitempty"`
	PPCredit          string `json:"ppCredit,omitempty"`
	CreditLimit       string `json:"creditLimit,omitempty"`
	TotalAsset        string `json:"totalAsset,omitempty"`
	TotalAssets       string `json:"totalAssets,omitempty"`
	MarginCallLMVSold string `json:"marginCallLMVSold,omitempty"`
	LMVNonMarginable  string `json:"lmvNonMarginable,omitempty"`
	EECredit          string `json:"eeCredit,omitempty"`
	TotalEquity       string `json:"totalEquity,omitempty"`
	EE90              string `json:"eE90,omitempty"`
	EE80              string `json:"eE80,omitempty"`
	EE70              string `json:"eE70,omitempty"`
	EE60              string `json:"eE60,omitempty"`
	EE50              string `json:"eE50,omitempty"`
	Dividend          string `json:"dividend,omitempty"`
	MaintenanceRatio  string `json:"maintenanceRatio,omitempty"`
	WarningRatio      string `json:"warningRatio,omitempty"`
}

type APILimitData struct {
	Endpoint string `json:"endpoint,omitempty"`
	Period   string `json:"period,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

// ──────────────────────────────── Stock Transfer ────────────────────────────────

type TransferableStockAccountData struct {
	Account            string              `json:"account,omitempty"`
	TransferableStocks []TransferableStock `json:"transferableStocks,omitempty"`
}

type TransferableStock struct {
	InstrumentID   string `json:"instrumentID,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	InstrumentType string `json:"instrumentType,omitempty"`
}

type StockTransferHistoryAccountData struct {
	Account                string                     `json:"account,omitempty"`
	StockTransferHistories []StockTransferHistoryData `json:"stockTransferHistories,omitempty"`
}

type StockTransferHistoryData struct {
	TransactionID      string `json:"transactionID,omitempty"`
	BeneficiaryAccount string `json:"beneficiaryAccount,omitempty"`
	InstrumentID       string `json:"instrumentID,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	DateTime           string `json:"dateTime,omitempty"`
	Status             string `json:"status,omitempty"`
	Remark             string `json:"remark,omitempty"`
	AuditRemark        string `json:"auditRemark,omitempty"`
}

// ──────────────────────────────── Rights ────────────────────────────────

type DividendsData struct {
	Account   string         `json:"account,omitempty"`
	Dividends []DividendData `json:"dividends,omitempty"`
}

type DividendData struct {
	StockDividend          string `json:"stockDividend,omitempty"`
	InstrumentID           string `json:"instrumentID,omitempty"`
	Quantity               string `json:"quantity,omitempty"`
	ExecutedRate           string `json:"executedRate,omitempty"`
	CloseDate              string `json:"closeDate,omitempty"`
	PaidDate               string `json:"paidDate,omitempty"`
	Amount                 string `json:"amount,omitempty"`
	Status                 string `json:"status,omitempty"`
	ReceivedQuantity       string `json:"receivedQuantity,omitempty"`
	IssueInstrument        string `json:"issueInstrument,omitempty"`
	DistributedFlag        string `json:"distributedFlag,omitempty"`
	PayableDate            string `json:"payableDate,omitempty"`
	SubscriptionPrice      string `json:"subscriptionPrice,omitempty"`
	SubscriptionAmount     string `json:"subscriptionAmount,omitempty"`
	SubscriptionQuantity   string `json:"subscriptionQuantity,omitempty"`
	SubscriptionPeriodFrom string `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo   string `json:"subscriptionPeriodTo,omitempty"`
	EntitlementID          string `json:"entitlementID,omitempty"`
	ExchangeID             string `json:"exchangeID,omitempty"`
}

type ExercisableQuantitiesData struct {
	Account               string                    `json:"account,omitempty"`
	ExercisableQuantities []ExercisableQuantityData `json:"exercisableQuantities,omitempty"`
}

type ExercisableQuantityData struct {
	EntitlementID               string `json:"entitlementID,omitempty"`
	InstrumentID                string `json:"instrumentID,omitempty"`
	SubscriptionPrice           string `json:"subscriptionPrice,omitempty"`
	ExecutedRateFrom            string `json:"executedRateFrom,omitempty"`
	SubscriptionPeriodFrom      string `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo        string `json:"subscriptionPeriodTo,omitempty"`
	ExerciseableQuantity        string `json:"exerciseableQuantity,omitempty"`
	ExerciseableReceiveQuantity string `json:"exerciseableReceiveQuantity,omitempty"`
	ExercisedReceiveQuantity    string `json:"exercisedReceiveQuantity,omitempty"`
	ExecutedRateTo              string `json:"executedRateTo,omitempty"`
	ExercisedQuantity           string `json:"exercisedQuantity,omitempty"`
	PayableDate                 string `json:"payableDate,omitempty"`
}

type RightsHistoriesData struct {
	Account                          string          `json:"account,omitempty"`
	OnlineRightSubscriptionHistories []RightsHistory `json:"onlineRightSubscriptionHistories,omitempty"`
}

type RightsHistory struct {
	TransactionID             string `json:"transactionID,omitempty"`
	DateTime                  string `json:"dateTime,omitempty"`
	InstrumentID              string `json:"instrumentID,omitempty"`
	RatioFrom                 string `json:"ratioFrom,omitempty"`
	SubscriptionPrice         string `json:"subscriptionPrice,omitempty"`
	SubscriptionPeriodFrom    string `json:"subscriptionPeriodFrom,omitempty"`
	SubscriptionPeriodTo      string `json:"subscriptionPeriodTo,omitempty"`
	ExercisedReceivedQuantity string `json:"exercisedReceivedQty,omitempty"`
	Amount                    string `json:"amount,omitempty"`
	Status                    string `json:"status,omitempty"`
	RatioTo                   string `json:"ratioTo,omitempty"`
	UnderlyingInstrumentID    string `json:"underlyingInstrumentID,omitempty"`
}

// ──────────────────────────────── Conditional Orders (FCO) ────────────────────────────────

type ConditionalOrderResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	FCOID   string `json:"fcoid,omitempty"`
	DCOID   string `json:"dcoid,omitempty"`
}

type ConditionalOrderPage[T any] struct {
	Message       string `json:"message,omitempty"`
	Status        int    `json:"status,omitempty"`
	Data          []T    `json:"data,omitempty"`
	FCOList       []T    `json:"fcoList,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	PageNumber    int    `json:"pageNumber,omitempty"`
	ItemsCount    int    `json:"itemsCount,omitempty"`
	PageCount     any    `json:"pageCount,omitempty"`
	TotalElements int    `json:"totalElements,omitempty"`
	TotalPages    int    `json:"totalPages,omitempty"`
	Page          int    `json:"page,omitempty"`
	Size          int    `json:"size,omitempty"`
}

type ConditionalTriggeredOrder struct {
	FCOID           string `json:"fcoId,omitempty"`
	Account         string `json:"account,omitempty"`
	Quantity        string `json:"quantity,omitempty"`
	Price           string `json:"price,omitempty"`
	InstrumentID    string `json:"instrumentId,omitempty"`
	Side            string `json:"side,omitempty"`
	OrderType       string `json:"orderType,omitempty"`
	IsMainOrder     bool   `json:"isMainOrder,omitempty"`
	IsAttachedOrder bool   `json:"isAttachedOrder,omitempty"`
	CreatedTime     string `json:"createdTime,omitempty"`
	UpdatedTime     string `json:"updatedTime,omitempty"`
	UniqueID        string `json:"uniqueId,omitempty"`
	OrderID         string `json:"orderId,omitempty"`
	MatchedQuantity string `json:"matchedQuantity,omitempty"`
	OSQuantity      string `json:"osQuantity,omitempty"`
	AveragePrice    string `json:"avgPrice,omitempty"`
	Status          string `json:"status,omitempty"`
	Detail          string `json:"detail,omitempty"`
}

type ConditionalOrderStatus struct {
	State  string `json:"state,omitempty"`
	Time   string `json:"time,omitempty"`
	Code   int    `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type ConditionalOrder struct {
	FCOID           string  `json:"fcoId,omitempty"`
	DCOID           string  `json:"dcoId,omitempty"`
	Username        string  `json:"username,omitempty"`
	UserID          string  `json:"userid,omitempty"`
	Account         string  `json:"account,omitempty"`
	AccountNo       string  `json:"accountNo,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	Price           any     `json:"price,omitempty"`
	PriceSlip       float64 `json:"priceSlip,omitempty"`
	InstrumentID    string  `json:"instrumentID,omitempty"`
	Side            string  `json:"side,omitempty"`
	Type            string  `json:"type,omitempty"`
	ProcessStatus   string  `json:"processStatus,omitempty"`
	OrderType       string  `json:"orderType,omitempty"`
	RunningMode     bool    `json:"runningMode,omitempty"`
	MainOrder       bool    `json:"mainOrder,omitempty"`
	AttachedOrder   bool    `json:"attachedOrder,omitempty"`
	CreatedDate     string  `json:"createdDate,omitempty"`
	CreatedTime     string  `json:"createdTime,omitempty"`
	UpdatedTime     string  `json:"updatedTime,omitempty"`
	MatchedQuantity int     `json:"matchedQuantity,omitempty"`
	PlaceOrder      bool    `json:"placeOrder,omitempty"`
	IsPlaceOrder    bool    `json:"isPlaceOrder,omitempty"`
	Status          string  `json:"status,omitempty"`
	StatusDetail    string  `json:"statusDetail,omitempty"`
	Detail          string  `json:"detail,omitempty"`
	Params          any     `json:"params,omitempty"`
	OrderID         string  `json:"orderID,omitempty"`
	OrderStatus     string  `json:"orderStatus,omitempty"`
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
