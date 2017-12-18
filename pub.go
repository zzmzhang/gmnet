package gmnet

import (
	//"fmt"
	//"net"
	"sync/atomic"
)

const (
	pubInit = iota
	pubStart
	pubStop
)

// Pub socket
type Pub struct {
	*listener
	buffer chan []byte
	status int32
}

// Send publish message
func (p *Pub) Send(msg []byte) (err error) {
	buffer <- msg
}

func (p *Pub) sendingLoop() {
	for pubStop != atomic.LoadInt32(&p.status) {
		b := <-p.buffer
		mu.RLock()
		defer mu.RUnlock()
		for k, v := range connHash {
			v.Write(b)
		}
	}
}

// NewSocket generate a new socket
func NewSocket() *Pub {
	p := &pub{buffer: make(chan []byte, 10),
		status: pubInit}
	go p.sendingLoop()
	return p
}
