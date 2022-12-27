package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/bigdata/fileio"
	"github.com/min65535/demo/dfm-test/inter/bigdata/timer"
	"io"
	"os"
	"os/signal"
	"syscall"
)

type Entity struct {
	signalChan chan os.Signal
	tm         timer.SvcTimer
}

func NewEntity(tm timer.SvcTimer) *Entity {
	return &Entity{
		signalChan: make(chan os.Signal, 1),
		tm:         tm,
	}
}

func (e *Entity) Start() {
	go e.tm.Start()
	signal.Notify(e.signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-e.signalChan
		fmt.Println("server get a signal:", s)
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			e.tm.Stop()
			fmt.Println("server exit")
			return
		default:
			return
		}
	}
}

type PublicInfo struct {
	Uid       uint64 `json:"uid"`
	Timestamp int64  `json:"timestamp"`
}

type BigEvent struct {
	EventName  string      `json:"event_name"`
	PublicInfo *PublicInfo `json:"public_info"`
	Extra      interface{} `json:"extra"`
}

var (
	spc      = flag.String("spc", "0 15 21 * * ?", "cron timer the given schedule")
	fromPath = flag.String("from", "/data/my/logs/", "data file from path")
	toPath   = flag.String("to", "/data/bigdata-storage/", "big data file path")
	file     = flag.String("file", "game.log", "game.log")
)

func sumTotalFuc() {
	filePathOrigin := fmt.Sprintf("%s/%s%s", "filePathDirOrigin", "mygame", "FileLast")
	fmt.Println("filePathOrigin: ", filePathOrigin)

	fi := fileio.NewFileIo(filePathOrigin)
	defer fi.Close()

	rd := fi.FileLineReader()
	for {
		ll, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}

		var param = &BigEvent{}
		if errJ := json.Unmarshal(ll, param); errJ != nil {
			continue
		}
		if param.PublicInfo != nil {
			// fmt.Println("param.PublicInfo.Timestamp: ", param.PublicInfo.Timestamp)
			// if yesTime21UnixMilli <= param.PublicInfo.Timestamp && param.PublicInfo.Timestamp < nowTime21UnixMilli {
			// 	fi.fileWrite(ll)
			// }
		}
	}
}

func main() {
	flag.Parse()
	fc := func() {
		sumTotalFuc()
	}
	// */1 * * * * ? 每秒
	// 0 0/1 * * * ? 每分
	// 0 15 21 * * ? 指定时间
	tm := timer.NewTm(*spc, fc)
	NewEntity(tm).Start()
}
