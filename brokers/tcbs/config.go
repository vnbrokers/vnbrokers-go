package tcbs

import (
	"net/http"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	ProductionBaseURL = "https://openapi.tcbs.com.vn"
	SITBaseURL        = "https://openapisit.tcbs.com.vn"
)

type Config struct {
	BaseURL       string
	AccessToken   string
	HTTPClient    *http.Client
	HTTPTransport transport.HTTPTransport
}

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = ProductionBaseURL
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	return c
}
