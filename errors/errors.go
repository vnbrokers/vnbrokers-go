package errors

import "fmt"

type Category string

const (
	CategoryAuth                  Category = "auth"
	CategoryBrokerRejected        Category = "broker_rejected"
	CategoryCapabilityUnsupported Category = "capability_unsupported"
	CategoryDecode                Category = "decode"
	CategoryNetwork               Category = "network"
	CategoryRateLimit             Category = "rate_limit"
	CategorySubscriptionClosed    Category = "subscription_closed"
)

type BrokerError struct {
	Category  Category
	Code      string
	Message   string
	Broker    string
	Operation string
	Retryable bool
	Raw       any
	Cause     error
}

func (e *BrokerError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Broker == "" {
		return e.Message
	}
	return fmt.Sprintf("%s: %s", e.Broker, e.Message)
}

func (e *BrokerError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

func UnsupportedCapability(broker string, capability string) *BrokerError {
	return &BrokerError{
		Category: CategoryCapabilityUnsupported,
		Broker:   broker,
		Code:     capability,
		Message:  fmt.Sprintf("%s does not support %s", broker, capability),
	}
}

func BrokerRejected(broker string, operation string, code string, message string, raw any) *BrokerError {
	if message == "" {
		message = "broker rejected request"
	}
	return &BrokerError{
		Category:  CategoryBrokerRejected,
		Broker:    broker,
		Operation: operation,
		Code:      code,
		Message:   message,
		Raw:       raw,
	}
}

func Decode(broker string, operation string, message string, raw any, cause error) *BrokerError {
	return &BrokerError{
		Category:  CategoryDecode,
		Broker:    broker,
		Operation: operation,
		Message:   message,
		Raw:       raw,
		Cause:     cause,
	}
}

func Auth(broker string, operation string, message string) *BrokerError {
	return &BrokerError{
		Category:  CategoryAuth,
		Broker:    broker,
		Operation: operation,
		Message:   message,
	}
}
