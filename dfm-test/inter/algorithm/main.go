package main

import (
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/algorithm/mersennetwister"
)

func main() {
	fmt.Println("--------start---------")
	mt := mersennetwister.NewMersenneTwister(90)
	for i := 0; i < 10; i++ {
		fmt.Println(mt.RandInt())
	}
	fmt.Println("--------end---------")
}
