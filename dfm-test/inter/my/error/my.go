package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
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

func qwe() {

}

func main() {
	// // go min()
	// // for {
	// // 	select {}
	// // }
	// c := make(chan int, 1)
	// // c <- 123
	// close(c)
	//
	// var res int
	// var ok bool
	//
	// res, ok = <-c
	// fmt.Println(res, ok)
	//
	// res, ok = <-c
	// fmt.Println(res, ok) // 此时ok为false
	//
	// var tb = &At{Bt: &Bt{LevelInfo: nil}}
	// // if tb.Bt.LevelInfo.Exp < 500 || tb.Bt.LevelInfo == nil {
	// if tb.Bt.LevelInfo == nil || tb.Bt.LevelInfo.Exp < 500 {
	// 	fmt.Println("1111")
	// }

	cv := make(chan interface{}, 1)

	go func() {
		time.Sleep(time.Second * 10)
		cv <- &errorString{
			s: "hhhhh",
		}
	}()

	// for {
	select {
	case res, ok := <-cv:
		if ok {
			if err, oks := res.(GameMatchError); oks {
				fmt.Println("err:", err.Error())
			}
			break
		}
	}
	// }

	fmt.Println("2222")
}
