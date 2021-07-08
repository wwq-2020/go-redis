package redis

import "github.com/wwq-2020/go.common/errors"

// errs
var (
	ErrGotUnexpectedMsgOnTrackingConn = errors.Std("got unexpected msg on tracking conn")
	ErrUnExpectedPublishMsg           = errors.Std("got unexpected publish msg")
)
