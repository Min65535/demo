package algorithm

import (
	"fmt"
	"testing"
)

func TestNewMersenneTwister(t *testing.T) {
	mt := NewMersenneTwister(0)
	for i := 0; i < 10; i++ {
		fmt.Println(mt.Rand())
	}
}
