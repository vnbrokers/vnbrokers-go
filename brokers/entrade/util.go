package entrade

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/vnbrokers/vnbrokers-go/domain"
)

func rawPayload(data any, raw []byte) domain.RawPayload {
	return domain.RawPayload{Source: "entrade", Data: data, Bytes: raw}
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
