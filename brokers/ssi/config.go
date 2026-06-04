package ssi

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Config struct {
	BaseURL        string
	ConsumerID     string
	ConsumerSecret string
	AccessToken    string
	PrivateKey     string
	DeviceID       string
	UserAgent      string
	ChannelID      string
	MarketID       string
	HTTPClient     *http.Client
	HTTPTransport  transport.HTTPTransport
	RequestID      func() string
}

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = "https://fc-tradeapi.ssi.com.vn"
	}
	if c.ChannelID == "" {
		c.ChannelID = "TA"
	}
	if c.MarketID == "" {
		c.MarketID = "VN"
	}
	if c.UserAgent == "" {
		c.UserAgent = "FCTrading"
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	if c.RequestID == nil {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		c.RequestID = func() string {
			return strconv.Itoa(rng.Intn(100000000))
		}
	}
	return c
}
