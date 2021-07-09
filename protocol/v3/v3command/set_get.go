package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetGet SetGet
type SetGet struct {
	key   string
	value string
}

// NewSetGet NewSetGet
func NewSetGet(key, value string) *SetGet {
	return &SetGet{
		key:   key,
		value: value,
	}
}

// EncodeTo EncodeTo
func (c *SetGet) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("set %s %s get\r\n", c.key, c.value)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
