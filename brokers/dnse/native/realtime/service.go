package realtime

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	sdkrealtime "github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const defaultSubscriptionBuffer = 128

type Dependencies struct {
	APIKey, APISecret   string
	StreamURL, Encoding string
	PongInterval        time.Duration
	WebSocketFactory    transport.WebSocketFactory
}

func (d Dependencies) URL() (string, error) {
	encoding, err := d.encoding()
	if err != nil {
		return "", err
	}
	parsed, err := url.Parse(d.StreamURL)
	if err != nil {
		return "", err
	}
	query := parsed.Query()
	query.Set("encoding", encoding)
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}
func (d Dependencies) encoding() (string, error) {
	encoding := strings.ToLower(strings.TrimSpace(d.Encoding))
	if encoding == "json" || encoding == "msgpack" {
		return encoding, nil
	}
	return "", fmt.Errorf("unsupported DNSE stream encoding: %s", d.Encoding)
}

func BuildAuthMessage(apiKey, apiSecret string, timestamp int64, nonce string) map[string]any {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	if nonce == "" {
		nonce = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	message := fmt.Sprintf("%s:%d:%s", apiKey, timestamp, nonce)
	mac := hmac.New(sha256.New, []byte(apiSecret))
	_, _ = mac.Write([]byte(message))
	return map[string]any{"action": "auth", "api_key": apiKey, "signature": hex.EncodeToString(mac.Sum(nil)), "timestamp": timestamp, "nonce": nonce}
}
func BuildPongMessage() map[string]any { return map[string]any{"action": "pong"} }

func Subscribe[T any](ctx context.Context, d Dependencies, subscribeMessage map[string]any, shouldPublish func(map[string]any) bool, decode func(map[string]any) (T, error)) (sdkrealtime.Subscription[T], error) {
	if d.APIKey == "" || d.APISecret == "" {
		return nil, sdkerrors.Auth("dnse", "realtime.subscribe", "DNSE realtime stream requires API key and API secret")
	}
	encoding, err := d.encoding()
	if err != nil {
		return nil, err
	}
	streamURL, err := d.URL()
	if err != nil {
		return nil, err
	}
	factory := d.WebSocketFactory
	if factory == nil {
		factory = transport.ConnectWebSocket
	}
	childCtx, cancel := context.WithCancel(ctx)
	var socket transport.WebSocketTransport
	var sendMu sync.Mutex
	send := func(ctx context.Context, message map[string]any) error {
		payload, err := encode(message, encoding)
		if err != nil {
			return err
		}
		sendMu.Lock()
		defer sendMu.Unlock()
		return socket.Send(ctx, payload)
	}
	subscription := sdkrealtime.NewQueueSubscription[T](defaultSubscriptionBuffer, func() error {
		cancel()
		if socket != nil {
			return socket.Close()
		}
		return nil
	})
	go func() {
		defer subscription.Close()
		subscription.PublishStatus(sdkrealtime.StatusConnecting)
		socket, err = factory(childCtx, streamURL)
		if err != nil {
			subscription.PublishStatus(sdkrealtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		subscription.PublishStatus(sdkrealtime.StatusConnected)
		subscription.PublishStatus(sdkrealtime.StatusAuthenticating)
		if err = send(childCtx, BuildAuthMessage(d.APIKey, d.APISecret, 0, "")); err != nil {
			subscription.PublishStatus(sdkrealtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		if err = send(childCtx, subscribeMessage); err != nil {
			subscription.PublishStatus(sdkrealtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		subscription.PublishStatus(sdkrealtime.StatusSubscribed)
		startPongLoop(childCtx, d.PongInterval, send, subscription, cancel)
		for {
			raw, receiveErr := socket.Receive(childCtx)
			if receiveErr != nil {
				if childCtx.Err() == nil {
					subscription.PublishStatus(sdkrealtime.StatusFailed)
					subscription.PublishError(receiveErr)
				}
				return
			}
			message, decodeErr := decodeMessage(raw, encoding)
			if decodeErr != nil {
				subscription.PublishError(sdkerrors.Decode("dnse", "realtime.decode", "failed to decode DNSE stream message", raw, decodeErr))
				continue
			}
			if strings.EqualFold(fmt.Sprint(message["action"]), "ping") {
				if sendErr := send(childCtx, BuildPongMessage()); sendErr != nil {
					if childCtx.Err() == nil {
						subscription.PublishStatus(sdkrealtime.StatusFailed)
						subscription.PublishError(sendErr)
						cancel()
					}
					return
				}
				continue
			}
			if !shouldPublish(message) {
				continue
			}
			event, decodeErr := decode(message)
			if decodeErr != nil {
				subscription.PublishError(decodeErr)
				continue
			}
			subscription.PublishEvent(event)
		}
	}()
	return subscription, nil
}

func startPongLoop[T any](ctx context.Context, interval time.Duration, send func(context.Context, map[string]any) error, subscription *sdkrealtime.QueueSubscription[T], cancel context.CancelFunc) {
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
				if err := send(ctx, BuildPongMessage()); err != nil {
					if ctx.Err() == nil {
						subscription.PublishStatus(sdkrealtime.StatusFailed)
						subscription.PublishError(err)
						cancel()
					}
					return
				}
			}
		}
	}()
}
func encode(message map[string]any, encoding string) ([]byte, error) {
	if encoding == "msgpack" {
		return msgpack.Marshal(message)
	}
	return json.Marshal(message)
}
func decodeMessage(message []byte, encoding string) (map[string]any, error) {
	out := map[string]any{}
	var err error
	if encoding == "msgpack" {
		err = msgpack.Unmarshal(message, &out)
	} else {
		err = json.Unmarshal(message, &out)
	}
	return out, err
}
