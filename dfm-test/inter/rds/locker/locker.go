package locker

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ErrTryLockLimit = errors.New("try lock limit")
)

type Locker interface {
	Lock(ctx context.Context, key interface{}) error
	Unlock(ctx context.Context, key interface{})
}

type LockOption interface {
	apply(opt *lockOption)
}

type lockOption struct {
	exp     time.Duration
	maxSpin int
	prefix  string
}

type OptionFunc func(*lockOption)

func (fn OptionFunc) apply(opt *lockOption) {
	fn(opt)
}

func WithExpiration(exp time.Duration) OptionFunc {
	return func(o *lockOption) {
		o.exp = exp
	}
}

func WithMaxSpin(maxSpin int) OptionFunc {
	return func(o *lockOption) {
		o.maxSpin = maxSpin
	}
}

func WithPrefix(prefix string) OptionFunc {
	return func(o *lockOption) {
		o.prefix = prefix
	}
}

type locker struct {
	rd  *redis.Client
	opt lockOption
}

func New(rd *redis.Client, opts ...LockOption) *locker {
	o := lockOption{
		exp:     30 * time.Second,
		maxSpin: 4,
	}
	for _, opt := range opts {
		opt.apply(&o)
	}
	l := &locker{
		rd:  rd,
		opt: o,
	}
	return l
}

func (l *locker) lockName(key interface{}) string {
	return l.opt.prefix + "_" + fmt.Sprintf("%v", key)
}

func (l *locker) Lock(ctx context.Context, key interface{}) error {
	if l.opt.maxSpin == 0 {
		return ErrTryLockLimit
	}
	var loopCount int
	for {
		set, err := l.rd.SetNX(ctx, l.lockName(key), true, l.opt.exp).Result()
		if err != nil {
			return fmt.Errorf("redis: setnx err: %w", err)
		}
		if set {
			break
		}
		loopCount++
		if loopCount >= l.opt.maxSpin {
			return ErrTryLockLimit
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (l *locker) Unlock(ctx context.Context, key interface{}) {
	l.rd.Del(ctx, l.lockName(key))
}
