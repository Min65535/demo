package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"runtime"
	"syscall"
)


type MemStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat() MemStatus {
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc

	//系统占用,仅linux/mac下有效
	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err == nil {
		mem.All = sysInfo.Totalram * uint64(syscall.Getpagesize())
		mem.Free = sysInfo.Freeram * uint64(syscall.Getpagesize())
		mem.Used = mem.All - mem.Free
	}
	return mem
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Press Test")
	fmt.Println("---------------------")
	i := 1
	for {
		go func() {
			for f := 0; f < 2*i; f++ {
				i++
			}
		}()
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		//convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, It's finished and the size of 'i' is ", i)
			return
		}
	}
}
