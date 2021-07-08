package v3

import (
	"bufio"
	"encoding/json"
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

func decodeNextBlobString(br *bufio.Reader) (string, error) {
	_, err := decodeNextInt(br)
	if err != nil {
		return "", errors.Trace(err)
	}
	partBytes, _, err := br.ReadLine()
	if err != nil {
		return "", errors.Trace(err)
	}
	return string(partBytes), nil
}

func decodeNextNumber(br *bufio.Reader) (int, error) {
	partInt, err := decodeNextInt(br)
	if err != nil {
		return 0, errors.Trace(err)
	}
	return partInt, nil
}

func decodeNextNumberString(br *bufio.Reader) (string, error) {
	partInt, err := decodeNextInt(br)
	if err != nil {
		return "", errors.Trace(err)
	}
	return strconv.Itoa(partInt), nil
}

func decodeNextArray(br *bufio.Reader) ([]interface{}, error) {
	length, err := decodeNextInt(br)
	if err != nil {
		return nil, errors.Trace(err)
	}
	var items []interface{}
	for i := 0; i < length; i++ {
		item, err := decodeNextItemString(br)
		if err != nil {
			return nil, errors.Trace(err)
		}
		items = append(items, item)
	}
	return items, nil
}

func decodeNextArrayString(br *bufio.Reader) (string, error) {
	array, err := decodeNextArray(br)
	if err != nil {
		return "", errors.Trace(err)
	}
	arrayBytes, err := json.Marshal(array)
	if err != nil {
		return "", errors.Trace(err)
	}
	return string(arrayBytes), nil
}

func decodeNextSimpleString(br *bufio.Reader) (string, error) {
	buf, _, err := br.ReadLine()
	if err != nil {
		return "", errors.Trace(err)
	}
	if len(buf) <= 1 {
		return "", errors.Trace(protocol.ErrInvalidProtocol)
	}
	return string(buf[1:]), nil
}

func decodeNextItemString(br *bufio.Reader) (string, error) {
	buf, err := br.Peek(1)
	if err != nil {
		return "", errors.Trace(err)
	}
	var part string
	switch buf[0] {
	case typeNumber:
		part, err = decodeNextNumberString(br)
	case typeBlobString:
		part, err = decodeNextBlobString(br)
	case typeArray:
		part, err = decodeNextArrayString(br)
	}
	return part, nil
}
