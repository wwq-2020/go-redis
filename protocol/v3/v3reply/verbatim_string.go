package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// VerbatimString VerbatimString
type VerbatimString struct {
	Val []byte
}

// NewVerbatimString NewVerbatimString
func NewVerbatimString() *VerbatimString {
	return &VerbatimString{}
}

// Type Type
func (r *VerbatimString) Type() protocol.ReplyType {
	return ReplyTypeVerbatimString
}

// DecodeFrom DecodeFrom
func (r *VerbatimString) DecodeFrom(br *bufio.Reader) error {
	return nil
}

// DecodeVerbatimString DecodeVerbatimString
func DecodeVerbatimString(br *bufio.Reader) (*VerbatimString, error) {
	r := NewVerbatimString()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
