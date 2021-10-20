package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	host := "10.8.0.13"
	go Ho(host, 10*time.Second) // 起一个协程去执行定时任务

	go Ho(host, 5*time.Second)
	stop := 0

	for {
		if stop, _ := fmt.Scan(&stop); stop == 1 {
			fmt.Println("my timer stop by fmt scan 1.")
			break
		}
	}
}

func ping(host string) {
	var b string
	goos := runtime.GOOS
	var cmd = &exec.Cmd{}
	switch goos {
	case "windows":
		cmd = exec.Command("ping", host, "-n", "2")

	case "linux":
		cmd = exec.Command("ping", host, "-c", "2")
	}
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("StdoutPipe err:", err.Error())
		return
	}
	b = string(stdout)
	log.Println("the result of ping:", b)
}

func Ho(host string, sec time.Duration) {
	timer := time.NewTimer(sec) // 新建一个Timer

	for {
		select {
		case <-timer.C:
			fmt.Println("my timer start tick tok in ", sec.String())
			ping(host)
			timer.Reset(sec) // 上一个when执行完毕重新设置
		}
	}
}
