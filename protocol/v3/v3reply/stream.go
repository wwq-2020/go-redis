package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Stream Stream
type Stream struct {
}

// NewStream NewStream
func NewStream() *Stream {
	return &Stream{}
}

// Type Type
func (r *Stream) Type() protocol.ReplyType {
	return ReplyTypeStream
}

// DecodeFrom DecodeFrom
func (r *Stream) DecodeFrom(br *bufio.Reader) error {
	return nil
}

// DecodeStream DecodeStream
func DecodeStream(br *bufio.Reader) (*Stream, error) {
	r := NewStream()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
