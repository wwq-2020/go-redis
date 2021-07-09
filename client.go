package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// vars
var (
	DefaultNetwork     = "tcp"
	DefaultAddr        = "127.0.0.1:6379"
	DefaultConnFactory = NewConn
	DefaultMaxConns    = 10
)

// Client Client
type Client interface {
	stringCommand
	listCommand
	hashCommand
	setCommand
	zsetCommand
	latencyCommand
	RoundTrip(ctx context.Context, req protocol.Command, resp protocol.Reply) error
}

type client struct {
	pool Pool
}

// ClientConf ClientConf
type ClientConf struct {
	Addr        string
	Network     string
	Password    string
	MaxConns    int
	ConnFactory ConnFactory
}

// vars
var (
	DefaultClientConf = &ClientConf{
		Addr:        DefaultAddr,
		Network:     DefaultNetwork,
		MaxConns:    DefaultMaxConns,
		ConnFactory: DefaultConnFactory,
	}
)

// ToPoolConf ToPoolConf
func (c *ClientConf) ToPoolConf() *PoolConf {
	dialer := NewDialer(c.Network, c.Addr, c.Password, c.ConnFactory)
	return &PoolConf{
		Dialer:   dialer,
		MaxConns: c.MaxConns,
	}
}

// Fill Fill
func (c *ClientConf) Fill() {
	if c.Network == "" {
		c.Network = DefaultNetwork
	}
	if c.Addr == "" {
		c.Network = DefaultAddr
	}
	if c.ConnFactory == nil {
		c.ConnFactory = NewConn
	}
	if c.MaxConns == 0 {
		c.MaxConns = DefaultMaxConns
	}
}

// NewClient NewClient
func NewClient(conf *ClientConf) Client {
	if conf == nil {
		conf = DefaultClientConf
	}
	conf.Fill()
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
