package main

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"math/rand"
	"strconv"
	"strings"
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

	checkArr := [][][]int{
		{{0, 2}, {0, 3}},
		{{-1, 1}, {1, 1}},
		{{-1, 1}, {-2, 1}},
		{{1, 1}, {2, 1}},
		{{0, -2}, {0, -3}},
		{{-1, -1}, {1, -1}},
		{{-1, -1}, {-2, -1}},
		{{1, -1}, {2, -1}},
		{{2, 0}, {3, 0}},
		{{1, -1}, {1, 1}},
		{{1, -1}, {1, -2}},
		{{1, 1}, {1, 2}},
		{{-2, 0}, {-3, 0}},
		{{-1, -1}, {-1, 1}},
		{{-1, -1}, {-1, -2}},
		{{-1, 1}, {-1, 2}},
	}

	fmt.Println("checkArr:", checkArr)

	var numArr []int
	arr1 := strings.Split("|5_10|5_20", "|")
	fmt.Println("arr1:", json.StringifyJson(arr1))
	fmt.Println("arr1.len:", len(arr1))
	for _, b := range arr1 {
		arr2 := strings.Split(b, "_")
		switch len(arr2) {
		case 2:
			n, _ := strconv.ParseFloat(arr2[1], 64)
			if n > 0 {
				numArr = append(numArr, int(n))
			}
		case 1:
			numArr = append(numArr, 0)
		}

	}
	fmt.Println("numArr:", numArr)
	var cost int
	numTotal := 2
	ln := len(numArr)
	if numTotal >= ln {
		cost = numArr[len(numArr)-1]
	} else if numTotal >= 0 && numTotal < ln {
		cost = numArr[numTotal]
	}

	fmt.Println("cost:", cost)

	rand.Seed(time.Now().UnixNano())
	ranChoiceNum := 2
	var directionMap = map[int]string{1: "U", 2: "D", 3: "L", 4: "R"}
	var res []string
	var str string
	for len(res) < ranChoiceNum {
		ranNum := rand.Intn(4) + 1
		if !strings.Contains(str, directionMap[ranNum]) {
			res = append(res, directionMap[ranNum])
			str = str + directionMap[ranNum]
		}
	}

	fmt.Println("res:", res)
	var rn []int
	for y := 3; y >= 3; y-- {
		rn = append(rn, y)
	}
	fmt.Println("rn:", rn)
	var rl []int
	for y := 3; y <= 3; y++ {
		rl = append(rl, y)
	}
	fmt.Println("rl:", rl)

	sdf := "udl"
	for i := 0; i < len(sdf); i++ {
		switch sdf[i] {
		case 'u':
			fmt.Println("-----u------")
		case 'd':
			fmt.Println("-----d------")
		case 'l':
			fmt.Println("-----l------")
		}
	}
}
