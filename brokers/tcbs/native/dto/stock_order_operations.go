package dto

type PlaceStockOrderRequest struct {
	AccountNo string `json:"-"`
	PlaceStockOrderBody
}

type UpdateStockOrderRequest struct {
	AccountNo string `json:"-"`
	OrderID   string `json:"-"`
	UpdateStockOrderBody
}

type CancelStockOrderRequest struct {
	AccountNo string `json:"-"`
	CancelStockOrderBody
}

type GetStockOrdersRequest struct {
	AccountNo string
}

type GetStockOrdersResponse StockOrderPage

type GetStockOrderRequest struct {
	AccountNo string
	OrderID   string
}

type GetStockOrderResponse StockOrderPage

type GetStockMatchingDetailsRequest struct {
	AccountNo string
}

type GetStockMatchingDetailsResponse StockMatchingDetails
