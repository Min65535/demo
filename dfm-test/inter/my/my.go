package main

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// func main() {
// 	host := "10.8.0.13"
// 	go Ho(host, 10*time.Second)
//
// 	go Ho(host, 5*time.Second)
// 	stop := 0
//
// 	for {
// 		if stop, _ := fmt.Scan(&stop); stop == 1 {
// 			fmt.Println("my timer stop by fmt scan 1.")
// 			break
// 		}
// 	}
// }
//
// func ping(host string) {
// 	var b string
// 	goos := runtime.GOOS
// 	var cmd = &exec.Cmd{}
// 	switch goos {
// 	case "windows":
// 		cmd = exec.Command("ping", host, "-n", "2")
//
// 	case "linux":
// 		cmd = exec.Command("ping", host, "-c", "2")
// 	}
// 	stdout, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Println("StdoutPipe err:", err.Error())
// 		return
// 	}
// 	b = string(stdout)
// 	log.Println("the result of ping:", b)
// }
//
// func Ho(host string, sec time.Duration) {
// 	timer := time.NewTimer(sec) // 新建一个Timer
//
// 	for {
// 		select {
// 		case <-timer.C:
// 			fmt.Println("my timer start tick tok in ", sec.String())
// 			ping(host)
// 			timer.Reset(sec) // 上一个when执行完毕重新设置
// 		}
// 	}
// }

type A struct {
	ElementB *B
}

type At struct {
	Bt *Bt
}

type Bt struct {
	LevelInfo *LevelInfo
}

type LevelInfo struct {
	Exp int32
}

type B struct {
	ElementA *A
}

func min() {
	a := A{}
	b := B{}

	a.ElementB = &b
	b.ElementA = &a
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	str, err := json.Marshal(a)
	if err != nil {
		return
	}
	fmt.Println("str:", string(str))
}

type GameMatchError interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

type Robot struct {
	task           chan int
	stop           chan bool
	outForkStop    chan bool
	normalStopFork chan int
	overtime       int32 // 毫秒
	robotExit      bool
	hasNovice      bool
}

func (r *Robot) mf() {
	if r.robotExit {
		return
	}
	tim := time.NewTimer(time.Duration(5) * time.Second) // 新建一个Timer	// 超时退出
	for {
		select {
		case <-r.outForkStop:
			r.stop <- true
			fmt.Println("out")
			return
		case <-r.normalStopFork:
			fmt.Println("normal out")
			return
		case <-tim.C:
			fmt.Println("time out")
			if r.robotExit {
				return
			}
			r.stop <- true
			return
		}
	}
}

func (r *Robot) notify() {
	for {
		if r.robotExit {
			return
		}
		select {
		// case tk, ok := <-r.task:
		case tk := <-r.task:
			// if ok {
			// 	if r.robotExit {
			// 		fmt.Println("111")
			// 		return
			// 	}
			// 	fmt.Println("tk:", tk)
			//
			// }
			if r.robotExit {
				fmt.Println("111")
				return
			}
			fmt.Println("tk:", tk)
			r.robotExit = true
			r.normalStopFork <- 1

			return
		case <-r.stop:
			r.robotExit = true
			time.Sleep(5 * time.Second)
			close(r.task)
			if _, ok := <-r.task; ok {
				fmt.Println("222")
			}
			fmt.Println("333")
			return
		}
	}
}
func m(ta chan bool) {
	time.Sleep(1 * time.Second)
	ta <- true
}

func main() {
	r := &Robot{
		task:           make(chan int, 1),
		stop:           make(chan bool),
		overtime:       1,
		outForkStop:    make(chan bool),
		normalStopFork: make(chan int),
	}

	go r.mf()

	// go m(r.outForkStop)
	go r.notify()
	go func() {
		time.Sleep(8 * time.Second)

		r.task <- 666

	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.QyLogger.Info("server get a signal ", zap.Any("s", s))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.QyLogger.Info("server exit")
			return
		default:
			return
		}
	}

}
