package v3reply

import (
	"bufio"
	"strconv"

	"github.com/wwq-2020/go-redis/protocol"
	"github.com/wwq-2020/go.common/errors"
)

func decodeNextInt(br *bufio.Reader) (int, error) {
	buf, _, err := br.ReadLine()
	if err != nil {
		return 0, errors.Trace(err)
	}
	if len(buf) <= 1 {
		return 0, errors.Trace(protocol.ErrInvalidProtocol)
	}
	val, err := strconv.Atoi(string(buf[1:]))
	if err != nil {
		return 0, errors.Trace(protocol.ErrInvalidProtocol)
	}
	return val, nil
}

func decodeNextFloat64(br *bufio.Reader) (float64, error) {
	buf, _, err := br.ReadLine()
	if err != nil {
		return 0, errors.Trace(err)
	}
	if len(buf) <= 1 {
		return 0, errors.Trace(protocol.ErrInvalidProtocol)
	}
	val, err := strconv.ParseFloat(string(buf[1:]), 64)
	if err != nil {
		return 0, errors.Trace(protocol.ErrInvalidProtocol)
	}
	return val, nil
}

func decodeNextBool(br *bufio.Reader) (bool, error) {
	buf, _, err := br.ReadLine()
	if err != nil {
		return false, errors.Trace(err)
	}
	if len(buf) <= 1 {
		return false, errors.Trace(protocol.ErrInvalidProtocol)
	}
	val, err := strconv.ParseBool(string(buf[1:]))
	if err != nil {
		return false, errors.Trace(protocol.ErrInvalidProtocol)
	}
	return val, nil
}
