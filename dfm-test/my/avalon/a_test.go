package avalon

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	var a = []int{1, 22, 3, 55, 6, 66, 21, 67, 454, 34, 24, 76, 80, 554, 2323, 8,}
	sort.Ints(a)
	fmt.Println("array:", a)
	fmt.Println(searchInts(a, 80))
}

func TestA(t *testing.T) {
	var a = []int{1, 22, 3, 55, 6, 66, 21, 67, 454, 34, 24, 76, 80, 81,554, 2323, 8,}
	sort.Ints(a)
	fmt.Println("array:", a)
	fmt.Println(binarySearch(a, 80))
}
