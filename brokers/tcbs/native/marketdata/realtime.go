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
	CapabilityRealtimeStockPrices                     core.Capability = "native.marketdata.realtime.stock_prices"
	CapabilityRealtimeStockTradeHistory               core.Capability = "native.marketdata.realtime.stock_trade_history"
	CapabilityRealtimeStockSupplyDemandOneMinute      core.Capability = "native.marketdata.realtime.stock_supply_demand_one_minute"
	CapabilityRealtimeStockSupplyDemandFifteenMinutes core.Capability = "native.marketdata.realtime.stock_supply_demand_fifteen_minutes"
	CapabilityRealtimeDerivativeBidPrices             core.Capability = "native.marketdata.realtime.derivative_bid_prices"
	CapabilityRealtimeDerivativeOfferPrices           core.Capability = "native.marketdata.realtime.derivative_offer_prices"
	CapabilityRealtimeDerivativeForeignTrading        core.Capability = "native.marketdata.realtime.derivative_foreign_trading"
	CapabilityRealtimeDerivativeBasePrices            core.Capability = "native.marketdata.realtime.derivative_base_prices"
	CapabilityRealtimeDerivativeMatchedPrices         core.Capability = "native.marketdata.realtime.derivative_matched_prices"
	CapabilityRealtimeDerivativeTickerMatches         core.Capability = "native.marketdata.realtime.derivative_ticker_matches"
	CapabilityRealtimeDerivativeIndexes               core.Capability = "native.marketdata.realtime.derivative_indexes"
)

type RealtimeService interface {
	SubscribeStockPrices(context.Context, dto.SubscribeStockPricesRequest) (realtime.Subscription[dto.RawMessage], error)

	SubscribeDerivativeBidPrices(context.Context, dto.SubscribeDerivativeBidPricesRequest) (realtime.Subscription[dto.BidPriceEvent], error)
	SubscribeDerivativeOfferPrices(context.Context, dto.SubscribeDerivativeOfferPricesRequest) (realtime.Subscription[dto.OfferPriceEvent], error)
	SubscribeDerivativeForeignTrading(context.Context, dto.SubscribeDerivativeForeignTradingRequest) (realtime.Subscription[dto.DerivativeForeignTradingEvent], error)
	SubscribeDerivativeBasePrices(context.Context, dto.SubscribeDerivativeBasePricesRequest) (realtime.Subscription[dto.DerivativeBasePriceEvent], error)
	SubscribeDerivativeMatchedPrices(context.Context, dto.SubscribeDerivativeMatchedPricesRequest) (realtime.Subscription[dto.DerivativeMatchedPriceEvent], error)
	SubscribeDerivativeTickerMatches(context.Context, dto.SubscribeDerivativeTickerMatchesRequest) (realtime.Subscription[dto.DerivativeTickerMatchEvent], error)
	SubscribeDerivativeIndexes(context.Context, dto.SubscribeDerivativeIndexesRequest) (realtime.Subscription[dto.DerivativeIndexEvent], error)

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

func (s *realtimeService) SubscribeStockPrices(ctx context.Context, _ dto.SubscribeStockPricesRequest) (realtime.Subscription[dto.RawMessage], error) {
	if err := s.dependencies.RequireCapability(CapabilityRealtimeStockPrices); err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	socket, err := s.dependencies.WebSocketFactory(childCtx, marketWebSocketURL(s.dependencies.BaseURL, "/ws/thesis/v1/stream/normal"), marketHeaders(s.dependencies))
	if err != nil {
		cancel()
		return nil, err
	}
	subscription := realtime.NewQueueSubscription[dto.RawMessage](128, func() error { cancel(); return socket.Close() })
	subscription.PublishStatus(realtime.StatusConnected)
	subscription.PublishStatus(realtime.StatusSubscribed)
	go func() {
		defer subscription.Close()
		for {
			payload, err := socket.Receive(childCtx)
			if err != nil {
				if childCtx.Err() == nil {
					subscription.PublishError(err)
				}
				return
			}
			subscription.PublishEvent(dto.RawMessage(append([]byte(nil), payload...)))
		}
	}()
	return subscription, nil
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

func (s *realtimeService) SubscribeDerivativeBidPrices(ctx context.Context, request dto.SubscribeDerivativeBidPricesRequest) (realtime.Subscription[dto.BidPriceEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeBidPrices, "bi", request.Symbols, func(payload []byte) (dto.BidPriceEvent, error) {
		var event dto.BidPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeOfferPrices(ctx context.Context, request dto.SubscribeDerivativeOfferPricesRequest) (realtime.Subscription[dto.OfferPriceEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeOfferPrices, "op", request.Symbols, func(payload []byte) (dto.OfferPriceEvent, error) {
		var event dto.OfferPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeForeignTrading(ctx context.Context, request dto.SubscribeDerivativeForeignTradingRequest) (realtime.Subscription[dto.DerivativeForeignTradingEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeForeignTrading, "fe", request.Symbols, func(payload []byte) (dto.DerivativeForeignTradingEvent, error) {
		var event dto.DerivativeForeignTradingEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeBasePrices(ctx context.Context, request dto.SubscribeDerivativeBasePricesRequest) (realtime.Subscription[dto.DerivativeBasePriceEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeBasePrices, "bp", request.Symbols, func(payload []byte) (dto.DerivativeBasePriceEvent, error) {
		var event dto.DerivativeBasePriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeMatchedPrices(ctx context.Context, request dto.SubscribeDerivativeMatchedPricesRequest) (realtime.Subscription[dto.DerivativeMatchedPriceEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeMatchedPrices, "mp", request.Symbols, func(payload []byte) (dto.DerivativeMatchedPriceEvent, error) {
		var event dto.DerivativeMatchedPriceEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeTickerMatches(ctx context.Context, request dto.SubscribeDerivativeTickerMatchesRequest) (realtime.Subscription[dto.DerivativeTickerMatchEvent], error) {
	return subscribeDerivative(ctx, s.dependencies, CapabilityRealtimeDerivativeTickerMatches, "tm", request.Symbols, func(payload []byte) (dto.DerivativeTickerMatchEvent, error) {
		var event dto.DerivativeTickerMatchEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func (s *realtimeService) SubscribeDerivativeIndexes(ctx context.Context, request dto.SubscribeDerivativeIndexesRequest) (realtime.Subscription[dto.DerivativeIndexEvent], error) {
	values := strings.Join(request.Indexes, ",")
	return subscribeMarket(ctx, s.dependencies, CapabilityRealtimeDerivativeIndexes, "/ws/thesis/v1/stream/derivative", "d|s|si|rt|"+values, "", "s|8", func(payload []byte) (dto.DerivativeIndexEvent, error) {
		var event dto.DerivativeIndexEvent
		err := json.Unmarshal(payload, &event)
		return event, err
	})
}

func subscribeOuranos[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, code string, tickers []string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	values := strings.Join(tickers, ",")
	return subscribeMarket(ctx, dependencies, capability, "/ws/ouranos/v1/stream", "d|st|"+code+"|"+values, "d|ut|"+code+"|"+values, code, decode)
}

func subscribeDerivative[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, code string, symbols []string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
	prefix := map[string]string{"bi": "s|23", "op": "s|24", "fe": "s|3", "bp": "s|4", "mp": "s|5", "tm": "s|21"}[code]
	return subscribeMarket(ctx, dependencies, capability, "/ws/thesis/v1/stream/derivative", "d|s|tk|"+code+"|"+strings.Join(symbols, ","), "", prefix, decode)
}

func subscribeMarket[T any](ctx context.Context, dependencies RealtimeDependencies, capability core.Capability, path, subscribeFrame, unsubscribeFrame, eventKind string, decode func([]byte) (T, error)) (realtime.Subscription[T], error) {
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
		fmt.Printf(">>> %s\n", message)
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
	go runMarketSubscription(childCtx, cancel, dependencies.PingInterval, capability, subscribeFrame, eventKind, socket, send, decode, subscription)
	return subscription, nil
}

func runMarketSubscription[T any](ctx context.Context, cancel context.CancelFunc, pingInterval time.Duration, capability core.Capability, subscribeFrame, eventKind string, socket transport.WebSocketTransport, send func(string) error, decode func([]byte) (T, error), subscription *realtime.QueueSubscription[T]) {
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
			if frame.kind != eventKind {
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

func parseMarketFrame(payload []byte) marketFrame {
	parts := strings.SplitN(string(payload), "|", 3)
	if len(parts) < 3 {
		return marketFrame{}
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
