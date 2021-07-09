package redis

import (
	"bufio"
	"net"

	"github.com/wwq-2020/go.common/errors"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go-redis/protocol/v3/v3command"
	"github.com/wwq-2020/go-redis/protocol/v3/v3reply"
)

// Conn Conn
type Conn struct {
	br *bufio.Reader
	bw *bufio.Writer
}

// NewConn NewConn
func NewConn(c net.Conn) *Conn {
	br := bufio.NewReaderSize(c, 1<<10)
	bw := bufio.NewWriterSize(c, 1<<10)
	return &Conn{
		br: br,
		bw: bw,
	}
}

// RoundTrip RoundTrip
func (c *Conn) RoundTrip(req protocol.Command, resp protocol.Reply) error {
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
func (c *Conn) Tracking(prefix string, callback func(key string)) error {
	var req protocol.Command
	var resp protocol.Reply
	req = v3command.NewHello(3)
	resp = v3reply.NewMap()
	if err := c.RoundTrip(req, resp); err != nil {
		return errors.Trace(err)
	}
	req = v3command.NewTracking(prefix)
	resp = v3reply.NewSimpleString()
	if err := c.RoundTrip(req, resp); err != nil {
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
