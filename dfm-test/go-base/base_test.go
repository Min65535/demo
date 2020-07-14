package go_base

import (
	"fmt"
	"testing"
)

func TestCheckOnce(t *testing.T) {
	checkOnce()
}

func TestCreatePointer(t *testing.T) {
	s := CreatePointer()
	fmt.Println("s:", s)
}

func TestLazyDefer(t *testing.T) {
	LazyDefer()
}
