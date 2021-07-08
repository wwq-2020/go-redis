package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type BlobStringMessage struct {
	Val string
}

func (m *BlobStringMessage) Type() int {
	return 0
}

func decodeBlobString(br *bufio.Reader) (Message, error) {
	_, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bytes, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &BlobStringMessage{
		Val: string(bytes),
	}, nil
}
