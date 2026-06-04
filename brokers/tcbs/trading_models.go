package tcbs

type PlaceOrderRequest struct {
	ExecType  string `json:"execType"`
	Price     int    `json:"price"`
	PriceType string `json:"priceType"`
	Quantity  int    `json:"quantity"`
	Symbol    string `json:"symbol"`
}

type PlaceOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type UpdateOrderRequest struct {
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}

type UpdateOrderResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type CancelOrderRequest struct {
	OrdersList []OrderIDRef `json:"ordersList"`
}

type OrderIDRef struct {
	OrderID string `json:"orderID"`
}

type CancelOrderResponse struct {
	Object     string         `json:"object"`
	PageSize   int            `json:"pageSize"`
	PageIndex  int            `json:"pageIndex"`
	TotalCount int            `json:"totalCount"`
	Data       []CancelResult `json:"data"`
}

type CancelResult struct {
	Object  string              `json:"object"`
	Details []CancelOrderDetail `json:"details"`
}

type CancelOrderDetail struct {
	Deleted      string `json:"deleted"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMesage"`
	OrderID      string `json:"orderID"`
}

type OrderSearchResponse struct {
	Object     string      `json:"object"`
	PageSize   int         `json:"pageSize"`
	PageIndex  int         `json:"pageIndex"`
	TotalCount int         `json:"totalCount"`
	Data       []OrderInfo `json:"data"`
}

type OrderInfo struct {
	Object      string  `json:"object"`
	AccountNo   string  `json:"accountNo"`
	OrderID     string  `json:"orderID"`
	ExecType    string  `json:"execType"`
	OrderQtty   float64 `json:"orderQtty"`
	ExecQtty    float64 `json:"execQtty"`
	CodeID      string  `json:"codeID"`
	Symbol      string  `json:"symbol"`
	PriceType   string  `json:"priceType"`
	TxTime      string  `json:"txtime"`
	TxDate      string  `json:"txdate"`
	ExpDate     string  `json:"expDate"`
	TimeType    string  `json:"timeType"`
	OrStatus    string  `json:"orStatus"`
	FeeAcr      float64 `json:"feeAcr"`
	LimitPrice  float64 `json:"limitPrice"`
	CancelQtty  float64 `json:"cancelQtty"`
	RemainQtty  float64 `json:"remainQtty"`
	Via         string  `json:"via"`
	QuotePrice  float64 `json:"quotePrice"`
	MatchPrice  float64 `json:"matchPrice"`
	TradePlace  string  `json:"tradePlace"`
	MatchType   string  `json:"matchType"`
	IsDisposal  string  `json:"isDisposal"`
	IsCancel    string  `json:"isCancel"`
	IsAmend     string  `json:"isAmend"`
	UserName    string  `json:"userName"`
	OrsOrderID  string  `json:"orsOrderID"`
	SecType     string  `json:"sectype"`
	IsFOOrder   string  `json:"isFOOrder"`
	OdTimeStamp string  `json:"odTimeStamp"`
	MatchAmount float64 `json:"matchAmount"`
	MMType      string  `json:"mmType"`
	BRatio      float64 `json:"bRatio"`
	TaxSellAmt  float64 `json:"taxSellAmout"`
}

type CommandMatchInformationResponse struct {
	Object     string                          `json:"object"`
	TotalCount int                             `json:"totalCount"`
	PageSize   int                             `json:"pageSize"`
	PageIndex  int                             `json:"pageIndex"`
	Data       []CommandMatchInformationDetail `json:"data"`
}

type CommandMatchInformationDetail struct {
	OrderID    string  `json:"orderId"`
	Side       string  `json:"side"`
	Symbol     string  `json:"symbol"`
	QuoteQtty  float64 `json:"quoteQtty"`
	QuotePrice float64 `json:"quotePrice"`
	TradeID    string  `json:"tradeId"`
	Qtty       float64 `json:"qtty"`
	Price      float64 `json:"price"`
	TimeExec   float64 `json:"timeExec"`
}

type PurchasingPowerResponse struct {
	Object          string  `json:"object"`
	AccountNo       string  `json:"accountNo"`
	CustodyID       string  `json:"custodyID"`
	Symbol          string  `json:"symbol"`
	Price           float64 `json:"price"`
	PP0             float64 `json:"pp0"`
	PPSE            float64 `json:"ppse"`
	AvailableTrade  float64 `json:"availableTrade"`
	PPSERef         float64 `json:"ppseref"`
	MaxBuyQuantity  float64 `json:"maxBuyQuantity"`
	RealMaxBuyQty   float64 `json:"realMaxBuyQuantity"`
	MinBuyQuantity  float64 `json:"minBuyQuantity"`
	MarginRatioLoan string  `json:"marginRatioLoan"`
	MarginPriceLoan string  `json:"marginPriceLoan"`
	RateBrkS        float64 `json:"rateBrkS"`
	RateBrkB        float64 `json:"rateBrkB"`
}

type MarginQuotaResponse struct {
	CustodyID     string  `json:"custodyID"`
	AccountNo     string  `json:"accountNo"`
	AFType        string  `json:"aftype"`
	VSDStatus     string  `json:"vsdStatus"`
	AccountStatus string  `json:"accountStatus"`
	MarginLimit   float64 `json:"marginLimit"`
	IsIA          string  `json:"isIA"`
	BankName      string  `json:"bankName"`
	BankAccount   string  `json:"bankAccount"`
	AccountType   string  `json:"accountType"`
}

type MarginAccountInfoResponse struct {
	AccountNo       string      `json:"accountNo"`
	RiskPolicy      *RiskPolicy `json:"riskPolicy,omitempty"`
	RTT             float64     `json:"rtt"`
	Outstanding     float64     `json:"outstanding"`
	AccruedInterest float64     `json:"accruedInterest"`
	DueAmount       float64     `json:"dueAmount"`
	OverdueAmount   float64     `json:"overdueAmount"`
	RiskStatus      *RiskStatus `json:"riskStatus,omitempty"`
	TotalFeeDebt    float64     `json:"totalFeeDebt"`
}

type RiskPolicy struct {
	MaintenanceMargin float64 `json:"maintenanceMargin"`
	InitialMargin     float64 `json:"initialMargin"`
	LiquidationMargin float64 `json:"liquidationMargin"`
}

type RiskStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type SeInfoDTO struct {
	Object    string             `json:"object"`
	AccountNo string             `json:"accountNo"`
	CustodyID string             `json:"custodyID"`
	FullName  string             `json:"fullName"`
	Stock     []StockHoldingInfo `json:"stock"`
}

type StockHoldingInfo struct {
	Symbol           string  `json:"symbol"`
	SecType          string  `json:"secType"`
	SecTypeName      string  `json:"secTypeName"`
	AvailableTrading float64 `json:"availableTrading"`
	Mortgaged        float64 `json:"mortgaged"`
	T0               float64 `json:"t0"`
	T1               float64 `json:"t1"`
	T2               float64 `json:"t2"`
	Blocked          float64 `json:"blocked"`
	SecuredQuantity  float64 `json:"securedQuantity"`
	SellRemain       float64 `json:"sellRemain"`
	ExercisedCA      float64 `json:"exercisedCA"`
	UnexercisedCA    float64 `json:"unexercisedCA"`
	StockDividend    float64 `json:"stockDividend"`
	CashDividend     float64 `json:"cashDividend"`
	WaitForTrade     float64 `json:"waitForTrade"`
	WaitForTransfer  float64 `json:"waitForTransfer"`
	WaitForWithdraw  float64 `json:"waitForWithdraw"`
	CurrentPrice     float64 `json:"currentPrice"`
	CostPrice        float64 `json:"costPrice"`
	SellExec         float64 `json:"sellExec"`
	OnHold           float64 `json:"onHold"`
	TotalQtty        float64 `json:"totalQtty"`
	Settlement       float64 `json:"settlement"`
}

type CashInvestmentResponse struct {
	Object     string           `json:"object"`
	TotalCount int              `json:"totalCount"`
	PageSize   int              `json:"pageSize"`
	PageIndex  int              `json:"pageIndex"`
	Data       []CashInvestment `json:"data"`
}

type CashInvestment struct {
	Object           string   `json:"object"`
	IAInfos          []IAInfo `json:"iaInfos"`
	PP0ForBF         float64  `json:"pp0forBF"`
	BankAvlBalanceBF float64  `json:"bankAvlBalanceBF"`
	BodBalance       float64  `json:"bodBalance"`
	CashBalance      float64  `json:"cashBalance"`
	AccountNo        string   `json:"accountNo"`
	CustodyID        string   `json:"custodyID"`
	FullName         string   `json:"fullName"`
	Balance          float64  `json:"balance"`
	AvlAdvanceAmount float64  `json:"avlAdvanceAmount"`
	BuyingAmount     float64  `json:"buyingAmount"`
	BlockAmount      float64  `json:"blockAmount"`
	CashDividend     float64  `json:"cashDevident"`
	BankAvlBalance   float64  `json:"bankAvlBalance"`
	BankBlockAmount  float64  `json:"bankBlockAmount"`
	AvlWithdraw      float64  `json:"avlWithdraw"`
	PP0              float64  `json:"pp0"`
	SecureAmtPO      float64  `json:"secureAmtPO"`
	BondBlockAmount  float64  `json:"bondBlockAmount"`
	MBlockAmount     float64  `json:"mBlockAmount"`
	FundBlockAmount  float64  `json:"fundBlockAmount"`
	AvalBondBlock    float64  `json:"avalBondBlockAmount"`
	DepoFee          float64  `json:"depoFee"`
	BCashDividend    float64  `json:"bCashDividend"`
	SCashDividend    float64  `json:"sCashDividend"`
	DSecured         float64  `json:"dsecured"`
	AdUsed           float64  `json:"adused"`
	MrUsed           float64  `json:"mrused"`
}

type IAInfo struct {
	Partner   string  `json:"partner"`
	Available float64 `json:"available"`
	Hold      float64 `json:"hold"`
}

type SupplementaryLoanPackageResponse struct {
	MarginSureViews []MarginSureView `json:"marginSureViews"`
	TPlus           *TPlusData       `json:"tplus,omitempty"`
}

type MarginSureView struct {
	ID              float64              `json:"id"`
	Name            string               `json:"name"`
	Code            string               `json:"code"`
	SubscriptionFee float64              `json:"subscriptionFee"`
	Status          string               `json:"status"`
	Proposals       []MarginSureProposal `json:"proposals"`
	Default         bool                 `json:"default"`
}

type MarginSureProposal struct {
	ID                       float64 `json:"id"`
	MarginInsuranceID        float64 `json:"marginInsuranceId"`
	InterestAdjustmentValue  float64 `json:"interestAdjustmentValue"`
	InterestPercentThreshold float64 `json:"interestPercentThreshold"`
	ThresholdType            string  `json:"thresholdType"`
}

type TPlusData struct {
	Data []TPlusPackage `json:"data"`
}

type TPlusPackage struct {
	FirstRate                                 float64       `json:"firstRate"`
	ID                                        float64       `json:"id"`
	Name                                      string        `json:"name"`
	Status                                    string        `json:"status"`
	UndueInterestType                         string        `json:"undueInterestType"`
	UndueLadderValue                          []TPlusLadder `json:"undueLadderValue"`
	OverdueInterest                           float64       `json:"overdueInterest"`
	ExtensionInterest                         float64       `json:"extensionInterest"`
	ExtensionInterestBeforeInterestSettlement float64       `json:"extensionInterestBeforeInterestSettlement"`
	InterestCalculationBasis                  float64       `json:"interestCalculationBasis"`
	UndueFee                                  float64       `json:"undueFee"`
	OverdueFee                                float64       `json:"overdueFee"`
	ExtensionFee                              float64       `json:"extensionFee"`
	DebtCollectionFee                         float64       `json:"debtCollectionFee"`
	Description                               string        `json:"description"`
	ValidFrom                                 string        `json:"validFrom"`
}

type TPlusLadder struct {
	ID        float64 `json:"id"`
	Rate      float64 `json:"rate"`
	StartDate float64 `json:"startDate"`
	DueDate   float64 `json:"dueDate"`
}

type LoanResponse struct {
	Size    int        `json:"size"`
	Content []LoanItem `json:"content"`
}

type LoanItem struct {
	OpeningDate        string   `json:"openingDate"`
	DueDate            string   `json:"dueDate"`
	RenewTime          int      `json:"renewTime"`
	MaxRenewTime       int      `json:"maxRenewTime"`
	IsRenewable        bool     `json:"isRenewable"`
	ReasonList         []string `json:"reasonList"`
	Symbol             string   `json:"symbol"`
	ID                 float64  `json:"id"`
	AccountNo          string   `json:"accountNo"`
	Principal          float64  `json:"principal"`
	RemainingPrincipal float64  `json:"remainingPrincipal"`
	Interest           float64  `json:"interest"`
	Rate               float64  `json:"rate"`
	Status             string   `json:"status"`
	LoanDays           int      `json:"loanDays"`
	MrxLoanID          float64  `json:"mrxLoanId"`
	Fee                float64  `json:"fee"`
	UndueLoanFee       float64  `json:"undueLoanFee"`
	PricingPolicyType  string   `json:"pricingPolicyType"`
}
