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
func (p *Pub) Send(msg []byte) {
	p.buffer <- msg
}

func (p *Pub) sendingLoop() {
	for pubStop != atomic.LoadInt32(&p.status) {
		b := <-p.buffer
		p.mu.RLock()
		defer p.mu.RUnlock()
		for _, v := range p.connHash {
			v.Write(b)
		}
	}
}

// NewSocket generate a new socket
func NewSocket() *Pub {
	p := &Pub{buffer: make(chan []byte, 10),
		status: pubInit}
	go p.sendingLoop()
	return p
}
