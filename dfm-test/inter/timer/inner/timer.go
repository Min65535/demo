package inner

import "time"

type Timer struct {
	Timer   *time.Timer
	EndTime time.Time
}

func NewTimer() *Timer {
	return &Timer{}
}

func (m *Timer) ResetTimer(timeout int32) {
	if m.Timer == nil {
		m.Timer = time.NewTimer(time.Millisecond * time.Duration(timeout))
	} else {
		m.Timer.Reset(time.Millisecond * time.Duration(timeout))
	}
	m.EndTime = time.Now().Add(time.Millisecond * time.Duration(timeout))
}

func (m *Timer) TimeRemaining() time.Duration {
	return m.EndTime.Sub(time.Now())
}

func (m *Timer) AddOperaTime(timeout int32) {
	timeOut := int32(m.TimeRemaining().Milliseconds()) + timeout
	m.ResetTimer(timeOut)
	m.EndTime = time.Now().Add(time.Millisecond * time.Duration(timeOut))
}

func (m *Timer) ActionTimer(millSecs int32) {
	m.ResetTimer(millSecs)
}
