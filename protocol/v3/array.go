package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

// ArragMessage ArragMessage
type ArragMessage struct {
	Items []Message
}

func (m *ArragMessage) Type() int {
	return 0
}

func decodeArray(br *bufio.Reader) (Message, error) {
	length, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	m := &ArragMessage{
		Items: make([]Message, 0, length),
	}
	for i := 0; i < length; i++ {
		item, err := Decode(br)
		if err != nil {
			return nil, errors.Trace(err)
		}
		m.Items = append(m.Items, item)
	}
	return m, nil
}
