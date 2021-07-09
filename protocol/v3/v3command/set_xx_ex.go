package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetXXEX SetXXEX
type SetXXEX struct {
	key     string
	value   string
	seconds int
}

// NewSetXXEX NewSetXXEX
func NewSetXXEX(key, value string, seconds int) *SetXXEX {
	return &SetXXEX{
		key:     key,
		value:   value,
		seconds: seconds,
	}
}

// EncodeTo EncodeTo
func (c *SetXXEX) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("set %s %s ex %d xx\r\n", c.key, c.value, c.seconds)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
