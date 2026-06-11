package dnse

import (
	"testing"
	"time"
)

func TestConfigDefaultsRealtimeValues(t *testing.T) {
	config := Config{}.withDefaults()
	if config.StreamURL != "wss://ws-openapi.dnse.com.vn/v1/stream?encoding=msgpack" {
		t.Fatalf("stream URL = %q", config.StreamURL)
	}
	if config.StreamEncoding != "msgpack" {
		t.Fatalf("encoding = %q", config.StreamEncoding)
	}
	if config.StreamPongInterval != 30*time.Second {
		t.Fatalf("pong interval = %s", config.StreamPongInterval)
	}
}
