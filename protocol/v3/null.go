package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type NullMessage struct {
}

func (m *NullMessage) Type() int {
	return 0
}

func decodeNull(br *bufio.Reader) (Message, error) {
	_, _, err := br.ReadLine()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &NullMessage{}, nil
}
