package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// Get Get
type Get struct {
	key string
}

// NewGet NewGet
func NewGet(key string) *Get {
	return &Get{
		key: key,
	}
}

// EncodeTo EncodeTo
func (c *Get) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, fmt.Sprintf("get %s\r\n", c.key)); err != nil {
		return errors.Trace(err)
	}
	return nil
}
