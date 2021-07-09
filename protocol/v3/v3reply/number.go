package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Number Number
type Number struct {
	Val int
}

// NewNumber NewNumber
func NewNumber() *Number {
	return &Number{}
}

// Type Type
func (r *Number) Type() protocol.ReplyType {
	return ReplyTypeNumber
}

// DecodeFrom DecodeFrom
func (r *Number) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	val, err := decodeNextInt(br)
	if err != nil {
		return errors.Trace(err)
	}

	r.Val = val
	return nil
}

// DecodeNumber DecodeNumber
func DecodeNumber(br *bufio.Reader) (*Number, error) {
	r := NewNumber()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
