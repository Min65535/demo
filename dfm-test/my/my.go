package main

import (
	"os/exec"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"errors"
	"fmt"
	"time"
	"log"
)

//获取可执行程序当前路径
func pathGet() (path string, err error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}

	path, err = filepath.Abs(file)
	if err != nil {
		return
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0 {
		err = errors.New(`没有找到字符“/”或者"\"`)
		return
	}
	path = string(path[0 : i+1])
	return
}

func main() {

	//go Love() // 起一个协程去执行定时任务
	go Love() // 起一个协程去执行定时任务

	stop := 0
	for {
		fmt.Scan(&stop)
		if stop == 1 {
			break
		}
	}
}

func ping() {
	var b string
	goos := runtime.GOOS
	switch goos {
	case "windows":
		cmd := exec.Command("ping", "172.16.0.30", "-n", "2")
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("StdoutPipe err:", err.Error())
			return
		}
		b = string(stdout)
	case "linux":
		cmd := exec.Command("ping", "172.16.0.30", "-c", "2")
		stdout, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("StdoutPipe err:", err.Error())
			return
		}
		b = string(stdout)
	}
	log.Println("the result of ping:", b)
}

func Love() {
	timer := time.NewTimer(1 * time.Minute) // 新建一个Timer
	//timer := time.NewTimer(2 * time.Second) // 新建一个Timer

	for {
		select {
		case <-timer.C:
			fmt.Println("my timer start")
			ping()
			timer.Reset(1 * time.Minute) // 上一个when执行完毕重新设置
			//timer.Reset(2 * time.Second) // 上一个when执行完毕重新设置
		}
	}
	return
}
