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
	for i := 0; i < 10; i++ {
		fmt.Println(mt.RandInt())
	}
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
}
