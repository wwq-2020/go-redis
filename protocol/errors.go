package protocol

import "github.com/wwq-2020/go.common/errors"

// vars
var (
	ErrInvalidProtocol = errors.Std("invalid protocol")
	ErrUnsupportedType = errors.Std("unsupported type")
	ErrNil             = errors.Std("nil")
)
