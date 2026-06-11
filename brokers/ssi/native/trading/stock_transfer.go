package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) Transferable(ctx context.Context, accountID string) (*dto.TransferableResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TransferableResponse](ctx, s, CapabilityTransferable, "/api/v2/stock/transferable", params)
}

func (s *service) StockTransferHistories(ctx context.Context, accountID string, startDate string, endDate string) (*dto.StockTransferHistoriesResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	return get[dto.StockTransferHistoriesResponse](ctx, s, CapabilityStockTransferHistories, "/api/v2/stock/transferHistories", params)
}

func (s *service) StockTransfer(ctx context.Context, body map[string]any) (*dto.StockTransferResponse, error) {
	return post[dto.StockTransferResponse](ctx, s, CapabilityStockTransfer, "/api/v2/stock/transfer", body)
}
