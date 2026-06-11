package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) Dividend(ctx context.Context, accountID string) (*dto.TradingResponse[dto.DividendsData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.DividendsData]](ctx, s, CapabilityDividend, "/api/v2/ors/dividend", params)
}

func (s *service) ExercisableQuantity(ctx context.Context, accountID string) (*dto.TradingResponse[dto.ExercisableQuantitiesData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.TradingResponse[dto.ExercisableQuantitiesData]](ctx, s, CapabilityExercisableQuantity, "/api/v2/ors/exercisableQuantity", params)
}

func (s *service) RightsHistories(ctx context.Context, accountID string, startDate string, endDate string) (*dto.TradingResponse[dto.RightsHistoriesData], error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	return get[dto.TradingResponse[dto.RightsHistoriesData]](ctx, s, CapabilityRightsHistories, "/api/v2/ors/histories", params)
}

func (s *service) CreateRight(ctx context.Context, body map[string]any) (*dto.TradingResponse[dto.TransactionResponse], error) {
	return post[dto.TradingResponse[dto.TransactionResponse]](ctx, s, CapabilityCreateRight, "/api/v2/ors/create", body)
}
