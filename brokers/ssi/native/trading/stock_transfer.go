package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) Transferable(ctx context.Context, accountID string) (*dto.TradingResponse[dto.TransferableStockAccountData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.TransferableStockAccountData]](ctx, s, CapabilityTransferable, "/api/v2/stock/transferable", params)
}

func (s *service) StockTransferHistories(ctx context.Context, accountID string, startDate string, endDate string) (*dto.TradingResponse[dto.StockTransferHistoryAccountData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	return get[dto.TradingResponse[dto.StockTransferHistoryAccountData]](ctx, s, CapabilityStockTransferHistories, "/api/v2/stock/transferHistories", params)
}

func (s *service) StockTransfer(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.TransactionResponse], error) {
	return post[dto.TradingResponse[dto.TransactionResponse]](ctx, s, CapabilityStockTransfer, "/api/v2/stock/transfer", body)
}
