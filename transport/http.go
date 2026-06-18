package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type HTTPRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	JSON    any
}

type HTTPResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       any
	Raw        []byte
}

type HTTPTransport interface {
	Send(context.Context, HTTPRequest) (HTTPResponse, error)
}

type HTTPClient struct {
	Client *http.Client
}

func NewHTTPClient(client *http.Client) *HTTPClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &HTTPClient{Client: client}
}

func (c *HTTPClient) Send(ctx context.Context, request HTTPRequest) (HTTPResponse, error) {
	var body io.Reader
	if request.JSON != nil {
		payload, err := json.Marshal(request.JSON)
		if err != nil {
			return HTTPResponse{}, err
		}
		body = bytes.NewReader(payload)
	}
	httpRequest, err := http.NewRequestWithContext(ctx, request.Method, request.URL, body)
	if err != nil {
		return HTTPResponse{}, err
	}
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}
	if httpRequest.Header.Get("Content-Type") == "" {
		httpRequest.Header.Set("Content-Type", "application/json")
	}
	httpResponse, err := c.Client.Do(httpRequest)
	if err != nil {
		return HTTPResponse{}, err
	}
	defer httpResponse.Body.Close()
	raw, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return HTTPResponse{}, err
	}
	var decoded any
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &decoded); err != nil {
			decoded = string(raw)
		}
	}
	headers := make(map[string]string, len(httpResponse.Header))
	for key, values := range httpResponse.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}
	return HTTPResponse{
		StatusCode: httpResponse.StatusCode,
		Headers:    headers,
		Body:       decoded,
		Raw:        raw,
	}, nil
}
