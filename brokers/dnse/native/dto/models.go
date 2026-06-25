package dto

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

type GetTradingTokenRequest struct {
	OTPType  string `json:"otpType"`
	Passcode string `json:"passcode"`
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type SendEmailOtpResponse = ErrorResponse

type TradingTokenResponse struct {
	TradingToken string `json:"tradingToken,omitempty"`
}

type TwoFAVerificationResponse = TradingTokenResponse

type AccountsResponse struct {
	Name        string    `json:"name,omitempty"`
	CustodyCode string    `json:"custodyCode,omitempty"`
	InvestorID  string    `json:"investorId,omitempty"`
	Accounts    []Account `json:"accounts,omitempty"`
}

type GetAccountsResponse = AccountsResponse

type Account struct {
	ID                string             `json:"id,omitempty"`
	DealAccount       bool               `json:"dealAccount,omitempty"`
	DerivativeAccount bool               `json:"derivativeAccount,omitempty"`
	Derivative        *DerivativeAccount `json:"derivative,omitempty"`
}

type DerivativeAccount struct {
	Status string `json:"status,omitempty"`
}

type AccountBalancesResponse struct {
	Stock      *StockBalance      `json:"stock,omitempty"`
	Derivative *DerivativeBalance `json:"derivative,omitempty"`
}

type GetAccountBalancesResponse = AccountBalancesResponse

type StockBalance struct {
	TotalCash             *decimal.Decimal `json:"totalCash,omitempty"`
	AvailableCash         *decimal.Decimal `json:"availableCash,omitempty"`
	DepositInterest       *decimal.Decimal `json:"depositInterest,omitempty"`
	TotalDebt             *decimal.Decimal `json:"totalDebt,omitempty"`
	DepositFeeAmount      *decimal.Decimal `json:"depositFeeAmount,omitempty"`
	WithdrawableCash      *decimal.Decimal `json:"withdrawableCash,omitempty"`
	CashDividendReceiving *decimal.Decimal `json:"cashDividendReceiving,omitempty"`
}

type DerivativeBalance struct {
	PendingDepositWithdraw *decimal.Decimal `json:"pendingDepositWithdraw,omitempty"`
	RemainSecure           *decimal.Decimal `json:"remainSecure,omitempty"`
	UsedSecure             *decimal.Decimal `json:"usedSecure,omitempty"`
	PendingSecure          *decimal.Decimal `json:"pendingSecure,omitempty"`
	HoldTaxAndFee          *decimal.Decimal `json:"holdTaxAndFee,omitempty"`
	TotalLoanDebt          *decimal.Decimal `json:"totalLoanDebt,omitempty"`
}

type CorporateActionHistoryResponse struct {
	AccountNo  string               `json:"accountNo,omitempty"`
	Data       *CorporateActionData `json:"data,omitempty"`
	Pagination *Pagination          `json:"pagination,omitempty"`
}

type GetCorporateActionHistoryResponse = CorporateActionHistoryResponse

type CorporateActionData struct {
	CashDividend   []CashDividend   `json:"cashDividend,omitempty"`
	StockDividend  []StockDividend  `json:"stockDividend,omitempty"`
	StockBonus     []StockBonus     `json:"stockBonus,omitempty"`
	RightsOffering []RightsOffering `json:"rightsOffering,omitempty"`
}

type CashDividend struct {
	ID              int              `json:"id,omitempty"`
	Symbol          string           `json:"symbol,omitempty"`
	CAStatus        string           `json:"caStatus,omitempty"`
	RecordDate      string           `json:"recordDate,omitempty"`
	ProcessDate     string           `json:"processDate,omitempty"`
	HoldingQuantity *decimal.Decimal `json:"holdingQuantity,omitempty"`
	DividendValue   *decimal.Decimal `json:"dividendValue,omitempty"`
	GrossAmount     *decimal.Decimal `json:"grossAmount,omitempty"`
	TaxAmount       *decimal.Decimal `json:"taxAmount,omitempty"`
	NetAmount       *decimal.Decimal `json:"netAmount,omitempty"`
}

type StockDividend struct {
	ID              int              `json:"id,omitempty"`
	Symbol          string           `json:"symbol,omitempty"`
	CAStatus        string           `json:"caStatus,omitempty"`
	RecordDate      string           `json:"recordDate,omitempty"`
	ProcessDate     string           `json:"processDate,omitempty"`
	HoldingQuantity *decimal.Decimal `json:"holdingQuantity,omitempty"`
	Ratio           string           `json:"ratio,omitempty"`
	Quantity        *decimal.Decimal `json:"quantity,omitempty"`
}

type StockBonus = StockDividend

type RightsOffering struct {
	ID                 int              `json:"id,omitempty"`
	Symbol             string           `json:"symbol,omitempty"`
	CAStatus           string           `json:"caStatus,omitempty"`
	RecordDate         string           `json:"recordDate,omitempty"`
	ProcessDate        string           `json:"processDate,omitempty"`
	HoldingQuantity    *decimal.Decimal `json:"holdingQuantity,omitempty"`
	Ratio              string           `json:"ratio,omitempty"`
	Quantity           *decimal.Decimal `json:"quantity,omitempty"`
	RightQuantity      *decimal.Decimal `json:"rightQuantity,omitempty"`
	RegisteredQuantity *decimal.Decimal `json:"registeredQuantity,omitempty"`
	RegisterFromDate   string           `json:"registerFromDate,omitempty"`
	RegisterToDate     string           `json:"registerToDate,omitempty"`
	TransferFromDate   string           `json:"transferFromDate,omitempty"`
	TransferToDate     string           `json:"transferToDate,omitempty"`
}

type Pagination struct {
	PageIndex    int `json:"pageIndex,omitempty"`
	PageSize     int `json:"pageSize,omitempty"`
	TotalRecords int `json:"totalRecords,omitempty"`
}

type OrdersResponse struct {
	Orders []Order `json:"orders,omitempty"`
}

type GetOrdersResponse = OrdersResponse

type OrdersHistoryResponse struct {
	AccountNo    string           `json:"accountNo,omitempty"`
	FillQuantity *decimal.Decimal `json:"fillQuantity,omitempty"`
	Total        int              `json:"total,omitempty"`
	Start        int              `json:"start,omitempty"`
	End          int              `json:"end,omitempty"`
	MarketType   string           `json:"marketType,omitempty"`
	Data         []Order          `json:"data,omitempty"`
}

type GetOrdersHistoryResponse = OrdersHistoryResponse

type Order struct {
	ID               int64             `json:"id,omitempty"`
	InvestorID       string            `json:"investorId,omitempty"`
	Side             string            `json:"side,omitempty"`
	AccountNo        string            `json:"accountNo,omitempty"`
	Symbol           string            `json:"symbol,omitempty"`
	Price            *decimal.Decimal  `json:"price,omitempty"`
	PriceSecure      *decimal.Decimal  `json:"priceSecure,omitempty"`
	AveragePrice     *decimal.Decimal  `json:"averagePrice,omitempty"`
	Quantity         *decimal.Decimal  `json:"quantity,omitempty"`
	FillQuantity     *decimal.Decimal  `json:"fillQuantity,omitempty"`
	CanceledQuantity *decimal.Decimal  `json:"canceledQuantity,omitempty"`
	LeaveQuantity    *decimal.Decimal  `json:"leaveQuantity,omitempty"`
	LastQuantity     *decimal.Decimal  `json:"lastQuantity,omitempty"`
	LastPrice        *decimal.Decimal  `json:"lastPrice,omitempty"`
	OrderType        string            `json:"orderType,omitempty"`
	OrderCategory    string            `json:"orderCategory,omitempty"`
	OrderStatus      string            `json:"orderStatus,omitempty"`
	LoanPackageID    int64             `json:"loanPackageId,omitempty"`
	MarketType       string            `json:"marketType,omitempty"`
	TransDate        string            `json:"transDate,omitempty"`
	TaxRate          *decimal.Decimal  `json:"taxRate,omitempty"`
	ExchangeFeeRate  *decimal.Decimal  `json:"exchangeFeeRate,omitempty"`
	FeeRate          *decimal.Decimal  `json:"feeRate,omitempty"`
	Error            string            `json:"error,omitempty"`
	Metadata         string            `json:"metadata,omitempty"`
	Reports          []ExecutionReport `json:"reports,omitempty"`
	CreatedDate      string            `json:"createdDate,omitempty"`
	ModifiedDate     string            `json:"modifiedDate,omitempty"`
}

type WrapperEvent struct {
	T         string `json:"T,omitempty"`
	Action    string `json:"action,omitempty"`
	Event     string `json:"event,omitempty"`
	Sequence  int64  `json:"sequence,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type WrapperOrderEvent struct {
	WrapperEvent
	Order Order `json:"order,omitempty"`
}

func (o *Order) UnmarshalJSON(data []byte) error {
	type orderAlias Order
	var wire struct {
		orderAlias
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	*o = Order(wire.orderAlias)
	if len(wire.ID) == 0 || string(wire.ID) == "null" {
		return nil
	}
	if err := json.Unmarshal(wire.ID, &o.ID); err == nil {
		return nil
	}
	var id string
	if err := json.Unmarshal(wire.ID, &id); err != nil {
		return fmt.Errorf("decode order id: %w", err)
	}
	if id == "" {
		return nil
	}
	parsed, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("decode order id %q: %w", id, err)
	}
	o.ID = parsed
	return nil
}

type OrderResponse = Order
type PlaceOrderResponse = Order
type ReplaceOrderResponse = Order
type CancelOrderResponse = Order
type ClosePositionResponse = Order
type OrderDetailResponse = Order
type ExecutionsResponse = Order
type GetOrderDetailResponse = Order
type GetExecutionsResponse = Order

type PositionsResponse struct {
	Positions  []Position `json:"positions,omitempty"`
	PageIndex  int        `json:"pageIndex,omitempty"`
	PageSize   int        `json:"pageSize,omitempty"`
	PageNumber int        `json:"pageNumber,omitempty"`
	Total      int        `json:"total,omitempty"`
}

type GetPositionsResponse = PositionsResponse

type PositionByIDResponse struct {
	Data *Position `json:"data,omitempty"`
}

type GetPositionByIDResponse = PositionByIDResponse
type GetPositionByIdResponse = PositionByIDResponse

type Position struct {
	ID                 int64            `json:"id,omitempty"`
	AccountNo          string           `json:"accountNo,omitempty"`
	Symbol             string           `json:"symbol,omitempty"`
	Status             string           `json:"status,omitempty"`
	LoanPackageID      int64            `json:"loanPackageId,omitempty"`
	Side               string           `json:"side,omitempty"`
	AccumulateQuantity *decimal.Decimal `json:"accumulateQuantity,omitempty"`
	TradeQuantity      *decimal.Decimal `json:"tradeQuantity,omitempty"`
	ClosedQuantity     *decimal.Decimal `json:"closedQuantity,omitempty"`
	CostPrice          *decimal.Decimal `json:"costPrice,omitempty"`
	MarketPrice        *decimal.Decimal `json:"marketPrice,omitempty"`
	BreakEvenPrice     *decimal.Decimal `json:"breakEvenPrice,omitempty"`
	OpenQuantity       *decimal.Decimal `json:"openQuantity,omitempty"`
	OverNightQuantity  *decimal.Decimal `json:"overNightQuantity,omitempty"`
	AverageClosePrice  *decimal.Decimal `json:"averageClosePrice,omitempty"`
	MarketType         string           `json:"marketType,omitempty"`
	CreatedDate        string           `json:"createdDate,omitempty"`
	ModifiedDate       string           `json:"modifiedDate,omitempty"`
	AverageCostPrice   *decimal.Decimal `json:"averageCostPrice,omitempty"`
}

type WrapperPositionEvent struct {
	WrapperEvent
	Position Position `json:"position,omitempty"`
}

type LoanPackagesResponse struct {
	SymbolType   string        `json:"symbolType,omitempty"`
	MarketType   string        `json:"marketType,omitempty"`
	LoanPackages []LoanPackage `json:"loanPackages,omitempty"`
}

type GetLoanPackagesResponse = LoanPackagesResponse

type LoanPackage struct {
	ID              int              `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	InitialRate     *decimal.Decimal `json:"initialRate,omitempty"`
	MaintenanceRate *decimal.Decimal `json:"maintenanceRate,omitempty"`
	LiquidRate      *decimal.Decimal `json:"liquidRate,omitempty"`
	TradingFee      *TradingFee      `json:"tradingFee,omitempty"`
}

type TradingFee struct {
	ID                           int                  `json:"id,omitempty"`
	Name                         string               `json:"name,omitempty"`
	Scope                        string               `json:"scope,omitempty"`
	Channel                      string               `json:"channel,omitempty"`
	SchemaType                   string               `json:"schemaType,omitempty"`
	CreatedDate                  string               `json:"createdDate,omitempty"`
	ModifiedDate                 string               `json:"modifiedDate,omitempty"`
	FixedTradingFee              int                  `json:"fixedTradingFee,omitempty"`
	FixedDailyCloseTradingFee    int                  `json:"fixedDailyCloseTradingFee,omitempty"`
	ProgressTradingFee           []ProgressTradingFee `json:"progressTradingFee,omitempty"`
	ProgressDailyCloseTradingFee []ProgressTradingFee `json:"progressDailyCloseTradingFee,omitempty"`
}

type ProgressTradingFee struct {
	FromQuantity *decimal.Decimal `json:"fromQuantity,omitempty"`
	ToQuantity   *decimal.Decimal `json:"toQuantity,omitempty"`
	Fee          *decimal.Decimal `json:"fee,omitempty"`
}

type PPSECredit struct {
	QMaxBuy  int              `json:"qmaxBuy,omitempty"`
	QMaxSell int              `json:"qmaxSell,omitempty"`
	Price    *decimal.Decimal `json:"price,omitempty"`
	PP0Buy   *decimal.Decimal `json:"pp0Buy,omitempty"`
	PP0Sell  *decimal.Decimal `json:"pp0Sell,omitempty"`
}

type GetPPSECreditResponse = PPSECredit
type GetPpseResponse = PPSECredit

type ExecutionReport struct {
	ID               int64            `json:"id,omitempty"`
	OrderID          int64            `json:"orderId,omitempty"`
	ExecID           string           `json:"execId,omitempty"`
	Side             string           `json:"side,omitempty"`
	AccountNo        string           `json:"accountNo,omitempty"`
	Symbol           string           `json:"symbol,omitempty"`
	Price            *decimal.Decimal `json:"price,omitempty"`
	Quantity         *decimal.Decimal `json:"quantity,omitempty"`
	OrderType        string           `json:"orderType,omitempty"`
	OrderCategory    string           `json:"orderCategory,omitempty"`
	OrderStatus      string           `json:"orderStatus,omitempty"`
	FillQuantity     *decimal.Decimal `json:"fillQuantity,omitempty"`
	LastQuantity     *decimal.Decimal `json:"lastQuantity,omitempty"`
	LastPrice        *decimal.Decimal `json:"lastPrice,omitempty"`
	AveragePrice     *decimal.Decimal `json:"averagePrice,omitempty"`
	TransDate        string           `json:"transDate,omitempty"`
	LeaveQuantity    *decimal.Decimal `json:"leaveQuantity,omitempty"`
	CanceledQuantity *decimal.Decimal `json:"canceledQuantity,omitempty"`
	MarketType       string           `json:"marketType,omitempty"`
	PriceSecure      *decimal.Decimal `json:"priceSecure,omitempty"`
	CreatedDate      string           `json:"createdDate,omitempty"`
	ModifiedDate     string           `json:"modifiedDate,omitempty"`
}

type InstrumentsResponse struct {
	Data     []Instrument `json:"data,omitempty"`
	Total    int          `json:"total,omitempty"`
	Page     int          `json:"page,omitempty"`
	PageSize int          `json:"pageSize,omitempty"`
}

type InstrumentsByFilterResponse = InstrumentsResponse
type GetInstrumentsResponse = InstrumentsResponse
type GetInstrumentsByFilterResponse = InstrumentsResponse

type ExpectedPrice struct {
	T                     string           `json:"T,omitempty"`
	MarketID              string           `json:"marketId,omitempty"`
	BoardID               string           `json:"boardId,omitempty"`
	Symbol                string           `json:"symbol,omitempty"`
	ISIN                  string           `json:"isin,omitempty"`
	ClosePrice            *decimal.Decimal `json:"closePrice,omitempty"`
	ExpectedTradePrice    *decimal.Decimal `json:"expectedTradePrice,omitempty"`
	ExpectedTradeQuantity int64            `json:"expectedTradeQuantity,omitempty"`
	Time                  StreamTimestamp  `json:"time,omitempty"`
}

type Foreign struct {
	T                            string           `json:"T,omitempty"`
	BoardID                      string           `json:"boardId,omitempty"`
	BuyTradedAmount              *decimal.Decimal `json:"buyTradedAmount,omitempty"`
	BuyVolume                    int64            `json:"buyVolume,omitempty"`
	ForeignInvestorTypeCode      string           `json:"foreignInvestorTypeCode,omitempty"`
	ForeignerBuyPossibleQuantity int64            `json:"foreignerBuyPossibleQuantity,omitempty"`
	ForeignerOrderLimitQuantity  int64            `json:"foreignerOrderLimitQuantity,omitempty"`
	MarketID                     string           `json:"marketId,omitempty"`
	MulticastReceiveTime         StreamTimestamp  `json:"multicastReceiveTime,omitempty"`
	SellTradedAmount             *decimal.Decimal `json:"sellTradedAmount,omitempty"`
	SellVolume                   int64            `json:"sellVolume,omitempty"`
	Symbol                       string           `json:"symbol,omitempty"`
	Time                         string           `json:"time,omitempty"`
	TotalBuyTradedAmount         *decimal.Decimal `json:"totalBuyTradedAmount,omitempty"`
	TotalBuyVolume               int64            `json:"totalBuyVolume,omitempty"`
	TotalSellTradedAmount        *decimal.Decimal `json:"totalSellTradedAmount,omitempty"`
	TotalSellVolume              int64            `json:"totalSellVolume,omitempty"`
	TradingSessionID             string           `json:"tradingSessionId,omitempty"`
	TransactTime                 string           `json:"transactTime,omitempty"`
}

type GetForeignTradingResponse struct {
	Foreigners    []Foreign `json:"foreigners,omitempty"`
	NextPageToken string    `json:"nextPageToken,omitempty"`
}

type EstimatedMarketIndex struct {
	T           string                   `json:"T,omitempty"`
	Action      string                   `json:"action,omitempty"`
	MarketIndex EstimatedMarketIndexData `json:"marketIndex,omitempty"`
	Timestamp   int64                    `json:"timestamp,omitempty"`
}

type EstimatedMarketIndexData struct {
	ChangedRatio                    *decimal.Decimal `json:"changedRatio,omitempty"`
	ChangedValue                    *decimal.Decimal `json:"changedValue,omitempty"`
	FluctuationDownIssueCount       int              `json:"fluctuationDownIssueCount,omitempty"`
	FluctuationSteadinessIssueCount int              `json:"fluctuationSteadinessIssueCount,omitempty"`
	FluctuationUpIssueCount         int              `json:"fluctuationUpIssueCount,omitempty"`
	GrossTradeAmount                *decimal.Decimal `json:"grossTradeAmount,omitempty"`
	IndexName                       string           `json:"indexName,omitempty"`
	Time                            string           `json:"time,omitempty"`
	TotalVolumeTraded               int64            `json:"totalVolumeTraded,omitempty"`
	ValueIndexes                    *decimal.Decimal `json:"valueIndexes,omitempty"`
}

type Instrument struct {
	Symbol          string   `json:"symbol,omitempty"`
	MarketID        string   `json:"marketId,omitempty"`
	SecurityGroupID string   `json:"securityGroupId,omitempty"`
	SymbolType      string   `json:"symbolType,omitempty"`
	ListedDate      string   `json:"listedDate,omitempty"`
	ShortName       string   `json:"shortName,omitempty"`
	Name            string   `json:"name,omitempty"`
	IndexName       []string `json:"indexName,omitempty"`
}

type QuotesResponse struct {
	Quotes        []Quote `json:"quotes,omitempty"`
	NextPageToken string  `json:"nextPageToken,omitempty"`
}

type LatestQuotesResponse = QuotesResponse
type GetQuotesResponse = QuotesResponse
type GetLatestQuotesResponse = QuotesResponse

type Quote struct {
	T                    string           `json:"T,omitempty"`
	Bid                  []PriceLevel     `json:"bid,omitempty"`
	BoardID              string           `json:"boardId,omitempty"`
	ISIN                 string           `json:"isin,omitempty"`
	MarketID             string           `json:"marketId,omitempty"`
	MulticastReceiveTime StreamTimestamp  `json:"multicastReceiveTime,omitempty"`
	Offer                []PriceLevel     `json:"offer,omitempty"`
	Symbol               string           `json:"symbol,omitempty"`
	Time                 StreamTimestamp  `json:"time,omitempty"`
	TotalOfferQtty       *decimal.Decimal `json:"totalOfferQtty,omitempty"`
	TotalBidQtty         *decimal.Decimal `json:"totalBidQtty,omitempty"`
}

type PriceLevel struct {
	Price    *decimal.Decimal `json:"price,omitempty"`
	Quantity *decimal.Decimal `json:"qtty,omitempty"`
}

type TradesResponse struct {
	Trades        []Trade `json:"trades,omitempty"`
	NextPageToken string  `json:"nextPageToken,omitempty"`
}

type LatestTradesResponse = TradesResponse
type HistoryTradesResponse = TradesResponse
type GetLatestTradesResponse = TradesResponse
type GetHistoryTradesResponse = TradesResponse

type Trade struct {
	T                    string           `json:"T,omitempty"`
	BoardID              string           `json:"boardId,omitempty"`
	GrossTradeAmount     *decimal.Decimal `json:"grossTradeAmount,omitempty"`
	HighestPrice         *decimal.Decimal `json:"highestPrice,omitempty"`
	ISIN                 string           `json:"isin,omitempty"`
	LowestPrice          *decimal.Decimal `json:"lowestPrice,omitempty"`
	MarketID             string           `json:"marketId,omitempty"`
	MatchPrice           *decimal.Decimal `json:"matchPrice,omitempty"`
	MatchQuantity        *decimal.Decimal `json:"matchQtty,omitempty"`
	MulticastReceiveTime StreamTimestamp  `json:"multicastReceiveTime,omitempty"`
	OpenPrice            *decimal.Decimal `json:"openPrice,omitempty"`
	Symbol               string           `json:"symbol,omitempty"`
	Time                 StreamTimestamp  `json:"time,omitempty"`
	TotalVolumeTraded    *decimal.Decimal `json:"totalVolumeTraded,omitempty"`
	TradingSessionID     string           `json:"tradingSessionId,omitempty"`
}

type TradeExtra struct {
	T                    string           `json:"T,omitempty"`
	AveragePrice         *decimal.Decimal `json:"avgPrice,omitempty"`
	BoardID              string           `json:"boardId,omitempty"`
	GrossTradeAmount     *decimal.Decimal `json:"grossTradeAmount,omitempty"`
	HighestPrice         *decimal.Decimal `json:"highestPrice,omitempty"`
	ISIN                 string           `json:"isin,omitempty"`
	LowestPrice          *decimal.Decimal `json:"lowestPrice,omitempty"`
	MarketID             string           `json:"marketId,omitempty"`
	MatchPrice           *decimal.Decimal `json:"matchPrice,omitempty"`
	MatchQuantity        *decimal.Decimal `json:"matchQtty,omitempty"`
	MulticastReceiveTime StreamTimestamp  `json:"multicastReceiveTime,omitempty"`
	OpenPrice            *decimal.Decimal `json:"openPrice,omitempty"`
	Side                 string           `json:"side,omitempty"`
	Symbol               string           `json:"symbol,omitempty"`
	Time                 StreamTimestamp  `json:"time,omitempty"`
	TotalVolumeTraded    *decimal.Decimal `json:"totalVolumeTraded,omitempty"`
	TradingSessionID     string           `json:"tradingSessionId,omitempty"`
}

type OHLC struct {
	T           string           `json:"T,omitempty"`
	Time        int64            `json:"time,omitempty"`
	Open        *decimal.Decimal `json:"open,omitempty"`
	High        *decimal.Decimal `json:"high,omitempty"`
	Low         *decimal.Decimal `json:"low,omitempty"`
	Close       *decimal.Decimal `json:"close,omitempty"`
	Volume      int64            `json:"volume,omitempty"`
	Symbol      string           `json:"symbol,omitempty"`
	Resolution  string           `json:"resolution,omitempty"`
	LastUpdated int64            `json:"lastUpdated,omitempty"`
	Type        string           `json:"type,omitempty"`
}

type OHLCClosed = OHLC

type OhlcResponse struct {
	T        []int              `json:"t,omitempty"`
	O        []*decimal.Decimal `json:"o,omitempty"`
	H        []*decimal.Decimal `json:"h,omitempty"`
	L        []*decimal.Decimal `json:"l,omitempty"`
	C        []*decimal.Decimal `json:"c,omitempty"`
	V        []*decimal.Decimal `json:"v,omitempty"`
	NextTime int                `json:"nextTime,omitempty"`
}

type GetOhlcResponse = OhlcResponse

type PriceSymbolCloseResponse struct {
	Prices []ClosePrice `json:"prices,omitempty"`
}

type GetPriceSymbolCloseResponse = PriceSymbolCloseResponse

type ClosePrice struct {
	MarketID   string           `json:"marketId,omitempty"`
	BoardID    string           `json:"boardId,omitempty"`
	ISIN       string           `json:"isin,omitempty"`
	Symbol     string           `json:"symbol,omitempty"`
	ClosePrice *decimal.Decimal `json:"closePrice,omitempty"`
	Time       string           `json:"time,omitempty"`
}

type SecdefResponse []SecurityDefinition
type SecurityDefinitionList = SecdefResponse
type GetSecdefResponse = SecdefResponse

func (items SecdefResponse) Find(symbol string) (*SecurityDefinition, bool) {
	for i := range items {
		if items[i].Symbol == symbol {
			return &items[i], true
		}
	}
	return nil, false
}

func (items SecdefResponse) FloorPrice(symbol string) (decimal.Decimal, bool) {
	item, ok := items.Find(symbol)
	if !ok || item.FloorPrice == nil {
		return decimal.Zero, false
	}
	return *item.FloorPrice, true
}

type SecurityDefinition struct {
	MarketID                        string           `json:"marketId,omitempty"`
	BoardID                         string           `json:"boardId,omitempty"`
	ISIN                            string           `json:"isin,omitempty"`
	Symbol                          string           `json:"symbol,omitempty"`
	ProductGrpID                    string           `json:"productGrpId,omitempty"`
	SecurityGroupID                 string           `json:"securityGroupId,omitempty"`
	BasicPrice                      *decimal.Decimal `json:"basicPrice,omitempty"`
	CeilingPrice                    *decimal.Decimal `json:"ceilingPrice,omitempty"`
	FloorPrice                      *decimal.Decimal `json:"floorPrice,omitempty"`
	SecurityStatus                  string           `json:"securityStatus,omitempty"`
	SymbolAdminStatusCode           string           `json:"symbolAdminStatusCode,omitempty"`
	SymbolTradingMethodStatusCode   string           `json:"symbolTradingMethodStatusCode,omitempty"`
	SymbolTradingSanctionStatusCode string           `json:"symbolTradingSanctionStatusCode,omitempty"`
	FinalTradeDate                  string           `json:"finalTradeDate,omitempty"`
	ListingDate                     string           `json:"listingDate,omitempty"`
	Time                            string           `json:"time,omitempty"`
}

type WorkingDatesResponse struct {
	WorkingDates []string `json:"workingDates,omitempty"`
}

type GetWorkingDatesResponse = WorkingDatesResponse

type CareByResponse struct {
	CareBy []CareByAccount `json:"careBy,omitempty"`
	Total  int             `json:"total,omitempty"`
}

type CareByAccount struct {
	AccountNo         string             `json:"accountNo,omitempty"`
	FullName          string             `json:"fullName,omitempty"`
	CustodyCode       string             `json:"custodyCode,omitempty"`
	UnderlyingNAV     *decimal.Decimal   `json:"underlyingNav,omitempty"`
	DerivativeNAV     *decimal.Decimal   `json:"derivativeNav,omitempty"`
	TotalNAV          *decimal.Decimal   `json:"totalNav,omitempty"`
	DealAccount       bool               `json:"dealAccount,omitempty"`
	DerivativeAccount bool               `json:"derivativeAccount,omitempty"`
	Derivative        *DerivativeAccount `json:"derivative,omitempty"`
	Permissions       []Permission       `json:"permissions,omitempty"`
}

type Permission struct {
	Product string `json:"product,omitempty"`
	Role    string `json:"role,omitempty"`
}
