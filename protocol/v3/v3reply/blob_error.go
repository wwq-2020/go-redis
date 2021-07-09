package v3reply

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// BlobError BlobError
type BlobError struct {
	Val []byte
}

// NewBlobError NewBlobError
func NewBlobError() *BlobError {
	return &BlobError{}
}

// Type Type
func (r *BlobError) Type() protocol.ReplyType {
	return ReplyTypeBlobError
}

// DecodeFrom DecodeFrom
func (r *BlobError) DecodeFrom(br *bufio.Reader) error {
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

// DecodeBlobError DecodeBlobError
func DecodeBlobError(br *bufio.Reader) (*BlobError, error) {
	r := NewBlobError()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
