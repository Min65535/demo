package apollo

import (
	"errors"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"regexp"
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
	fd := ^d
	fmt.Printf("%016b\n", fd)
	fmt.Printf("%016b\n", d)
	dXOfd := d ^ (^d)
	fmt.Printf("%016b\n", dXOfd)
	fmt.Printf("%016b\n", dXOfd&fd)

	fmt.Printf("%016b\n", d&(^d))
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%016b [A]\n", a)
	fmt.Printf("%016b [B]\n", b)
	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}

func TestMyApolloDemo3(t *testing.T) {
	var cids1 = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var cids2 = []string{"10", "11", "12", "13", "14", "15", "16"}
	var cidArr []string
	cidArr = append(cidArr, cids1...)
	cidArr = append(cidArr, cids2...)
	lgd := len(cidArr)
	var size = 5
	if lgd > size {
		var arr [][]string
		doz := lgd / size
		fmt.Println("doz:", doz)
		eqa := lgd % size
		fmt.Println("eqa:", eqa)
		if eqa != 0 {
			doz = doz + 1
		}
		for i := 0; i < doz; i++ {
			fmt.Println("i:", i)
			if i == doz-1 {
				arr = append(arr, cidArr[size*i:])
			} else {
				fmt.Println("size#sss:", cidArr[size*i:size*(i+1)])
				arr = append(arr, cidArr[size*i:size*(i+1)])
			}
		}
		fmt.Println("arr:", json.StringifyJson(arr))
		fmt.Println("arr.len:", len(arr))
	}
}

func checkEmail(email string) error {
	pattern := `\w+(([-+.]\w+)*(\*){4})@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	if !regexp.MustCompile(pattern).MatchString(email) {
		return errors.New("请输入合法邮箱")
	}
	return nil
}
func TestMyApolloDemo4(t *testing.T) {
	var email = "298****@qq.com"
	if err := checkEmail(email); err != nil {
		t.Log("err:", err.Error())
	}
}
