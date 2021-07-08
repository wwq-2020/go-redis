package redis

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"

	v3 "github.com/wwq-2020/go-redis/protocol/v3"
)

// Conn Conn
type Conn struct {
	br       *bufio.Reader
	tracking bool
}

func newConn(br *bufio.Reader) *Conn {
	return &Conn{
		br: br,
	}
}

// Tracking Tracking
func (c *Conn) Tracking(callback func(key string)) error {
	msg, err := v3.Decode(c.br)
	if err != nil {
		return errors.Trace(err)
	}
	pushMsg, ok := msg.(*v3.PushMessage)
	if !ok {
		return errors.Trace(ErrGotUnexpectedMsgOnTrackingConn)
	}
	if len(pushMsg.Items) != 2 {
		return errors.Trace(ErrUnExpectedPublishMsg)
	}
	blobStringMessage, ok := pushMsg.Items[0].(*v3.BlobStringMessage)
	if !ok {
		return errors.Trace(ErrUnExpectedPublishMsg)
	}
	if blobStringMessage.Val != "invalidate" {
		return errors.Trace(ErrUnExpectedPublishMsg)
	}
	arrayMessage, ok := pushMsg.Items[1].(*v3.ArragMessage)
	if !ok {
		return errors.Trace(ErrUnExpectedPublishMsg)
	}
	for _, item := range arrayMessage.Items {
		blobStringMessage, ok := item.(*v3.BlobStringMessage)
		if !ok {
			return errors.Trace(ErrUnExpectedPublishMsg)
		}
		callback(blobStringMessage.Val)
	}
	return nil
}
