package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetStockPurchasingPower(ctx context.Context, request dto.GetStockPurchasingPowerRequest) (*dto.GetStockPurchasingPowerResponse, error) {
	return do[dto.GetStockPurchasingPowerResponse](s, ctx, CapabilityGetStockPurchasingPower, "GET", "/aion/v1/accounts/"+escaped(request.AccountNo)+"/ppse", url.Values{}, nil)
}

func (s *service) GetStockPurchasingPowerBySymbol(ctx context.Context, request dto.GetStockPurchasingPowerBySymbolRequest) (*dto.GetStockPurchasingPowerBySymbolResponse, error) {
	path := "/aion/v1/accounts/" + escaped(request.AccountNo) + "/ppse/" + escaped(request.Symbol)
	return do[dto.GetStockPurchasingPowerBySymbolResponse](s, ctx, CapabilityGetStockPurchasingPowerBySymbol, "GET", path, url.Values{}, nil)
}

func (s *service) GetStockPurchasingPowerBySymbolPrice(ctx context.Context, request dto.GetStockPurchasingPowerBySymbolPriceRequest) (*dto.GetStockPurchasingPowerBySymbolPriceResponse, error) {
	path := "/aion/v1/accounts/" + escaped(request.AccountNo) + "/ppse/" + escaped(request.Symbol) + "/" + escaped(request.Price)
	return do[dto.GetStockPurchasingPowerBySymbolPriceResponse](s, ctx, CapabilityGetStockPurchasingPowerBySymbolPrice, "GET", path, url.Values{}, nil)
}

func (s *service) GetMarginQuota(ctx context.Context, request dto.GetMarginQuotaRequest) (*dto.GetMarginQuotaResponse, error) {
	return do[dto.GetMarginQuotaResponse](s, ctx, CapabilityGetMarginQuota, "GET", "/aion/v1/customers/"+escaped(request.CustodyID)+"/accounts", url.Values{}, nil)
}

func (s *service) GetMarginAccountInformation(ctx context.Context, request dto.GetMarginAccountInformationRequest) (*dto.GetMarginAccountInformationResponse, error) {
	return do[dto.GetMarginAccountInformationResponse](s, ctx, CapabilityGetMarginAccountInformation, "GET", "/hydros/v1/account/"+escaped(request.AccountNo)+"/risk", url.Values{}, nil)
}

func (s *service) GetSupplementaryLoanPackages(ctx context.Context, request dto.GetSupplementaryLoanPackagesRequest) (*dto.GetSupplementaryLoanPackagesResponse, error) {
	path := "/campaign-management/v1/margin/subscription/" + escaped(request.AccountNo) + "/addons/detail"
	return do[dto.GetSupplementaryLoanPackagesResponse](s, ctx, CapabilityGetSupplementaryLoanPackages, "GET", path, url.Values{}, nil)
}

func (s *service) GetLoans(ctx context.Context, request dto.GetLoansRequest) (*dto.GetLoansResponse, error) {
	return do[dto.GetLoansResponse](s, ctx, CapabilityGetLoans, "GET", "/khaos/v1/loan/"+escaped(request.AccountNo), url.Values{}, nil)
}

func (s *service) GetStockAssets(ctx context.Context, request dto.GetStockAssetsRequest) (*dto.GetStockAssetsResponse, error) {
	return do[dto.GetStockAssetsResponse](s, ctx, CapabilityGetStockAssets, "GET", "/aion/v1/accounts/"+escaped(request.AccountNo)+"/se", url.Values{}, nil)
}

func (s *service) GetCashInvestments(ctx context.Context, request dto.GetCashInvestmentsRequest) (*dto.GetCashInvestmentsResponse, error) {
	return do[dto.GetCashInvestmentsResponse](s, ctx, CapabilityGetCashInvestments, "GET", "/aion/v1/accounts/"+escaped(request.AccountNo)+"/cashInvestments", url.Values{}, nil)
}

func (s *service) GetCashStatements(ctx context.Context, request dto.GetCashStatementsRequest) (*dto.GetCashStatementsResponse, error) {
	query := url.Values{}
	set(query, "acctno", request.AccountNo)
	set(query, "fromDate", request.FromDate)
	set(query, "toDate", request.ToDate)
	set(query, "pageSize", request.PageSize)
	set(query, "pageIndex", request.PageIndex)
	setOptional(query, "transactionCode", request.TransactionCode)
	return do[dto.GetCashStatementsResponse](s, ctx, CapabilityGetCashStatements, "GET", "/erebos/v2/digital/trans-hist-cashStatements", query, nil)
}

func (s *service) GetMarginInformation(ctx context.Context, request dto.GetMarginInformationRequest) (*dto.GetMarginInformationResponse, error) {
	query := url.Values{}
	set(query, "acctno", request.AccountNo)
	set(query, "fromdate", request.FromDate)
	set(query, "toDate", request.ToDate)
	set(query, "page", request.Page)
	set(query, "size", request.Size)
	set(query, "custodycd", request.CustodyCode)
	return do[dto.GetMarginInformationResponse](s, ctx, CapabilityGetMarginInformation, "GET", "/erebos/v2/digital/margin-info", query, nil)
}
