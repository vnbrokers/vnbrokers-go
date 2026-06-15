package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) TransferBetweenSubaccounts(ctx context.Context, request dto.TransferBetweenSubaccountsRequest) (*dto.TransferBetweenSubaccountsResponse, error) {
	return do[dto.TransferBetweenSubaccountsResponse](s, ctx, CapabilityTransferBetweenSubaccounts, "POST", "/physis/v1/stock/transfer", url.Values{}, request)
}

func (s *service) WithdrawDerivativeMargin(ctx context.Context, request dto.WithdrawDerivativeMarginRequest) (*dto.WithdrawDerivativeMarginResponse, error) {
	return do[dto.WithdrawDerivativeMarginResponse](s, ctx, CapabilityWithdrawDerivativeMargin, "POST", "/khronos/v1/cash/withdraw/update", url.Values{}, request)
}

func (s *service) DepositDerivativeMargin(ctx context.Context, request dto.DepositDerivativeMarginRequest) (*dto.DepositDerivativeMarginResponse, error) {
	return do[dto.DepositDerivativeMarginResponse](s, ctx, CapabilityDepositDerivativeMargin, "POST", "/khronos/v1/cash/deposit/update", url.Values{}, request)
}
