package gmnet

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

const (
	sessionInit = iota
	sessionStart
	sessionStop
)

type listener struct {
	addr     string
	status   int32
	stopSig  chan bool
	ln       net.Listener
	connHash map[string]*net.Conn
}

func newListener() (s *listener) {
	s = &listener{status: sessionInit}
	return
}

func (l *listener) listening() {
	var timeDelay time.Duration
	for sessionStop != atomic.LoadInt32(&l.status) {
		c, err := l.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if timeDelay == 0 {
					timeDelay = 5 * time.Millisecond
				} else {
					timeDelay *= 2
					if timeDelay > time.Second {
						timeDelay = time.Second
					}
				}
				continue
			}
			panic(fmt.Sprintf("listen fatal error:%+v", err))
		}

		remoteAddr := c.RemoteAddr().String()
		l.connHash[remoteAddr] = &c
	}
}

func (l *listener) setAddress(addr string) (err error) {
	l.addr = addr
	if l.addr != "" {
		l.ln, err = net.Listen("tcp", l.addr)
		if err != nil {
			return err
		}

		go l.listening()
	}
	return
}
