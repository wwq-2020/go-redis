package protocol

import "bufio"

// Command Command
type Command interface {
	EncodeTo(bw *bufio.Writer) error
}
