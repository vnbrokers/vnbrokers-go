package trading

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) GetDerivativeCash(ctx context.Context, request dto.GetDerivativeCashRequest) (*dto.GetDerivativeCashResponse, error) {
	query := url.Values{}
	set(query, "accountId", request.AccountID)
	set(query, "subAccountId", request.SubAccountID)
	set(query, "getType", request.GetType)
	fmt.Printf("query: %s\n", query.Encode())
	return do[dto.GetDerivativeCashResponse](s, ctx, CapabilityGetDerivativeCash, "GET", "/khronos/v1/account/status", query, nil)
}

func (s *service) GetClosedDerivativePositions(ctx context.Context, request dto.GetClosedDerivativePositionsRequest) (*dto.GetClosedDerivativePositionsResponse, error) {
	query := url.Values{}
	set(query, "accountId", request.AccountID)
	set(query, "subAccountId", request.SubAccountID)
	set(query, "symbol", request.Symbol)
	set(query, "pageNo", strconv.FormatInt(request.PageNo, 10))
	set(query, "pageSize", strconv.FormatInt(request.PageSize, 10))
	return do[dto.GetClosedDerivativePositionsResponse](s, ctx, CapabilityGetClosedDerivativePositions, "GET", "/khronos/v1/account/portfolio/position/close", query, nil)
}

func (s *service) GetOpenDerivativePositions(ctx context.Context, request dto.GetOpenDerivativePositionsRequest) (*dto.GetOpenDerivativePositionsResponse, error) {
	query := url.Values{}
	set(query, "accountId", request.AccountID)
	set(query, "subAccountId", request.SubAccountID)
	return do[dto.GetOpenDerivativePositionsResponse](s, ctx, CapabilityGetOpenDerivativePositions, "GET", "/khronos/v1/account/portfolio/status", query, nil)
}
