package dto

import "encoding/json"

type GetStockPurchasingPowerRequest struct{ AccountNo string }
type GetStockPurchasingPowerResponse StockPurchasingPower

type GetStockPurchasingPowerBySymbolRequest struct {
	AccountNo string
	Symbol    string
}
type GetStockPurchasingPowerBySymbolResponse StockPurchasingPower

type GetStockPurchasingPowerBySymbolPriceRequest struct {
	AccountNo string
	Symbol    string
	Price     string
}
type GetStockPurchasingPowerBySymbolPriceResponse StockPurchasingPower

type GetMarginQuotaRequest struct{ CustodyID string }
type GetMarginQuotaResponse []MarginQuota

type GetMarginAccountInformationRequest struct{ AccountNo string }
type GetMarginAccountInformationResponse MarginAccountInformation

func (r *GetMarginAccountInformationResponse) UnmarshalJSON(data []byte) error {
	var empty []MarginAccountInformation
	if err := json.Unmarshal(data, &empty); err == nil && len(empty) == 0 {
		*r = GetMarginAccountInformationResponse{}
		return nil
	}

	var response MarginAccountInformation
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	*r = GetMarginAccountInformationResponse(response)
	return nil
}

type GetSupplementaryLoanPackagesRequest struct{ AccountNo string }
type GetSupplementaryLoanPackagesResponse SupplementaryLoanPackages

type GetLoansRequest struct{ AccountNo string }
type GetLoansResponse Loans

type GetStockAssetsRequest struct{ AccountNo string }
type GetStockAssetsResponse StockAssets

type GetCashInvestmentsRequest struct{ AccountNo string }
type GetCashInvestmentsResponse CashInvestments

type GetCashStatementsRequest struct {
	AccountNo       string
	FromDate        string
	ToDate          string
	PageSize        string
	PageIndex       string
	TransactionCode string
}
type GetCashStatementsResponse CashStatements

type GetMarginInformationRequest struct {
	AccountNo   string
	FromDate    string
	ToDate      string
	Page        string
	Size        string
	CustodyCode string
}
type GetMarginInformationResponse MarginInformation
