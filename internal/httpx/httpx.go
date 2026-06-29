package httpx

import (
	"encoding/json"
	"net/url"
	"strings"

	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

// URL joins a base URL, path, and optional query values for adapter HTTP requests.
func URL(baseURL, path string, query url.Values) string {
	endpoint := strings.TrimRight(baseURL, "/") + path
	if encoded := query.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	return endpoint
}

// DecodeResponse decodes a transport response body into the requested DTO type.
func DecodeResponse[T any](broker, operation, message string, response transport.HTTPResponse) (*T, error) {
	payload := response.Raw
	if len(payload) == 0 {
		var err error
		payload, err = json.Marshal(response.Body)
		if err != nil {
			return nil, err
		}
	}

	result := new(T)
	if err := json.Unmarshal(payload, result); err != nil {
		return nil, sdkerrors.Decode(broker, operation, message, response.Body, err)
	}
	return result, nil
}
