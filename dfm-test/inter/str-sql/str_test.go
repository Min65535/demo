package str_sql

import (
	"fmt"
	"testing"
)

func TestGetSqlStr(t *testing.T) {
	str, err := GetSqlStr(82343000, 82344000)
	if err != nil {
		fmt.Println("Err:", err.Error())
		return
	}
	fmt.Println("str:", str)
}

func TestTimerMy(t *testing.T) {
	TimerMy()
}

func TestGetSqlStr2(t *testing.T) { // 1, 12, 23321, 9832, 345,
	// var uis = []uint64{4, 1341234235, 5655, 656, 894, 55267, 77222, 757, 458454, 12, 22, 33, 44, 55, 777, 232, 332, 131, 331, 133, 132, 623, 44, 47, 57, 58, 59, 533, 56, 6, 78678, 77, 78, 878, 8, 89, 999564, 78578, 345453}
	var uis = []uint64{4, 1341234235, 5655}
	fmt.Println("len:", len(uis))
	term := len(uis) / 20
	tail := len(uis) % 20
	var origin []uint64
	var head, end int
	if term > 0 {
		for i := 0; i < term; i++ {
			end = head + 20
			newUs := uis[head:end]
			fmt.Println("i, newUs: ", i, newUs)
			fmt.Println("i, newUs len: ", i, len(newUs))
			origin = append(origin, newUs...)
			fmt.Println("i, origin: ", i, origin)
			head = end
		}
	}
	if tail > 0 {
		tailArr := uis[end:]
		fmt.Println("end, len, tailArr: ", end, len(tailArr), tailArr)
		origin = append(origin, tailArr...)
	}
	fmt.Println("origin latest: ", origin)
	fmt.Println("origin latest len: ", len(origin))
}
