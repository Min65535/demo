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
	mt := NewMersenneTwister(90)
	for i := 0; i < 10; i++ {
		fmt.Println(mt.RandInt())
	}
	fmt.Println("--------end---------")
}
