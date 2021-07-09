package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol/v3/v3command"
	"github.com/wwq-2020/go-redis/protocol/v3/v3reply"
	"github.com/wwq-2020/go.common/errors"
)

type stringCommand interface {
	Set(ctx context.Context, key, value string) error
	SetGet(ctx context.Context, key, value string) (string, error)
	SetNX(ctx context.Context, key, value string) (bool, error)
	SetNXEX(ctx context.Context, key, value string, seconds int) (bool, error)
	SetXX(ctx context.Context, key, value string) (bool, error)
	SetXXEX(ctx context.Context, key, value string, seconds int) (bool, error)
	SetEX(ctx context.Context, key, value string, seconds int) (bool, error)
	Get(ctx context.Context, key string) (string, error)
}

// Set Set
func (c *client) Set(ctx context.Context, key, value string) error {
	req := v3command.NewSet(key, value)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return errors.Trace(err)
	}
	if resp.Err != nil {
		return resp.Err
	}
	return nil
}

// SetGet SetGet
func (c *client) SetGet(ctx context.Context, key, value string) (string, error) {
	req := v3command.NewSetGet(key, value)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return "", errors.Trace(err)
	}
	if resp.Err != nil {
		return "", errors.Trace(resp.Err)
	}
	return string(resp.Val), nil
}

// SetNX SetNX
func (c *client) SetNX(ctx context.Context, key, value string) (bool, error) {
	req := v3command.NewSetNX(key, value)
	resp := v3reply.NewNumber()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return false, errors.Trace(err)
	}
	return resp.Val == 1, nil
}

// SetNXEX SetNXEX
func (c *client) SetNXEX(ctx context.Context, key, value string, seconds int) (bool, error) {
	req := v3command.NewSetNXEX(key, value, seconds)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return false, errors.Trace(err)
	}
	if resp.Err != nil {
		return false, errors.Trace(resp.Err)
	}
	return resp.Success, nil
}

// SetXX SetXX
func (c *client) SetXX(ctx context.Context, key, value string) (bool, error) {
	req := v3command.NewSetXX(key, value)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return false, errors.Trace(err)
	}
	if resp.Err != nil {
		return false, errors.Trace(resp.Err)
	}
	return resp.Success, nil
}

// SetXXEX SetXXEX
func (c *client) SetXXEX(ctx context.Context, key, value string, seconds int) (bool, error) {
	req := v3command.NewSetXXEX(key, value, seconds)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return false, errors.Trace(err)
	}
	if resp.Err != nil {
		return false, errors.Trace(resp.Err)
	}
	return resp.Success, nil
}

// SetEX SetEX
func (c *client) SetEX(ctx context.Context, key, value string, seconds int) (bool, error) {
	req := v3command.NewSetEX(key, value, seconds)
	resp := v3reply.NewNumber()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return false, errors.Trace(err)
	}
	return resp.Val == 1, nil
}

// Get Get
func (c *client) Get(ctx context.Context, key string) (string, error) {
	req := v3command.NewGet(key)
	resp := v3reply.NewStatus()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return "", errors.Trace(err)
	}
	if resp.Err != nil {
		return "", errors.Trace(resp.Err)
	}
	return string(resp.Val), nil
}
