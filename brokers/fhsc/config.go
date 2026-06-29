package fhsc

import (
	"net/http"
	"time"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

const ProductionBaseURL = "https://open-api.fhsc.com.vn"

type Config struct {
	BaseURL                string
	APIKey                 string
	APISecret              string
	UserID                 int64
	TwoFactorToken         string
	OpenAPISkillVersion    string
	OpenAPIOperatingSystem string
	OpenAPIAgent           string
	UserAgent              string
	Now                    func() time.Time
	Nonce                  func() string
	HTTPClient             *http.Client
	HTTPTransport          transport.HTTPTransport
}

func (Config) BrokerName() string { return "fhsc" }

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = ProductionBaseURL
	}
	if c.OpenAPISkillVersion == "" {
		c.OpenAPISkillVersion = "vnbrokers-go"
	}
	if c.OpenAPIOperatingSystem == "" {
		c.OpenAPIOperatingSystem = "linux"
	}
	if c.OpenAPIAgent == "" {
		c.OpenAPIAgent = "vnbrokers-go"
	}
	if c.UserAgent == "" {
		c.UserAgent = "vnbrokers-go/fhsc"
	}
	if c.Now == nil {
		c.Now = time.Now
	}
	if c.Nonce == nil {
		c.Nonce = randomNonce
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	return c
}
