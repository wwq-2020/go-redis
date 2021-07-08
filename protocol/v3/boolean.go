package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

// NumberMessage NumberMessage
type BooleanMessage struct {
	Val bool
}

func (m *BooleanMessage) Type() int {
	return 0
}

func decodeBoolean(br *bufio.Reader) (Message, error) {
	val, err := decodeNextBool(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &BooleanMessage{
		Val: val,
	}, nil
}
