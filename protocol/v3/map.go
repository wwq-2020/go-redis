package v3

import (
	"bufio"

	"github.com/wwq-2020/go.common/errors"
)

// MapMessage MapMessage
type MapMessage struct {
	m map[Message]Message
}

func decodeMap(br *bufio.Reader) (Message, error) {
	length, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	m := &MapMessage{
		m: make(map[Message]Message),
	}
	for i := 0; i < length; i++ {
		k, v, err := decodeMapItem(br)
		if err != nil {
			return nil, errors.Trace(err)
		}
		m.m[k] = v
	}
	return nil, nil
}

func decodeMapItem(br *bufio.Reader) (Message, Message, error) {
	key, err := Decode(br)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	value, err := Decode(br)
	if err != nil {
		return nil, nil, errors.Trace(err)
	}
	return key, value, nil
}
