package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Map Map
type Map struct {
	Items map[protocol.Reply]protocol.Reply
}

// NewMap NewMap
func NewMap() *Map {
	return &Map{
		Items: make(map[protocol.Reply]protocol.Reply),
	}
}

// Type Type
func (r *Map) Type() protocol.ReplyType {
	return ReplyTypeMap
}

// DecodeFrom DecodeFrom
func (r *Map) DecodeFrom(br *bufio.Reader) error {
	length, err := decodeNextInt(br)
	if err != nil {
		return errors.Trace(err)
	}
	for i := 0; i < length; i++ {
		k, v, err := decodeMapItem(br)
		if err != nil {
			return errors.Trace(err)
		}
		r.Items[k] = v
	}
	return nil
}

// DecodeMap DecodeMap
func DecodeMap(br *bufio.Reader) (*Map, error) {
	r := NewMap()
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return nil, nil
}

func decodeMapItem(br *bufio.Reader) (protocol.Reply, protocol.Reply, error) {
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
