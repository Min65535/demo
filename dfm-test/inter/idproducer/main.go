package main

import (
	"fmt"
	"sync"
	"time"
)

// GetTime 获得零点时间
func GetTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// GetDays 获得时间天数差
func GetDays(start, end time.Time) int {
	return int(GetTime(start).Sub(GetTime(end)).Hours() / 24)
}

type IDProducer interface {
	GetID() string
}

type RoomIDProducer struct {
	day   time.Time
	index int64
	mutex sync.Mutex
}

func NewRoomIDProducer() *RoomIDProducer {
	return &RoomIDProducer{
		day:   time.Now(),
		index: 0,
		mutex: sync.Mutex{},
	}
}

func (r *RoomIDProducer) updateDayAndIndex() {
	now := time.Now()
	if GetDays(now, r.day) > 0 {
		r.day = now
		r.index = 1
		return
	}
	r.index++
}

func (r *RoomIDProducer) GetID() string {
	r.mutex.Lock()
	r.updateDayAndIndex()
	num := r.index
	r.mutex.Unlock()
	return fmt.Sprintf(r.day.Format("20120102")+"-%d", num)
}

func main() {
	var gen = NewRoomIDProducer()
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("1#gen.GetID():", gen.GetID())
		}()
		go func() {
			fmt.Println("2#gen.GetID():", gen.GetID())
		}()
		go func() {
			fmt.Println("3#gen.GetID():", gen.GetID())
		}()
		go func() {
			fmt.Println("4#gen.GetID():", gen.GetID())
		}()
		go func() {
			fmt.Println("5#gen.GetID():", gen.GetID())
		}()
		go func() {
			fmt.Println("6#gen.GetID():", gen.GetID())
		}()
	}
	time.Sleep(time.Second * 5)
}
