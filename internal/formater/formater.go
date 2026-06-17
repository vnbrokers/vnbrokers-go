package formater

import (
	"encoding/json"
	"fmt"
)

func Pretty(v any) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}
	return string(data)
}
