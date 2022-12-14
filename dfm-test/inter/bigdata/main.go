package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/bigdata/timer"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
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

const (
	// /data/bigdata-storage/20221202
	MyName   = "mygame_"
	FileLast = ".log"
)

type PublicInfo struct {
	Uid       uint64 `json:"uid"`
	Timestamp int64  `json:"timestamp"`
}

type BigEvent struct {
	EventName  string      `json:"event_name"`
	PublicInfo *PublicInfo `json:"public_info"`
	Extra      interface{} `json:"extra"`
}

type FileIo struct {
	first bool
	fn    *os.File
	lock  sync.Mutex
}

func NewFileIo(newFile string) *FileIo {
	fn, first, sign := fileIsExistOrCreate(newFile)
	if !sign {
		return nil
	}
	return &FileIo{
		first: first,
		fn:    fn,
	}
}

func (fi *FileIo) close() {
	fi.fn.Close()
}

func (fi *FileIo) fileWrite(ll []byte) {
	fi.lock.Lock()
	defer fi.lock.Unlock()
	if fi.first {
		fi.fn.WriteString(fmt.Sprintf("%s", string(ll)))
		fi.first = false
	} else {
		fi.fn.WriteString(fmt.Sprintf("\n%s", string(ll)))
	}
}

func (fi *FileIo) fileRange(name string, yesTime21UnixMilli, nowTime21UnixMilli int64) {
	f, err3 := os.OpenFile(name, os.O_RDONLY, 0666)
	if err3 != nil {
		return
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		ll, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}
		// fmt.Println("ll:", string(ll))
		var param = &BigEvent{}
		if errJ := json.Unmarshal(ll, param); errJ != nil {
			continue
		}
		if param.PublicInfo != nil {
			// fmt.Println("param.PublicInfo.Timestamp: ", param.PublicInfo.Timestamp)
			if yesTime21UnixMilli <= param.PublicInfo.Timestamp && param.PublicInfo.Timestamp < nowTime21UnixMilli {
				fi.fileWrite(ll)
			}
		}
	}
}

// getTime 获得零点时间
func getTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func fileFuc(date string) {
	now := time.Now()
	if date != "" {
		var err error
		now, err = time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println("now err: ", err)
			return
		}
	}
	today := now.Format("20060102")
	fmt.Println("today: ", today)
	yesterdayTime := now.AddDate(0, 0, -1)
	yesterday := yesterdayTime.Format("20060102")
	fmt.Println("yesterday: ", yesterday)
	filePathYes := fmt.Sprintf("%s%s%s/%s%s%s", *fromPath, MyName, yesterday, MyName, yesterday, FileLast)
	fmt.Println("filePathYes: ", filePathYes)

	filePathTo := fmt.Sprintf("%s%s%s/%s%s%s", *fromPath, MyName, today, MyName, today, FileLast)
	fmt.Println("filePathTo: ", filePathTo)

	filePathDirTrue := fmt.Sprintf("%s%s", *toPath, today)
	fmt.Println("filePathDirTrue: ", filePathDirTrue)
	if err := dirCheck(filePathDirTrue); err != nil {
		fmt.Println("filePathDirTrue create err: ", err)
		return
	}
	filePathTrue := fmt.Sprintf("%s/%s%s", filePathDirTrue, "matchgame", FileLast)
	fmt.Println("filePathTrue: ", filePathTrue)

	yesTime21 := getTime(yesterdayTime).Add(time.Duration(21) * time.Hour)
	fmt.Println("yesTime21: ", yesTime21.Format("20060102150405"))
	yesTime21UnixMilli := yesTime21.UnixMilli()
	fmt.Println("yesTime21UnixMilli: ", yesTime21UnixMilli)

	nowTime21 := getTime(now).Add(time.Duration(21) * time.Hour)
	fmt.Println("nowTime21: ", nowTime21.Format("20060102150405"))
	nowTime21UnixMilli := nowTime21.UnixMilli()
	fmt.Println("nowTime21UnixMilli: ", nowTime21UnixMilli)
	fi := NewFileIo(filePathTrue)
	defer fi.close()
	if dirOrFileExist(filePathYes) {
		fmt.Println("dirOrFileExist(filePathYes) true")
		fi.fileRange(filePathYes, yesTime21UnixMilli, nowTime21UnixMilli)
		if dirOrFileExist(filePathYes) {
			execute("gzip", filePathYes)
		}
	}

	if dirOrFileExist(filePathTo) {
		fmt.Println("dirOrFileExist(filePathTo) true")
		fi.fileRange(filePathTo, yesTime21UnixMilli, nowTime21UnixMilli)
	}

	if dirOrFileExist(filePathTrue) {
		execute("gzip", filePathTrue)
	}
	execute("ls", "-alh")
}

func dirOrFileExist(dir string) bool {
	_, err1 := os.Stat(dir)
	if err1 != nil {
		if os.IsNotExist(err1) {
			return false
		}
	}
	return true
}

func dirCheck(dir string) error {
	path, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	_, err1 := os.Stat(path)
	if err1 != nil {
		if os.IsNotExist(err1) {
			err2 := os.Mkdir(path, os.ModePerm)
			if err2 != nil {
				return err2
			} else {
				return nil
			}
		} else {
			return err1
		}
	}
	return nil
}

func fileIsExistOrCreate(name string) (*os.File, bool, bool) {
	var firstCreate bool
	path, err := filepath.Abs(name)
	if err != nil {
		// fmt.Println("Abs err: ", err)
		return nil, firstCreate, false
	}

	// for unix
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return nil, firstCreate, false
	}

	_, err1 := os.Stat(name)
	if err1 != nil {
		if os.IsNotExist(err1) {
			f, err2 := os.Create(name)
			if err2 != nil {
				return nil, firstCreate, false
			} else {
				firstCreate = true
				return f, firstCreate, true
			}
		}
	} else {
		f, err3 := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err3 != nil {
			return nil, firstCreate, false
		}
		return f, firstCreate, true
	}
	return nil, firstCreate, false
}

func execute(name string, args ...string) {
	cmd := exec.Command(name, args...)
	ot, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("execute#name:%s, args:%v, err:%v", name, args, err))
		return
	}
	fmt.Println("execute output: ", string(ot))
}

var (
	spc      = flag.String("spc", "0 15 21 * * ?", "cron timer the given schedule")
	fromPath = flag.String("from", "/data/my/logs/", "data file from path")
	toPath   = flag.String("to", "/data/bigdata-storage/", "big data file path")
	date     = flag.String("date", "", "2022-12-13")
)

func main() {
	flag.Parse()
	fc := func() {
		fileFuc(*date)
	}
	// */1 * * * * ? 每秒
	// 0 0/1 * * * ? 每分
	// 0 15 21 * * ? 指定时间
	tm := timer.NewTm(*spc, fc)
	NewEntity(tm).Start()
}
