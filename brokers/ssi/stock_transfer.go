package ssi

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/core"
)

type TradingStockTransferService struct {
	broker *Broker
}

func (s *TradingStockTransferService) Transferable(
	ctx context.Context,
	accountID string,
) (TradingResponse[TransferableStockAccount], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingStockTransfer); err != nil {
		return TradingResponse[TransferableStockAccount]{}, err
	}
	var response TradingResponse[TransferableStockAccount]
	err := s.broker.getAndDecode(ctx, "trading.stock_transfers.transferable", "/api/v2/stock/transferable", accountParams(accountID), &response)
	return response, err
}

func (s *TradingStockTransferService) Histories(
	ctx context.Context,
	request StockTransferHistoriesRequest,
) (TradingResponse[StockTransferHistoryAccount], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingStockTransfer); err != nil {
		return TradingResponse[StockTransferHistoryAccount]{}, err
	}
	params := accountParams(request.AccountID)
	addDateRange(params, request.StartDate, request.EndDate)
	var response TradingResponse[StockTransferHistoryAccount]
	err := s.broker.getAndDecode(ctx, "trading.stock_transfers.histories", "/api/v2/stock/transferHistories", params, &response)
	return response, err
}

func (s *TradingStockTransferService) TransferWithRequest(
	ctx context.Context,
	request StockTransferRequest,
) (TradingResponse[TransactionResponse], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingStockTransfer); err != nil {
		return TradingResponse[TransactionResponse]{}, err
	}
	body := map[string]any{
		"account":            request.AccountID,
		"beneficiaryAccount": request.BeneficiaryAccount,
		"exchangeID":         request.ExchangeID,
		"instrumentID":       request.Symbol,
		"quantity":           numberValue(&request.Quantity),
		"code":               request.Code,
	}
	var response TradingResponse[TransactionResponse]
	err := s.broker.postAndDecode(ctx, "trading.stock_transfers.transfer", "/api/v2/stock/transfer", body, &response)
	return response, err
}
