package gmnet

import (
	"fmt"
)

type sock interface {
	Recv(p []byte) (n int, err error)
	Send(p []byte) (n int, err error)
}

type socketBase struct {
}
