package redis

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/wwq-2020/go.common/log"
)

func TestConn(t *testing.T) {
	c, err := net.Dial("tcp4", "127.0.0.1:6377")
	if err != nil {
		panic(err)
	}
	cc := NewConn(c)
	ch := make(chan string, 2)

	onInvalid := func(key string) {
		fmt.Println(key)
		ch <- key
	}
	go func() {
		log.Error(cc.Tracking("a", onInvalid))
	}()
	select {
	case <-time.After(time.Second * 10):
		panic("timeout")
	case <-ch:
	}

}
