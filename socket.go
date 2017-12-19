package gmnet

import (
	"net"
)

type onRecvHandler func(data []byte, c net.Conn)

type sock interface {
	Recv(p []byte) (n int, err error)
	Send(p []byte) (n int, err error)
}

type socketBase struct {
}
