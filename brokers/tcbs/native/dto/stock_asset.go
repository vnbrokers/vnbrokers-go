// Code generated from TCBS OpenAPI; DO NOT EDIT.

package dto

import "github.com/shopspring/decimal"

type MarginQuota struct {
	AccountNo     string          `json:"accountNo"`
	AccountStatus string          `json:"accountStatus"`
	AccountType   string          `json:"accountType"`
	Aftype        string          `json:"aftype"`
	BankAccount   string          `json:"bankAccount"`
	BankName      string          `json:"bankName"`
	CustodyID     string          `json:"custodyID"`
	IsIA          string          `json:"isIA"`
	MarginLimit   decimal.Decimal `json:"marginLimit"`
	VSDStatus     string          `json:"vsdStatus"`
}

type MarginAccountInformation struct {
	AccountNo       string                             `json:"accountNo"`
	AccruedInterest decimal.Decimal                    `json:"accruedInterest"`
	DueAmount       decimal.Decimal                    `json:"dueAmount"`
	Outstanding     decimal.Decimal                    `json:"outstanding"`
	OverdueAmount   decimal.Decimal                    `json:"overdueAmount"`
	RiskPolicy      MarginAccountInformationRiskPolicy `json:"riskPolicy"`
	RiskStatus      MarginAccountInformationRiskStatus `json:"riskStatus"`
	RTT             decimal.Decimal                    `json:"rtt"`
	TotalFeeDebt    decimal.Decimal                    `json:"totalFeeDebt"`
}

type MarginAccountInformationRiskPolicy struct {
	InitialMargin     decimal.Decimal `json:"initialMargin"`
	LiquidationMargin decimal.Decimal `json:"liquidationMargin"`
	MaintenanceMargin decimal.Decimal `json:"maintenanceMargin"`
	NumDayMarginCall  decimal.Decimal `json:"numDayMarginCall"`
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
	ID              decimal.Decimal      `json:"id"`
	Name            string               `json:"name"`
	Proposals       []MarginSureProposal `json:"proposals"`
	Status          string               `json:"status"`
	SubscriptionFee decimal.Decimal      `json:"subscriptionFee"`
}

type MarginSureProposal struct {
	ID                       decimal.Decimal `json:"id"`
	InterestAdjustmentValue  decimal.Decimal `json:"interestAdjustmentValue"`
	InterestPercentThreshold decimal.Decimal `json:"interestPercentThreshold"`
	MarginInsuranceID        decimal.Decimal `json:"marginInsuranceId"`
	ThresholdType            string          `json:"thresholdType"`
}

type TPlusPackages struct {
	Data []TPlusPackage `json:"data"`
}

type TPlusPackage struct {
	CreatedAt                                 string             `json:"createdAt"`
	DebtCollectionFee                         decimal.Decimal    `json:"debtCollectionFee"`
	Description                               string             `json:"description"`
	Discountable                              bool               `json:"discountable"`
	ExtensionFee                              decimal.Decimal    `json:"extensionFee"`
	ExtensionInterest                         decimal.Decimal    `json:"extensionInterest"`
	ExtensionInterestBeforeInterestSettlement decimal.Decimal    `json:"extensionInterestBeforeInterestSettlement"`
	FirstRate                                 decimal.Decimal    `json:"firstRate"`
	ID                                        decimal.Decimal    `json:"id"`
	InterestCalculationBasis                  decimal.Decimal    `json:"interestCalculationBasis"`
	IwealthPartnerInterest                    decimal.Decimal    `json:"iwealthPartnerInterest"`
	MinRateConstraint                         bool               `json:"minRateConstraint"`
	Name                                      string             `json:"name"`
	OverdueFee                                decimal.Decimal    `json:"overdueFee"`
	OverdueInterest                           decimal.Decimal    `json:"overdueInterest"`
	Status                                    string             `json:"status"`
	UndueFee                                  decimal.Decimal    `json:"undueFee"`
	UndueInterestType                         string             `json:"undueInterestType"`
	UndueLadderValue                          []InterestRateTier `json:"undueLadderValue"`
	UpdatedAt                                 string             `json:"updatedAt"`
	ValidFrom                                 string             `json:"validFrom"`
}

type InterestRateTier struct {
	DueDate   decimal.Decimal `json:"dueDate"`
	ID        decimal.Decimal `json:"id"`
	Rate      decimal.Decimal `json:"rate"`
	StartDate decimal.Decimal `json:"startDate"`
}

type Loans struct {
	Content []Loan          `json:"content"`
	Size    decimal.Decimal `json:"size"`
}

type Loan struct {
	AccountNo          string          `json:"accountNo"`
	Addons             []LoanAddon     `json:"addons"`
	DueDate            string          `json:"dueDate"`
	Fee                decimal.Decimal `json:"fee"`
	ID                 decimal.Decimal `json:"id"`
	Insurance          LoanInsurance   `json:"insurance"`
	Interest           decimal.Decimal `json:"interest"`
	IsRenewable        bool            `json:"isRenewable"`
	LoanDays           decimal.Decimal `json:"loanDays"`
	MaxRenewTime       decimal.Decimal `json:"maxRenewTime"`
	MrxLoanID          decimal.Decimal `json:"mrxLoanId"`
	OpeningDate        string          `json:"openingDate"`
	PricingPolicyType  string          `json:"pricingPolicyType"`
	Principal          decimal.Decimal `json:"principal"`
	Rate               decimal.Decimal `json:"rate"`
	ReasonList         []string        `json:"reasonList"`
	RemainingPrincipal decimal.Decimal `json:"remainingPrincipal"`
	RenewTime          decimal.Decimal `json:"renewTime"`
	Status             string          `json:"status"`
	Symbol             string          `json:"symbol"`
	UndueLoanFee       decimal.Decimal `json:"undueLoanFee"`
}

type LoanAddon struct {
	Discounts []LoanDiscount  `json:"discounts"`
	LoanID    decimal.Decimal `json:"loanId"`
	Others    []LoanPolicy    `json:"others"`
}

type LoanDiscount struct {
	Code                  string          `json:"code"`
	EndDate               string          `json:"endDate"`
	OrigDiscountLoanAmt   decimal.Decimal `json:"origDiscountLoanAmt"`
	PreferentialRate      decimal.Decimal `json:"preferentialRate"`
	RemainDiscountLoanAmt decimal.Decimal `json:"remainDiscountLoanAmt"`
	StartDate             string          `json:"startDate"`
}

type LoanPolicy struct {
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Name        string          `json:"name"`
	Rate        decimal.Decimal `json:"rate"`
	Symbols     []string        `json:"symbols"`
}

type LoanInsurance struct {
	AccountNo                        string          `json:"accountNo"`
	AdjustStrikePrice                decimal.Decimal `json:"adjustStrikePrice"`
	DecreasedInterest                decimal.Decimal `json:"decreasedInterest"`
	DecreasedPrice                   decimal.Decimal `json:"decreasedPrice"`
	EndDate                          string          `json:"endDate"`
	IncreasePrice                    decimal.Decimal `json:"increasePrice"`
	IncreasedInterest                decimal.Decimal `json:"increasedInterest"`
	InsuranceCode                    string          `json:"insuranceCode"`
	InsuranceFee                     decimal.Decimal `json:"insuranceFee"`
	InsuranceName                    string          `json:"insuranceName"`
	InterestPercentageLowerThreshold decimal.Decimal `json:"interestPercentageLowerThreshold"`
	InterestPercentageUpperThreshold decimal.Decimal `json:"interestPercentageUpperThreshold"`
	InterestRateLowerAdjustmentValue decimal.Decimal `json:"interestRateLowerAdjustmentValue"`
	InterestRateUpperAdjustmentValue decimal.Decimal `json:"interestRateUpperAdjustmentValue"`
	LoanID                           decimal.Decimal `json:"loanId"`
	OriginalInterest                 decimal.Decimal `json:"originalInterest"`
	StartDate                        string          `json:"startDate"`
	Symbol                           string          `json:"symbol"`
}

type StockPurchasingPower struct {
	AccountNo          string          `json:"accountNo"`
	CustodyID          string          `json:"custodyID"`
	MarginPriceLoan    decimal.Decimal `json:"marginPriceLoan"`
	MarginRatioLoan    decimal.Decimal `json:"marginRatioLoan"`
	MaxBuyQuantity     decimal.Decimal `json:"maxBuyQuantity"`
	MinBuyQuantity     decimal.Decimal `json:"minBuyQuantity"`
	PP0                decimal.Decimal `json:"pp0"`
	PPSE               decimal.Decimal `json:"ppse"`
	PPSERef            decimal.Decimal `json:"ppseref"`
	Price              decimal.Decimal `json:"price"`
	RateBrkB           decimal.Decimal `json:"rateBrkB"`
	RateBrkS           decimal.Decimal `json:"rateBrkS"`
	RealMaxBuyQuantity decimal.Decimal `json:"realMaxBuyQuantity"`
	Symbol             string          `json:"symbol"`
}

type StockAssets struct {
	AccountNo string             `json:"accountNo"`
	CustodyID string             `json:"custodyID"`
	FullName  string             `json:"fullName"`
	Object    string             `json:"object"`
	Stock     []StockHoldingInfo `json:"stock"`
}

type StockHoldingInfo struct {
	AvailableTrading decimal.Decimal `json:"availableTrading"`
	Blocked          decimal.Decimal `json:"blocked"`
	CashDividend     decimal.Decimal `json:"cashDividend"`
	CostPrice        decimal.Decimal `json:"costPrice"`
	CurrentPrice     decimal.Decimal `json:"currentPrice"`
	ExercisedCA      decimal.Decimal `json:"exercisedCA"`
	Mortgaged        decimal.Decimal `json:"mortgaged"`
	OnHold           decimal.Decimal `json:"onHold"`
	SecType          string          `json:"secType"`
	SecTypeName      string          `json:"secTypeName"`
	SecuredQuantity  decimal.Decimal `json:"securedQuantity"`
	SellExec         decimal.Decimal `json:"sellExec"`
	SellRemain       decimal.Decimal `json:"sellRemain"`
	Settlement       string          `json:"settlement"`
	StockDividend    decimal.Decimal `json:"stockDividend"`
	Symbol           string          `json:"symbol"`
	T0               decimal.Decimal `json:"t0"`
	T1               decimal.Decimal `json:"t1"`
	T2               decimal.Decimal `json:"t2"`
	TotalQtty        decimal.Decimal `json:"totalQtty"`
	UnexercisedCA    decimal.Decimal `json:"unexercisedCA"`
	WaitForTrade     decimal.Decimal `json:"waitForTrade"`
	WaitForTransfer  decimal.Decimal `json:"waitForTransfer"`
	WaitForWithdraw  decimal.Decimal `json:"waitForWithdraw"`
}

type CashInvestments struct {
	Data       []CashInvestment `json:"data"`
	Object     string           `json:"object"`
	PageIndex  decimal.Decimal  `json:"pageIndex"`
	PageSize   decimal.Decimal  `json:"pageSize"`
	TotalCount decimal.Decimal  `json:"totalCount"`
}

type CashInvestment struct {
	AccountNo           string                  `json:"accountNo"`
	Adused              decimal.Decimal         `json:"adused"`
	AvalBondBlockAmount decimal.Decimal         `json:"avalBondBlockAmount"`
	AvlAdvanceAmount    decimal.Decimal         `json:"avlAdvanceAmount"`
	AvlWithdraw         decimal.Decimal         `json:"avlWithdraw"`
	BCashDividend       decimal.Decimal         `json:"bCashDividend"`
	Balance             decimal.Decimal         `json:"balance"`
	BankAvlBalance      decimal.Decimal         `json:"bankAvlBalance"`
	BankAvlBalanceBF    decimal.Decimal         `json:"bankAvlBalanceBF"`
	BankBlockAmount     decimal.Decimal         `json:"bankBlockAmount"`
	BlockAmount         decimal.Decimal         `json:"blockAmount"`
	BodBalance          decimal.Decimal         `json:"bodBalance"`
	BondBlockAmount     decimal.Decimal         `json:"bondBlockAmount"`
	BuyingAmount        decimal.Decimal         `json:"buyingAmount"`
	CashBalance         decimal.Decimal         `json:"cashBalance"`
	CashDevident        decimal.Decimal         `json:"cashDevident"`
	CustodyID           string                  `json:"custodyID"`
	DepoFee             decimal.Decimal         `json:"depoFee"`
	Dsecured            decimal.Decimal         `json:"dsecured"`
	FullName            string                  `json:"fullName"`
	FundBlockAmount     decimal.Decimal         `json:"fundBlockAmount"`
	IaInfos             []InvestmentAccountInfo `json:"iaInfos"`
	MBlockAmount        decimal.Decimal         `json:"mBlockAmount"`
	Mrused              decimal.Decimal         `json:"mrused"`
	Object              string                  `json:"object"`
	PP0                 decimal.Decimal         `json:"pp0"`
	Pp0forBF            decimal.Decimal         `json:"pp0forBF"`
	SCashDividend       decimal.Decimal         `json:"sCashDividend"`
	SecureAmtPO         decimal.Decimal         `json:"secureAmtPO"`
}

type CashStatements struct {
	Response CashStatementsResponse `json:"response"`
}

type CashStatementsResponse struct {
	Data              []CashStatement `json:"data"`
	PageIndex         decimal.Decimal `json:"pageIndex"`
	PageSize          decimal.Decimal `json:"pageSize"`
	TotalCount        decimal.Decimal `json:"totalCount"`
	TotalCreditAmount decimal.Decimal `json:"totalCreditAmount"`
	TotalDebitAmount  decimal.Decimal `json:"totalDebitAmount"`
}

type CashStatement struct {
	AccountNo       string          `json:"accountNo"`
	BusinessDate    string          `json:"businessDate"`
	CreditAmount    decimal.Decimal `json:"creditAmount"`
	CustodyID       string          `json:"custodyID"`
	DebitAmount     decimal.Decimal `json:"debitAmount"`
	Descriptions    string          `json:"descriptions"`
	TransactionCode string          `json:"transactionCode"`
	TransactionName string          `json:"transactionName"`
	TransactionNum  string          `json:"transactionNum"`
	TransationDate  string          `json:"transationDate"`
}

type MarginInformation struct {
	Response MarginInformationResponse `json:"response"`
}

type MarginInformationResponse struct {
	Data      []MarginInformationEntry `json:"data"`
	TotalPage decimal.Decimal          `json:"totalPage"`
	TotalRow  decimal.Decimal          `json:"totalRow"`
}

type MarginInformationEntry struct {
	IntAmount            decimal.Decimal `json:"intAmount"`
	IntPaid              decimal.Decimal `json:"intPaid"`
	OverDueDate          string          `json:"overDueDate"`
	PaidFee              decimal.Decimal `json:"paidFee"`
	PaidInterestFee      decimal.Decimal `json:"paidInterestFee"`
	PrinPaid             decimal.Decimal `json:"prinPaid"`
	PrintAmount          decimal.Decimal `json:"printAmount"`
	Rate2                decimal.Decimal `json:"rate2"`
	ReleaseDate          string          `json:"releaseDate"`
	ReleasedAmount       decimal.Decimal `json:"releasedAmount"`
	ReleasedDay          decimal.Decimal `json:"releasedDay"`
	RemainingFee         decimal.Decimal `json:"remainingFee"`
	RemainingInterestFee decimal.Decimal `json:"remainingInterestFee"`
}

type InvestmentAccountInfo struct {
	Available decimal.Decimal `json:"available"`
	Hold      decimal.Decimal `json:"hold"`
	Partner   string          `json:"partner"`
}
