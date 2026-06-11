package dnse

import (
	"context"
	"fmt"

	"github.com/vnbrokers/vnbrokers-go/realtime"
)

func startMarketDataSubscription[T any](ctx context.Context, broker *Broker, channel string, symbols []string, publisher streamPublisher[T]) (realtime.Subscription[T], error) {
	return startRealtimeSubscription(ctx, broker, buildMarketDataSubscribeMessage(channel, symbols), isMarketDataPayload, publisher)
}
func buildMarketDataChannel(kind, boardID, resolution, indexName, encoding string) string {
	if encoding == "" {
		encoding = "msgpack"
	}
	switch kind {
	case "tick", "tick_extra", "top_price", "expected_price", "security_definition", "foreign":
		if boardID == "" {
			boardID = "G1"
		}
		return fmt.Sprintf("%s.%s.%s", kind, boardID, encoding)
	case "ohlc", "ohlc_closed":
		if resolution == "" {
			return ""
		}
		return fmt.Sprintf("%s.%s.%s", kind, resolution, encoding)
	case "market_index":
		if indexName == "" {
			return ""
		}
		return fmt.Sprintf("market_index.%s.%s", indexName, encoding)
	default:
		return ""
	}
}
func buildMarketDataSubscribeMessage(channel string, symbols []string) map[string]any {
	item := map[string]any{"name": channel}
	if symbols != nil {
		values := make([]any, len(symbols))
		for i, symbol := range symbols {
			values[i] = symbol
		}
		item["symbols"] = values
	}
	return map[string]any{"action": "subscribe", "channels": []any{item}}
}
func marketDataMessageData(message map[string]any) map[string]any {
	if data, ok := message["data"].(map[string]any); ok {
		return data
	}
	return message
}
func isMarketDataPayload(message map[string]any) bool {
	data := marketDataMessageData(message)
	return data["T"] != nil || data["symbol"] != nil
}
