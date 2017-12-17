package gmnet

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connector struct {
	addrs    []string
	connHash sync.Map
	onRecv   onRecvHandler
}

func newConnector() (con *Connector) {
	con := &connector{}
	go func(con *Connector) {
		for {
			key := <-con.closeSig
			con.ConnectRoutine()
		}
	}(con)
}

func (con *connector) Connect(addr string) {
	con.addrs = append(con.addrs, addr)
	go ConnectRoutine(addr)
}

func (con *connector) ConnectRoutine(addr string) {
	conn, err := net.Dial("tcp", addr)

	timeDelay := time.Second
	for err != nil {
		fmt.Printf("failed to connect to %s\n", addr)
		time.Sleep(timeDelay)
		conn, err = net.Dial("tcp", addr)
	}
	fmt.Printf("connect to %s\n", addr)
	connHash.Store(addr, &conn)
}

func (con *connector) ConnSession(c *net.Conn, haskKey string, closeSig chan string) {
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
