package brokerage

import (
	"context"
	"testing"

	"github.com/vnbrokers/vnbrokers-go/brokers/dnse/native/dto"
)

func TestBrokerageServiceExposesCareByAccounts(t *testing.T) {
	var service Service
	if false {
		_, _ = service.GetCareByAccounts(context.Background(), dto.GetCareByAccountsRequest{})
	}
}
