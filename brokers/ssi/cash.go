package ssi

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/core"
)

type TradingCashService struct {
	broker *Broker
}

func (s *TradingCashService) CashInAdvanceAmount(
	ctx context.Context,
	accountID string,
) (TradingResponse[CashInAdvanceAmountData], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[CashInAdvanceAmountData]{}, err
	}
	var response TradingResponse[CashInAdvanceAmountData]
	err := s.broker.getAndDecode(ctx, "trading.cash.cash_in_advance_amount", "/api/v2/cash/cashInAdvanceAmount", accountParams(accountID), &response)
	return response, err
}

func (s *TradingCashService) UnsettledSoldTransactions(
	ctx context.Context,
	request CashUnsettledSoldTransactionsRequest,
) (TradingResponse[UnsettledSoldTransactionsData], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[UnsettledSoldTransactionsData]{}, err
	}
	params := accountParams(request.AccountID)
	setOptionalString(params, "settleDate", request.SettleDate)
	var response TradingResponse[UnsettledSoldTransactionsData]
	err := s.broker.getAndDecode(ctx, "trading.cash.unsettled_sold_transactions", "/api/v2/cash/unsettleSoldTransaction", params, &response)
	return response, err
}

func (s *TradingCashService) TransferHistories(
	ctx context.Context,
	request CashTransferHistoriesRequest,
) (TradingResponse[CashTransferHistoriesData], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[CashTransferHistoriesData]{}, err
	}
	params := dateRangeParams(request.AccountID, "fromDate", request.FromDate, "toDate", request.ToDate)
	var response TradingResponse[CashTransferHistoriesData]
	err := s.broker.getAndDecode(ctx, "trading.cash.transfer_histories", "/api/v2/cash/transferHistories", params, &response)
	return response, err
}

func (s *TradingCashService) CashInAdvanceHistories(
	ctx context.Context,
	request CashInAdvanceHistoriesRequest,
) (TradingResponse[CashInAdvanceHistoriesData], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[CashInAdvanceHistoriesData]{}, err
	}
	params := dateRangeParams(request.AccountID, "fromDate", request.FromDate, "toDate", request.ToDate)
	var response TradingResponse[CashInAdvanceHistoriesData]
	err := s.broker.getAndDecode(ctx, "trading.cash.cash_in_advance_histories", "/api/v2/cash/cashInAdvanceHistories", params, &response)
	return response, err
}

func (s *TradingCashService) EstimateCashInAdvanceFee(
	ctx context.Context,
	request CashInAdvanceFeeRequest,
) (TradingResponse[CashInAdvanceFeeData], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[CashInAdvanceFeeData]{}, err
	}
	params := accountParams(request.AccountID)
	if !request.CIAAmount.IsZero() {
		params.Set("ciaAmount", request.CIAAmount.String())
	}
	if !request.ReceiveAmount.IsZero() {
		params.Set("receiveAmount", request.ReceiveAmount.String())
	}
	var response TradingResponse[CashInAdvanceFeeData]
	err := s.broker.getAndDecode(ctx, "trading.cash.estimate_cash_in_advance_fee", "/api/v2/cash/estCashInAdvanceFee", params, &response)
	return response, err
}

func (s *TradingCashService) VSDCashDWWithRequest(
	ctx context.Context,
	request VSDCashDWRequest,
) (TradingResponse[TransactionResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[TransactionResponse]{}, err
	}
	body := map[string]any{
		"account": request.AccountID,
		"amount":  numberValue(&request.Amount),
		"type":    request.Type,
		"remark":  request.Remark,
		"code":    request.Code,
	}
	var response TradingResponse[TransactionResponse]
	err := s.broker.postAndDecode(ctx, "trading.cash.vsd_cash_dw", "/api/v2/cash/vsdCashDW", body, &response)
	return response, err
}

func (s *TradingCashService) TransferInternalWithRequest(
	ctx context.Context,
	request CashTransferInternalRequest,
) (TradingResponse[TransactionResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[TransactionResponse]{}, err
	}
	body := map[string]any{
		"account":            request.AccountID,
		"beneficiaryAccount": request.BeneficiaryAccount,
		"amount":             numberValue(&request.Amount),
		"remark":             request.Remark,
		"code":               request.Code,
	}
	var response TradingResponse[TransactionResponse]
	err := s.broker.postAndDecode(ctx, "trading.cash.transfer_internal", "/api/v2/cash/transferInternal", body, &response)
	return response, err
}

func (s *TradingCashService) CreateCashInAdvanceWithRequest(
	ctx context.Context,
	request CreateCashInAdvanceRequest,
) (TradingResponse[TransactionResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingCashTransfer); err != nil {
		return TradingResponse[TransactionResponse]{}, err
	}
	body := map[string]any{
		"account": request.AccountID,
		"code":    request.Code,
	}
	setOptionalDecimalBody(body, "ciaAmount", request.CIAAmount)
	setOptionalDecimalBody(body, "receiveAmount", request.ReceiveAmount)
	var response TradingResponse[TransactionResponse]
	err := s.broker.postAndDecode(ctx, "trading.cash.create_cash_in_advance", "/api/v2/cash/createCashInAdvance", body, &response)
	return response, err
}

func addDateRange(params url.Values, startDate string, endDate string) {
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
}
