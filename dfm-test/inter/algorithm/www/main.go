package main

import (
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/algorithm"
	"time"
)

func main() {
	fmt.Println("--------start---------")
	mt := algorithm.NewMersenneTwister(90)
	for i := 0; i < 10; i++ {
		fmt.Println(mt.RandInt())
	}
	fmt.Println("--------end---------")
	time.Sleep(10 * time.Second)
	for {

	}
}
