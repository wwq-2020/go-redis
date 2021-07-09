package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Boolean Boolean
type Boolean struct {
	Val bool
}

// NewBoolean NewBoolean
func NewBoolean() *Boolean {
	return &Boolean{}
}

// Type Type
func (r *Boolean) Type() protocol.ReplyType {
	return ReplyTypeBoolean
}

// DecodeFrom DecodeFrom
func (r *Boolean) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	val, err := decodeNextBool(br)
	if err != nil {
		return errors.Trace(err)
	}
	r.Val = val
	return nil
}

// DecodeBoolean DecodeBoolean
func DecodeBoolean(br *bufio.Reader) (*Boolean, error) {
	r := NewBoolean()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
