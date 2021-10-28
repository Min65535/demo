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

func TestDeferAndReturn(t *testing.T) {
	fmt.Println("DeferAndReturn:", DeferAndReturn())
	fmt.Println("DeferAndReturn1:", DeferAndReturn1())
	fmt.Println("DeferAndReturn2:", DeferAndReturn2())
	fmt.Println("DeferAndReturn3:", DeferAndReturn3())
}
