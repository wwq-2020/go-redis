package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// BigNumber BigNumber
type BigNumber struct {
}

// NewBigNumber NewBigNumber
func NewBigNumber() *BigNumber {
	return &BigNumber{}
}

// Type Type
func (r *BigNumber) Type() protocol.ReplyType {
	return ReplyTypeBigNumber
}

// DecodeFrom DecodeFrom
func (r *BigNumber) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// DecodeBigNumber DecodeBigNumber
func DecodeBigNumber(br *bufio.Reader) (*BigNumber, error) {
	r := NewBigNumber()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
