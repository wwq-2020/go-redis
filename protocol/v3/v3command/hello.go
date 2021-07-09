package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// Hello Hello
type Hello struct {
	version int
}

// NewHello NewHello
func NewHello(version int) *Hello {
	return &Hello{
		version: version,
	}
}

// EncodeTo EncodeTo
func (c *Hello) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("hello %d\r\n", c.version)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
