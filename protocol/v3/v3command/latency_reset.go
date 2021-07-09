package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyReset LatencyReset
type LatencyReset struct {
	events []string
}

// NewLatencyReset NewLatencyReset
func NewLatencyReset(events ...string) *LatencyReset {
	return &LatencyReset{
		events: events,
	}
}

// EncodeTo EncodeTo
func (c *LatencyReset) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency reset\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
