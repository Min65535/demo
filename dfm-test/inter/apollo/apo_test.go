package apollo

import (
	"fmt"
	"testing"
	"time"
)

func TestMyApolloDemo(t *testing.T) {
	x := [3]int{1, 2, 3}
	go func(arr [3]int) {
		arr[0] = 7
		fmt.Println("x arr:", arr) // prints [7 2 3]
	}(x)

	time.Sleep(1 * time.Second)
	fmt.Println(x) // prints [1 2 3]

	y := []int{4, 5, 6}
	go func(arr []int) {
		arr[0] = 8
		fmt.Println("y arr:", arr) // prints [8 5 6]
	}(y)

	time.Sleep(1 * time.Second)
	fmt.Println(y) // prints [8 5 6]

	t.Skip()
	MyApolloDemo()
}

func TestMyApolloDemo2(t *testing.T) {
	var d uint8 = 2
	fd:=^d
	fmt.Printf("%016b\n",fd)
	fmt.Printf("%016b\n", d)
	dXOfd := d ^ (^d)
	fmt.Printf("%016b\n", dXOfd)
	fmt.Printf("%016b\n", dXOfd&fd)

	fmt.Printf("%016b\n", d&(^d))
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%016b [A]\n",a)
	fmt.Printf("%016b [B]\n",b)
	fmt.Printf("%08b (NOT B)\n",^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff)
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
}
