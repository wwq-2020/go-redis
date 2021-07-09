package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Status Status
type Status struct {
	Success bool
	Err     error
	Val     []byte
}

// NewStatus NewStatus
func NewStatus() *Status {
	return &Status{}
}

// Type Type
func (r *Status) Type() protocol.ReplyType {
	return ReplyStatus
}

// DecodeFrom DecodeFrom
func (r *Status) DecodeFrom(br *bufio.Reader) error {
	reply, err := Decode(br)
	if err != nil {
		return errors.Trace(err)
	}
	switch t := reply.(type) {
	case *SimpleError:
		msg := string(t.Val)
		if msg == "" {
			return errors.Trace(protocol.ErrNil)
		}
		r.Err = errors.New(msg)
	case *SimpleString:
		r.Success = true
	case *BlobString:
		r.Val = t.Val
	case *Null:
		r.Err = errors.Trace(protocol.ErrNil)
	default:
		return errors.Trace(protocol.ErrInvalidProtocol)
	}
	return nil
}

// DecodeStatus DecodeStatus
func DecodeStatus(br *bufio.Reader) (*Status, error) {
	r := NewStatus()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
