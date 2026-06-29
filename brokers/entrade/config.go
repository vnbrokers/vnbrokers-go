package entrade

import (
	"net/http"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Config struct {
	BaseURL       string
	AuthBaseURL   string
	Token         string
	HTTPClient    *http.Client
	HTTPTransport transport.HTTPTransport
}

func (Config) BrokerName() string { return "entrade" }

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = "https://services-staging.entrade.com.vn/papertrade-entrade-api"
	}
	if c.AuthBaseURL == "" {
		c.AuthBaseURL = "https://services.entrade.com.vn/entrade-api"
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	return c
}
