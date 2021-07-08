package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type BlobErrorMessage struct {
	Val string
}

func (m *BlobErrorMessage) Type() int {
	return 0
}

func decodeBlobError(br *bufio.Reader) (Message, error) {
	_, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bytes, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &BlobErrorMessage{
		Val: string(bytes),
	}, nil
}
