package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyGraph LatencyGraph
type LatencyGraph struct {
	event string
}

// NewLatencyGraph NewLatencyGraph
func NewLatencyGraph(event string) *LatencyGraph {
	return &LatencyGraph{
		event: event,
	}
}

// EncodeTo EncodeTo
func (c *LatencyGraph) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency graph\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
