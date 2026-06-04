package ssi

import (
	"encoding/json"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func UnmarshalRawPayload(payload domain.RawPayload, out any) error {
	if len(payload.Bytes) > 0 {
		return json.Unmarshal(payload.Bytes, out)
	}
	bytes, err := json.Marshal(payload.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}

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

type TokenResponse struct {
	ConsumerID     string `json:"consumerID,omitempty"`
	ConsumerSecret string `json:"consumerSecret,omitempty"`
	TwoFactorType  int    `json:"twoFactorType,omitempty"`
	Code           string `json:"code,omitempty"`
	IsSave         bool   `json:"isSave,omitempty"`
}

type AccessTokenRequest struct {
	TwoFactorType int
	Code          string
	IsSave        bool
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}

type MaxBuyQuantityRequest struct {
	AccountID string
	Symbol    string
	Price     decimal.Decimal
}

type PlaceOrderRequest struct {
	domain.PlaceOrderRequest
	MarketID  string
	ChannelID string
	Code      string
	DeviceID  string
	UserAgent string
	RequestID string
}

type CancelOrderRequest struct {
	AccountID string
	OrderID   string
	Symbol    string
	Side      domain.OrderSide
	MarketID  string
	Code      string
	DeviceID  string
	UserAgent string
	RequestID string
}

type ModifyOrderRequest struct {
	AccountID string
	OrderID   string
	Symbol    string
	Side      domain.OrderSide
	MarketID  string
	Price     decimal.Decimal
	Quantity  int
	Code      string
	DeviceID  string
	UserAgent string
	RequestID string
}

type Security struct {
	Market      string `json:"Market,omitempty"`
	Symbol      string `json:"Symbol,omitempty"`
	StockName   string `json:"StockName,omitempty"`
	StockEnName string `json:"StockEnName,omitempty"`
}

type SecurityDetailGroup struct {
	RType        string           `json:"RType,omitempty"`
	ReportDate   string           `json:"ReportDate,omitempty"`
	TotalNoSym   string           `json:"TotalNoSym,omitempty"`
	RepeatedInfo []SecurityDetail `json:"RepeatedInfo,omitempty"`
}

type SecurityDetail struct {
	ISIN               any    `json:"Isin,omitempty"`
	Symbol             string `json:"Symbol,omitempty"`
	SymbolName         string `json:"SymbolName,omitempty"`
	SymbolEngName      string `json:"SymbolEngName,omitempty"`
	SecType            string `json:"SecType,omitempty"`
	MarketID           string `json:"MarketId,omitempty"`
	Exchange           string `json:"Exchange,omitempty"`
	Issuer             any    `json:"Issuer,omitempty"`
	LotSize            string `json:"LotSize,omitempty"`
	IssueDate          string `json:"IssueDate,omitempty"`
	MaturityDate       string `json:"MaturityDate,omitempty"`
	FirstTradingDate   string `json:"FirstTradingDate,omitempty"`
	LastTradingDate    string `json:"LastTradingDate,omitempty"`
	ContractMultiplier string `json:"ContractMultiplier,omitempty"`
	SettlementMethod   string `json:"SettlMethod,omitempty"`
	Underlying         any    `json:"Underlying,omitempty"`
	PutOrCall          any    `json:"PutOrCall,omitempty"`
	ExercisePrice      string `json:"ExercisePrice,omitempty"`
	ExerciseStyle      string `json:"ExerciseStyle,omitempty"`
	ExerciseRatio      string `json:"ExcerciseRatio,omitempty"`
	ListedShare        string `json:"ListedShare,omitempty"`
}

type Index struct {
	IndexCode string `json:"IndexCode,omitempty"`
	IndexName string `json:"IndexName,omitempty"`
	Exchange  string `json:"Exchange,omitempty"`
}

type IndexComponentGroup struct {
	IndexCode      string           `json:"IndexCode,omitempty"`
	IndexName      string           `json:"IndexName,omitempty"`
	Exchange       string           `json:"Exchange,omitempty"`
	TotalSymbolNo  string           `json:"TotalSymbolNo,omitempty"`
	IndexComponent []IndexComponent `json:"IndexComponent,omitempty"`
}

type IndexComponent struct {
	ISIN        string `json:"Isin,omitempty"`
	StockSymbol string `json:"StockSymbol,omitempty"`
}

type IndexDailyData struct {
	IndexID        string           `json:"IndexId,omitempty"`
	IndexValue     *decimal.Decimal `json:"IndexValue,omitempty"`
	TradingDate    string           `json:"TradingDate,omitempty"`
	Time           any              `json:"Time,omitempty"`
	Change         *decimal.Decimal `json:"Change,omitempty"`
	RatioChange    *decimal.Decimal `json:"RatioChange,omitempty"`
	TotalTrade     *decimal.Decimal `json:"TotalTrade,omitempty"`
	TotalMatchVol  *decimal.Decimal `json:"TotalMatchVol,omitempty"`
	TotalMatchVal  *decimal.Decimal `json:"TotalMatchVal,omitempty"`
	TypeIndex      any              `json:"TypeIndex,omitempty"`
	IndexName      string           `json:"IndexName,omitempty"`
	Advances       *decimal.Decimal `json:"Advances,omitempty"`
	NoChanges      *decimal.Decimal `json:"NoChanges,omitempty"`
	Declines       *decimal.Decimal `json:"Declines,omitempty"`
	Ceilings       *decimal.Decimal `json:"Ceilings,omitempty"`
	Floors         *decimal.Decimal `json:"Floors,omitempty"`
	TotalDealVol   *decimal.Decimal `json:"TotalDealVol,omitempty"`
	TotalDealVal   *decimal.Decimal `json:"TotalDealVal,omitempty"`
	TotalVol       *decimal.Decimal `json:"TotalVol,omitempty"`
	TotalVal       *decimal.Decimal `json:"TotalVal,omitempty"`
	TradingSession string           `json:"TradingSession,omitempty"`
}

type SecurityDailyPrice struct {
	TradingDate         string           `json:"TradingDate,omitempty"`
	PriceChange         *decimal.Decimal `json:"PriceChange,omitempty"`
	PerPriceChange      *decimal.Decimal `json:"PerPriceChange,omitempty"`
	CeilingPrice        *decimal.Decimal `json:"CeilingPrice,omitempty"`
	FloorPrice          *decimal.Decimal `json:"FloorPrice,omitempty"`
	RefPrice            *decimal.Decimal `json:"RefPrice,omitempty"`
	OpenPrice           *decimal.Decimal `json:"OpenPrice,omitempty"`
	HighestPrice        *decimal.Decimal `json:"HighestPrice,omitempty"`
	LowestPrice         *decimal.Decimal `json:"LowestPrice,omitempty"`
	ClosePrice          *decimal.Decimal `json:"ClosePrice,omitempty"`
	AveragePrice        *decimal.Decimal `json:"AveragePrice,omitempty"`
	ClosePriceAdjusted  *decimal.Decimal `json:"ClosePriceAdjusted,omitempty"`
	TotalMatchVol       *decimal.Decimal `json:"TotalMatchVol,omitempty"`
	TotalMatchVal       *decimal.Decimal `json:"TotalMatchVal,omitempty"`
	TotalDealVal        *decimal.Decimal `json:"TotalDealVal,omitempty"`
	TotalDealVol        *decimal.Decimal `json:"TotalDealVol,omitempty"`
	ForeignBuyVolTotal  *decimal.Decimal `json:"ForeignBuyVolTotal,omitempty"`
	ForeignCurrentRoom  *decimal.Decimal `json:"ForeignCurrentRoom,omitempty"`
	ForeignSellVolTotal *decimal.Decimal `json:"ForeignSellVolTotal,omitempty"`
	ForeignBuyValTotal  *decimal.Decimal `json:"ForeignBuyValTotal,omitempty"`
	ForeignSellValTotal *decimal.Decimal `json:"ForeignSellValTotal,omitempty"`
	TotalBuyTrade       *decimal.Decimal `json:"TotalBuyTrade,omitempty"`
	TotalBuyTradeVol    *decimal.Decimal `json:"TotalBuyTradeVol,omitempty"`
	TotalSellTrade      *decimal.Decimal `json:"TotalSellTrade,omitempty"`
	TotalSellTradeVol   *decimal.Decimal `json:"TotalSellTradeVol,omitempty"`
	NetBuySellVol       *decimal.Decimal `json:"NetBuySellVol,omitempty"`
	NetBuySellVal       *decimal.Decimal `json:"NetBuySellVal,omitempty"`
	TotalTradedVol      *decimal.Decimal `json:"TotalTradedVol,omitempty"`
	TotalTradedValue    *decimal.Decimal `json:"TotalTradedValue,omitempty"`
	Symbol              string           `json:"Symbol,omitempty"`
	Time                any              `json:"Time,omitempty"`
}

type OHLC struct {
	Symbol      string           `json:"Symbol,omitempty"`
	Market      string           `json:"Market,omitempty"`
	TradingDate string           `json:"TradingDate,omitempty"`
	Time        any              `json:"Time,omitempty"`
	Open        *decimal.Decimal `json:"Open,omitempty"`
	High        *decimal.Decimal `json:"High,omitempty"`
	Low         *decimal.Decimal `json:"Low,omitempty"`
	Close       *decimal.Decimal `json:"Close,omitempty"`
	Volume      *decimal.Decimal `json:"Volume,omitempty"`
	Value       *decimal.Decimal `json:"Value,omitempty"`
}

type OrderRequestResponse struct {
	Account      string           `json:"account,omitempty"`
	InstrumentID string           `json:"instrumentID,omitempty"`
	MarketID     string           `json:"marketID,omitempty"`
	Market       string           `json:"market,omitempty"`
	BuySell      string           `json:"buySell,omitempty"`
	OrderType    string           `json:"orderType,omitempty"`
	Price        *decimal.Decimal `json:"price,omitempty"`
	Quantity     *decimal.Decimal `json:"quantity,omitempty"`
	RequestID    string           `json:"requestID,omitempty"`
	OrderID      string           `json:"orderID,omitempty"`
	ChannelID    string           `json:"channelID,omitempty"`
	Code         string           `json:"code,omitempty"`
	DeviceID     string           `json:"deviceId,omitempty"`
	UserAgent    string           `json:"userAgent,omitempty"`
}

type Order struct {
	UniqueID     string           `json:"uniqueID,omitempty"`
	OrderID      string           `json:"orderID,omitempty"`
	BuySell      string           `json:"buySell,omitempty"`
	Price        *decimal.Decimal `json:"price,omitempty"`
	Quantity     *decimal.Decimal `json:"quantity,omitempty"`
	FilledQty    *decimal.Decimal `json:"filledQty,omitempty"`
	OrderStatus  string           `json:"orderStatus,omitempty"`
	MarketID     string           `json:"marketID,omitempty"`
	InputTime    string           `json:"inputTime,omitempty"`
	ModifiedTime string           `json:"modifiedTime,omitempty"`
	InstrumentID string           `json:"instrumentID,omitempty"`
	OrderType    string           `json:"orderType,omitempty"`
	CancelQty    *decimal.Decimal `json:"cancelQty,omitempty"`
	AveragePrice *decimal.Decimal `json:"avgPrice,omitempty"`
	IsForceSell  string           `json:"isForcesell,omitempty"`
	IsShortSell  string           `json:"isShortsell,omitempty"`
	RejectReason string           `json:"rejectReason,omitempty"`
}

type OrderHistoryData struct {
	OrderHistories []Order `json:"orderHistories,omitempty"`
}

type StockAccountBalance struct {
	Account             string           `json:"account,omitempty"`
	CashBalance         *decimal.Decimal `json:"cashbal,omitempty"`
	CashOnHold          *decimal.Decimal `json:"cashonhold,omitempty"`
	SecureAmount        *decimal.Decimal `json:"secureamount,omitempty"`
	Withdrawable        *decimal.Decimal `json:"withdrawable,omitempty"`
	ReceivingCashT1     *decimal.Decimal `json:"receivingcasht1,omitempty"`
	ReceivingCashT2     *decimal.Decimal `json:"receivingcasht2,omitempty"`
	MatchedBuyVolume    *decimal.Decimal `json:"matchedbuyvolume,omitempty"`
	MatchedSellVolume   *decimal.Decimal `json:"matchedsellvolume,omitempty"`
	UnmatchedBuyVolume  *decimal.Decimal `json:"unmatchedbuyvolume,omitempty"`
	UnmatchedSellVolume *decimal.Decimal `json:"unmatchedsellvolume,omitempty"`
	PaidCashT1          *decimal.Decimal `json:"paidcasht1,omitempty"`
	PaidCashT2          *decimal.Decimal `json:"paidcasht2,omitempty"`
	CIA                 *decimal.Decimal `json:"cia,omitempty"`
	Debt                *decimal.Decimal `json:"debt,omitempty"`
	PurchasingPower     *decimal.Decimal `json:"purchasingpower,omitempty"`
	TotalAsset          *decimal.Decimal `json:"totalasset,omitempty"`
}

type DerivativeAccountBalance struct {
	Account               string           `json:"account,omitempty"`
	AccountBalance        *decimal.Decimal `json:"accountbalance,omitempty"`
	Fee                   *decimal.Decimal `json:"fee,omitempty"`
	Commission            *decimal.Decimal `json:"commission,omitempty"`
	Interest              *decimal.Decimal `json:"interest,omitempty"`
	Loan                  *decimal.Decimal `json:"loan,omitempty"`
	DeliveryAmount        *decimal.Decimal `json:"deliveryamount,omitempty"`
	FloatingPL            *decimal.Decimal `json:"floatingpl,omitempty"`
	TotalPL               *decimal.Decimal `json:"totalpl,omitempty"`
	Marginable            *decimal.Decimal `json:"marginable,omitempty"`
	Depositable           *decimal.Decimal `json:"depositable,omitempty"`
	RCCall                *decimal.Decimal `json:"rccall,omitempty"`
	Withdrawable          *decimal.Decimal `json:"withdrawable,omitempty"`
	NonCashDrawableRCCall *decimal.Decimal `json:"noncashdrawablerccall,omitempty"`
	InternalAssets        []any            `json:"internalassets,omitempty"`
	ExchangeAssets        []any            `json:"exchangeassets,omitempty"`
	InternalMargin        []any            `json:"internalmargin,omitempty"`
	ExchangeMargin        []any            `json:"exchangemargin,omitempty"`
	NAV                   *decimal.Decimal `json:"nav,omitempty"`
	OrigMarginRatio       *decimal.Decimal `json:"origMarginRatio,omitempty"`
}

type StockPortfolio struct {
	TotalMarketValue *decimal.Decimal `json:"totalMarketValue,omitempty"`
	StockPositions   []StockPosition  `json:"stockPositions,omitempty"`
}

type StockPosition struct {
	MarketID     string           `json:"marketID,omitempty"`
	InstrumentID string           `json:"instrumentID,omitempty"`
	OnHand       *decimal.Decimal `json:"onHand,omitempty"`
	Block        *decimal.Decimal `json:"block,omitempty"`
	Bonus        *decimal.Decimal `json:"bonus,omitempty"`
	BuyT0        *decimal.Decimal `json:"buyT0,omitempty"`
	BuyT1        *decimal.Decimal `json:"buyT1,omitempty"`
	BuyT2        *decimal.Decimal `json:"buyT2,omitempty"`
	SellT0       *decimal.Decimal `json:"sellT0,omitempty"`
	SellT1       *decimal.Decimal `json:"sellT1,omitempty"`
	SellT2       *decimal.Decimal `json:"sellT2,omitempty"`
	AveragePrice *decimal.Decimal `json:"avgPrice,omitempty"`
	Mortgage     *decimal.Decimal `json:"mortgage,omitempty"`
	HoldForTrade *decimal.Decimal `json:"holdForTrade,omitempty"`
	MarketPrice  *decimal.Decimal `json:"marketPrice,omitempty"`
}

type DerivativePositions struct {
	Account        string               `json:"account,omitempty"`
	OpenPositions  []DerivativePosition `json:"openPositions,omitempty"`
	ClosePositions []DerivativePosition `json:"closePositions,omitempty"`
}

type DerivativePosition struct {
	MarketID     string           `json:"marketID,omitempty"`
	InstrumentID string           `json:"instrumentID,omitempty"`
	LongQty      *decimal.Decimal `json:"longQty,omitempty"`
	ShortQty     *decimal.Decimal `json:"shortQty,omitempty"`
	Net          *decimal.Decimal `json:"net,omitempty"`
	BidAvgPrice  *decimal.Decimal `json:"bidAvgPrice,omitempty"`
	AskAvgPrice  *decimal.Decimal `json:"askAvgPrice,omitempty"`
	TradePrice   *decimal.Decimal `json:"tradePrice,omitempty"`
	MarketPrice  *decimal.Decimal `json:"marketPrice,omitempty"`
	FloatingPL   *decimal.Decimal `json:"floatingPL,omitempty"`
	TradingPL    *decimal.Decimal `json:"tradingPL,omitempty"`
}

type AccountAsset struct {
	Account          string           `json:"account,omitempty"`
	TotalAsset       *decimal.Decimal `json:"totalAsset,omitempty"`
	PurchasingPower  *decimal.Decimal `json:"purchasingPower,omitempty"`
	MarginRatio      *decimal.Decimal `json:"marginRatio,omitempty"`
	MaintenanceRatio *decimal.Decimal `json:"maintenanceRatio,omitempty"`
	WarningRatio     *decimal.Decimal `json:"warningRatio,omitempty"`
}

type MaxBuyQuantity struct {
	Account         string           `json:"account,omitempty"`
	MaxBuyQty       *decimal.Decimal `json:"maxbuyqty,omitempty"`
	MarginRatio     *decimal.Decimal `json:"marginRatio,omitempty"`
	PurchasingPower *decimal.Decimal `json:"purchasingPower,omitempty"`
	OrigMarginRatio *decimal.Decimal `json:"origMarginRatio,omitempty"`
}

type MaxSellQuantity struct {
	Account    string           `json:"account,omitempty"`
	MaxSellQty *decimal.Decimal `json:"maxSellQty,omitempty"`
}

type APILimit struct {
	Endpoint string `json:"endpoint,omitempty"`
	Period   string `json:"period,omitempty"`
	Limit    int    `json:"limit,omitempty"`
}

type TransferableStockAccount struct {
	Account            string              `json:"account,omitempty"`
	TransferableStocks []TransferableStock `json:"transferableStocks,omitempty"`
}

type TransferableStock struct {
	InstrumentID   string           `json:"instrumentID,omitempty"`
	Quantity       *decimal.Decimal `json:"quantity,omitempty"`
	InstrumentType string           `json:"instrumentType,omitempty"`
}

type StockTransferHistoryAccount struct {
	Account                string                 `json:"account,omitempty"`
	StockTransferHistories []StockTransferHistory `json:"stockTransferHistories,omitempty"`
}

type StockTransferHistory struct {
	TransactionID      string           `json:"transactionID,omitempty"`
	BeneficiaryAccount string           `json:"beneficiaryAccount,omitempty"`
	InstrumentID       string           `json:"instrumentID,omitempty"`
	Quantity           *decimal.Decimal `json:"quantity,omitempty"`
	DateTime           string           `json:"dateTime,omitempty"`
	Status             string           `json:"status,omitempty"`
	Remark             string           `json:"remark,omitempty"`
	AuditRemark        string           `json:"auditRemark,omitempty"`
}

type CashTransferHistory struct {
	TransactionID   string           `json:"transactionID,omitempty"`
	SenderAccount   string           `json:"senderAccount,omitempty"`
	ReceiverAccount string           `json:"receiverAccount,omitempty"`
	Amount          *decimal.Decimal `json:"amount,omitempty"`
	DateTime        string           `json:"dateTime,omitempty"`
	Status          string           `json:"status,omitempty"`
	Remark          string           `json:"remark,omitempty"`
}

type CashInAdvance struct {
	Account       string           `json:"account,omitempty"`
	TransactionID string           `json:"transactionID,omitempty"`
	Amount        *decimal.Decimal `json:"amount,omitempty"`
	SettleDate    string           `json:"settleDate,omitempty"`
	InstrumentID  string           `json:"instrumentID,omitempty"`
	CIAAmount     *decimal.Decimal `json:"ciaAmount,omitempty"`
	Fee           *decimal.Decimal `json:"fee,omitempty"`
	DateTime      string           `json:"dateTime,omitempty"`
	Status        string           `json:"status,omitempty"`
}

type MaxCashInAdvance struct {
	Account         string           `json:"account,omitempty"`
	AvailableAmount *decimal.Decimal `json:"availableAmount,omitempty"`
	Fee             *decimal.Decimal `json:"fee,omitempty"`
}

type CashInAdvanceFee struct {
	Account          string           `json:"account,omitempty"`
	CIAAmount        *decimal.Decimal `json:"ciaAmount,omitempty"`
	Fee              *decimal.Decimal `json:"fee,omitempty"`
	ReceivableAmount *decimal.Decimal `json:"receivableAmount,omitempty"`
}

type PurchasableRight struct {
	Account           string           `json:"account,omitempty"`
	InstrumentID      string           `json:"instrumentID,omitempty"`
	AvailableQuantity *decimal.Decimal `json:"availableQuantity,omitempty"`
	RightRatio        string           `json:"rightRatio,omitempty"`
	RegisterFromDate  string           `json:"registerFromDate,omitempty"`
	RegisterToDate    string           `json:"registerToDate,omitempty"`
}

type PurchasableRightQuantity struct {
	Account             string           `json:"account,omitempty"`
	InstrumentID        string           `json:"instrumentID,omitempty"`
	ExercisableQuantity *decimal.Decimal `json:"exercisableQuantity,omitempty"`
}

type PurchasableRightHistory struct {
	TransactionID string           `json:"transactionID,omitempty"`
	Account       string           `json:"account,omitempty"`
	InstrumentID  string           `json:"instrumentID,omitempty"`
	Quantity      *decimal.Decimal `json:"quantity,omitempty"`
	DateTime      string           `json:"dateTime,omitempty"`
	Status        string           `json:"status,omitempty"`
}

type ConditionalOrder struct {
	FCOID         string           `json:"fcoId,omitempty"`
	Account       string           `json:"account,omitempty"`
	InstrumentID  string           `json:"instrumentID,omitempty"`
	Side          string           `json:"side,omitempty"`
	Type          string           `json:"type,omitempty"`
	ProcessStatus string           `json:"processStatus,omitempty"`
	CreatedDate   string           `json:"createdDate,omitempty"`
	OrderID       string           `json:"orderID,omitempty"`
	Quantity      *decimal.Decimal `json:"quantity,omitempty"`
	Price         *decimal.Decimal `json:"price,omitempty"`
	OrderStatus   string           `json:"orderStatus,omitempty"`
}
