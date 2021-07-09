package v3command

import (
	"bufio"
	"io"

	"github.com/wwq-2020/go.common/errors"
)

// LatencyLatest LatencyLatest
type LatencyLatest struct {
}

// NewLatencyLatest NewLatencyLatest
func NewLatencyLatest() *LatencyLatest {
	return &LatencyLatest{}
}

// EncodeTo EncodeTo
func (c *LatencyLatest) EncodeTo(bw *bufio.Writer) error {
	if _, err := io.WriteString(bw, "latency latest\r\n"); err != nil {
		return errors.Trace(err)
	}
	return nil
}
