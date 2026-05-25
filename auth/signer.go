package auth

import (
	"context"

	"github.com/vnbrokers/vnbrokers-go/transport"
)

type Signer interface {
	Sign(context.Context, transport.HTTPRequest) (transport.HTTPRequest, error)
}
