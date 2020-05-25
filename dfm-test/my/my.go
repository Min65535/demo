package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	go Ha() // 起一个协程去执行定时任务
	stop := 0

	// signal
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		if stop, _ := fmt.Scan(&stop); stop == 1 {
			fmt.Println("my timer stop by fmt scan 1.")
			break
		}

		//s := <-c
		//log.Println(fmt.Sprintf("server get a signal %s", s.String()))
		//switch s {
		//case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		//	//log.Println("server gracefully shutdown")
		//	//Stop()
		//	fmt.Println("my timer stop by terminal fir.")
		//	return
		//case syscall.SIGHUP:
		//	fmt.Println("my timer stop by terminal sec.")
		//default:
		//	return
		//}
	}
}

func ping() {
	var b string
	goos := runtime.GOOS
	switch goos {
	case "windows":
		cmd := exec.Command("ping", "localhost", "-n", "2")
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("StdoutPipe err:", err.Error())
			return
		}
		b = string(stdout)
	case "linux":
		cmd := exec.Command("ping", "localhost", "-c", "2")
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("StdoutPipe err:", err.Error())
			return
		}
		b = string(stdout)
	}
	log.Println("the result of ping:", b)
}

func Ha() {
	timer := time.NewTimer(10 * time.Second) // 新建一个Timer
	//timer := time.NewTimer(1 * time.Minute) // 新建一个Timer
	//timer := time.NewTimer(2 * time.Second) // 新建一个Timer

	for {
		select {
		case <-timer.C:
			fmt.Println("my timer start.")
			ping()
			timer.Reset(10 * time.Second) // 上一个when执行完毕重新设置
			//timer.Reset(2 * time.Second) // 上一个when执行完毕重新设置
		}
	}
}
