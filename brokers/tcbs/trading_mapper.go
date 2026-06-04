package tcbs

import (
	"strings"

	"github.com/vnbrokers/vnbrokers-go/domain"
)

func MapOrderInfo(order OrderInfo) domain.Order {
	price := optionalDecimal(order.LimitPrice)
	return domain.Order{
		Broker:    "tcbs",
		AccountID: order.AccountNo,
		OrderID:   order.OrderID,
		Symbol:    order.Symbol,
		Side:      mapOrderSide(order.ExecType),
		Type:      mapOrderType(order.PriceType),
		Quantity:  decimalFrom(order.OrderQtty),
		Status:    mapOrderStatus(order.OrStatus),
		Price:     price,
		Raw:       rawPayload(order, nil),
	}
}

func MapOrders(response OrderSearchResponse) []domain.Order {
	orders := make([]domain.Order, 0, len(response.Data))
	for _, order := range response.Data {
		orders = append(orders, MapOrderInfo(order))
	}
	return orders
}

func MapStockHoldings(accountNo string, response SeInfoDTO) []domain.Position {
	if accountNo == "" {
		accountNo = response.AccountNo
	}
	positions := make([]domain.Position, 0, len(response.Stock))
	for _, stock := range response.Stock {
		currentPrice := optionalDecimal(stock.CurrentPrice)
		positions = append(positions, domain.Position{
			AccountID:         accountNo,
			Symbol:            stock.Symbol,
			Quantity:          decimalFrom(stock.TotalQtty),
			AvailableQuantity: decimalFrom(stock.AvailableTrading),
			AveragePrice:      optionalDecimal(stock.CostPrice),
			MarketValue:       currentPrice,
			Raw:               rawPayload(stock, nil),
		})
	}
	return positions
}

func MapCashBalance(accountNo string, response CashInvestmentResponse) domain.Balance {
	balance := domain.Balance{
		AccountID: accountNo,
		Currency:  "VND",
		Raw:       rawPayload(response, nil),
	}
	if len(response.Data) == 0 {
		return balance
	}
	item := response.Data[0]
	if item.AccountNo != "" {
		balance.AccountID = item.AccountNo
	}
	balance.CashTotal = optionalDecimal(item.CashBalance)
	balance.CashAvailable = optionalDecimal(item.AvlWithdraw)
	balance.BuyingPower = optionalDecimal(item.PP0)
	return balance
}

func MapStockMatchEvent(message map[string]any) domain.OrderEvent {
	return domain.OrderEvent{
		Broker:         "tcbs",
		AccountID:      stringify(message["accountNo"]),
		OrderID:        firstString(message, "orderID", "orderId", "orderNo"),
		Symbol:         stringify(message["symbol"]),
		Status:         mapOrderStatus(firstString(message, "orStatus", "status", "orderStatus")),
		RawStatus:      firstString(message, "orStatus", "status", "orderStatus"),
		FilledQuantity: firstString(message, "execQtty", "qtty", "matchVolume"),
		ReceivedAt:     firstString(message, "txTime", "txtime", "time", "timeExec"),
		Raw:            rawPayload(message, nil),
	}
}

func MapStockOrderEvent(message StockOrderRealtimeMessage) domain.OrderEvent {
	return domain.OrderEvent{
		Broker:         "tcbs",
		AccountID:      message.AccountNo,
		OrderID:        message.OrderID,
		Symbol:         message.Symbol,
		Status:         mapOrderStatus(message.OrStatus),
		RawStatus:      message.OrStatus,
		FilledQuantity: "",
		ReceivedAt:     message.TxTime,
		Raw:            rawPayload(message, nil),
	}
}

func mapOrderSide(value string) domain.OrderSide {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "NB", "B", "BUY":
		return domain.OrderSideBuy
	case "NS", "S", "SELL":
		return domain.OrderSideSell
	default:
		return ""
	}
}

func mapOrderType(value string) domain.OrderType {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "LO", "LIMIT":
		return domain.OrderTypeLimit
	case "MP", "MTL", "MOK", "MAK", "ATO", "ATC", "MARKET":
		return domain.OrderTypeMarket
	default:
		return domain.OrderTypeUnknown
	}
}

func mapOrderStatus(value string) domain.OrderStatus {
	normalized := strings.ToLower(strings.TrimSpace(value))
	switch {
	case normalized == "":
		return domain.OrderStatusUnknown
	case normalized == "8" || normalized == "11" || normalized == "a":
		return domain.OrderStatusPending
	case normalized == "c":
		return domain.OrderStatusPendingCancel
	case normalized == "2" || normalized == "10" || normalized == "s":
		if normalized == "s" {
			return domain.OrderStatusFilled
		}
		return domain.OrderStatusAccepted
	case normalized == "3":
		return domain.OrderStatusCancelled
	case normalized == "4" || normalized == "12":
		return domain.OrderStatusFilled
	case normalized == "0":
		return domain.OrderStatusRejected
	case normalized == "5":
		return domain.OrderStatusUnknown
	case strings.Contains(normalized, "match") || strings.Contains(normalized, "filled"):
		return domain.OrderStatusFilled
	case strings.Contains(normalized, "cancel"):
		return domain.OrderStatusCancelled
	case strings.Contains(normalized, "reject"):
		return domain.OrderStatusRejected
	case strings.Contains(normalized, "pending"):
		return domain.OrderStatusPending
	default:
		return domain.OrderStatusUnknown
	}
}

func firstString(message map[string]any, keys ...string) string {
	for _, key := range keys {
		if value := stringify(message[key]); value != "" {
			return value
		}
	}
	return ""
}
