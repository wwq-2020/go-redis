package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type SimpleStringMessage struct {
	Val string
}

func (m *SimpleStringMessage) Type() int {
	return 0
}

func decodeSimpleString(br *bufio.Reader) (Message, error) {
	bytes, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &SimpleStringMessage{
		Val: string(bytes),
	}, nil
}
