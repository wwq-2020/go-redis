package v3command

import "bufio"

// String String
type String struct{}

// NewString NewString
func NewString() *String {
	return &String{}
}

// EncodeTo EncodeTo
func (c *String) EncodeTo(bw *bufio.Writer) error {
	return nil
}
