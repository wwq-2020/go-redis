package protocol

import "bufio"

// Reply Reply
type Reply interface {
	Type() ReplyType
	DecodeFrom(br *bufio.Reader) error
}

// ReplyType ReplyType
type ReplyType string
