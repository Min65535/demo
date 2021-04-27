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
	fmt.Println(x) // prints [7 2 3]

	y := []int{4, 5, 6}
	go func(arr []int) {
		arr[0] = 8
		fmt.Println("y arr:", arr) // prints [7 2 3]
	}(y)

	time.Sleep(1 * time.Second)
	fmt.Println(y) // prints [7 2 3]

	t.Skip()
	MyApolloDemo()
}
