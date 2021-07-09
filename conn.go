package redis

import (
	"bufio"
	"context"
	"net"

	"github.com/wwq-2020/go.common/errors"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go-redis/protocol/v3/v3command"
	"github.com/wwq-2020/go-redis/protocol/v3/v3reply"
)

// Dialer Dialer
type Dialer func(ctx context.Context) (Conn, error)

// Conn Conn
type Conn interface {
	RoundTrip(ctx context.Context, req protocol.Command, resp protocol.Reply) error
}

// Conn Conn
type conn struct {
	br  *bufio.Reader
	bw  *bufio.Writer
	raw net.Conn
}

// NewConn NewConn
func NewConn(raw net.Conn) Conn {
	br := bufio.NewReaderSize(raw, 1<<10)
	bw := bufio.NewWriterSize(raw, 1<<10)
	return &conn{
		br:  br,
		bw:  bw,
		raw: raw,
	}
}

// RoundTrip RoundTrip
func (c *conn) RoundTrip(ctx context.Context, req protocol.Command, resp protocol.Reply) error {
	d, ok := ctx.Deadline()
	if ok {
		c.raw.SetDeadline(d)
	}
	if err := req.EncodeTo(c.bw); err != nil {
		return errors.Trace(err)
	}
	if err := c.bw.Flush(); err != nil {
		return errors.Trace(err)
	}
	if err := resp.DecodeFrom(c.br); err != nil {
		return errors.Trace(err)
	}
	return nil
}

// Tracking Tracking
func (c *conn) Tracking(prefix string, callback func(key string)) error {
	var req protocol.Command
	var resp protocol.Reply
	req = v3command.NewHello(3)
	resp = v3reply.NewMap()
	if err := c.RoundTrip(context.TODO(), req, resp); err != nil {
		return errors.Trace(err)
	}
	req = v3command.NewTracking(prefix)
	resp = v3reply.NewSimpleString()
	if err := c.RoundTrip(context.TODO(), req, resp); err != nil {
		return errors.Trace(err)
	}
	for {
		pushReply, err := v3reply.DecodePush(c.br)
		if err != nil {
			return errors.Trace(err)
		}
		if len(pushReply.Items) != 2 {
			return errors.Trace(ErrUnExpectedPublishMsg)
		}
		BlobString, ok := pushReply.Items[0].(*v3reply.BlobString)
		if !ok {
			return errors.Trace(ErrUnExpectedPublishMsg)
		}
		if string(BlobString.Val) != "invalidate" {
			return errors.Trace(ErrUnExpectedPublishMsg)
		}
		Array, ok := pushReply.Items[1].(*v3reply.Array)
		if !ok {
			return errors.Trace(ErrUnExpectedPublishMsg)
		}
		for _, item := range Array.Items {
			BlobString, ok := item.(*v3reply.BlobString)
			if !ok {
				return errors.Trace(ErrUnExpectedPublishMsg)
			}
			callback(string(BlobString.Val))
		}
	}
}
