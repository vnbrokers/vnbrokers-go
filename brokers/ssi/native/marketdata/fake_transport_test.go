package marketdata_test

import (
	"context"
	"sync"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type fakeHTTPTransport struct {
	mu        sync.Mutex
	requests  []transport.HTTPRequest
	responses []transport.HTTPResponse
}

func (f *fakeHTTPTransport) Send(_ context.Context, request transport.HTTPRequest) (transport.HTTPResponse, error) {
	f.mu.Lock()
	f.requests = append(f.requests, request)
	f.mu.Unlock()
	var response transport.HTTPResponse
	if len(f.responses) > 0 {
		response = f.responses[0]
		f.responses = f.responses[1:]
	}
	return response, nil
}
