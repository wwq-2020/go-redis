package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Null Null
type Null struct {
}

// NewNull NewNull
func NewNull() *Null {
	return &Null{}
}

// Type Type
func (r *Null) Type() protocol.ReplyType {
	return ReplyTypeNull
}

// DecodeFrom DecodeFrom
func (r *Null) DecodeFrom(br *bufio.Reader) error {
	_, _, err := br.ReadLine()
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// DecodeNull DecodeNull
func DecodeNull(br *bufio.Reader) (*Null, error) {
	r := NewNull()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
