package redis

import (
	"context"
)

// Auth Auth
func Auth(ctx context.Context, conn Conn, password string) error {
	if password == "" {
		return nil
	}
	return nil
}
