package gmnet

import (
	"fmt"
	"testing"
)

func TestSubFunc(t *testing.T) {
	fmt.Println("TestSubFunc")
}

func TestMain(m *testing.M) {
	// TO DO Setup
	retCode := m.Run()
	// TO DO TearDown
}
