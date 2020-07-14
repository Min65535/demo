package singleton

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type singleton struct{}

//懒汉单例 (缺点：非线程安全。当正在创建时，有线程来访问此时ins = nil就会再创建，单例类就会有多个实例)
var ins *singleton

func GetInsLazy() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

//饿汉单例
var insHun = &singleton{}

func GetInsHun() *singleton {
	return insHun
}

//双重锁
var insDouble *singleton
var mu sync.Mutex

func GetInsDoubleMutex() *singleton {
	if insDouble == nil {
		mu.Lock()
		defer mu.Unlock()
		if insDouble == nil {
			insDouble = &singleton{}
		}
	}
	return insDouble
}

//线程安全只跑一次
var insOnce *singleton
var once sync.Once

func GetInsOnce() *singleton {
	once.Do(func() {
		insOnce = &singleton{}
	})
	return insOnce
}

func checkOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			once.Do(onceBody)
			done <- i
		}
		quit <- 1
	}()

	for i := 13; i < 33; i++ {
		go func(x int) {
			once.Do(onceBody)
			done <- x
		}(i)
	}

	for {
		select {
		case x := <-done:
			fmt.Println("x:", x)
		case <-quit:
			return
		case <-time.After(time.Second * 5):
			fmt.Println("timeout 5")
			return
		}
	}
}

type M struct {
	X string
	I int
	F float64
	V map[string]string
	//A []string
}

func CreatePointer() bool {
	//return new(int)

	var q = M{X: "as", I: 1, F: 0.7, V: map[string]string{"ss": "ss"}}
	var a = M{X: "as", I: 1, F: 0.7, V: map[string]string{"ss": "ss"}}
	return reflect.DeepEqual(q, a)
}
