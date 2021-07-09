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

// PoolConf PoolConf
type PoolConf struct {
	Dialer   Dialer
	MaxConns int
}

type pool struct {
	dialer   Dialer
	maxConns int
}

// NewPool NewPool
func NewPool(conf *PoolConf) Pool {
	return &pool{
		dialer:   conf.Dialer,
		maxConns: conf.MaxConns,
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
