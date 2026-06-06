package signalr

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestConnectionDataNormalizesHubNames(t *testing.T) {
	client := NewClient("https://example.com/signalr", []string{" FcTradingHub ", "", "FcMarketDataV2Hub"})

	data, err := client.connectionData()
	if err != nil {
		t.Fatalf("connectionData: %v", err)
	}
	if data != `[{"name":"fctradinghub"},{"name":"fcmarketdatav2hub"}]` {
		t.Fatalf("connectionData = %s", data)
	}
}

func TestNegotiateBuildsClassicSignalRRequest(t *testing.T) {
	var gotURL *url.URL
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotURL = r.URL
		if r.Header.Get("Authorization") != "Bearer token" {
			t.Fatalf("authorization = %s", r.Header.Get("Authorization"))
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ConnectionId":      "connection-id",
			"ConnectionToken":   "connection-token",
			"ProtocolVersion":   "2.1",
			"TryWebSockets":     true,
			"KeepAliveTimeout":  20,
			"DisconnectTimeout": 30,
		})
	}))
	defer server.Close()

	client := NewClient(server.URL+"/signalr", []string{"FcTradingHub"})
	client.Headers.Set("Authorization", "Bearer token")
	client.Query.Set("access_token", "query-token")

	if err := client.Negotiate(context.Background()); err != nil {
		t.Fatalf("negotiate: %v", err)
	}
	if gotURL.Path != "/signalr/negotiate" {
		t.Fatalf("path = %s", gotURL.Path)
	}
	query := gotURL.Query()
	if query.Get("clientProtocol") != "2.1" ||
		query.Get("connectionData") != `[{"name":"fctradinghub"}]` ||
		query.Get("access_token") != "query-token" {
		t.Fatalf("query = %s", gotURL.RawQuery)
	}
	if client.connToken != "connection-token" || client.connID != "connection-id" {
		t.Fatalf("connection = %s %s", client.connToken, client.connID)
	}
}

func TestStartBuildsClassicSignalRRequest(t *testing.T) {
	var gotURL *url.URL
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotURL = r.URL
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"Response":"started"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL+"/signalr", []string{"FcTradingHub"})
	client.connToken = "connection-token"
	client.Query.Set("someKey", "someValue")

	if err := client.Start(context.Background()); err != nil {
		t.Fatalf("start: %v", err)
	}
	query := gotURL.Query()
	if gotURL.Path != "/signalr/start" ||
		query.Get("transport") != "webSockets" ||
		query.Get("connectionToken") != "connection-token" ||
		query.Get("someKey") != "someValue" {
		t.Fatalf("url = %s", gotURL.String())
	}
}

func TestDispatchCallsRegisteredHandler(t *testing.T) {
	client := NewClient("https://example.com/signalr", []string{"FcTradingHub"})
	called := false
	client.On("FcTradingHub", "orderUpdate", func(args []json.RawMessage) {
		called = true
		if string(args[0]) != `"26060500251341"` {
			t.Fatalf("arg = %s", args[0])
		}
	})

	client.dispatch([]byte(`{"M":[{"H":"fctradinghub","M":"OrderUpdate","A":["26060500251341"]}]}`))

	if !called {
		t.Fatalf("handler was not called")
	}
}

func TestNegotiateUnauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	client := NewClient(server.URL+"/signalr", []string{"FcTradingHub"})
	if err := client.Negotiate(context.Background()); err == nil {
		t.Fatalf("expected unauthorized error")
	}
}

func TestToWebSocketURL(t *testing.T) {
	got, err := toWebSocketURL("https://example.com/v2.0/signalr/connect?a=b")
	if err != nil {
		t.Fatalf("toWebSocketURL: %v", err)
	}
	if got != "wss://example.com/v2.0/signalr/connect?a=b" {
		t.Fatalf("url = %s", got)
	}
}
