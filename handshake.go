package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol/v3/v3command"
	"github.com/wwq-2020/go-redis/protocol/v3/v3reply"
	"github.com/wwq-2020/go.common/errors"
)

// HandShake HandShake
func HandShake(ctx context.Context, conn Conn) error {
	req := v3command.NewHello(3)
	resp := v3reply.NewMap()
	if err := conn.RoundTrip(ctx, req, resp); err != nil {
		return errors.Trace(err)
	}
	return nil
}
