package ssi

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/vnbrokers/vnbrokers-go/internal/signalr"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

type SignalRClient interface {
	Connect(context.Context) error
	Invoke(string, string, ...any) error
	On(string, string, func([]json.RawMessage))
	OnError(func(error))
	SetHeader(string, string)
	SetQuery(string, string)
	Close() error
}

type SignalRFactory func(baseURL string, hubs []string) SignalRClient

type Config struct {
	BaseURL               string
	DataBaseURL           string
	ConsumerID            string
	DataConsumerSecret    string
	TradingConsumerSecret string
	DataToken             string
	TradingToken          string
	PrivateKey            string
	DeviceID              string
	UserAgent             string
	ChannelID             string
	MarketID              string
	HTTPClient            *http.Client
	HTTPTransport         transport.HTTPTransport
	RequestID             func() string
	TradingStreamURL      string
	MarketDataStreamURL   string
	SignalRFactory        SignalRFactory
}

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = "https://fc-tradeapi.ssi.com.vn"
	}
	if c.DataBaseURL == "" {
		c.DataBaseURL = "https://fc-data.ssi.com.vn"
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
	if c.TradingStreamURL == "" {
		c.TradingStreamURL = "https://fc-tradehub.ssi.com.vn/v2.0/signalr"
	}
	if c.MarketDataStreamURL == "" {
		c.MarketDataStreamURL = "https://fc-datahub.ssi.com.vn/v2.0/signalr"
	}
	if c.SignalRFactory == nil {
		c.SignalRFactory = func(baseURL string, hubs []string) SignalRClient {
			return signalr.NewClient(baseURL, hubs)
		}
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
