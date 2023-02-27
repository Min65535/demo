package logger

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	data_logger "github.com/min65535/demo/dfm-test/inter/data_logger"
	"math/rand"
	"testing"
	"time"
)

// PublicInfo   test用
type PublicInfo struct {
	Uid       uint64 `json:"uid"`
	Timestamp int64  `json:"timestamp"`
}

func (p *PublicInfo) GetTimestamp() {
	p.Timestamp = time.Now().UnixMilli()
}

type Event struct {
	Name       string      `json:"name"`
	PublicInfo *PublicInfo `json:"public_info"`
	Extra      interface{} `json:"extra"`
	M          map[uint64]int32
}

func (be *Event) GetTimeStamp() int64 {
	if be != nil && be.PublicInfo != nil {
		return be.PublicInfo.Timestamp
	}
	return 0
}

func (be *Event) GetBytes() []byte {
	if be != nil {
		return json.StringifyJsonToBytes(be)
	}
	return nil
}

func getAbsIds(startId, endId, currentId, runTimes int32) []int32 {
	rand.Seed(time.Now().UnixNano())
	allId := make(map[int32]int32)
	for i := startId; i <= endId; i++ {
		allId[i] = i
	}
	fmt.Println("allId:", allId)
	var res []int32
	var x int32
	for x = 1; x <= runTimes; x++ {
		var tmr []int32
		z := x + currentId
		if _, ok := allId[z]; ok {
			tmr = append(tmr, z)
		}
		n := x*(-1) + currentId
		if _, ok := allId[n]; ok {
			tmr = append(tmr, n)
		}
		if len(tmr) > 1 {
			rand.Shuffle(len(tmr), func(i, j int) { tmr[i], tmr[j] = tmr[j], tmr[i] })
		}
		if len(tmr) > 0 {
			res = append(res, tmr...)
		}
	}
	return res
}

func TestDisableColors(t *testing.T) {
	e := &Event{
		Name: "111",
	}
	m := map[uint64]int32{1: 1}
	fmt.Println("m:", m)
	if e.M == nil || len(e.M) == 0 {
		e.M = m
	}
	fmt.Println("e:", e)
	var startId, endId, current int32
	startId = 14
	endId = 24
	current = 18
	idsRes := getAbsIds(startId, endId, current, 111)
	fmt.Println("idsRes:", idsRes)
}

func TestFileIsExist(t *testing.T) {
	t.Skip()
	data_logger.Init(data_logger.WithDirRoot("/mnt/d/code/go/src/demo/dfm-test/inter/filedata/data1"), data_logger.WithLogName("data"))
	data_logger.Info(&Event{
		Name: "hhhh",
		PublicInfo: &PublicInfo{
			Uid:       9527,
			Timestamp: time.Now().UnixMilli(),
		},
		Extra: map[string]interface{}{"test": 1},
	})
	data_logger.Info(&Event{
		Name: "hhhh",
		PublicInfo: &PublicInfo{
			Uid:       9528,
			Timestamp: time.Now().UnixMilli(),
		},
		Extra: map[string]interface{}{"test": 2},
	})
	yesDay := time.Now().AddDate(0, 0, -1)
	data_logger.Info(&Event{
		Name: "hhhh",
		PublicInfo: &PublicInfo{
			Uid:       9528,
			Timestamp: yesDay.UnixMilli(),
		},
		Extra: map[string]interface{}{"test": 3},
	})
}
