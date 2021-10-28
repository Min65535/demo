package main

import (
	"fmt"
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

func run() {
	data0 := make(map[int]int)
	data1 := make(map[int]int)
	data2 := make(map[int]int)
	data3 := make(map[int]int)
	data4 := make(map[int]int)
	data5 := make(map[int]int)
	data6 := make(map[int]int)
	data7 := make(map[int]int)
	data8 := make(map[int]int)
	data9 := make(map[int]int)
	i := 1
	var rate = 0.15
	for {
		i++
		data0[i] = i
		data1[i] = i
		data2[i] = i
		data3[i] = i
		data4[i] = i
		data5[i] = i
		data6[i] = i
		data7[i] = i
		data8[i] = i
		data9[i] = i
		if float64(MemStat().Free)/float64(MemStat().All) < rate {
			fmt.Println("the current free rate is:", float64(MemStat().Free)/float64(MemStat().All))
			return
		}
		fmt.Println("i:", i)
	}
}

func main() {
	fmt.Println("Simple Memory Press Test")
	fmt.Println("---------------------")
	run()
}
