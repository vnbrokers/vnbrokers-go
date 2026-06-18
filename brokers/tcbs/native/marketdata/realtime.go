package marketdata

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/vnbrokers/vnbrokers-go/brokers/tcbs/native/dto"
	"github.com/vnbrokers/vnbrokers-go/core"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const (
	CapabilityRealtimeStockTradeHistory               core.Capability = "native.marketdata.realtime.stock_trade_history"
	CapabilityRealtimeStockSupplyDemandOneMinute      core.Capability = "native.marketdata.realtime.stock_supply_demand_one_minute"
	CapabilityRealtimeStockSupplyDemandFifteenMinutes core.Capability = "native.marketdata.realtime.stock_supply_demand_fifteen_minutes"
	CapabilityRealtimeBidPrices                       core.Capability = "native.marketdata.realtime.bid_prices"
	CapabilityRealtimeOfferPrices                     core.Capability = "native.marketdata.realtime.offer_prices"
	CapabilityRealtimeForeignTrading                  core.Capability = "native.marketdata.realtime.foreign_trading"
	CapabilityRealtimeBasePrices                      core.Capability = "native.marketdata.realtime.base_prices"
	CapabilityRealtimeMatchedPrices                   core.Capability = "native.marketdata.realtime.matched_prices"
	CapabilityRealtimeTickerMatches                   core.Capability = "native.marketdata.realtime.ticker_matches"
	CapabilityRealtimeIndexes                         core.Capability = "native.marketdata.realtime.indexes"
	CapabilityRealtimeRaw                             core.Capability = "native.marketdata.realtime.raw"
	CapabilityRealtimePutThroughAdvertisements        core.Capability = "native.marketdata.realtime.put_through_advertisements"
	CapabilityRealtimePutThroughMatches               core.Capability = "native.marketdata.realtime.put_through_matches"
)

type RealtimeService interface {
	SubscribeRawEvent(ctx context.Context, request dto.SubscribeRawRequest) (realtime.Subscription[dto.RawEvent], error)

	// /ws/thesis/v1/stream/normal+derivative
	SubscribeBidPrices(context.Context, dto.SubscribeBidPricesRequest) (realtime.Subscription[dto.BidPriceEvent], error)
	SubscribeOfferPrices(context.Context, dto.SubscribeOfferPricesRequest) (realtime.Subscription[dto.OfferPriceEvent], error)
	SubscribeMarketIndexes(context.Context, dto.SubscribeMarketIndexesRequest) (realtime.Subscription[dto.MarketIndexEvent], error)
	SubscribePutThroughAdvertisements(context.Context, dto.SubscribePutThroughAdvertisementsRequest) (realtime.Subscription[dto.PutThroughAdvertisementEvent], error)
	SubscribePutThroughMatches(context.Context, dto.SubscribePutThroughMatchesRequest) (realtime.Subscription[dto.PutThroughMatchedEvent], error)
	//
	SubscribeForeignTrading(context.Context, dto.SubscribeForeignTradingRequest) (realtime.Subscription[dto.ForeignTradingEvent], error)
	SubscribeBasePrices(context.Context, dto.SubscribeBasePricesRequest) (realtime.Subscription[dto.BasePriceEvent], error)
	SubscribeMatchedPrices(context.Context, dto.SubscribeMatchedPricesRequest) (realtime.Subscription[dto.MatchedPriceEvent], error)
	SubscribeTickerMatches(context.Context, dto.SubscribeTickerMatchesRequest) (realtime.Subscription[dto.TickerMatchEvent], error)

	// 5.7.9 /ws/ouranos/v1/stream
	SubscribeStockTradeHistory(context.Context, dto.SubscribeStockTradeHistoryRequest) (realtime.Subscription[dto.StockTradeHistoryEvent], error)
	SubscribeStockSupplyDemandOneMinute(context.Context, dto.SubscribeStockSupplyDemandOneMinuteRequest) (realtime.Subscription[dto.StockSupplyDemandEvent], error)
	SubscribeStockSupplyDemandFifteenMinutes(context.Context, dto.SubscribeStockSupplyDemandFifteenMinutesRequest) (realtime.Subscription[dto.StockSupplyDemandEvent], error)
}

type RealtimeDependencies struct {
	BaseURL           string
	AccessToken       func() string
	Headers           func(bool, bool) map[string]string
	RequireCapability func(core.Capability) error
	WebSocketFactory  func(context.Context, string, map[string]string) (transport.WebSocketTransport, error)
	PingInterval      time.Duration
}

type realtimeService struct {
	dependencies RealtimeDependencies
}

func NewRealtimeService(dependencies RealtimeDependencies) RealtimeService {
	if dependencies.PingInterval == 0 {
		dependencies.PingInterval = 2 * time.Second
	}
	return &realtimeService{dependencies: dependencies}
}

func (s *realtimeService) SubscribeRawEvent(ctx context.Context, request dto.SubscribeRawRequest) (realtime.Subscription[dto.RawEvent], error) {
	return subscribeRaw(ctx, s.dependencies, CapabilityRealtimeRaw, "bp+bi+tm+mp+op+fe", func(payload []byte) (dto.RawEvent, error) {
		var event dto.RawEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeStockTradeHistory(ctx context.Context, request dto.SubscribeStockTradeHistoryRequest) (realtime.Subscription[dto.StockTradeHistoryEvent], error) {
	return subscribeOuranos(ctx, s.dependencies, CapabilityRealtimeStockTradeHistory, "C001", request.Tickers, func(payload []byte) (dto.StockTradeHistoryEvent, error) {
		var event dto.StockTradeHistoryEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeStockSupplyDemandOneMinute(ctx context.Context, request dto.SubscribeStockSupplyDemandOneMinuteRequest) (realtime.Subscription[dto.StockSupplyDemandEvent], error) {
	return subscribeOuranos(ctx, s.dependencies, CapabilityRealtimeStockSupplyDemandOneMinute, "C002S60", request.Tickers, decodeStockSupplyDemand)
}

func (s *realtimeService) SubscribeStockSupplyDemandFifteenMinutes(ctx context.Context, request dto.SubscribeStockSupplyDemandFifteenMinutesRequest) (realtime.Subscription[dto.StockSupplyDemandEvent], error) {
	return subscribeOuranos(ctx, s.dependencies, CapabilityRealtimeStockSupplyDemandFifteenMinutes, "C002S900", request.Tickers, decodeStockSupplyDemand)
}

func (s *realtimeService) SubscribeBidPrices(ctx context.Context, request dto.SubscribeBidPricesRequest) (realtime.Subscription[dto.BidPriceEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeBidPrices, "bi", request.Symbols, func(payload []byte) (dto.BidPriceEvent, error) {
		var event dto.BidPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeOfferPrices(ctx context.Context, request dto.SubscribeOfferPricesRequest) (realtime.Subscription[dto.OfferPriceEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeOfferPrices, "op", request.Symbols, func(payload []byte) (dto.OfferPriceEvent, error) {
		var event dto.OfferPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeForeignTrading(ctx context.Context, request dto.SubscribeForeignTradingRequest) (realtime.Subscription[dto.ForeignTradingEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeForeignTrading, "fe", request.Symbols, func(payload []byte) (dto.ForeignTradingEvent, error) {
		var event dto.ForeignTradingEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeBasePrices(ctx context.Context, request dto.SubscribeBasePricesRequest) (realtime.Subscription[dto.BasePriceEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeBasePrices, "bp", request.Symbols, func(payload []byte) (dto.BasePriceEvent, error) {
		var event dto.BasePriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeMatchedPrices(ctx context.Context, request dto.SubscribeMatchedPricesRequest) (realtime.Subscription[dto.MatchedPriceEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeMatchedPrices, "mp", request.Symbols, func(payload []byte) (dto.MatchedPriceEvent, error) {
		var event dto.MatchedPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeTickerMatches(ctx context.Context, request dto.SubscribeTickerMatchesRequest) (realtime.Subscription[dto.TickerMatchEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimeTickerMatches, "tm", request.Symbols, func(payload []byte) (dto.TickerMatchEvent, error) {
		var event dto.TickerMatchEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeMarketIndexes(ctx context.Context, request dto.SubscribeMarketIndexesRequest) (realtime.Subscription[dto.MarketIndexEvent], error) {
	values := strings.Join(request.Indexes, ",")
	return subscribeMarket(ctx, s.dependencies, CapabilityRealtimeIndexes, "/ws/thesis/v1/stream/derivative", "d|s|si|rt|"+values, "", []string{"s|8"}, func(payload []byte) (dto.MarketIndexEvent, error) {
		var event dto.MarketIndexEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}
func (s *realtimeService) SubscribePutThroughAdvertisements(ctx context.Context, request dto.SubscribePutThroughAdvertisementsRequest) (realtime.Subscription[dto.PutThroughAdvertisementEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimePutThroughAdvertisements, "pta", request.Symbols, func(payload []byte) (dto.PutThroughAdvertisementEvent, error) {
		var event dto.PutThroughAdvertisementEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}
func (s *realtimeService) SubscribePutThroughMatches(ctx context.Context, request dto.SubscribePutThroughMatchesRequest) (realtime.Subscription[dto.PutThroughMatchedEvent], error) {
	return subscribeThesis(ctx, s.dependencies, CapabilityRealtimePutThroughMatches, "ptm", request.Symbols, func(payload []byte) (dto.PutThroughMatchedEvent, error) {
		var event dto.PutThroughMatchedEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func subscribeOuranos[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, code string, tickers []string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	values := strings.Join(tickers, ",")
	return subscribeMarket(ctx, dependencies, capability, "/ws/ouranos/v1/stream", "d|st|"+code+"|"+values, "d|ut|"+code+"|"+values, []string{code}, decode)
}

func subscribeThesis[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, code string, symbols []string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	receivedPrefixes := map[string][]string{
		"bi":  {"s|1", "s|28", "s|23"},
		"op":  {"s|2", "s|29", "s|24"},
		"fe":  {"s|3", "s|30", "s|25"},
		"bp":  {"s|4", "s|31", "s|26"},
		"mp":  {"s|5", "s|32", "s|27"},
		"tm":  {"s|6", "s|18", "s|21"},
		"rt":  {"s|8"},
		"pta": {"s|16"},
		"ptm": {"s|17"},
	}
	prefixes := receivedPrefixes[code]

	channel := "d|s|tk|" + code + "|" + strings.Join(symbols, ",")
	if len(symbols) == 0 {
		channel = "d|s|ro|" + code + "|1,2,3,4,5,7,8,9"
	}
	switch code {
	case "rt":
		channel = "d|s|si|" + code + "|" + strings.Join(symbols, ",")
	case "pta", "ptm":
		channel = "d|s|pt|" + code + "|" + strings.Join(symbols, ",")
	}
	return subscribeMarket(ctx, dependencies, capability, "/ws/thesis/v1/stream/derivative", channel, "", prefixes, decode)
}

func subscribeRaw[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, code string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	prefix := map[string]string{"bi": "s|23", "op": "s|24", "fe": "s|3", "bp": "s|4", "mp": "s|5", "tm": "s|21", "ptm": "s|17"}[code]
	return subscribeMarket(ctx, dependencies, capability, "/ws/thesis/v1/stream/normal", "d|s|pt|ptm+pta|1,2,3,4,5", "", []string{prefix}, decode)
}

func subscribeMarket[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, path, subscribeFrame, unsubscribeFrame string, prefixes []string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	if err := dependencies.RequireCapability(capability); err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	socket, err := dependencies.WebSocketFactory(childCtx, marketWebSocketURL(dependencies.BaseURL, path), marketHeaders(dependencies))
	if err != nil {
		cancel()
		return nil, err
	}
	var sendMu sync.Mutex
	send := func(message string) error {
		// fmt.Printf(">>> %s\n", message)
		sendMu.Lock()
		defer sendMu.Unlock()
		return socket.Send(childCtx, transport.WebSocketMessage(message))
	}
	subscription := realtime.NewQueueSubscription[T](128, func() error {
		if unsubscribeFrame != "" {
			_ = send(unsubscribeFrame)
		}
		cancel()
		return socket.Close()
	})
	subscription.PublishStatus(realtime.StatusConnected)
	subscription.PublishStatus(realtime.StatusAuthenticating)
	if err := send("d|a|||" + base64.StdEncoding.EncodeToString([]byte(dependencies.AccessToken()))); err != nil {
		_ = subscription.Close()
		return nil, err
	}
	go runMarketSubscription(childCtx, cancel, dependencies.PingInterval, capability, subscribeFrame, prefixes, socket, send, decode, subscription)
	return subscription, nil
}

func runMarketSubscription[T any](ctx context.Context, cancel context.CancelFunc, pingInterval time.Duration, capability core.Capability, subscribeFrame string, prefixes []string, socket transport.WebSocketTransport, send func(string) error, decode func([]byte) (T, error), subscription *realtime.QueueSubscription[T]) {
	defer subscription.Close()
	authenticated := false
	subscribed := false
	pingStarted := false
	subscribe := func() bool {
		if !authenticated || subscribed {
			return true
		}
		if err := send(subscribeFrame); err != nil {
			subscription.PublishError(err)
			return false
		}
		subscribed = true
		subscription.PublishStatus(realtime.StatusSubscribed)
		return true
	}
	for {
		payload, err := socket.Receive(ctx)
		if err != nil {
			if ctx.Err() == nil {
				subscription.PublishError(err)
			}
			return
		}
		frame := parseMarketFrame(payload)
		switch frame.kind {
		case "d|0":

			if !marketAuthSucceeded(frame.payload) {
				subscription.PublishError(sdkerrors.Auth("tcbs", string(capability), "TCBS websocket authentication failed"))
				return
			}
			authenticated = true
			if !subscribe() {
				return
			}
		case "d|33":
			if !pingStarted {
				pingStarted = true
				startMarketPingLoop(ctx, cancel, pingInterval, send, subscription)
			}
		case "d|34":
			if err := send("d|p|||"); err != nil {
				subscription.PublishError(err)
				return
			}
		case "d|pi": // /ws/ouranos/v1/stream
			if err := send("d|po"); err != nil {
				subscription.PublishError(err)
				return
			}
		default:
			if !marketFrameKindMatches(frame.kind, prefixes) {
				err := fmt.Errorf("unexpected event kind: %s", frame.kind)
				subscription.PublishError(
					sdkerrors.Decode(
						"tcbs",
						string(capability),
						"unexpected TCBS websocket event kind",
						payload,
						err,
					),
				)
				continue
			}
			event, err := decode(frame.payload)
			if err != nil {
				subscription.PublishError(sdkerrors.Decode("tcbs", string(capability), "decode TCBS websocket event", payload, err))
				continue
			}
			subscription.PublishEvent(event)
		}
	}
}

type marketFrame struct {
	kind    string
	payload []byte
}

func marketFrameKindMatches(kind string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if kind == prefix {
			return true
		}
	}
	return false
}

func parseMarketFrame(payload []byte) marketFrame {
	parts := strings.SplitN(string(payload), "|", 3)
	if len(parts) < 3 {
		return marketFrame{kind: string(payload)}
	}
	if parts[0] == "d" || parts[0] == "s" {
		return marketFrame{kind: parts[0] + "|" + parts[1], payload: []byte(parts[2])}
	}
	return marketFrame{kind: parts[0], payload: []byte(parts[2])}
}

func marketAuthSucceeded(payload []byte) bool {
	message := struct {
		Success bool `json:"success"`
	}{}
	return json.Unmarshal(payload, &message) == nil && message.Success
}

func decodeStockSupplyDemand(payload []byte) (dto.StockSupplyDemandEvent, error) {
	var event dto.StockSupplyDemandEvent
	err := json.Unmarshal(payload, &event)
	return event, err
}

func marketHeaders(dependencies RealtimeDependencies) map[string]string {
	if dependencies.Headers == nil {
		return map[string]string{}
	}
	return dependencies.Headers(true, false)
}

func marketWebSocketURL(baseURL, path string) string {
	endpoint := strings.TrimRight(baseURL, "/") + path
	endpoint = strings.Replace(endpoint, "https://", "wss://", 1)
	return strings.Replace(endpoint, "http://", "ws://", 1)
}

func startMarketPingLoop[T any](ctx context.Context, cancel context.CancelFunc, interval time.Duration, send func(string) error, subscription *realtime.QueueSubscription[T]) {
	if interval <= 0 {
		return
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := send("d|p|||"); err != nil {
					if ctx.Err() == nil {
						subscription.PublishError(err)
						cancel()
					}
					return
				}
			}
		}
	}()
}
