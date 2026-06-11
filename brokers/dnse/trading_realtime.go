package dnse

func buildStreamSubscribeOrdersMessage(marketType, encoding string) map[string]any {
	return buildTradingSubscribeMessage("order", marketType, encoding)
}
func buildStreamSubscribePositionsMessage(marketType, encoding string) map[string]any {
	return buildTradingSubscribeMessage("position", marketType, encoding)
}
func buildTradingSubscribeMessage(kind, marketType, encoding string) map[string]any {
	if encoding == "" {
		encoding = "msgpack"
	}
	return map[string]any{"action": "subscribe", "channels": []any{map[string]any{"name": kind + "." + marketType + "." + encoding, "symbols": []any{}}}}
}
func tradingMessageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		return data
	}
	if message["T"] == "eo" {
		if order, ok := message["order"].(map[string]any); ok {
			return order
		}
	}
	if message["T"] == "ep" {
		if position, ok := message["position"].(map[string]any); ok {
			return position
		}
	}
	if order, ok := message["order"].(map[string]any); ok {
		return order
	}
	if position, ok := message["position"].(map[string]any); ok {
		return position
	}
	return message
}
func isTradingStreamPayload(message map[string]any) bool {
	if message["T"] == "eo" {
		_, ok := message["order"].(map[string]any)
		return ok
	}
	if message["T"] == "ep" {
		_, ok := message["position"].(map[string]any)
		return ok
	}
	data := tradingMessageData(message)
	return data["accountNo"] != nil && (data["id"] != nil || data["orderId"] != nil || data["symbol"] != nil)
}
