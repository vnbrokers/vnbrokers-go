package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) Dividend(ctx context.Context, accountID string) (*dto.DividendResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.DividendResponse](ctx, s, CapabilityDividend, "/api/v2/ors/dividend", params)
}

func (s *service) ExercisableQuantity(ctx context.Context, accountID string) (*dto.ExercisableQuantityResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	return get[dto.ExercisableQuantityResponse](ctx, s, CapabilityExercisableQuantity, "/api/v2/ors/exercisableQuantity", params)
}

func (s *service) RightsHistories(ctx context.Context, accountID string, startDate string, endDate string) (*dto.RightsHistoriesResponse, error) {
	params := url.Values{}
	params.Set("account", accountID)
	params.Set("startDate", startDate)
	params.Set("endDate", endDate)
	return get[dto.RightsHistoriesResponse](ctx, s, CapabilityRightsHistories, "/api/v2/ors/histories", params)
}

func (s *service) CreateRight(ctx context.Context, body map[string]any) (*dto.RightsCreateResponse, error) {
	return post[dto.RightsCreateResponse](ctx, s, CapabilityCreateRight, "/api/v2/ors/create", body)
}
