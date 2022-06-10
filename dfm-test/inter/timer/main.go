package main

import (
	"fmt"
	"github.com/min65535/demo/dfm-test/inter/timer/inner"
	"time"
)

func main() {
	tr := inner.NewTimer()
	syncStartChan := make(chan bool, 1)
	gg := make(chan bool, 1)

	time.Sleep(time.Duration(2000) * time.Millisecond)
	syncStartChan <- true

	tr.ResetTimer(7000)
	for {
		select {
		case _, ok := <-tr.Timer.C:
			if ok {
				fmt.Println("hahahhahaha111")
				gg <- true
			}
		case <-syncStartChan:
			fmt.Println("hahahhahaha2222")
			tr.ResetTimer(3000)
			fmt.Println("hahahhahaha2222 end")
		case <-gg:
			fmt.Println("gg")
			return
		}
	}
}
