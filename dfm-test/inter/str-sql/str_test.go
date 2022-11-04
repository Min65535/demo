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
