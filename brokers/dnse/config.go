package dnse

import (
	"net/http"
	"time"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Config struct {
	BaseURL               string
	StreamURL             string
	APIKey                string
	APISecret             string
	TradingToken          string
	StreamEncoding        string
	StreamPongInterval    time.Duration
	MarketType            string
	OrderCategory         string
	LoanPackageID         *int
	PositionsPageSize     int
	MarketDataSymbolLimit int
	MarketDataBoardID     string
	CandleMarketType      string
	CandleLookbackSeconds int64
	HTTPClient            *http.Client
	HTTPTransport         transport.HTTPTransport
	WebSocketFactory      transport.WebSocketFactory
}

func (Config) BrokerName() string { return "dnse" }

func (c Config) withDefaults() Config {
	if c.BaseURL == "" {
		c.BaseURL = "https://openapi.dnse.com.vn"
	}
	if c.StreamURL == "" {
		c.StreamURL = "wss://ws-openapi.dnse.com.vn/v1/stream?encoding=msgpack"
	}
	if c.StreamEncoding == "" {
		c.StreamEncoding = "msgpack"
	}
	if c.StreamPongInterval == 0 {
		c.StreamPongInterval = 30 * time.Second
	}
	if c.MarketType == "" {
		c.MarketType = "DERIVATIVE"
	}
	if c.OrderCategory == "" {
		c.OrderCategory = "NORMAL"
	}
	if c.PositionsPageSize == 0 {
		c.PositionsPageSize = 20
	}
	if c.MarketDataSymbolLimit == 0 {
		c.MarketDataSymbolLimit = 1000
	}
	if c.MarketDataBoardID == "" {
		c.MarketDataBoardID = "G1"
	}
	if c.CandleMarketType == "" {
		c.CandleMarketType = "STOCK"
	}
	if c.CandleLookbackSeconds == 0 {
		c.CandleLookbackSeconds = 86400
	}
	if c.HTTPTransport == nil {
		c.HTTPTransport = transport.NewHTTPClient(c.HTTPClient)
	}
	if c.WebSocketFactory == nil {
		c.WebSocketFactory = transport.ConnectWebSocket
	}
	return c
}
