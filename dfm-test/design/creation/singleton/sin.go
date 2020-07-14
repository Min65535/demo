package singleton

import (
	"sync"
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
