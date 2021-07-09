package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// SimpleError SimpleError
type SimpleError struct {
	Val []byte
}

// NewSimpleError NewSimpleError
func NewSimpleError() *SimpleError {
	return &SimpleError{}
}

// Type Type
func (r *SimpleError) Type() protocol.ReplyType {
	return ReplyTypeSimpleError
}

// DecodeFrom DecodeFrom
func (r *SimpleError) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	bytes, _, err := br.ReadLine()
	if err != nil {
		return errors.Trace(err)
	}
	r.Val = bytes
	return nil
}

// DecodeSimpleError DecodeSimpleError
func DecodeSimpleError(br *bufio.Reader) (*SimpleError, error) {
	r := NewSimpleError()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
