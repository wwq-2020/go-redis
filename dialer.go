package redis

import (
	"context"
	"net"

	"github.com/wwq-2020/go.common/errors"
)

// NewDialer NewDialer
func NewDialer(netDialer *net.Dialer, network, addr, password string, connFactory ConnFactory) Dialer {
	return func(ctx context.Context) (Conn, error) {
		c, err := netDialer.DialContext(ctx, network, addr)
		if err != nil {
			return nil, errors.Trace(err)
		}
		conn := connFactory(c)
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
