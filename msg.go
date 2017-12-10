package gmnet

import (
	"fmt"
)

type MSG struct {
	content []byte
}

func init() {

}

func (m *MSG) Init(size, int) {
	fmt.Println("hello world")
	fmt.Println("hello world")
	m.content = make([]byte, size)
}
