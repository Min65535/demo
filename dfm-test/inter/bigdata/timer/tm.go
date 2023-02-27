package timer

import (
	"github.com/robfig/cron"
)

type SvcTimer interface {
	Start()
	Stop()
}

type tm struct {
	spc string
	cro *cron.Cron
}

func (t *tm) Start() {
	t.cro.Start()
}

func (t *tm) Stop() {
	t.cro.Stop()
}

func NewTm(spc string, fc func()) SvcTimer {
	c := cron.New()
	c.AddFunc(spc, fc)
	return &tm{spc: spc, cro: c}
}
