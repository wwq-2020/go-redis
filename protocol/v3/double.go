package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

// NumberMessage NumberMessage
type DoubleMessage struct {
	Val float64
}

func (m *DoubleMessage) Type() int {
	return 0
}

func decodeDouble(br *bufio.Reader) (Message, error) {
	val, err := decodeNextFloat64(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &DoubleMessage{
		Val: val,
	}, nil
}
