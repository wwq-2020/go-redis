package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Set Set
type Set struct {
}

// NewSet NewSet
func NewSet() *Set {
	return &Set{}
}

// Type Type
func (r *Set) Type() protocol.ReplyType {
	return ReplyTypeSet
}

// DecodeFrom DecodeFrom
func (r *Set) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// DecodeSet DecodeSet
func DecodeSet(br *bufio.Reader) (*Set, error) {
	r := NewSet()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return nil, nil
}
