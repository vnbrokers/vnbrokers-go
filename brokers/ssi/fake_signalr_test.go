package ssi

import (
	"context"
	"encoding/json"
	"sync"
)

type fakeSignalRClient struct {
	mu       sync.Mutex
	handlers map[string]map[string]func([]json.RawMessage)
	onError  func(error)
	headers  map[string]string
	query    map[string]string
	invokes  []fakeSignalRInvoke
	closed   bool
}

type fakeSignalRInvoke struct {
	hub    string
	method string
	args   []any
}

func newFakeSignalRClient() *fakeSignalRClient {
	return &fakeSignalRClient{
		handlers: map[string]map[string]func([]json.RawMessage){},
		headers:  map[string]string{},
		query:    map[string]string{},
	}
}

func (c *fakeSignalRClient) Connect(context.Context) error { return nil }
func (c *fakeSignalRClient) Invoke(hub string, method string, args ...any) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.invokes = append(c.invokes, fakeSignalRInvoke{hub: hub, method: method, args: args})
	return nil
}
func (c *fakeSignalRClient) On(hub string, method string, fn func([]json.RawMessage)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.handlers[hub] == nil {
		c.handlers[hub] = map[string]func([]json.RawMessage){}
	}
	c.handlers[hub][method] = fn
}
func (c *fakeSignalRClient) OnError(fn func(error))             { c.onError = fn }
func (c *fakeSignalRClient) SetHeader(key string, value string) { c.headers[key] = value }
func (c *fakeSignalRClient) SetQuery(key string, value string)  { c.query[key] = value }
func (c *fakeSignalRClient) Close() error                       { c.closed = true; return nil }
