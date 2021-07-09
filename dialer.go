package redis

import (
	"context"
	"net"

	"github.com/wwq-2020/go.common/errors"
)

// NewDialer NewDialer
func NewDialer(network, addr, password string) Dialer {
	dialer := net.Dialer{}
	return func(ctx context.Context) (Conn, error) {
		c, err := dialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, errors.Trace(err)
		}
		conn := NewConn(c)
		if err := HandShake(ctx, conn); err != nil {
			c.Close()
			return nil, errors.Trace(err)
		}
		if err := Auth(ctx, conn, password); err != nil {
			c.Close()
			return nil, errors.Trace(err)
		}
		return conn, nil
	}
}
