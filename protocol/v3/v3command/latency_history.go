package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyHistory LatencyHistory
type LatencyHistory struct {
	event string
}

// NewLatencyHistory NewLatencyHistory
func NewLatencyHistory(event string) *LatencyHistory {
	return &LatencyHistory{
		event: event,
	}
}

// EncodeTo EncodeTo
func (c *LatencyHistory) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency help\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
