package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

// NumberMessage NumberMessage
type NumberMessage struct {
	Val int
}

func (m *NumberMessage) Type() int {
	return 0
}

func decodeNumber(br *bufio.Reader) (Message, error) {
	val, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &NumberMessage{
		Val: val,
	}, nil
}
