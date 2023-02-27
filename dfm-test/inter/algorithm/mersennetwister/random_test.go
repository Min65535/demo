package mersennetwister

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
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
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("go routine RandInt: ", mt.RandInt())
			fmt.Println("go routine GetUsageCount:", mt.GetUsageCount())
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("RandInt:", mt.RandInt())
			fmt.Println("GetUsageCount:", mt.GetUsageCount())
		}
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("--------end---------")
}

func TestMersenneTwister_ShuffleArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inp := rand.Intn(100)
	fmt.Println("input:", inp)
	fmt.Println("--------start---------")
	mt := NewMersenneTwister(inp)
	var in = []int{0, 1, 3, 6, 7, 9}
	arr, _ := Marshal(in)
	for i := 0; i < 10; i++ {
		mt.ShuffleArray(arr)
		fmt.Println("arr:", json.StringifyJson(arr))
	}
	fmt.Println("last arr:", json.StringifyJson(arr))
	var newIn []int
	arr.Unmarshal(&newIn)
	fmt.Println("newIn:", newIn)

	fmt.Println("--------end---------")
}

func TestMersenneTwister_GetRandomArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	inp := rand.Intn(100)
	fmt.Println("input:", inp)
	fmt.Println("--------start---------")
	mt := NewMersenneTwister(inp)
	var in = []int{0, 1, 3, 6, 7, 9}
	arr, _ := Marshal(in)
	for i := 0; i < 10; i++ {
		res := mt.GetRandomArray(arr, 3)
		fmt.Println("res:", json.StringifyJson(res))
	}
	fmt.Println("--------end---------")
	type inner struct {
		My int
	}
	var data [][]*inner
	if len(data) == 0 {
		fmt.Println("22222")
	}
	for x := 0; x < 7; x++ {
		var res2 []*inner
		for y := 0; y < 7; y++ {

			res2 = append(res2, &inner{My: x*10 + y})
		}
		data = append(data, res2)
	}
	fmt.Println("data:", json.StringifyJson(data))
	fmt.Println("data len:", len(data))

	var data1 [][]*inner
	if err := json.ParseJson(json.StringifyJson(data), &data1); err != nil {
		fmt.Println("err:", err.Error())
	}

	fmt.Println("data1:", json.StringifyJson(data1))
}
