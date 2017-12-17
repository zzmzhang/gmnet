package gmnet

import (
	"fmt"
	"net"
	"sync.atomic"
)

const (
	subInit = iota
	subStart
	subStop
)

type sub struct {
	buffer chan []byte
	*connector
	status int32
}

func (s *sub) onRecv(data []byte, c *net.Conn) {
	s.buffer <- data
}

func NewSubscribe() (s *sub) {
	s := &sub{buffer: make(chan []byte, 10)}
	go s.receiveLoop()
}
