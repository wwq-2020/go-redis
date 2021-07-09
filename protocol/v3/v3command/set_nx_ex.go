package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetNXEX SetNXEX
type SetNXEX struct {
	key     string
	value   string
	seconds int
}

// NewSetNXEX NewSetNXEX
func NewSetNXEX(key, value string, seconds int) *SetNXEX {
	return &SetNXEX{
		key:     key,
		value:   value,
		seconds: seconds,
	}
}

// EncodeTo EncodeTo
func (c *SetNXEX) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("set %s %s ex %d nx\r\n", c.key, c.value, c.seconds)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
