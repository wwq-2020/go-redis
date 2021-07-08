package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type SimpleErrorMessage struct {
	Val string
}

func (m *SimpleErrorMessage) Type() int {
	return 0
}

func decodeSimpleError(br *bufio.Reader) (Message, error) {
	bytes, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &SimpleErrorMessage{
		Val: string(bytes),
	}, nil
}
