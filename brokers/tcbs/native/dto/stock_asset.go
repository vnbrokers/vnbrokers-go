// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

type MarginQuota struct {
	AccountNo     string  `json:"accountNo"`
	AccountStatus string  `json:"accountStatus"`
	AccountType   string  `json:"accountType"`
	Aftype        string  `json:"aftype"`
	BankAccount   string  `json:"bankAccount"`
	BankName      string  `json:"bankName"`
	CustodyID     string  `json:"custodyID"`
	IsIA          string  `json:"isIA"`
	MarginLimit   float64 `json:"marginLimit"`
	VSDStatus     string  `json:"vsdStatus"`
}

type MarginAccountInformation struct {
	AccountNo       string                             `json:"accountNo"`
	AccruedInterest float64                            `json:"accruedInterest"`
	DueAmount       float64                            `json:"dueAmount"`
	Outstanding     float64                            `json:"outstanding"`
	OverdueAmount   float64                            `json:"overdueAmount"`
	RiskPolicy      MarginAccountInformationRiskPolicy `json:"riskPolicy"`
	RiskStatus      MarginAccountInformationRiskStatus `json:"riskStatus"`
	RTT             float64                            `json:"rtt"`
	TotalFeeDebt    float64                            `json:"totalFeeDebt"`
}

type MarginAccountInformationRiskPolicy struct {
	InitialMargin     float64 `json:"initialMargin"`
	LiquidationMargin float64 `json:"liquidationMargin"`
	MaintenanceMargin float64 `json:"maintenanceMargin"`
}

type MarginAccountInformationRiskStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type SupplementaryLoanPackages struct {
	MarginSureViews []MarginSurePackage `json:"marginSureViews"`
	Tplus           TPlusPackages       `json:"tplus"`
}

type MarginSurePackage struct {
	Code            string               `json:"code"`
	Default         bool                 `json:"default"`
	ID              float64              `json:"id"`
	Name            string               `json:"name"`
	Proposals       []MarginSureProposal `json:"proposals"`
	Status          string               `json:"status"`
	SubscriptionFee float64              `json:"subscriptionFee"`
}

type MarginSureProposal struct {
	ID                       float64 `json:"id"`
	InterestAdjustmentValue  float64 `json:"interestAdjustmentValue"`
	InterestPercentThreshold float64 `json:"interestPercentThreshold"`
	MarginInsuranceID        float64 `json:"marginInsuranceId"`
	ThresholdType            string  `json:"thresholdType"`
}

type TPlusPackages struct {
	Data []TPlusPackage `json:"data"`
}

type TPlusPackage struct {
	CreatedAt                                 string             `json:"createdAt"`
	DebtCollectionFee                         float64            `json:"debtCollectionFee"`
	Description                               string             `json:"description"`
	Discountable                              bool               `json:"discountable"`
	ExtensionFee                              float64            `json:"extensionFee"`
	ExtensionInterest                         float64            `json:"extensionInterest"`
	ExtensionInterestBeforeInterestSettlement float64            `json:"extensionInterestBeforeInterestSettlement"`
	FirstRate                                 float64            `json:"firstRate"`
	ID                                        float64            `json:"id"`
	InterestCalculationBasis                  float64            `json:"interestCalculationBasis"`
	IwealthPartnerInterest                    float64            `json:"iwealthPartnerInterest"`
	MinRateConstraint                         bool               `json:"minRateConstraint"`
	Name                                      string             `json:"name"`
	OverdueFee                                float64            `json:"overdueFee"`
	OverdueInterest                           float64            `json:"overdueInterest"`
	Status                                    string             `json:"status"`
	UndueFee                                  float64            `json:"undueFee"`
	UndueInterestType                         string             `json:"undueInterestType"`
	UndueLadderValue                          []InterestRateTier `json:"undueLadderValue"`
	UpdatedAt                                 string             `json:"updatedAt"`
	ValidFrom                                 string             `json:"validFrom"`
}

type InterestRateTier struct {
	DueDate   float64 `json:"dueDate"`
	ID        float64 `json:"id"`
	Rate      float64 `json:"rate"`
	StartDate float64 `json:"startDate"`
}

type Loans struct {
	Content []Loan  `json:"content"`
	Size    float64 `json:"size"`
}

type Loan struct {
	AccountNo          string        `json:"accountNo"`
	Addons             []LoanAddon   `json:"addons"`
	DueDate            string        `json:"dueDate"`
	Fee                float64       `json:"fee"`
	ID                 float64       `json:"id"`
	Insurance          LoanInsurance `json:"insurance"`
	Interest           float64       `json:"interest"`
	IsRenewable        bool          `json:"isRenewable"`
	LoanDays           float64       `json:"loanDays"`
	MaxRenewTime       float64       `json:"maxRenewTime"`
	MrxLoanID          float64       `json:"mrxLoanId"`
	OpeningDate        string        `json:"openingDate"`
	PricingPolicyType  string        `json:"pricingPolicyType"`
	Principal          float64       `json:"principal"`
	Rate               float64       `json:"rate"`
	ReasonList         []string      `json:"reasonList"`
	RemainingPrincipal float64       `json:"remainingPrincipal"`
	RenewTime          float64       `json:"renewTime"`
	Status             string        `json:"status"`
	Symbol             string        `json:"symbol"`
	UndueLoanFee       float64       `json:"undueLoanFee"`
}

type LoanAddon struct {
	Discounts []LoanDiscount `json:"discounts"`
	LoanID    float64        `json:"loanId"`
	Others    []LoanPolicy   `json:"others"`
}

type LoanDiscount struct {
	Code                  string  `json:"code"`
	EndDate               string  `json:"endDate"`
	OrigDiscountLoanAmt   float64 `json:"origDiscountLoanAmt"`
	PreferentialRate      float64 `json:"preferentialRate"`
	RemainDiscountLoanAmt float64 `json:"remainDiscountLoanAmt"`
	StartDate             string  `json:"startDate"`
}

type LoanPolicy struct {
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Rate        float64  `json:"rate"`
	Symbols     []string `json:"symbols"`
}

type LoanInsurance struct {
	AccountNo                        string  `json:"accountNo"`
	AdjustStrikePrice                float64 `json:"adjustStrikePrice"`
	DecreasedInterest                float64 `json:"decreasedInterest"`
	DecreasedPrice                   float64 `json:"decreasedPrice"`
	EndDate                          string  `json:"endDate"`
	IncreasePrice                    float64 `json:"increasePrice"`
	IncreasedInterest                float64 `json:"increasedInterest"`
	InsuranceCode                    string  `json:"insuranceCode"`
	InsuranceFee                     float64 `json:"insuranceFee"`
	InsuranceName                    string  `json:"insuranceName"`
	InterestPercentageLowerThreshold float64 `json:"interestPercentageLowerThreshold"`
	InterestPercentageUpperThreshold float64 `json:"interestPercentageUpperThreshold"`
	InterestRateLowerAdjustmentValue float64 `json:"interestRateLowerAdjustmentValue"`
	InterestRateUpperAdjustmentValue float64 `json:"interestRateUpperAdjustmentValue"`
	LoanID                           float64 `json:"loanId"`
	OriginalInterest                 float64 `json:"originalInterest"`
	StartDate                        string  `json:"startDate"`
	Symbol                           string  `json:"symbol"`
}

type StockPurchasingPower struct {
	AccountNo          string  `json:"accountNo"`
	CustodyID          string  `json:"custodyID"`
	MarginPriceLoan    float64 `json:"marginPriceLoan"`
	MarginRatioLoan    float64 `json:"marginRatioLoan"`
	MaxBuyQuantity     float64 `json:"maxBuyQuantity"`
	MinBuyQuantity     float64 `json:"minBuyQuantity"`
	PP0                float64 `json:"pp0"`
	PPSE               float64 `json:"ppse"`
	PPSERef            float64 `json:"ppseref"`
	Price              float64 `json:"price"`
	RateBrkB           string  `json:"rateBrkB"`
	RateBrkS           string  `json:"rateBrkS"`
	RealMaxBuyQuantity float64 `json:"realMaxBuyQuantity"`
	Symbol             string  `json:"symbol"`
}

type StockAssets struct {
	AccountNo string             `json:"accountNo"`
	CustodyID string             `json:"custodyID"`
	FullName  string             `json:"fullName"`
	Object    string             `json:"object"`
	Stock     []StockHoldingInfo `json:"stock"`
}

type StockHoldingInfo struct {
	AvailableTrading float64 `json:"availableTrading"`
	Blocked          float64 `json:"blocked"`
	CashDividend     float64 `json:"cashDividend"`
	CostPrice        float64 `json:"costPrice"`
	CurrentPrice     float64 `json:"currentPrice"`
	ExercisedCA      float64 `json:"exercisedCA"`
	Mortgaged        float64 `json:"mortgaged"`
	OnHold           float64 `json:"onHold"`
	SecType          string  `json:"secType"`
	SecTypeName      string  `json:"secTypeName"`
	SecuredQuantity  float64 `json:"securedQuantity"`
	SellExec         float64 `json:"sellExec"`
	SellRemain       float64 `json:"sellRemain"`
	Settlement       float64 `json:"settlement"`
	StockDividend    float64 `json:"stockDividend"`
	Symbol           string  `json:"symbol"`
	T0               float64 `json:"t0"`
	T1               float64 `json:"t1"`
	T2               float64 `json:"t2"`
	TotalQtty        float64 `json:"totalQtty"`
	UnexercisedCA    float64 `json:"unexercisedCA"`
	WaitForTrade     float64 `json:"waitForTrade"`
	WaitForTransfer  float64 `json:"waitForTransfer"`
	WaitForWithdraw  float64 `json:"waitForWithdraw"`
}

type CashInvestments struct {
	Data       []CashInvestment `json:"data"`
	Object     string           `json:"object"`
	PageIndex  int64            `json:"pageIndex"`
	PageSize   int64            `json:"pageSize"`
	TotalCount int64            `json:"totalCount"`
}

type CashInvestment struct {
	AccountNo           string                  `json:"accountNo"`
	Adused              float64                 `json:"adused"`
	AvalBondBlockAmount float64                 `json:"avalBondBlockAmount"`
	AvlAdvanceAmount    float64                 `json:"avlAdvanceAmount"`
	AvlWithdraw         float64                 `json:"avlWithdraw"`
	BCashDividend       float64                 `json:"bCashDividend"`
	Balance             float64                 `json:"balance"`
	BankAvlBalance      float64                 `json:"bankAvlBalance"`
	BankAvlBalanceBF    float64                 `json:"bankAvlBalanceBF"`
	BankBlockAmount     float64                 `json:"bankBlockAmount"`
	BlockAmount         float64                 `json:"blockAmount"`
	BodBalance          float64                 `json:"bodBalance"`
	BondBlockAmount     float64                 `json:"bondBlockAmount"`
	BuyingAmount        float64                 `json:"buyingAmount"`
	CashBalance         float64                 `json:"cashBalance"`
	CashDevident        float64                 `json:"cashDevident"`
	CustodyID           string                  `json:"custodyID"`
	DepoFee             float64                 `json:"depoFee"`
	Dsecured            float64                 `json:"dsecured"`
	FullName            string                  `json:"fullName"`
	FundBlockAmount     float64                 `json:"fundBlockAmount"`
	IaInfos             []InvestmentAccountInfo `json:"iaInfos"`
	MBlockAmount        float64                 `json:"mBlockAmount"`
	Mrused              float64                 `json:"mrused"`
	Object              string                  `json:"object"`
	PP0                 float64                 `json:"pp0"`
	Pp0forBF            float64                 `json:"pp0forBF"`
	SCashDividend       float64                 `json:"sCashDividend"`
	SecureAmtPO         float64                 `json:"secureAmtPO"`
}

type CashStatements struct {
	Response CashStatementsResponse `json:"response"`
}

type CashStatementsResponse struct {
	Data              []CashStatement `json:"data"`
	PageIndex         int64           `json:"pageIndex"`
	PageSize          int64           `json:"pageSize"`
	TotalCount        int64           `json:"totalCount"`
	TotalCreditAmount int64           `json:"totalCreditAmount"`
	TotalDebitAmount  int64           `json:"totalDebitAmount"`
}

type CashStatement struct {
	AccountNo       string  `json:"accountNo"`
	BusinessDate    string  `json:"businessDate"`
	CreditAmount    float64 `json:"creditAmount"`
	CustodyID       string  `json:"custodyID"`
	DebitAmount     float64 `json:"debitAmount"`
	Descriptions    string  `json:"descriptions"`
	TransactionCode string  `json:"transactionCode"`
	TransactionName string  `json:"transactionName"`
	TransactionNum  string  `json:"transactionNum"`
	TransationDate  string  `json:"transationDate"`
}

type MarginInformation struct {
	Response MarginInformationResponse `json:"response"`
}

type MarginInformationResponse struct {
	Data      []MarginInformationEntry `json:"data"`
	TotalPage int64                    `json:"totalPage"`
	TotalRow  int64                    `json:"totalRow"`
}

type MarginInformationEntry struct {
	IntAmount            float64 `json:"intAmount"`
	IntPaid              float64 `json:"intPaid"`
	OverDueDate          string  `json:"overDueDate"`
	PaidFee              float64 `json:"paidFee"`
	PaidInterestFee      float64 `json:"paidInterestFee"`
	PrinPaid             float64 `json:"prinPaid"`
	PrintAmount          float64 `json:"printAmount"`
	Rate2                float64 `json:"rate2"`
	ReleaseDate          string  `json:"releaseDate"`
	ReleasedAmount       float64 `json:"releasedAmount"`
	ReleasedDay          float64 `json:"releasedDay"`
	RemainingFee         float64 `json:"remainingFee"`
	RemainingInterestFee float64 `json:"remainingInterestFee"`
}

type InvestmentAccountInfo struct {
	Available float64 `json:"available"`
	Hold      float64 `json:"hold"`
	Partner   string  `json:"partner"`
}
