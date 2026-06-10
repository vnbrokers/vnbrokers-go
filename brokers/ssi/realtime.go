package ssi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/vnbrokers/vnbrokers-go/domain"
	sdkerrors "github.com/vnbrokers/vnbrokers-go/errors"
	"github.com/vnbrokers/vnbrokers-go/realtime"
)

const (
	ssiTradingHub    = "BroadcastHubV2"
	ssiMarketDataHub = "FcMarketDataV2Hub"
)

type ssiRealtimeEnvelope struct {
	DataType string `json:"DataType"`
	Content  string `json:"Content"`
}

type ssiBroadcastEnvelope struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type ssiRealtimeMessage struct {
	DataType string
	Data     map[string]any
	Raw      []byte
}

func startSSISignalRSubscription[T any](
	ctx context.Context,
	broker *Broker,
	accessToken string,
	baseURL string,
	hub string,
	configure func(SignalRClient, *realtime.QueueSubscription[T]),
	start func(SignalRClient) error,
) (realtime.Subscription[T], error) {
	if accessToken == "" {
		return nil, sdkerrors.Auth("ssi", "realtime.subscribe", "SSI realtime requires an access token")
	}
	childCtx, cancel := context.WithCancel(ctx)
	client := broker.config.SignalRFactory(baseURL, []string{hub})
	client.SetHeader("Authorization", "Bearer "+accessToken)

	subscription := realtime.NewQueueSubscription[T](128, func() error {
		cancel()
		return client.Close()
	})
	subscription.PublishStatus(realtime.StatusConnecting)
	client.OnError(func(err error) {
		if childCtx.Err() == nil {
			subscription.PublishError(err)
		}
	})
	configure(client, subscription)

	if err := client.Connect(childCtx); err != nil {
		subscription.PublishStatus(realtime.StatusFailed)
		subscription.PublishError(err)
		_ = subscription.Close()
		return nil, err
	}
	subscription.PublishStatus(realtime.StatusConnected)
	if start != nil {
		if err := start(client); err != nil {
			subscription.PublishStatus(realtime.StatusFailed)
			subscription.PublishError(err)
			_ = subscription.Close()
			return nil, err
		}
	}
	subscription.PublishStatus(realtime.StatusSubscribed)

	go func() {
		<-childCtx.Done()
		_ = subscription.Close()
	}()
	return subscription, nil
}

func registerSSIHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	hub string,
	method string,
	mapEvent func(ssiRealtimeMessage) T,
) {
	client.On(hub, method, func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeSSIRealtimeArg(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w", hub, method, err))
				continue
			}
			subscription.PublishEvent(mapEvent(message))
		}
	})
}

func registerSSIJSONHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	hub string,
	method string,
) {
	client.On(hub, method, func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeSSIRealtimeArg(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w", hub, method, err))
				continue
			}
			var event T
			if err := json.Unmarshal(message.Raw, &event); err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.%s: %w %s", hub, method, err, message.Raw))
				continue
			}
			subscription.PublishEvent(event)
		}
	})
}

func registerSSIBroadcastHandler[T any](
	client SignalRClient,
	subscription *realtime.QueueSubscription[T],
	eventTypes []string,
	mapEvent func(ssiRealtimeMessage) T,
) {
	client.On(ssiTradingHub, "Broadcast", func(args []json.RawMessage) {
		for _, arg := range args {
			message, err := decodeSSIRealtimeArg(arg)
			if err != nil {
				subscription.PublishError(fmt.Errorf("ssi realtime decode %s.Broadcast: %w", ssiTradingHub, err))
				continue
			}
			if !containsFold(eventTypes, message.DataType) {
				continue
			}
			subscription.PublishEvent(mapEvent(message))
		}
	})
}

func containsFold(values []string, target string) bool {
	for _, value := range values {
		if strings.EqualFold(value, target) {
			return true
		}
	}
	return false
}

func decodeSSIRealtimeArg(arg json.RawMessage) (ssiRealtimeMessage, error) {
	payload := []byte(arg)
	var text string
	if json.Unmarshal(arg, &text) == nil {
		payload = []byte(text)
	}

	var broadcast ssiBroadcastEnvelope
	if json.Unmarshal(payload, &broadcast) == nil && broadcast.Type != "" && len(broadcast.Data) > 0 && string(broadcast.Data) != "null" {
		data := map[string]any{}
		if err := json.Unmarshal(broadcast.Data, &data); err != nil {
			return ssiRealtimeMessage{}, err
		}
		return ssiRealtimeMessage{DataType: broadcast.Type, Data: data, Raw: broadcast.Data}, nil
	}

	var envelope ssiRealtimeEnvelope
	if json.Unmarshal(payload, &envelope) == nil && envelope.Content != "" {
		content := []byte(envelope.Content)
		data := map[string]any{}
		if err := json.Unmarshal(content, &data); err != nil {
			return ssiRealtimeMessage{}, err
		}
		return ssiRealtimeMessage{DataType: envelope.DataType, Data: data, Raw: content}, nil
	}

	data := map[string]any{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return ssiRealtimeMessage{}, err
	}
	return ssiRealtimeMessage{Data: data, Raw: payload}, nil
}

func ssiValue(data map[string]any, keys ...string) any {
	for _, key := range keys {
		if value, ok := data[key]; ok {
			return value
		}
	}
	for existing, value := range data {
		for _, key := range keys {
			if strings.EqualFold(existing, key) {
				return value
			}
		}
	}
	return nil
}

func ssiString(data map[string]any, keys ...string) string {
	return stringify(ssiValue(data, keys...))
}

func ssiInt64(data map[string]any, keys ...string) int64 {
	value := ssiValue(data, keys...)
	switch typed := value.(type) {
	case float64:
		return int64(typed)
	case int64:
		return typed
	case int:
		return int64(typed)
	case json.Number:
		out, _ := typed.Int64()
		return out
	default:
		out, _ := strconv.ParseInt(fmt.Sprint(value), 10, 64)
		return out
	}
}

func ssiBool(data map[string]any, keys ...string) bool {
	value := ssiValue(data, keys...)
	if typed, ok := value.(bool); ok {
		return typed
	}
	out, _ := strconv.ParseBool(fmt.Sprint(value))
	return out
}

func ssiRawPayload(message ssiRealtimeMessage) domain.RawPayload {
	return domain.RawPayload{Source: "ssi", Data: message.Data, Bytes: message.Raw}
}
