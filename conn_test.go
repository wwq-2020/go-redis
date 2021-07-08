package redis

import (
	"bufio"
	"io"
	"net"
	"testing"
	"time"

	v3 "github.com/wwq-2020/go-redis/protocol/v3"
)

func TestConn(t *testing.T) {
	c, err := net.Dial("tcp4", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(c)
	cc := newConn(br)
	_, err = io.WriteString(c, "hello 3\r\n")
	if err != nil {
		panic(err)
	}
	v3.Decode(br)
	_, err = io.WriteString(c, "client tracking on bcast\r\n")
	if err != nil {
		panic(err)
	}
	v3.Decode(br)
	ch := make(chan string, 2)
	onInvalid := func(key string) {
		ch <- key
	}
	_, err = io.WriteString(c, "set a 1\r\n")
	if err != nil {
		panic(err)
	}
	v3.Decode(br)
	err = cc.Tracking(onInvalid)
	if err != nil {
		panic(err)
	}
	select {
	case <-time.After(time.Second * 2):
		panic("timeout")
	case <-ch:
	}
}
