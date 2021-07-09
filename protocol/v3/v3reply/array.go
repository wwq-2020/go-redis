package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Array Array
type Array struct {
	Items []protocol.Reply
}

// NewArray NewArray
func NewArray() *Array {
	return &Array{}
}

// Type Type
func (r *Array) Type() protocol.ReplyType {
	return ReplyArray
}

// DecodeFrom DecodeFrom
func (r *Array) DecodeFrom(br *bufio.Reader) error {
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

// DecodeArray DecodeArray
func DecodeArray(br *bufio.Reader) (*Array, error) {
	r := NewArray()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
