package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// SetEX SetEX
type SetEX struct {
	key     string
	value   string
	seconds int
}

// NewSetEX NewSetEX
func NewSetEX(key, value string, seconds int) *SetEX {
	return &SetEX{
		key:     key,
		value:   value,
		seconds: seconds,
	}
}

// EncodeTo EncodeTo
func (c *SetEX) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("setex %s %d %s\r\n", c.key, c.seconds, c.value)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
