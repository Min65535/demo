package avalon

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"sort"
	"strconv"
	"strings"
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
	var indexes []int

	indexes[10] = 2

	fmt.Println("index:", indexes)
}

func TestShuffle(t *testing.T) {
	sl := []int64{1, 3, 7, 8}
	fmt.Println(sl)
	Shuffle(sl)
	fmt.Println(sl)

	Shuffle2(sl)
	fmt.Println(sl)
	nw := time.Now()
	fmt.Println("now unix:", nw.Unix())
	fmt.Println("now UnixNano:", nw.UnixNano())

	for _, v := range strings.Split("2", "_") {
		fmt.Println("v: ", v)
		gdID, _ := strconv.Atoi(v)
		fmt.Println("gd: ", gdID)
	}
}

func TestMoneyTotal(t *testing.T) {
	ta := MoneyTotal(123267.61, 59644, 0.02, 30)
	fmt.Println("total:", int64(ta))

	bal := ta / 480
	fmt.Println("bal:", int64(bal))

	no := time.Now()
	noUnix := time.Now().Unix()
	fmt.Println("no:", no)
	fmt.Println("noUnix:", noUnix)
	bef := no.Add(time.Duration(-5) * time.Minute)
	befUnix := bef.Unix()
	fmt.Println("bef:", bef)
	fmt.Println("befUnix:", befUnix)
}
