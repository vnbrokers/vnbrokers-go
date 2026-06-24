package dto

type GetOrderHistoryRequest struct {
	SubAccountID string
	FromDate     string
	ToDate       string
	Page         int64
	OrderStatus  string
	Symbol       string
}
type GetOrderHistoryResponse OrderHistoryResult

type GetOrderBookDetailRequest struct {
	SubAccountID string
	OrderID      string
	CacheControl string
}
type GetOrderBookDetailResponse OrderBookEntry

type GetOrderBookRequest struct {
	SubAccountID string
	CacheControl string
}
type GetOrderBookResponse []OrderBookEntry

type PlaceOrderRequest struct {
	SubAccountID string
	Body         CreateOrderRequest
}
type PlaceOrderResponse []OrderResponse

type CancelOrderOperationRequest struct {
	SubAccountID string
	OrderID      string
	Body         CancelOrderRequest
}
type CancelOrderOperationResponse []OrderResponse

type UpdateOrderOperationRequest struct {
	SubAccountID string
	OrderID      string
	Body         UpdateOrderRequest
}
type UpdateOrderOperationResponse []OrderResponse
