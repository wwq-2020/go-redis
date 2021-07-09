package redis

import (
	"context"

	"github.com/wwq-2020/go.common/errors"
)

// Pool Pool
type Pool interface {
	Get(ctx context.Context) (Conn, error)
	Put(Conn)
}

type pool struct {
	dialer Dialer
	size   int
}

// NewPool NewPool
func NewPool(size int, dialer Dialer) Pool {
	return &pool{
		size:   size,
		dialer: dialer,
	}
}

func (p *pool) Get(ctx context.Context) (Conn, error) {
	conn, err := p.dialer(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return conn, nil
}

func (p *pool) Put(conn Conn) {

}
