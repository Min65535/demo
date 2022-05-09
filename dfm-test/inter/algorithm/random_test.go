package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewMersenneTwister(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inp := rand.Intn(100)
	fmt.Println("input:", inp)
	fmt.Println("--------start---------")
	mt := NewMersenneTwister(inp)
	for i := 0; i < 100; i++ {
		fmt.Println(mt.Rand())
	}
	fmt.Println("--------end---------")
}
