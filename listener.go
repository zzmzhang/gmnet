package gmnet

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	sessionInit = iota
	sessionStart
	sessionStop
)

type listener struct {
	addr    string
	status  int32
	stopSig chan bool
	ln      net.Listener

	mu       sync.RWMutex
	connHash map[string]*net.Conn

	closeSig chan string

	onRecv onRecvHandler
}

func newListener() (l *listener) {
	l = &listener{status: sessionInit}

	go func(l *listener) {
		for {
			key := <-l.closeSig
			l.mu.Lock()
			delete(connHash, key)
			l.mu.Unlock()
		}
	}(l)
	return
}

type onRecvHandler func(data []byte, c *net.Conn)

func (l *listener) GetConnHash(m map[string]*net.Conn) {
	mu.RLock()
	defer mu.RUnlock()
	for k, v := range connHash {
		m[k] = v
	}
}

func (l *listener) ConnSession(c *net.Conn, haskKey string, closeSig chan string) {
	for {
		// TODO handler multi message after the format of message is confirmed.
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("invalid disconnection err:%+v\n", err)
			}
			c.Close()
			closeSig <- hashKey
			return
		}
		msg := buf[:n]
		onRecv(msg, c)
		fmt.Printf("connSession recv:%+v\n", msg)
	}
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
		go l.ConnSession(&c, remoteAddr, closeSig)
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
