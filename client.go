package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Client Client
type Client struct {
	pool Pool
}

// Conf Conf
type Conf struct {
	Addr    string
	Network string
}

// New New
func New(pool Pool) *Client {
	return &Client{
		pool: pool,
	}
}

// Do Do
func (c *Client) Do(ctx context.Context, req protocol.Command, resp protocol.Reply) error {
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
