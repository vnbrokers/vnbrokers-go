package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/ssi/native/dto"
)

func (s *service) FcoNewOrder(ctx context.Context, body map[string]any) (*dto.ConditionalOrderResponse, error) {
	return post[dto.ConditionalOrderResponse](ctx, s, CapabilityFcoNewOrder, "/api/v2/fco/neworder", body)
}

func (s *service) FcoCancelOrder(ctx context.Context, body map[string]any) (*dto.ConditionalOrderResponse, error) {
	return post[dto.ConditionalOrderResponse](ctx, s, CapabilityFcoCancelOrder, "/api/v2/fco/cancelorder", body)
}

func (s *service) FcoOrderBook(ctx context.Context, fcoID string, pageIndex int, pageSize int) (*dto.ConditionalOrderPage[dto.ConditionalTriggeredOrder], error) {
	params := url.Values{}
	setOptionalString(params, "fcoId", fcoID)
	setOptionalInt(params, "pageIndex", pageIndex)
	setOptionalInt(params, "pageSize", pageSize)
	return get[dto.ConditionalOrderPage[dto.ConditionalTriggeredOrder]](ctx, s, CapabilityFcoOrderBook, "/api/v2/fco/orderbook", params)
}

func (s *service) FcoStatusHistory(ctx context.Context, fcoID string, pageIndex int, pageSize int) (*dto.ConditionalOrderPage[dto.ConditionalOrderStatus], error) {
	params := url.Values{}
	setOptionalString(params, "fcoId", fcoID)
	setOptionalInt(params, "pageIndex", pageIndex)
	setOptionalInt(params, "pageSize", pageSize)
	return get[dto.ConditionalOrderPage[dto.ConditionalOrderStatus]](ctx, s, CapabilityFcoStatusHistory, "/api/v2/fco/statusHistory", params)
}

func (s *service) FcoList(ctx context.Context, params url.Values) (*dto.ConditionalOrderPage[dto.ConditionalOrder], error) {
	if params == nil {
		params = url.Values{}
	}
	return get[dto.ConditionalOrderPage[dto.ConditionalOrder]](ctx, s, CapabilityFcoList, "/api/v2/fco/list", params)
}

func fcoPageParams(fcoID string, pageIndex int, pageSize int) url.Values {
	params := url.Values{}
	setOptionalString(params, "fcoId", fcoID)
	setOptionalInt(params, "pageIndex", pageIndex)
	setOptionalInt(params, "pageSize", pageSize)
	return params
}
