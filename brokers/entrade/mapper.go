package entrade

import "github.com/vnbrokers/vnbrokers-go/domain"

func MapDerivatives(payload map[string]any) []domain.Symbol {
	items, _ := payload["data"].([]any)
	symbols := make([]domain.Symbol, 0, len(items))
	for _, item := range items {
		derivative, ok := item.(map[string]any)
		if !ok || derivative["symbol"] == nil {
			continue
		}
		symbols = append(symbols, domain.Symbol{
			Symbol:      stringify(derivative["symbol"]),
			DisplayName: stringify(derivative["type"]),
			Raw:         rawPayload(derivative, nil),
		})
	}
	return symbols
}
