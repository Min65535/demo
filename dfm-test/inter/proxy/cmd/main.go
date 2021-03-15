package main

import "fmt"

func main() {
	var x uint64 = 3
	// if x >= 1<<32 {
	// 	x >>= 32
	// 	fmt.Println("sss")
	// }

	x >>= 32
	fmt.Println("x:", x)
}
