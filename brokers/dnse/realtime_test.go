package dnse

import "testing"

func TestMarketDataSubscribeMessageSupportsMultipleSymbols(t *testing.T) {
	message := BuildMarketDataSubscribeMessage("top_price.G1.json", []string{"ACB", "HPG"})
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)
	symbols := channel["symbols"].([]any)

	if len(symbols) != 2 {
		t.Fatalf("symbols len = %d", len(symbols))
	}
}

func TestMarketDataSubscribeMessageSupportsAllSymbols(t *testing.T) {
	message := BuildMarketDataSubscribeMessage("top_price.G1.json", nil)
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)

	if _, ok := channel["symbols"]; ok {
		t.Fatalf("all-symbol subscribe should omit symbols")
	}
}

func TestTradingSubscribeOrdersMessageUsesEncoding(t *testing.T) {
	message := BuildStreamSubscribeOrdersMessage("STOCK", "msgpack")
	channels := message["channels"].([]any)
	channel := channels[0].(map[string]any)

	if channel["name"] != "order.STOCK.msgpack" {
		t.Fatalf("channel = %s", channel["name"])
	}
}
