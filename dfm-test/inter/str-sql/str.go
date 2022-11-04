package str_sql

import (
	"errors"
	"fmt"
)

func GetSqlStr(fromId, toId int64) (string, error) {
	if fromId > toId {
		return "", errors.New("bad id")
	}
	str := "insert toptop_match.robot_record(`uid`,create_time) values "
	by := ""
	for fromId <= toId {
		by = by + fmt.Sprintf(`(%d,now()),`, fromId)
		fromId++
	}
	by = by[:len(by)-1]
	fmt.Println("by:", by)
	str = str + by + ";"
	return str, nil
}
