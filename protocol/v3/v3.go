package v3

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Message Message
type Message interface {
	Type() int
}

// Decode Decode
func Decode(br *bufio.Reader) (Message, error) {
	buf, err := br.Peek(4)
	if err != nil {
		return nil, errors.Trace(err)
	}
	var msg Message
	switch buf[0] {
	case typeBlobString:
		if string(buf[:5]) == typeStream {
			panic("unimplemented")
			break
		}
		msg, err = decodeBlobString(br)
	case typeSimpleString:
		msg, err = decodeSimpleString(br)
	case typeSimpleError:
		msg, err = decodeSimpleError(br)
	case typeNumber:
		msg, err = decodeNumber(br)
	case typeNull:
		msg, err = decodeNull(br)
	case typeDouble:
		msg, err = decodeDouble(br)
	case typeBoolean:
		msg, err = decodeBoolean(br)
	case typeBlobError:
		msg, err = decodeBlobError(br)
	case typeVerbatimString:
		panic("unimplemented")
	case typeBigNumber:
		panic("unimplemented")
	case typeArray:
		msg, err = decodeArray(br)
	case typeMap:
		msg, err = decodeMap(br)
	case typeSet:
		panic("unimplemented")
	case typeAttribute:
		panic("unimplemented")
	case typePush:
		msg, err = decodePush(br)
	default:
		err = errors.Trace(protocol.ErrUnsupportedType)
	}
	if err != nil {
		return nil, errors.Trace(err)
	}
	return msg, nil
}
