package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

type PushMessage struct {
	Items []Message
}

func (m *PushMessage) Type() int {
	return 0
}

func decodePush(br *bufio.Reader) (Message, error) {
	length, err := decodeNextInt(br)
	if err != nil {
		return nil, nil
	}
	m := &PushMessage{
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
