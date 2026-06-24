package fhsc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func expectObject(value any) map[string]any {
	if value == nil {
		return nil
	}
	object, _ := value.(map[string]any)
	return object
}

func stringify(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func cloneHeaders(headers map[string]string) map[string]string {
	if len(headers) == 0 {
		return map[string]string{}
	}
	out := make(map[string]string, len(headers))
	for key, value := range headers {
		out[key] = value
	}
	return out
}

func randomNonce() string {
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		return ""
	}
	return hex.EncodeToString(buffer)
}
