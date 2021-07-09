package v3reply

import (
	"bufio"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

// Decode Decode
func Decode(br *bufio.Reader) (protocol.Reply, error) {
	buf, err := br.Peek(1)
	if err != nil {
		return nil, errors.Trace(err)
	}
	var r protocol.Reply
	switch buf[0] {
	case typeBlobString:
		buf, err := br.Peek(5)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if string(buf) == typeStream {
			r = NewStream()
			break
		}
		r = NewBlobString()
	case typeSimpleString:
		r = NewSimpleString()
	case typeSimpleError:
		r = NewSimpleError()
	case typeNumber:
		r = NewNumber()
	case typeNull:
		r = NewNull()
	case typeDouble:
		r = NewDouble()
	case typeBoolean:
		r = NewBoolean()
	case typeBlobError:
		r = NewBlobError()
	case typeVerbatimString:
		r = NewVerbatimString()
	case typeBigNumber:
		r = NewBigNumber()
	case typeArray:
		r = NewArray()
	case typeMap:
		r = NewBlobString()
	case typeSet:
		r = NewSet()
	case typeAttribute:
		r = NewBlobString()
	case typePush:
		r = NewBlobString()
	default:
		return nil, errors.Trace(protocol.ErrUnsupportedType)
	}
	if err := r.DecodeFrom(br); err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}
