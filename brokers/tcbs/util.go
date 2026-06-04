package tcbs

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func rawPayload(data any, raw []byte) domain.RawPayload {
	return domain.RawPayload{Source: "tcbs", Data: data, Bytes: raw}
}

func expectObject(value any) map[string]any {
	object, ok := value.(map[string]any)
	if !ok {
		return map[string]any{}
	}
	return object
}

func stringify(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func cloneHeaders(headers map[string]string) map[string]string {
	out := make(map[string]string, len(headers))
	for key, value := range headers {
		out[key] = value
	}
	return out
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
