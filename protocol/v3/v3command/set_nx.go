package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetNX SetNX
type SetNX struct {
	key   string
	value string
}

// NewSetNX NewSetNX
func NewSetNX(key, value string) *SetNX {
	return &SetNX{
		key:   key,
		value: value,
	}
}

// EncodeTo EncodeTo
func (c *SetNX) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("setnx %s %s\r\n", c.key, c.value)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
