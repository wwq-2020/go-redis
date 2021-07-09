package v3command

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// Tracking Tracking
type Tracking struct {
	prefix string
}

// NewTracking NewTracking
func NewTracking(prefix string) *Tracking {
	return &Tracking{
		prefix: prefix,
	}
}

// EncodeTo EncodeTo
func (c *Tracking) EncodeTo(bw *bufio.Writer) error {
	if c.prefix != "" {
		_, err := io.WriteString(bw, fmt.Sprintf("client tracking on bcast prefix %s\r\n", c.prefix))
		if err != nil {
			return errors.Trace(err)
		}
		return nil
	}
	_, err := io.WriteString(bw, "client tracking on bcast\r\n")
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}
