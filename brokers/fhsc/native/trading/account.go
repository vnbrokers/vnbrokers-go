package trading

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vnbrokers/vnbrokers-go/brokers/fhsc/native/dto"
)

func (s *service) GetAccountSummary(ctx context.Context, request dto.GetAccountSummaryRequest) (*dto.SummaryAccountResponse, error) {
	return do[dto.SummaryAccountResponse](s, ctx, CapabilityGetAccountSummary, "GET", "/trading/accounts/"+escaped(request.SubAccountID)+"/summary", url.Values{}, nil)
}

func (s *service) GetUserAssetsSummary(ctx context.Context, request dto.GetUserAssetsSummaryRequest) (*dto.UserAssetsSummaryResponse, error) {
	query := url.Values{}
	setOptional(query, "cache-control", request.CacheControl)
	path := "/users/v3/users/" + strconv.FormatInt(request.UserID, 10) + "/assets/summary"
	return do[dto.UserAssetsSummaryResponse](s, ctx, CapabilityGetUserAssetsSummary, "GET", path, query, nil)
}

func (s *service) GetPnLToday(ctx context.Context, request dto.GetPnLTodayRequest) (*dto.PnLTodayResponse, error) {
	query := url.Values{}
	setOptional(query, "sub-account-id", request.SubAccountID)
	path := "/trading/pnl-today/" + strconv.FormatInt(request.UserID, 10)
	return do[dto.PnLTodayResponse](s, ctx, CapabilityGetPnLToday, "GET", path, query, nil)
}

func (s *service) GetPortfolio(ctx context.Context, request dto.GetPortfolioRequest) (*dto.GetPortfolioResponse, error) {
	query := url.Values{}
	setOptional(query, "cache-control", request.CacheControl)
	path := "/trading/v2/sub-accounts/" + escaped(request.SubAccountID) + "/portfolio"
	return do[dto.GetPortfolioResponse](s, ctx, CapabilityGetPortfolio, "GET", path, query, nil)
}

func (s *service) GetAvailableTrade(ctx context.Context, request dto.GetAvailableTradeRequest) (*dto.AvailableTradeResult, error) {
	query := url.Values{}
	set(query, "orderSide", request.OrderSide)
	set(query, "symbol", request.Symbol)
	set(query, "quotePrice", strconv.FormatInt(request.QuotePrice, 10))
	path := "/trading/v2/accounts/" + escaped(request.SubAccountID) + "/available-trade"
	return do[dto.AvailableTradeResult](s, ctx, CapabilityGetAvailableTrade, "GET", path, query, nil)
}

func (s *service) GetUserRights(ctx context.Context, request dto.GetUserRightsRequest) (*[]dto.UserRight, error) {
	query := url.Values{}
	setOptional(query, "fromDate", request.FromDate)
	setOptional(query, "toDate", request.ToDate)
	setOptional(query, "catType", request.CatType)
	setOptional(query, "isCom", request.IsCom)
	setOptional(query, "symbol", request.Symbol)
	setOptional(query, "status", request.Status)
	path := "/trading/v5/account/" + escaped(request.SubAccountID) + "/user-rights"
	return do[[]dto.UserRight](s, ctx, CapabilityGetUserRights, "GET", path, query, nil)
}

func (s *service) GetMarketSession(ctx context.Context, request dto.GetMarketSessionRequest) (*dto.ExchangeSessionInfo, error) {
	query := url.Values{}
	set(query, "exchange", request.Exchange)
	return do[dto.ExchangeSessionInfo](s, ctx, CapabilityGetMarketSession, "GET", "/trading/market/session", query, nil)
}
