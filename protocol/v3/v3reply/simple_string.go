package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// SimpleString SimpleString
type SimpleString struct {
	Val []byte
}

// NewSimpleString NewSimpleString
func NewSimpleString() *SimpleString {
	return &SimpleString{}
}

// Type Type
func (r *SimpleString) Type() protocol.ReplyType {
	return ReplyTypeSimpleString
}

// DecodeFrom DecodeFrom
func (r *SimpleString) DecodeFrom(br *bufio.Reader) error {
	bytes, _, err := br.ReadLine()
	if err != nil {
		return errors.Trace(err)
	}
	r.Val = bytes
	return nil
}

// DecodeSimpleString DecodeSimpleString
func DecodeSimpleString(br *bufio.Reader) (*SimpleString, error) {
	r := NewSimpleString()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
