package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) CashInAdvanceAmount(ctx context.Context, accountID string) (*dto.CashInAdvanceAmountResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.CashInAdvanceAmountResponse](ctx, s, CapabilityCashInAdvanceAmount, "/api/v2/cash/cashInAdvanceAmount", params)
}

func (s *service) UnsettleSoldTransaction(ctx context.Context, accountID string, settleDate string) (*dto.UnsettleSoldTransactionResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "settleDate", settleDate)
	return get[dto.UnsettleSoldTransactionResponse](ctx, s, CapabilityUnsettleSoldTransaction, "/api/v2/cash/unsettleSoldTransaction", params)
}

func (s *service) TransferHistories(ctx context.Context, accountID string, fromDate string, toDate string) (*dto.CashTransferHistoriesResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "fromDate", fromDate)
	setOptionalString(params, "toDate", toDate)
	return get[dto.CashTransferHistoriesResponse](ctx, s, CapabilityTransferHistories, "/api/v2/cash/transferHistories", params)
}

func (s *service) CashInAdvanceHistories(ctx context.Context, accountID string, fromDate string, toDate string) (*dto.CashInAdvanceHistoriesResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "fromDate", fromDate)
	setOptionalString(params, "toDate", toDate)
	return get[dto.CashInAdvanceHistoriesResponse](ctx, s, CapabilityCashInAdvanceHistories, "/api/v2/cash/cashInAdvanceHistories", params)
}

func (s *service) EstCashInAdvanceFee(ctx context.Context, accountID string, ciaAmount string, receiveAmount string) (*dto.EstCashInAdvanceFeeResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "ciaAmount", ciaAmount)
	setOptionalString(params, "receiveAmount", receiveAmount)
	return get[dto.EstCashInAdvanceFeeResponse](ctx, s, CapabilityEstCashInAdvanceFee, "/api/v2/cash/estCashInAdvanceFee", params)
}

func (s *service) VSDCashDW(ctx context.Context, accountID string, amount string, vsdType string, remark string, code string) (*dto.VSDCashDWResponse, error) {
	body := map[string]any{
		"account": accountID,
		"amount":  amount,
		"type":    vsdType,
		"remark":  remark,
		"code":    code,
	}
	return post[dto.VSDCashDWResponse](ctx, s, CapabilityVSDCashDW, "/api/v2/cash/vsdCashDW", body)
}

func (s *service) TransferInternal(ctx context.Context, accountID string, beneficiaryAccount string, amount string, remark string, code string) (*dto.TransferInternalResponse, error) {
	body := map[string]any{
		"account":            accountID,
		"beneficiaryAccount": beneficiaryAccount,
		"amount":             amount,
		"remark":             remark,
		"code":               code,
	}
	return post[dto.TransferInternalResponse](ctx, s, CapabilityTransferInternal, "/api/v2/cash/transferInternal", body)
}

func (s *service) CreateCashInAdvance(ctx context.Context, accountID string, ciaAmount string, receiveAmount string, code string) (*dto.CreateCashInAdvanceResponse, error) {
	body := map[string]any{
		"account": accountID,
		"code":    code,
	}
	setOptionalStringMap(body, "ciaAmount", ciaAmount)
	setOptionalStringMap(body, "receiveAmount", receiveAmount)
	return post[dto.CreateCashInAdvanceResponse](ctx, s, CapabilityCreateCashInAdvance, "/api/v2/cash/createCashInAdvance", body)
}

func setOptionalStringMap(m map[string]any, key string, value string) {
	if value != "" {
		m[key] = value
	}
}
