package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyDoctor LatencyDoctor
type LatencyDoctor struct {
}

// NewLatencyDoctor NewLatencyDoctor
func NewLatencyDoctor() *LatencyDoctor {
	return &LatencyDoctor{}
}

// EncodeTo EncodeTo
func (c *LatencyDoctor) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency doctor\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
