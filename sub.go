package gmnet

import (
	//"fmt"
	"net"
	//"sync/atomic"
)

const (
	subInit = iota
	subStart
	subStop
)

// Sub socket
type Sub struct {
	buffer chan []byte
	*connector
	status int32
}

func (s *Sub) onRecv(data []byte, c *net.Conn) {
	s.buffer <- data
}

func (s *Sub) Recv(data []byte) (n int, err error) {
	data = <-s.buffer
	return len(data), nil
}

func (s *Sub) Send(data []byte) (n int, err error) {
	return 0, nil
}

// NewSubscribe returns a new subscriber
func NewSubscribe() *Sub {
	s := &Sub{buffer: make(chan []byte, 10)}
	return s
}
