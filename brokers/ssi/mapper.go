package ssi

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func MapStockBalance(balance StockAccountBalance) domain.Balance {
	return domain.Balance{
		AccountID:     balance.Account,
		CashAvailable: balance.Withdrawable,
		CashTotal:     balance.CashBalance,
		BuyingPower:   balance.PurchasingPower,
		Currency:      "VND",
		Raw:           rawPayload(balance, nil),
	}
}

func MapOrders(orders []Order) []domain.Order {
	out := make([]domain.Order, 0, len(orders))
	for _, order := range orders {
		out = append(out, MapOrder(order))
	}
	return out
}

func MapOrder(order Order) domain.Order {
	return domain.Order{
		Broker:   "ssi",
		OrderID:  order.OrderID,
		Symbol:   order.InstrumentID,
		Side:     MapSide(order.BuySell),
		Type:     MapOrderType(order.OrderType),
		Quantity: decimalValue(order.Quantity),
		Status:   MapOrderStatus(order.OrderStatus),
		Price:    order.Price,
		Raw:      rawPayload(order, nil),
	}
}

func MapPlaceOrderResponse(response OrderRequestResponse) domain.PlaceOrderResponse {
	return domain.PlaceOrderResponse{
		OrderID: response.OrderID,
		Status:  domain.OrderStatusPending,
		Raw:     rawPayload(response, nil),
	}
}

func MapStockPortfolios(accountID string, portfolios []StockPortfolio) []domain.Position {
	positions := []domain.Position{}
	for _, portfolio := range portfolios {
		for _, position := range portfolio.StockPositions {
			positions = append(positions, MapStockPosition(accountID, position))
		}
	}
	return positions
}

func MapStockPosition(accountID string, position StockPosition) domain.Position {
	quantity := decimalValue(position.OnHand)
	marketValue := position.MarketPrice
	if marketValue != nil {
		value := quantity.Mul(*marketValue)
		marketValue = &value
	}
	return domain.Position{
		AccountID:         accountID,
		Symbol:            position.InstrumentID,
		Quantity:          quantity,
		AvailableQuantity: quantity,
		AveragePrice:      position.AveragePrice,
		MarketValue:       marketValue,
		Raw:               rawPayload(position, nil),
	}
}

func MapSide(raw string) domain.OrderSide {
	switch strings.ToUpper(raw) {
	case "B":
		return domain.OrderSideBuy
	case "S":
		return domain.OrderSideSell
	default:
		return ""
	}
}

func MapOrderType(raw string) domain.OrderType {
	switch strings.ToUpper(raw) {
	case "LO":
		return domain.OrderTypeLimit
	default:
		return domain.OrderTypeUnknown
	}
}

func MapOrderStatus(raw string) domain.OrderStatus {
	normalized := strings.NewReplacer("_", "", "-", "", " ", "").Replace(strings.ToUpper(raw))
	switch normalized {
	case "QU", "QUEUE", "PENDING":
		return domain.OrderStatusPending
	case "RS", "ACCEPTED", "NEW":
		return domain.OrderStatusAccepted
	case "PF", "PARTIALLYFILLED":
		return domain.OrderStatusPartiallyFilled
	case "FF", "FILLED", "FULLFILLED":
		return domain.OrderStatusFilled
	case "PC", "PENDINGCANCEL":
		return domain.OrderStatusPendingCancel
	case "CA", "CANCELLED", "CANCELED":
		return domain.OrderStatusCancelled
	case "RJ", "REJECTED":
		return domain.OrderStatusRejected
	default:
		return domain.OrderStatusUnknown
	}
}

func ssiSide(side domain.OrderSide) string {
	switch side {
	case domain.OrderSideBuy:
		return "B"
	case domain.OrderSideSell:
		return "S"
	default:
		return ""
	}
}

func ssiOrderType(orderType domain.OrderType) string {
	switch orderType {
	case domain.OrderTypeLimit:
		return "LO"
	case domain.OrderTypeMarket:
		return "MP"
	default:
		return string(orderType)
	}
}

func decimalValue(value *decimal.Decimal) decimal.Decimal {
	if value == nil {
		return decimal.Zero
	}
	return *value
}

func decimalFrom(value any) decimal.Decimal {
	if value == nil {
		return decimal.Zero
	}
	out, err := decimal.NewFromString(fmt.Sprint(value))
	if err != nil {
		return decimal.Zero
	}
	return out
}

func optionalDecimal(value any) *decimal.Decimal {
	if value == nil {
		return nil
	}
	out := decimalFrom(value)
	return &out
}

func numberValue(value *decimal.Decimal) any {
	if value == nil {
		return nil
	}
	if value.Equal(value.Truncate(0)) {
		return value.IntPart()
	}
	float, _ := value.Float64()
	return float
}
