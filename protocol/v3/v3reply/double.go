package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Double Double
type Double struct {
	Val float64
}

// NewDouble NewDouble
func NewDouble() *Double {
	return &Double{}
}

// Type Type
func (r *Double) Type() protocol.ReplyType {
	return ReplyTypeDouble
}

// DecodeFrom DecodeFrom
func (r *Double) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	val, err := decodeNextFloat64(br)
	if err != nil {
		return errors.Trace(err)
	}
	r.Val = val
	return nil
}

// DecodeDouble DecodeDouble
func DecodeDouble(br *bufio.Reader) (*Double, error) {
	r := NewDouble()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
