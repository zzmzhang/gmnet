package gmnet

import (
	"fmt"
	"net"
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

func (p *Pub) onRecv(data []byte, c *net.Conn) {
	fmt.Println(data)
}

func (p *Pub) Recv(data []byte) (n int, err error) {
	return 0, nil
}

func (p *Pub) Send(msg []byte) (n int, err error) {
	p.buffer <- msg
	return len(msg), nil
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
