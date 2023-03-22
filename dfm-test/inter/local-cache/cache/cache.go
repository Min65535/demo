package cache

import (
	"container/list"
	"log"

	"fmt"

	"strings"

	"sync"

	"time"
)

type Entity[T comparable] struct {
	list *list.List

	m map[T]*list.Element

	lock sync.RWMutex
}

type entry[T comparable] struct {
	Key T

	Unix int64
}

// 暂时无数量限制

func NewEntity[T comparable]() *Entity[T] {

	return &Entity[T]{

		list: list.New(),

		m: make(map[T]*list.Element, 1024),

		lock: sync.RWMutex{},
	}

}

func (lru *Entity[T]) Update(key T) {

	lru.lock.Lock()

	defer lru.lock.Unlock()

	ele, ok := lru.m[key]

	if ok {

		eleEntry := ele.Value.(*entry[T])

		eleEntry.Unix = time.Now().Unix()

		lru.list.MoveToBack(ele)

	} else {

		ent := &entry[T]{Key: key, Unix: time.Now().Unix()}

		ele = lru.list.PushBack(ent)

		lru.m[key] = ele

	}

}

// CheckTimeout  timeoutSec 单位second / timeoutCallBack 回调函数 key为过期key

func (lru *Entity[T]) CheckTimeout(timeoutSec int64, timeoutCallBack func(key T)) {

	lru.lock.Lock()

	defer lru.lock.Unlock()

	now := time.Now().Unix()

	for cur := lru.list.Front(); cur != nil; cur = lru.list.Front() {

		curEntry := cur.Value.(*entry[T])

		if now-curEntry.Unix < timeoutSec {

			return

		}

		delete(lru.m, curEntry.Key)

		lru.list.Remove(cur)

		go timeoutCallBack(curEntry.Key)

	}

}

func (lru *Entity[T]) ExistKey(key T) bool {

	lru.lock.RLock()

	defer lru.lock.RUnlock()

	_, ok := lru.m[key]

	return ok

}

func (lru *Entity[T]) Print() {

	log.Println(lru)
}

func (lru *Entity[T]) String() string {

	ans := ""

	for cur := lru.list.Front(); cur != nil; cur = cur.Next() {

		ent := cur.Value.(*entry[T])

		ans = fmt.Sprintf("%s (key=%v unix=%d) -->", ans, ent.Key, ent.Unix)

	}

	return strings.TrimSuffix(ans, "-->")

}

//Issue a POSTrequest to this endpoint to update a Creative Set’s name or to create a new Creative Set. To create a new Creative Set, don’t fill in the creative_set_idin the request URL nor the request body. A Creative Set is a collection of assets that can be combined together to generate many ads so that you do not need to upload each asset multiple times. Creative Sets also include information about the languages that the underlying assets are in.