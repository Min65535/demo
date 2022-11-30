package logger

import (
	"github.com/dipperin/go-ms-toolkit/json"
	data_logger "github.com/min65535/demo/dfm-test/inter/data_logger"
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
