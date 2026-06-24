// Code generated from Finhay OpenAPI v2; DO NOT EDIT.

package dto

type DebtSummary struct {
	// Total debt
	Total float64 `json:"total,omitempty"`
	// Secured loan amount
	SecureAmount float64 `json:"secure_amount,omitempty"`
	// Advanced amount
	AdvanceAmt float64 `json:"advance_amt,omitempty"`
	// SMS fee amount
	SmsFeeAmt float64 `json:"sms_fee_amt,omitempty"`
	// CIDEPO fee accrual
	CidepoFeeAcr float64 `json:"cidepo_fee_acr,omitempty"`
	// Deposit owed
	OweDeposit float64 `json:"owe_deposit,omitempty"`
}

type GetPortfolioResponse struct {
	Portfolio []PortfolioEntry `json:"portfolio,omitempty"`
	// Whether the data is from a cached snapshot
	IsSnapshot bool `json:"is_snapshot,omitempty"`
}

type MoneySummary struct {
	// Total cash
	Total float64 `json:"total,omitempty"`
	// Cash balance in account
	CiBalance float64 `json:"ci_balance,omitempty"`
	// Dividend cash pending
	CaReceiving float64 `json:"ca_receiving,omitempty"`
	// Other blocked funds
	EmkAmt float64 `json:"emk_amt,omitempty"`
	// Cash to receive
	ReceivingAmt float64 `json:"receiving_amt,omitempty"`
	// Available withdrawable amount
	Baldefovd float64 `json:"baldefovd,omitempty"`
}

type OrderHistoryEntry struct {
	// Order ID
	OrderID   string `json:"order_id,omitempty"`
	AccountID string `json:"account_id,omitempty"`
	// Transaction date
	TransactionDate string `json:"transaction_date,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	// Order side Enum: BUY, SELL, UNKNOWN
	OrderSide string `json:"order_side,omitempty"`
	// Ordered quantity
	OrderQuantity int64 `json:"order_quantity,omitempty"`
	// Limit (quote) price
	LimitPrice int64 `json:"limit_price,omitempty"`
	// Market price at order time
	MarketPrice string `json:"market_price,omitempty"`
	// Executed quantity
	ExecuteQuantity int64 `json:"execute_quantity,omitempty"`
	// Executed price
	ExecutePrice int64 `json:"execute_price,omitempty"`
	// Enum: WAITING_TO_ACTIVATE, WAITING_TO_SEND, SENDING, SENT, FIXING, FIXED, CANCELLED, CANCELLING, MATCHED, EXPIRED, MATCHED_ALL, REJECTED, COMPLETED, FAILED, EXPIRED_ACTIVATION_TIME, EXPIRED_AUTHORIZATION, RECEIVED, INTERNAL_SENDING, INTERNAL_CANCELLED, SENDING_TO_CORE, CANCEL_BY_STOCK_EVENT, WAIT_FOR_ACCEPTING, ACCEPTED, ACCEPTING, REJECTING, UNKNOWN
	OrderStatus string `json:"order_status,omitempty"`
	// Fee amount
	FeeAmount int64 `json:"fee_amount,omitempty"`
	// Tax amount
	TaxAmount int64 `json:"tax_amount,omitempty"`
	// Executed amount
	ExecuteAmount int64 `json:"execute_amount,omitempty"`
	// Order type Enum: LO, MP, ATO, ATC, MAK, MOK, MTL, PLO, FOK, FAK
	OrderType string `json:"order_type,omitempty"`
}

type OrderHistoryResult struct {
	Data []OrderHistoryEntry `json:"data,omitempty"`
}

type PnLEntry struct {
	// Profit/loss amount
	PnL float64 `json:"pnl,omitempty"`
	// Profit/loss rate
	PnLRate float64 `json:"pnl_rate,omitempty"`
}

type PnLSummary struct {
	Stock PnLEntry `json:"stock,omitempty"`
	Fund  PnLEntry `json:"fund,omitempty"`
}

type PnLTodayResponse struct {
	// Sub-account ID or `ALL`
	SubAccountID string `json:"sub_account_id,omitempty"`
	// Today's PnL amount (VND)
	PnLAmount int64 `json:"pnl_amount,omitempty"`
	// Today's PnL rate (percentage)
	PnLRate float64 `json:"pnl_rate,omitempty"`
}

type PortfolioEntry struct {
	SubAccountID string `json:"sub_account_id,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	// Type of securities
	SecuritiesType string `json:"securities_type,omitempty"`
	// Total shares held
	Total int64 `json:"total,omitempty"`
	// Shares available to sell
	Available int64 `json:"available,omitempty"`
	// Blocked shares
	Blocked int64 `json:"blocked,omitempty"`
	// Mortgaged shares
	Mortgage int64 `json:"mortgage,omitempty"`
	// VSD mortgaged shares
	VsdMortgage int64 `json:"vsd_mortgage,omitempty"`
	// Restricted shares
	Restrict       int64 `json:"restrict,omitempty"`
	ReceivingRight int64 `json:"receiving_right,omitempty"`
	// Receiving shares T+0
	ReceivingT0 int64 `json:"receiving_t0,omitempty"`
	// Receiving shares T+1
	ReceivingT1 int64 `json:"receiving_t1,omitempty"`
	// Receiving shares T+2
	ReceivingT2    int64 `json:"receiving_t2,omitempty"`
	MatchingAmount int64 `json:"matching_amount,omitempty"`
	Withdraw       int64 `json:"withdraw,omitempty"`
	// Average cost price
	CostPrice int64 `json:"cost_price,omitempty"`
	// Basic (reference) price
	BasicPrice int64 `json:"basic_price,omitempty"`
	// Whether the stock can be sold
	IsSellable bool   `json:"is_sellable,omitempty"`
	Custodycd  string `json:"custodycd,omitempty"`
	// Profit/Loss amount
	PnLAmount int64 `json:"pnl_amount,omitempty"`
	// Profit/Loss rate (percentage)
	PnLRate float64 `json:"pnl_rate,omitempty"`
	// Close (current) price
	ClosePrice int64 `json:"close_price,omitempty"`
	// Total cost value
	CostPriceAmount int64 `json:"cost_price_amount,omitempty"`
	// Total basic value
	BasicPriceAmount int64 `json:"basic_price_amount,omitempty"`
	// Total PnL
	TotalPnL  int64 `json:"total_pnl,omitempty"`
	SendingT0 int64 `json:"sending_t0,omitempty"`
	SendingT1 int64 `json:"sending_t1,omitempty"`
	SendingT2 int64 `json:"sending_t2,omitempty"`
	// Whether the stock has recent news
	HasNewestNews bool `json:"has_newest_news,omitempty"`
	// Deprecated field
	Trade int64 `json:"trade,omitempty"`
}

type ProductsSummary struct {
	// Total product value
	Total float64 `json:"total,omitempty"`
	// Stock value
	Stock float64 `json:"stock,omitempty"`
	// Fund value
	Fund float64 `json:"fund,omitempty"`
	// Saving amount
	Saving *float64 `json:"saving,omitempty"`
	// HayBond value
	Bond float64 `json:"bond,omitempty"`
	// Hay0 NAV
	Hay0 float64 `json:"hay0,omitempty"`
	// Hay0 interest
	Hay0Interest float64 `json:"hay0_interest,omitempty"`
	// Hay0 amount being deposited
	Hay0Depositing float64 `json:"hay0_depositing,omitempty"`
	// Hay0 amount being withdrawn
	Hay0Withdrawing float64 `json:"hay0_withdrawing,omitempty"`
}

type SummaryAccountResponse struct {
	// Available balance
	Balance         int64 `json:"balance,omitempty"`
	CiBalance       int64 `json:"ci_balance,omitempty"`
	TdBalance       int64 `json:"td_balance,omitempty"`
	InterestBalance int64 `json:"interest_balance,omitempty"`
	CaReceiving     int64 `json:"ca_receiving,omitempty"`
	// Receiving amount T+1
	ReceivingT1 int64 `json:"receiving_t1,omitempty"`
	// Receiving amount T+2
	ReceivingT2 int64 `json:"receiving_t2,omitempty"`
	// Receiving amount T+3
	ReceivingT3 int64 `json:"receiving_t3,omitempty"`
	// Securities amount
	SecuritiesAmt int64 `json:"securities_amt,omitempty"`
	// Total debt amount
	TotalDebtAmt int64 `json:"total_debt_amt,omitempty"`
	SecureAmt    int64 `json:"secure_amt,omitempty"`
	TrfBuyAmt    int64 `json:"trf_buy_amt,omitempty"`
	// Margin amount
	MarginAmt int64 `json:"margin_amt,omitempty"`
	// T+0 debt amount
	T0DebtAmt int64 `json:"t0_debt_amt,omitempty"`
	// Advanced amount
	AdvancedAmt  int64 `json:"advanced_amt,omitempty"`
	DfDebtAmt    int64 `json:"df_debt_amt,omitempty"`
	TdDebtAmt    int64 `json:"td_debt_amt,omitempty"`
	CidepoFeeAcr int64 `json:"cidepo_fee_acr,omitempty"`
	// Net asset value
	NetAssetValue int64 `json:"net_asset_value,omitempty"`
	// Margin credit limit
	MrcrLimit        int64 `json:"mrcr_limit,omitempty"`
	DebtAmt          int64 `json:"debt_amt,omitempty"`
	AdvanceMaxAmtFee int64 `json:"advance_max_amt_fee,omitempty"`
	ReceivingAmt     int64 `json:"receiving_amt,omitempty"`
	// Margin rate
	MarginRate    int64 `json:"margin_rate,omitempty"`
	SmsFeeAmt     int64 `json:"sms_fee_amt,omitempty"`
	IbrokerFeeAmt int64 `json:"ibroker_fee_amt,omitempty"`
	HoldBalance   int64 `json:"hold_balance,omitempty"`
	MriRate       int64 `json:"mri_rate,omitempty"`
	MrmRate       int64 `json:"mrm_rate,omitempty"`
	CidepoFee     int64 `json:"cidepo_fee,omitempty"`
	TdIntAmt      int64 `json:"td_int_amt,omitempty"`
	AddVND        int64 `json:"add_vnd,omitempty"`
	AddVND1       int64 `json:"add_vnd_1,omitempty"`
	// Core bank code
	CoreBank string `json:"core_bank,omitempty"`
	// Bank account number
	BankacctNo string `json:"bankacct_no,omitempty"`
	// Bank name
	BankName  string `json:"bank_name,omitempty"`
	EmkAmt    int64  `json:"emk_amt,omitempty"`
	Baldefovd string `json:"baldefovd,omitempty"`
	// Maximum margin credit limit
	MrcrLimitMax int64 `json:"mrcr_limit_max,omitempty"`
}

type UserAssetsSummaryResponse struct {
	// Total net asset value
	NetAssetValue float64         `json:"net_asset_value,omitempty"`
	Products      ProductsSummary `json:"products,omitempty"`
	Money         MoneySummary    `json:"money,omitempty"`
	Debt          DebtSummary     `json:"debt,omitempty"`
	PnL           PnLSummary      `json:"pnl,omitempty"`
}

type UserRight struct {
	AccountID        string `json:"accountId,omitempty"`
	DepositoryNumber string `json:"depositoryNumber,omitempty"`
	FullName         string `json:"fullName,omitempty"`
	// Number of shares owned at record date
	OwnNumberOfShare int64 `json:"ownNumberOfShare,omitempty"`
	// Right ratio (e.g. "1:5")
	Ratio string `json:"ratio,omitempty"`
	// Corporate action status Enum: WAIT_FOR_REVIEW, WAIT_EXECUTED, WAITING_FOR_APPROVE, NAVIGATING, COMPLETED, ALLOCATION_COMPLETED, CANCELED, EMPTY_STOCK, APPROVED, REJECTED, READY_FOR_TRADE, READY_FOR_CLOSE, STOCK_ALLOTTED, MONEY_ALLOTTED, PARTIALLY_DONE, VERIFIED, DELETED, REGISTERED
	Status string `json:"status,omitempty"`
	// User registration status Enum: UNREGISTER, REGISTERED, EXPIRED, RECEIVED, PARTIAL_REGISTERED, PENDING, PARTIAL_RECEIVED, WAIT_REGISTER, REGISTERED_V2, WAIT_RECEIVE, WAIT_STOCK, PARTIAL_RECEIVED_V2, UNKNOWN
	UserRightRegisterStatus string `json:"userRightRegisterStatus,omitempty"`
	// Corporate action master ID
	CaMastID string `json:"caMastId,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	// Target symbol (for stock conversion events)
	ToSymbol               *string `json:"toSymbol,omitempty"`
	NumberOfWaitingStock   int64   `json:"numberOfWaitingStock,omitempty"`
	WaitingAmountForReturn int64   `json:"waitingAmountForReturn,omitempty"`
	// Right offering rate
	RightOffRate string `json:"rightOffRate,omitempty"`
	// Subscription price (for stock rights)
	BuyPrice int64 `json:"buyPrice,omitempty"`
	// Whether registration is allowed
	AllowRegister     bool   `json:"allowRegister,omitempty"`
	TotalStocksCanBuy int64  `json:"totalStocksCanBuy,omitempty"`
	TotalPayAmount    int64  `json:"totalPayAmount,omitempty"`
	StartDate         string `json:"startDate,omitempty"`
	FinishDate        string `json:"finishDate,omitempty"`
	StartDateTransfer string `json:"startDateTransfer,omitempty"`
	// Record date / Ex-date
	ActionDate         string `json:"actionDate,omitempty"`
	FinishDateTransfer string `json:"finishDateTransfer,omitempty"`
	// Used for sorting (descending)
	ReportDate string `json:"reportDate,omitempty"`
	// Type of corporate action Enum: OTC_BOND_INTEREST, SHAREHOLDER_MEETING, SOLICIT_SHAREHOLDER_OPINIONS, CASH_DIVIDEND, STOCK_DIVIDEND, STOCK_RIGHT, BOND_INTEREST, BOND_PRINCIPAL_AND_INTEREST, CONVERT_BONDS_TO_STOCKS, CONVERT_STOCKS_TO_OTHER_STOCKS, BONUS_SHARES, VOTING_RIGHTS, CONVERTIBLE_BOND, TRANSFER_PENDING_STOCKS, WARRANT_DIVIDENDS
	TypeValue      string `json:"type,omitempty"`
	StockName      string `json:"stockName,omitempty"`
	TotalVolume    int64  `json:"totalVolume,omitempty"`
	TotalValue     int64  `json:"totalValue,omitempty"`
	ReferencePrice int64  `json:"referencePrice,omitempty"`
	ClosePrice     int64  `json:"closePrice,omitempty"`
	CeilingPrice   int64  `json:"ceilingPrice,omitempty"`
	FloorPrice     int64  `json:"floorPrice,omitempty"`
	PredictPrice   int64  `json:"predictPrice,omitempty"`
	ChangePrice    int64  `json:"changePrice,omitempty"`
	// Whether the user has read this right notification
	HasRead bool `json:"hasRead,omitempty"`
}
