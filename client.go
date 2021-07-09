package redis

import (
	"context"
	"net"
	"time"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// vars
var (
	DefaultNetwork       = "tcp"
	DefaultAddr          = "127.0.0.1:6379"
	DefaultConnFactory   = NewConn
	DefaultMaxConns      = 10
	DefaultDialTimeout   = time.Second * 5
	DefaultDialKeepalive = time.Second * 30
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
	Addr          string
	Network       string
	Password      string
	MaxConns      int
	ConnFactory   ConnFactory
	DialTimeout   time.Duration
	DialKeepalive time.Duration
}

// vars
var (
	DefaultClientConf = &ClientConf{
		Addr:          DefaultAddr,
		Network:       DefaultNetwork,
		MaxConns:      DefaultMaxConns,
		ConnFactory:   DefaultConnFactory,
		DialTimeout:   DefaultDialTimeout,
		DialKeepalive: DefaultDialKeepalive,
	}
)

// ToPoolConf ToPoolConf
func (c *ClientConf) ToPoolConf() *PoolConf {
	netDialer := &net.Dialer{
		Timeout:   c.DialTimeout,
		KeepAlive: c.DialKeepalive,
	}
	dialer := NewDialer(netDialer, c.Network, c.Addr, c.Password, c.ConnFactory)
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
	if c.DialTimeout == 0 {
		c.DialTimeout = DefaultDialTimeout
	}
	if c.DialKeepalive == 0 {
		c.DialKeepalive = DefaultDialKeepalive
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
