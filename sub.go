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

// NewSubscribe returns a new subscriber
func NewSubscribe() *Sub {
	s := &Sub{buffer: make(chan []byte, 10)}
	return s
}
