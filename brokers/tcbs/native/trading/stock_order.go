package trading

import (
	"context"
	"net/url"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
)

func (s *service) PlaceStockOrder(ctx context.Context, request dto.PlaceStockOrderRequest) (*dto.PlaceStockOrderResponse, error) {
	path := "/akhlys/v1/accounts/" + escaped(request.AccountNo) + "/orders"
	return do[dto.PlaceStockOrderResponse](s, ctx, CapabilityPlaceStockOrder, "POST", path, url.Values{}, request)
}

func (s *service) UpdateStockOrder(ctx context.Context, request dto.UpdateStockOrderRequest) (*dto.UpdateStockOrderResponse, error) {
	path := "/akhlys/v1/accounts/" + escaped(request.AccountNo) + "/orders/" + escaped(request.OrderID)
	return do[dto.UpdateStockOrderResponse](s, ctx, CapabilityUpdateStockOrder, "PUT", path, url.Values{}, request)
}

func (s *service) CancelStockOrder(ctx context.Context, request dto.CancelStockOrderRequest) (*dto.CancelStockOrderResponse, error) {
	path := "/akhlys/v1/accounts/" + escaped(request.AccountNo) + "/cancel-orders"
	return do[dto.CancelStockOrderResponse](s, ctx, CapabilityCancelStockOrder, "PUT", path, url.Values{}, request)
}

func (s *service) GetStockOrders(ctx context.Context, request dto.GetStockOrdersRequest) (*dto.GetStockOrdersResponse, error) {
	path := "/aion/v1/accounts/" + escaped(request.AccountNo) + "/orders"
	return do[dto.GetStockOrdersResponse](s, ctx, CapabilityGetStockOrders, "GET", path, url.Values{}, nil)
}

func (s *service) GetStockOrder(ctx context.Context, request dto.GetStockOrderRequest) (*dto.GetStockOrderResponse, error) {
	path := "/aion/v1/accounts/" + escaped(request.AccountNo) + "/orders/" + escaped(request.OrderID)
	return do[dto.GetStockOrderResponse](s, ctx, CapabilityGetStockOrder, "GET", path, url.Values{}, nil)
}

func (s *service) GetStockMatchingDetails(ctx context.Context, request dto.GetStockMatchingDetailsRequest) (*dto.GetStockMatchingDetailsResponse, error) {
	path := "/aion/v1/accounts/" + escaped(request.AccountNo) + "/matching-details"
	return do[dto.GetStockMatchingDetailsResponse](s, ctx, CapabilityGetStockMatchingDetails, "GET", path, url.Values{}, nil)
}
