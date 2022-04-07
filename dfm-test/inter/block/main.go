package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randSliceValue(xs []int) ([]int, int) {
	var res []int
	var num int
	length := len(xs)
	index := rand.Intn(length)
	num = xs[index]
	if index > 0 && index < length-1 {
		front := xs[:index]
		end := xs[index+1:]
		res = append(res, front...)
		res = append(res, end...)
	}
	if index == 0 {
		res = append(res, xs[1:]...)
	}
	if index == length-1 {
		res = append(res, xs[:index]...)
	}
	return res, num
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var xs = []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		resArr, resNum := randSliceValue(xs)
		fmt.Println("resArr:", resArr)
		fmt.Println("resNum:", resNum)
	}
	// resArr, resNum := randSliceValue(xs)
	// fmt.Println("resArr:", resArr)
	// fmt.Println("resNum:", resNum)
}
