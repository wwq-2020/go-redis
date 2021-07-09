package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Client Client
type Client interface {
	stringCommand
	RoundTrip(ctx context.Context, req protocol.Command, resp protocol.Reply) error
}

type client struct {
	pool Pool
}

// ClientConf ClientConf
type ClientConf struct {
	Addr     string
	Network  string
	Password string
	MaxConns int
}

// vars
var (
	DefaultClientConf = &ClientConf{
		Addr:     "127.0.0.1:6379",
		Network:  "tcp",
		MaxConns: 10,
	}
)

// ToPoolConf ToPoolConf
func (c *ClientConf) ToPoolConf() *PoolConf {
	dialer := NewDialer(c.Network, c.Addr, c.Password)
	return &PoolConf{
		Dialer:   dialer,
		MaxConns: c.MaxConns,
	}
}

// NewClient NewClient
func NewClient(conf *ClientConf) Client {
	pool := NewPool(conf.ToPoolConf())
	return &client{
		pool: pool,
	}
}

// DefaultClient DefaultClient
func DefaultClient() Client {
	return NewClient(DefaultClientConf)
}

// Do Do
func (c *client) RoundTrip(ctx context.Context, req protocol.Command, resp protocol.Reply) error {
	conn, err := c.pool.Get(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	defer c.pool.Put(conn)
	if err := conn.RoundTrip(ctx, req, resp); err != nil {
		return errors.Trace(err)
	}
	return nil
}
