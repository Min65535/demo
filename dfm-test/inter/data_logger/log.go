package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//
// func InitLogger(name string) {
// 	switch qyenv.GetDockerEnv() {
// 	case "", "0":
// 	default:
// 		// logDir := "/data/match/logs/toptop_match_data/"
// 		logDir := "/mnt/d/code/go/src/demo/dfm-test/inter/filedata/data/"
// 		if err := os.MkdirAll(logDir, 0755); err != nil {
// 			panic(err)
// 		}
// 		byS := "SS"
// 		if _, err := os.Create(logDir + name + ".log"); err != nil {
// 			fmt.Println("ccccc: ", err)
// 		}
// 		if err := os.WriteFile(logDir+name+".log", []byte(byS), 755); err != nil {
// 			fmt.Println("wwww: ", err)
// 		}
// 	}
// }

// GetTime 获得零点时间
func GetTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// GetDays 获得时间天数差
func GetDays(start, end time.Time) int {
	return int(GetTime(start).Sub(GetTime(end)).Hours() / 24)
}

const (
	DayStr = "2006-01-02"
)

type TimeSpecial interface {
	GetTimeStamp() int64
	GetBytes() []byte
}

type Logger interface {
	Info(an any)
}

type Formatter interface {
	Format(msg []byte) ([]byte, error)
}

type myLogFileCfg struct {
	dir      string
	fileName string
	now      time.Time
}

type Option func(*myLogFileCfg)

var DataLogger = &logger{mw: &MutexWrap{}}

func WithDirRoot(dir string) Option {
	return func(my *myLogFileCfg) {
		DirCheck(dir)
		my.dir = dir
	}
}

func WithLogName(name string) Option {
	return func(my *myLogFileCfg) {
		my.fileName = name
	}
}

func Init(opts ...Option) {
	c := &myLogFileCfg{now: time.Now()}
	for _, v := range opts {
		v(c)
		if c.dir != "" {
			if err := DirCheck(c.dir); err != nil {
				panic(fmt.Sprintf("DirCheck err: %v", err))
			}
		}
	}
	DataLogger.format = &MatchFormatter{}
	DataLogger.cfg = c
}

type logger struct {
	cfg    *myLogFileCfg
	mw     *MutexWrap
	format Formatter
}

type MutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *MutexWrap) Lock() {
	if !mw.disabled {
		mw.lock.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.disabled {
		mw.lock.Unlock()
	}
}

func (mw *MutexWrap) Disable() {
	mw.disabled = true
}

func (m *logger) Info(an any) {
	m.mw.Lock()
	defer m.mw.Unlock()
	inn, ok := an.(TimeSpecial)
	if !ok {
		fmt.Println("TimeSpecial no ok")
		return
	}
	bts := inn.GetBytes()
	if bts == nil {
		fmt.Println("TimeSpecial no bts")
		return
	}
	tm := inn.GetTimeStamp()
	if tm == 0 {
		fmt.Println("TimeSpecial no tm")
		return
	}
	tmt := time.UnixMilli(tm)
	dayLogStr := GetTime(tmt).Format(DayStr)
	fileName := fmt.Sprintf("%s-%s.log", m.cfg.fileName, dayLogStr)
	fileFullPath := fmt.Sprintf("%s/%s", m.cfg.dir, fileName)
	// if GetDays(tmt, m.cfg.now) > 0 {
	// 	m.cfg.now = tmt
	// }
	// if GetDays(tmt, time.Now()) > 0 {
	// 	m.cfg.now = tmt
	// }

	f, first, sign := FileIsExistOrCreate(fileFullPath)
	if !sign {
		return
	}
	defer f.Close()
	if first {
		f.WriteString(fmt.Sprintf("%s", string(bts)))
	} else {
		f.WriteString(fmt.Sprintf("\n%s", string(bts)))
	}
}

func Info(an any) {
	go DataLogger.Info(an)
}

func DirCheck(dir string) error {
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

func FileIsExistOrCreate(name string) (*os.File, bool, bool) {
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
