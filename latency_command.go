package redis

import (
	"context"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go-redis/protocol/v3/v3command"
	"github.com/wwq-2020/go-redis/protocol/v3/v3reply"
	"github.com/wwq-2020/go.common/errors"
)

type latencyCommand interface {
	LatencyDoctor(ctx context.Context) (string, error)
	LatencyGraph(ctx context.Context, event string) (string, error)
	LatencyHelp(ctx context.Context, event string) ([]string, error)
	LatencyHistory(ctx context.Context, event string) ([]*LatencyHistory, error)
	LatencyLatest(ctx context.Context) ([]*LatencyLatest, error)
	LatencyReset(ctx context.Context, events ...string) (int, error)
}

func (c *client) LatencyDoctor(ctx context.Context) (string, error) {
	req := v3command.NewLatencyDoctor()
	resp := v3reply.NewBlobString()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return "", errors.Trace(err)
	}
	return string(resp.Val), nil
}

func (c *client) LatencyGraph(ctx context.Context, event string) (string, error) {
	req := v3command.NewLatencyGraph(event)
	resp := v3reply.NewBlobString()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return "", errors.Trace(err)
	}
	return string(resp.Val), nil
}

func (c *client) LatencyHelp(ctx context.Context, event string) ([]string, error) {
	req := v3command.NewLatencyHelp()
	resp := v3reply.NewArray()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return nil, errors.Trace(err)
	}
	rets := make([]string, 0, len(resp.Items))
	for _, item := range resp.Items {
		each, ok := item.(*v3reply.BlobString)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		rets = append(rets, string(each.Val))
	}
	return rets, nil
}

// LatencyHistory LatencyHistory
type LatencyHistory struct {
	Timestamp int64
	Latency   int
}

func (c *client) LatencyHistory(ctx context.Context, event string) ([]*LatencyHistory, error) {
	req := v3command.NewLatencyHistory(event)
	resp := v3reply.NewArray()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return nil, errors.Trace(err)
	}
	rets := make([]*LatencyHistory, 0, len(resp.Items))
	for _, item := range resp.Items {
		each, ok := item.(*v3reply.Array)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		if len(each.Items) != 2 {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		timestampPart, ok := each.Items[0].(*v3reply.Number)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		latencyPart, ok := each.Items[0].(*v3reply.Number)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		rets = append(rets, &LatencyHistory{
			Timestamp: int64(timestampPart.Val),
			Latency:   latencyPart.Val,
		})
	}
	return rets, nil
}

// LatencyLatest LatencyLatest
type LatencyLatest struct {
	Event     string
	Timestamp int64
	Latency   int
	AllTime   int
}

func (c *client) LatencyLatest(ctx context.Context) ([]*LatencyLatest, error) {
	req := v3command.NewLatencyLatest()
	resp := v3reply.NewArray()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return nil, errors.Trace(err)
	}
	rets := make([]*LatencyLatest, 0, len(resp.Items))
	for _, item := range resp.Items {
		each, ok := item.(*v3reply.Array)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		if len(each.Items) != 2 {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		eventPart, ok := each.Items[0].(*v3reply.BlobString)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		timestampPart, ok := each.Items[0].(*v3reply.Number)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		latencyPart, ok := each.Items[0].(*v3reply.Number)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		allTimePart, ok := each.Items[0].(*v3reply.Number)
		if !ok {
			return nil, errors.Trace(protocol.ErrInvalidProtocol)
		}
		rets = append(rets, &LatencyLatest{
			Event:     string(eventPart.Val),
			Timestamp: int64(timestampPart.Val),
			Latency:   latencyPart.Val,
			AllTime:   allTimePart.Val,
		})
	}
	return rets, nil
}

func (c *client) LatencyReset(ctx context.Context, events ...string) (int, error) {
	req := v3command.NewLatencyReset(events...)
	resp := v3reply.NewNumber()
	if err := c.RoundTrip(ctx, req, resp); err != nil {
		return 0, errors.Trace(err)
	}
	return resp.Val, nil
}
