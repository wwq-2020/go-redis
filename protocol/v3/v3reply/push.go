package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Push Push
type Push struct {
	Items []protocol.Reply
}

// NewPush NewPush
func NewPush() *Push {
	return &Push{}
}

// Type Type
func (r *Push) Type() protocol.ReplyType {
	return ReplyTypePush
}

// DecodeFrom DecodeFrom
func (r *Push) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	length, err := decodeNextInt(br)
	if err != nil {
		return errors.Trace(err)
	}
	for i := 0; i < length; i++ {
		item, err := Decode(br)
		if err != nil {
			return errors.Trace(err)
		}
		r.Items = append(r.Items, item)
	}
	return nil
}

// DecodePush DecodePush
func DecodePush(br *bufio.Reader) (*Push, error) {
	r := NewPush()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
