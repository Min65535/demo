package avalon

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var a = []int{1, 22, 3, 55, 6, 66, 21, 67, 454, 34, 24, 76, 80, 554, 2323, 8}
	sort.Ints(a)
	fmt.Println("array:", a)
	fmt.Println(searchInts(a, 80))
}

func TestA(t *testing.T) {
	var a = []int{1, 22, 3, 55, 6, 66, 21, 67, 454, 34, 24, 76, 80, 81, 554, 2323, 8}
	sort.Ints(a)
	fmt.Println("array:", a)
	fmt.Println(binarySearch(a, 80))
}

type A struct {
	ElementB *B
}

type B struct {
	ElementA *A
}

func TestA1(t *testing.T) {
	a := A{}
	b := B{}

	a.ElementB = &b
	b.ElementA = &a
	json.Marshal(a)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		x := rand.Int31n(7)
		fmt.Println("x:", x)
	}
	var blkCalList []int
	for bb := 0; bb < 7; bb++ {
		blkCalList = append(blkCalList, bb)
	}
	
	fmt.Println("blkCalList:", blkCalList)

}
