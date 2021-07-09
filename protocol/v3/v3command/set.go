package v3command

import "bufio"

// Set Set
type Set struct{}

// NewSet NewSet
func NewSet() *Set {
	return &Set{}
}

// EncodeTo EncodeTo
func (c *Set) EncodeTo(bw *bufio.Writer) error {
	return nil
}
