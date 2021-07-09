package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyHelp LatencyHelp
type LatencyHelp struct {
}

// NewLatencyHelp NewLatencyHelp
func NewLatencyHelp() *LatencyHelp {
	return &LatencyHelp{}
}

// EncodeTo EncodeTo
func (c *LatencyHelp) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency help\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
