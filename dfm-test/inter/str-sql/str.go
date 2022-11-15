package str_sql

import (
	"errors"
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/timer/inner"
	"time"
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

func TimerMy() {
	tr := inner.NewTimer()
	now := time.Now()
	nowUnix := now.Unix()
	latestStart := 1667491200
	nextStart := time.Unix(int64(latestStart), 0).Add(time.Second * time.Duration(2592000))
	nextStartUnix := nextStart.Unix()
	fmt.Println("nextStartUnix: ", nextStartUnix)
	remainUnix := nextStartUnix - nowUnix
	fmt.Println("remainUnix: ", remainUnix)
	tr.ResetTimer(int32(remainUnix))
	for {
		select {
		case <-tr.Timer.C:
			fmt.Println("do do do")
			// if err = tuc.DoUpdateUsersExpAndGameSeasonRcds(); err != nil {
			// 	log.Errorf("SeasonRun#DoUpdateUsersExpAndGameSeasonRcds#err:%s", err.Error())
			// }
			tr.ResetTimer(int32(2592000))
		}
	}
}
