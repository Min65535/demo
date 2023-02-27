package logger

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime"
)

type config struct {
	fields          log.Fields // 默认字段
	fieldsFunc      []func(*log.Entry) log.Fields
	hook            []log.Hook
	disableColors   bool
	fullTimestamp   bool
	reportCaller    bool
	level           log.Level
	timestampFormat string
	logPath         string
	maxSize         int // 保存的最大文件记录数量
	maxAge          int // 最大保存天数
	maxBackups      int // 最多保存文件数量
}

type formatter struct {
	c  *config
	lf log.Formatter
}

func (f *formatter) Format(e *log.Entry) ([]byte, error) {
	for k, v := range f.c.fields {
		e.Data[k] = v
	}

	for _, f := range f.c.fieldsFunc {
		fields := f(e)
		for k, v := range fields {
			e.Data[k] = v
		}
	}

	return f.lf.Format(e)
}

type Option func(*config)

// WithFields 添加默认字段
func WithFields(fields log.Fields) Option {
	return func(c *config) {
		c.fields = fields
	}
}

// WithFieldsFunc 补充额外字段
func WithFieldsFunc(f func(*log.Entry) log.Fields) Option {
	return func(c *config) {
		c.fieldsFunc = append(c.fieldsFunc, f)
	}
}

// WithHook 添加钩子
func WithHook(hook log.Hook) Option {
	return func(c *config) {
		c.hook = append(c.hook, hook)
	}
}

// WithHooks 设置多个钩子
func WithHooks(hooks []log.Hook) Option {
	return func(c *config) {
		c.hook = hooks
	}
}

// WithLogPath 指定日志路径
func WithLogPath(path string) Option {
	return func(c *config) {
		c.logPath = path
	}
}

// DisableColors 是否启动彩色模式
func DisableColors(v bool) Option {
	return func(c *config) {
		c.disableColors = v
	}
}

// Level 日志打印等级
func Level(level log.Level) Option {
	return func(c *config) {
		c.level = level
	}
}

// ReportCaller 是否添加调用信息字段
func ReportCaller(v bool) Option {
	return func(c *config) {
		c.reportCaller = v
	}
}

// TimestampFormat 设置日志时间格式
func TimestampFormat(v string) Option {
	return func(c *config) {
		c.timestampFormat = v
	}
}

// WithMaxSize 日志文件最大
func WithMaxSize(size int) Option {
	return func(c *config) {
		c.maxSize = size
	}
}

// WithMaxAge 日志文件最大天数
func WithMaxAge(age int) Option {
	return func(c *config) {
		c.maxAge = age
	}
}

func WithMaxBackups(backups int) Option {
	return func(c *config) {
		c.maxBackups = backups
	}
}

var (
	// std is the name of the standard logger in stdlib `log`
	stdMatch = log.New()
)

func Init(opts ...Option) {
	c := &config{
		disableColors:   true,
		fullTimestamp:   true,
		reportCaller:    true,
		level:           log.DebugLevel,
		timestampFormat: "2006-01-02T15:04:05.000Z0700",
		logPath:         "/data/match/logs/my_match_data/stdout.log",
		// maxSize:         1024,
		maxAge: 7,
		// maxBackups:      10,
		// hook:            []log.Hook{TraceHook{}},
	}

	for _, v := range opts {
		v(c)
	}

	stdMatch.SetFormatter(&formatter{
		c:  c,
		lf: &MatchFormatter{},
	})

	stdMatch.SetReportCaller(c.reportCaller)
	stdMatch.SetLevel(c.level)
	for _, hook := range c.hook {
		stdMatch.AddHook(hook)
	}

	if runtime.GOOS == "linux" {
		option := &lumberjack.Logger{
			Filename: c.logPath,
			// MaxSize:  c.maxSize, // megabytes
			// MaxBackups: c.maxBackups,
			MaxAge:   c.maxAge, // days
			Compress: false,    // disabled by default
		}
		stdMatch.SetOutput(option)
	}
}
