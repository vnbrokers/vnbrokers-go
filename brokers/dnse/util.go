package dnse

import "fmt"

func stringify(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}
