package ssi

import (
	"context"
	"net/url"
	"strconv"

	"github.com/shopspring/decimal"
)

func (b *Broker) getAndDecode(
	ctx context.Context,
	operation string,
	path string,
	params url.Values,
	out any,
) error {
	return b.sendAndDecode(ctx, operation, "GET", b.query(path, params), nil, false, out)
}

func (b *Broker) postAndDecode(
	ctx context.Context,
	operation string,
	path string,
	body any,
	out any,
) error {
	return b.sendAndDecode(ctx, operation, "POST", b.url(path), body, true, out)
}

func accountParams(accountID string) url.Values {
	params := url.Values{}
	params.Set("account", accountID)
	return params
}

func dateRangeParams(accountID string, fromKey string, fromDate string, toKey string, toDate string) url.Values {
	params := accountParams(accountID)
	params.Set(fromKey, fromDate)
	params.Set(toKey, toDate)
	return params
}

func setOptionalString(params url.Values, key string, value string) {
	if value != "" {
		params.Set(key, value)
	}
}

func setOptionalInt(params url.Values, key string, value int) {
	if value > 0 {
		params.Set(key, strconv.Itoa(value))
	}
}

func setOptionalDecimalBody(body map[string]any, key string, value decimal.Decimal) {
	if !value.IsZero() {
		body[key] = numberValue(&value)
	}
}
