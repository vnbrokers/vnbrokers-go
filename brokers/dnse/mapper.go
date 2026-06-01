package dnse

import (
	"strings"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func MapAccounts(payload map[string]any) []domain.Account {
	rawAccounts, _ := payload["accounts"].([]any)
	displayName := stringify(payload["name"])
	accounts := make([]domain.Account, 0, len(rawAccounts))
	for _, item := range rawAccounts {
		account, ok := item.(map[string]any)
		if !ok || account["id"] == nil {
			continue
		}
		accounts = append(accounts, domain.Account{
			Broker:      "dnse",
			AccountID:   stringify(account["id"]),
			DisplayName: displayName,
			Raw:         rawPayload(payload, nil),
		})
	}
	return accounts
}

func MapBalance(accountID string, payload map[string]any) domain.Balance {
	stock, _ := payload["stock"].(map[string]any)
	available := optionalDecimal(stock["availableCash"])
	return domain.Balance{
		AccountID:     accountID,
		CashAvailable: available,
		CashTotal:     optionalDecimal(stock["totalCash"]),
		BuyingPower:   available,
		Currency:      "VND",
		Raw:           rawPayload(payload, nil),
	}
}

func MapOrder(payload map[string]any) domain.Order {
	return domain.Order{
		Broker:    "dnse",
		AccountID: stringify(payload["accountNo"]),
		OrderID:   stringify(payload["id"]),
		Symbol:    stringify(payload["symbol"]),
		Side:      MapSide(stringify(payload["side"])),
		Type:      MapOrderType(stringify(payload["orderType"])),
		Quantity:  decimalFrom(payload["quantity"]),
		Status:    MapOrderStatus(stringify(payload["orderStatus"])),
		Price:     optionalDecimal(payload["price"]),
		Raw:       rawPayload(payload, nil),
	}
}

func MapOrders(payload map[string]any) []domain.Order {
	return mapOrderList(payload["orders"])
}

func MapOrderHistory(payload map[string]any) []domain.Order {
	return mapOrderList(payload["data"])
}

func mapOrderList(value any) []domain.Order {
	items, _ := value.([]any)
	orders := make([]domain.Order, 0, len(items))
	for _, item := range items {
		order, ok := item.(map[string]any)
		if !ok {
			continue
		}
		orders = append(orders, MapOrder(order))
	}
	return orders
}

func MapPlaceOrderResponse(payload map[string]any) domain.PlaceOrderResponse {
	return domain.PlaceOrderResponse{
		OrderID: stringify(payload["id"]),
		Status:  MapOrderStatus(stringify(payload["orderStatus"])),
		Raw:     rawPayload(payload, nil),
	}
}

func MapPositions(payload map[string]any) []domain.Position {
	items, _ := payload["positions"].([]any)
	positions := make([]domain.Position, 0, len(items))
	for _, item := range items {
		position, ok := item.(map[string]any)
		if !ok {
			continue
		}
		positions = append(positions, MapPosition(position))
	}
	return positions
}

func MapPosition(payload map[string]any) domain.Position {
	if data, ok := payload["data"].(map[string]any); ok {
		payload = data
	}
	quantity := decimalFrom(payload["openQuantity"])
	if quantity.IsZero() && payload["openQuantity"] == nil {
		quantity = decimalFrom(payload["tradeQuantity"])
	}
	marketPrice := optionalDecimal(payload["marketPrice"])
	var marketValue *decimal.Decimal
	if marketPrice != nil {
		value := quantity.Mul(*marketPrice)
		marketValue = &value
	}
	return domain.Position{
		AccountID:         stringify(payload["accountNo"]),
		Symbol:            stringify(payload["symbol"]),
		Quantity:          quantity,
		AvailableQuantity: quantity,
		AveragePrice:      optionalDecimal(firstNonNil(payload["averageCostPrice"], payload["costPrice"])),
		MarketValue:       marketValue,
		Raw:               rawPayload(payload, nil),
	}
}

func MapSymbols(payload map[string]any) []domain.Symbol {
	items, _ := payload["data"].([]any)
	symbols := make([]domain.Symbol, 0, len(items))
	for _, item := range items {
		symbol, ok := item.(map[string]any)
		if !ok || symbol["symbol"] == nil {
			continue
		}
		symbols = append(symbols, domain.Symbol{
			Symbol:      stringify(symbol["symbol"]),
			Exchange:    stringify(symbol["marketId"]),
			DisplayName: stringify(firstNonNil(symbol["name"], symbol["shortName"])),
			Raw:         rawPayload(symbol, nil),
		})
	}
	return symbols
}

func MapQuote(symbol string, payload map[string]any) domain.Quote {
	trades, _ := payload["trades"].([]any)
	trade := map[string]any{}
	if len(trades) > 0 {
		trade, _ = trades[0].(map[string]any)
	}
	return domain.Quote{
		Symbol:     symbol,
		LastPrice:  optionalDecimal(trade["price"]),
		ReceivedAt: stringify(trade["time"]),
		Raw:        rawPayload(payload, nil),
	}
}

func MapCandles(symbol string, interval string, payload map[string]any) []domain.Candle {
	times, _ := payload["t"].([]any)
	opens, _ := payload["o"].([]any)
	highs, _ := payload["h"].([]any)
	lows, _ := payload["l"].([]any)
	closes, _ := payload["c"].([]any)
	volumes, _ := payload["v"].([]any)
	candles := make([]domain.Candle, 0, len(times))
	for i := range times {
		candles = append(candles, domain.Candle{
			Symbol:   symbol,
			Interval: interval,
			OpenedAt: stringify(at(times, i)),
			Open:     decimalFrom(at(opens, i)),
			High:     decimalFrom(at(highs, i)),
			Low:      decimalFrom(at(lows, i)),
			Close:    decimalFrom(at(closes, i)),
			Volume:   decimalFrom(at(volumes, i)),
			Raw:      rawPayload(payload, nil),
		})
	}
	return candles
}

func MapOrderEvent(message map[string]any) domain.OrderEvent {
	data := tradingMessageData(message)
	return domain.OrderEvent{
		Broker:         "dnse",
		AccountID:      stringify(data["accountNo"]),
		OrderID:        stringify(firstNonNil(data["id"], data["orderId"])),
		Symbol:         stringify(data["symbol"]),
		Status:         MapOrderStatus(stringify(data["orderStatus"])),
		RawStatus:      stringify(data["orderStatus"]),
		FilledQuantity: stringify(data["fillQuantity"]),
		ReceivedAt:     stringify(firstNonNil(data["createdDate"], data["modifiedDate"])),
		Raw:            rawPayload(message, nil),
	}
}

func MapSide(raw string) domain.OrderSide {
	switch raw {
	case "NB":
		return domain.OrderSideBuy
	case "NS":
		return domain.OrderSideSell
	default:
		return ""
	}
}

func MapOrderType(raw string) domain.OrderType {
	switch raw {
	case "LO":
		return domain.OrderTypeLimit
	case "MP", "MTL", "MOK", "MAK":
		return domain.OrderTypeMarket
	default:
		return domain.OrderTypeUnknown
	}
}

func MapOrderStatus(raw string) domain.OrderStatus {
	normalized := strings.NewReplacer("_", "", "-", "", " ", "").Replace(strings.ToUpper(raw))
	switch normalized {
	case "PENDING", "PENDINGNEW":
		return domain.OrderStatusPending
	case "NEW", "ACCEPTED":
		return domain.OrderStatusAccepted
	case "PARTIALLYFILLED":
		return domain.OrderStatusPartiallyFilled
	case "FILLED":
		return domain.OrderStatusFilled
	case "PENDINGCANCEL":
		return domain.OrderStatusPendingCancel
	case "CANCELLED", "CANCELED":
		return domain.OrderStatusCancelled
	case "REJECTED":
		return domain.OrderStatusRejected
	default:
		return domain.OrderStatusUnknown
	}
}

func dnseSide(side domain.OrderSide) string {
	if side == domain.OrderSideBuy {
		return "NB"
	}
	if side == domain.OrderSideSell {
		return "NS"
	}
	return ""
}

func dnseOrderType(orderType domain.OrderType) string {
	if orderType == domain.OrderTypeLimit {
		return "LO"
	}
	if orderType == domain.OrderTypeMarket {
		return "MTL"
	}
	return string(orderType)
}

func firstNonNil(values ...any) any {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

func at(values []any, index int) any {
	if index < 0 || index >= len(values) {
		return nil
	}
	return values[index]
}
