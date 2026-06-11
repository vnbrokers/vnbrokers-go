package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) CashInAdvanceAmount(ctx context.Context, accountID string) (*dto.TradingResponse[dto.CashInAdvanceAmountData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.CashInAdvanceAmountData]](ctx, s, CapabilityCashInAdvanceAmount, "/api/v2/cash/cashInAdvanceAmount", params)
}

func (s *service) UnsettleSoldTransaction(ctx context.Context, accountID string, settleDate string) (*dto.TradingResponse[dto.UnsettledSoldTransactionsData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "settleDate", settleDate)
	return get[dto.TradingResponse[dto.UnsettledSoldTransactionsData]](ctx, s, CapabilityUnsettleSoldTransaction, "/api/v2/cash/unsettleSoldTransaction", params)
}

func (s *service) TransferHistories(ctx context.Context, accountID string, fromDate string, toDate string) (*dto.TradingResponse[dto.TransferHistoriesData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "fromDate", fromDate)
	setOptionalString(params, "toDate", toDate)
	return get[dto.TradingResponse[dto.TransferHistoriesData]](ctx, s, CapabilityTransferHistories, "/api/v2/cash/transferHistories", params)
}

func (s *service) CashInAdvanceHistories(ctx context.Context, accountID string, fromDate string, toDate string) (*dto.TradingResponse[dto.CashInAdvanceHistoriesData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "fromDate", fromDate)
	setOptionalString(params, "toDate", toDate)
	return get[dto.TradingResponse[dto.CashInAdvanceHistoriesData]](ctx, s, CapabilityCashInAdvanceHistories, "/api/v2/cash/cashInAdvanceHistories", params)
}

func (s *service) EstCashInAdvanceFee(ctx context.Context, accountID string, ciaAmount string, receiveAmount string) (*dto.TradingResponse[dto.EstimateCashInAdvanceFeeData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	setOptionalString(params, "ciaAmount", ciaAmount)
	setOptionalString(params, "receiveAmount", receiveAmount)
	return get[dto.TradingResponse[dto.EstimateCashInAdvanceFeeData]](ctx, s, CapabilityEstCashInAdvanceFee, "/api/v2/cash/estCashInAdvanceFee", params)
}

func (s *service) VSDCashDW(ctx context.Context, accountID string, amount string, vsdType string, remark string, code string) (*dto.TradingResponse[dto.TransactionResponse], error) {
	body := map[string]any{
		"account": accountID,
		"amount":  amount,
		"type":    vsdType,
		"remark":  remark,
		"code":    code,
	}
	return post[dto.TradingResponse[dto.TransactionResponse]](ctx, s, CapabilityVSDCashDW, "/api/v2/cash/vsdCashDW", body)
}

func (s *service) TransferInternal(ctx context.Context, accountID string, beneficiaryAccount string, amount string, remark string, code string) (*dto.TradingResponse[dto.TransactionResponse], error) {
	body := map[string]any{
		"account":            accountID,
		"beneficiaryAccount": beneficiaryAccount,
		"amount":             amount,
		"remark":             remark,
		"code":               code,
	}
	return post[dto.TradingResponse[dto.TransactionResponse]](ctx, s, CapabilityTransferInternal, "/api/v2/cash/transferInternal", body)
}

func (s *service) CreateCashInAdvance(ctx context.Context, accountID string, ciaAmount string, receiveAmount string, code string) (*dto.TradingResponse[dto.TransactionResponse], error) {
	body := map[string]any{
		"account": accountID,
		"code":    code,
	}
	setOptionalStringMap(body, "ciaAmount", ciaAmount)
	setOptionalStringMap(body, "receiveAmount", receiveAmount)
	return post[dto.TradingResponse[dto.TransactionResponse]](ctx, s, CapabilityCreateCashInAdvance, "/api/v2/cash/createCashInAdvance", body)
}

func setOptionalStringMap(m map[string]any, key string, value string) {
	if value != "" {
		m[key] = value
	}
}
