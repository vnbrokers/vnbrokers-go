package ssi

import "context"

type TradingRightsService struct {
	broker *Broker
}

func (s *TradingRightsService) Dividends(
	ctx context.Context,
	accountID string,
) (TradingResponse[DividendsData], error) {
	var response TradingResponse[DividendsData]
	err := s.broker.getAndDecode(ctx, "trading.rights.dividends", "/api/v2/ors/dividend", accountParams(accountID), &response)
	return response, err
}

func (s *TradingRightsService) ExercisableQuantity(
	ctx context.Context,
	accountID string,
) (TradingResponse[ExercisableQuantitiesData], error) {
	var response TradingResponse[ExercisableQuantitiesData]
	err := s.broker.getAndDecode(ctx, "trading.rights.exercisable_quantity", "/api/v2/ors/exercisableQuantity", accountParams(accountID), &response)
	return response, err
}

func (s *TradingRightsService) Histories(
	ctx context.Context,
	request RightsHistoriesRequest,
) (TradingResponse[RightsHistoriesData], error) {
	params := accountParams(request.AccountID)
	addDateRange(params, request.StartDate, request.EndDate)
	var response TradingResponse[RightsHistoriesData]
	err := s.broker.getAndDecode(ctx, "trading.rights.histories", "/api/v2/ors/histories", params, &response)
	return response, err
}

func (s *TradingRightsService) CreateWithRequest(
	ctx context.Context,
	request RightsCreateRequest,
) (TradingResponse[TransactionResponse], error) {
	body := map[string]any{
		"account":       request.AccountID,
		"instrumentID":  request.Symbol,
		"entitlementID": request.EntitlementID,
		"quantity":      numberValue(&request.Quantity),
		"amount":        numberValue(&request.Amount),
		"code":          request.Code,
	}
	var response TradingResponse[TransactionResponse]
	err := s.broker.postAndDecode(ctx, "trading.rights.create", "/api/v2/ors/create", body, &response)
	return response, err
}
