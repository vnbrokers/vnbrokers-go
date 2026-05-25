package dnse

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
	"github.com/vnbrokers/vnbrokers-go/transport"
)

const defaultSubscriptionBuffer = 128

type streamPublisher[T any] func(*realtime.QueueSubscription[T], map[string]any)

func BuildStreamAuthMessage(apiKey string, apiSecret string, timestamp int64, nonce string) map[string]any {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	if nonce == "" {
		nonce = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	message := fmt.Sprintf("%s:%d:%s", apiKey, timestamp, nonce)
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(message))
	return map[string]any{
		"action":    "auth",
		"api_key":   apiKey,
		"signature": hex.EncodeToString(mac.Sum(nil)),
		"timestamp": timestamp,
		"nonce":     nonce,
	}
}

func startRealtimeSubscription[T any](
	ctx context.Context,
	broker *Broker,
	subscribeMessage map[string]any,
	shouldPublish func(map[string]any) bool,
	publisher streamPublisher[T],
) (realtime.Subscription[T], error) {
	if broker.config.APIKey == "" || broker.config.APISecret == "" {
		return nil, errors.Auth("dnse", "realtime.subscribe", "DNSE realtime stream requires API key and API secret")
	}
	encoding, err := streamEncoding(broker.config)
	if err != nil {
		return nil, err
	}
	streamURL, err := streamURL(broker.config)
	if err != nil {
		return nil, err
	}
	childCtx, cancel := context.WithCancel(ctx)
	var socket transport.WebSocketTransport
	subscription := realtime.NewQueueSubscription[T](defaultSubscriptionBuffer, func() error {
		cancel()
		if socket != nil {
			return socket.Close()
		}
		return nil
	})
	go func() {
		defer subscription.Close()
		subscription.PublishStatus(realtime.StatusConnecting)
		var connectErr error
		socket, connectErr = broker.config.WebSocketFactory(childCtx, streamURL)
		if connectErr != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(connectErr)
			return
		}
		subscription.PublishStatus(realtime.StatusConnected)
		subscription.PublishStatus(realtime.StatusAuthenticating)
		authPayload, err := encodeStreamMessage(
			BuildStreamAuthMessage(broker.config.APIKey, broker.config.APISecret, 0, ""),
			encoding,
		)
		if err != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		if err := socket.Send(childCtx, authPayload); err != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		subscribePayload, err := encodeStreamMessage(subscribeMessage, encoding)
		if err != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		if err := socket.Send(childCtx, subscribePayload); err != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(err)
			return
		}
		subscription.PublishStatus(realtime.StatusSubscribed)
		for {
			message, err := socket.Receive(childCtx)
			if err != nil {
				if childCtx.Err() == nil {
					subscription.PublishStatus(realtime.StatusFailed)
					subscription.PublishError(err)
				}
				return
			}
			decoded, err := decodeStreamMessage(message, encoding)
			if err != nil {
				subscription.PublishError(errors.Decode("dnse", "realtime.decode", "failed to decode DNSE stream message", message, err))
				continue
			}
			if !shouldPublish(decoded) {
				continue
			}
			publisher(subscription, decoded)
		}
	}()
	return subscription, nil
}

func streamEncoding(config Config) (string, error) {
	encoding := strings.ToLower(strings.TrimSpace(config.StreamEncoding))
	if encoding == "json" || encoding == "msgpack" {
		return encoding, nil
	}
	return "", fmt.Errorf("unsupported DNSE stream encoding: %s", config.StreamEncoding)
}

func streamURL(config Config) (string, error) {
	encoding, err := streamEncoding(config)
	if err != nil {
		return "", err
	}
	parsed, err := url.Parse(config.StreamURL)
	if err != nil {
		return "", err
	}
	query := parsed.Query()
	query.Set("encoding", encoding)
	parsed.RawQuery = query.Encode()
	return parsed.String(), nil
}

func encodeStreamMessage(message map[string]any, encoding string) ([]byte, error) {
	if encoding == "msgpack" {
		return msgpack.Marshal(message)
	}
	return json.Marshal(message)
}

func decodeStreamMessage(message []byte, encoding string) (map[string]any, error) {
	out := map[string]any{}
	var err error
	if encoding == "msgpack" {
		err = msgpack.Unmarshal(message, &out)
	} else {
		err = json.Unmarshal(message, &out)
	}
	return out, err
}
