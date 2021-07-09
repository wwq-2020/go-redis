package v3reply

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// BlobString BlobString
type BlobString struct {
	Val []byte
}

// NewBlobString NewBlobString
func NewBlobString() *BlobString {
	return &BlobString{}
}

// Type Type
func (r *BlobString) Type() protocol.ReplyType {
	return ReplyTypeBlobString
}

// DecodeFrom DecodeFrom
func (r *BlobString) DecodeFrom(br *bufio.Reader) error {
	_, err := br.Peek(1)
	if err != nil {
		return errors.Trace(err)
	}
	length, err := decodeNextInt(br)
	if err != nil {
		return errors.Trace(err)
	}
	buf := make([]byte, length+2)
	if _, err := io.ReadFull(br, buf); err != nil {
		return errors.Trace(err)
	}
	r.Val = buf[:length]
	return nil
}

// DecodeBlobString DecodeBlobString
func DecodeBlobString(br *bufio.Reader) (*BlobString, error) {
	r := NewBlobString()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
