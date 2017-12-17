package gmnet

import (
	"fmt"
)

type connector struct {
	addrs []string
}

func (con *connector) Connect(addr string) {
	con.addrs = append[con.addrs, addr]
}
