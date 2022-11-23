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

func TestDeferAndReturn1(t *testing.T) {
	for x := int32(0); x < 7; x++ {
		for y := int32(0); y < 7; y++ {

			fmt.Printf("[%d,%d]\n", x, y)
		}
	}
}
