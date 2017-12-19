package gmnet

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type connector struct {
	addrs    []string
	connHash sync.Map
	onRecv   onRecvHandler

	closeSig chan string
}

func newConnector() (con *connector) {
	con = &connector{}
	go func(con *connector) {
		for {
			key := <-con.closeSig
			con.ConnectRoutine(key)
		}
	}(con)
	return con
}

func (con *connector) Connect(addr string) {
	con.addrs = append(con.addrs, addr)
	go con.ConnectRoutine(addr)
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
	con.connHash.Store(addr, &conn)
}

func (con *connector) ConnSession(c net.Conn, hashKey string, closeSig chan string) {
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
	con.onRecv(msg, &c)
	fmt.Printf("connSession recv:%+v\n", msg)
}
