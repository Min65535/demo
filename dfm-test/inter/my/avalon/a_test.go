package avalon

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
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
	// a := A{}
	// b := B{}
	//
	// a.ElementB = &b
	// b.ElementA = &a
	// json.Marshal(a)
	//
	// rand.Seed(time.Now().UnixNano())
	//
	// for i := 0; i < 100; i++ {
	// 	x := rand.Int31n(7)
	// 	fmt.Println("x:", x)
	// }
	// var blkCalList []int
	// for bb := 0; bb < 7; bb++ {
	// 	blkCalList = append(blkCalList, bb)
	// }
	//
	// fmt.Println("blkCalList:", blkCalList)

	n1 := time.Now()
	fmt.Println("n1:", n1.Format("2006-01-02 15:04:05"))
	date := time.Date(n1.Year(), n1.Month(), n1.Day(), 0, 0, 0, 0, time.Local)
	dateTimestamp := date.Unix()
	fmt.Println("dateTimestamp unix:", dateTimestamp)
	fmt.Println("dateTimestamp 2 Unix:", time.Unix(dateTimestamp, 0).Local())
	dateNext := date.Add(time.Duration(1) * time.Second * 60 * 60 * 24)
	dateNextTimestamp := dateNext.Unix()
	fmt.Println("dateNextTimestamp unix:", dateNextTimestamp)
	fmt.Println("dateNextTimestamp 2 Unix:", time.Unix(dateNextTimestamp, 0).Local())

	st := "[]"
	var data []int64
	if err := json.ParseJson(st, &data); err != nil {
		fmt.Println("err:", err.Error())
	}

}
