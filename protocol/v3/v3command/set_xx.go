package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetXX SetXX
type SetXX struct {
	key   string
	value string
}

// NewSetXX NewSetXX
func NewSetXX(key, value string) *SetXX {
	return &SetXX{
		key:   key,
		value: value,
	}
}

// EncodeTo EncodeTo
func (c *SetXX) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("set %s %s xx\r\n", c.key, c.value)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
