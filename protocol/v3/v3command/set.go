package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// Set Set
type Set struct {
	key   string
	value string
}

// NewSet NewSet
func NewSet(key, value string) *Set {
	return &Set{
		key:   key,
		value: value,
	}
}

// EncodeTo EncodeTo
func (c *Set) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("set %s %s\r\n", c.key, c.value)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
