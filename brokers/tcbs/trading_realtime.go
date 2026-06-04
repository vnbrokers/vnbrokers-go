package tcbs

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/core"
	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

const stockMatchEndpoint = "/ws/aither"

type TradingRealtimeService struct {
	broker *Broker
}

func (s *TradingRealtimeService) SubscribeStockMatches(
	ctx context.Context,
) (realtime.Subscription[domain.OrderEvent], error) {
	if err := s.broker.RequireCapability(core.CapabilityTradingRealtimeOrders); err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	subscription := realtime.NewQueueSubscription[domain.OrderEvent](128, func() error {
		cancel()
		return nil
	})
	subscription.PublishStatus(realtime.StatusConnecting)
	socket, err := s.broker.config.WebSocketFactory(childCtx, s.broker.wsURL(stockMatchEndpoint), s.broker.headers(true, false))
	if err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription = realtime.NewQueueSubscription[domain.OrderEvent](128, func() error {
		cancel()
		return socket.Close()
	})
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
			message := map[string]any{}
			if err := json.Unmarshal(payload, &message); err != nil {
				subscription.PublishError(sdkerrors.Decode("tcbs", "trading.realtime.stock_matches", "decode TCBS websocket message", payload, err))
				continue
			}
			if !isStockMatchMessage(message) {
				continue
			}
			subscription.PublishEvent(MapStockMatchEvent(message))
		}
	}()
	return subscription, nil
}

func (b *Broker) wsURL(path string) string {
	out := b.url(path)
	out = strings.Replace(out, "https://", "wss://", 1)
	out = strings.Replace(out, "http://", "ws://", 1)
	return out
}

func isStockMatchMessage(message map[string]any) bool {
	return firstString(message, "orderID", "orderId", "orderNo") != "" || stringify(message["symbol"]) != ""
}
