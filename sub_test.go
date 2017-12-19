package gmnet

import (
	"fmt"
	"os"
	"testing"
)

func TestSubFunc(t *testing.T) {
	fmt.Println("TestSubFunc")
	sub := NewSubscribe()
	for i := 1; i < 10; i++ {
		msg := <-sub.buffer
		fmt.Println(msg)
	}
}

func TestMain(m *testing.M) {
	// TO DO Setup
	retCode := m.Run()
	// TO DO TearDown
	os.Exit(retCode)
}
