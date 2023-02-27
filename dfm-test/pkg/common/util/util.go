package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// 获取可执行程序当前路径
func BinaryPathGet() (path string, err error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}

	path, err = filepath.Abs(file)
	if err != nil {
		return
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}

	i := strings.LastIndex(path, "/")
	if i < 0 {
		err = errors.New(`没有找到字符“/”或者"\"`)
		return
	}
	path = path[0 : i+1]
	return
}

func PrintReq() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不打印有文件的请求
		if c.ContentType() == "multipart/form-data" {
			return
		}
		params, _ := ioutil.ReadAll(c.Request.Body)
		if err := c.Request.Body.Close(); err != nil {
			panic(err)
		}
		// 这里有意思，呵呵
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(params))
		paramsStr := string(params)
		log.QyLogger.Info("[Get Request]", zap.String("req_url", c.Request.RequestURI), zap.String("params", paramsStr))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func PrintResp() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		respStr := blw.body.String()
		if len(respStr) < 500 {
			log.QyLogger.Info("[Response Result]", zap.String("req_url", c.Request.RequestURI), zap.String("resp", respStr))
		} else {
			log.QyLogger.Info("[Response Partial Result]", zap.String("req_url", c.Request.RequestURI), zap.String("resp", respStr[0:500]))
		}
	}
}

// 处理 go 中 log 到 /var/log/qy 中
func InitLogger(rename ...string) {
	switch qyenv.GetDockerEnv() {
	case "", "0":
	default:
		logNameSuffix := "docker"
		if len(rename) == 1 {
			logNameSuffix += "_" + rename[0]
		}
		dockerLogDir := filepath.Join("/var/log/qy", logNameSuffix)
		if err := os.MkdirAll(dockerLogDir, 0755); err != nil {
			panic(err)
		}
		podName := os.Getenv("HOSTNAME")
		if podName == "" {
			panic("can't get HOSTNAME from env")
		}
		log.InitLoggerWithCaller(zapcore.DebugLevel, dockerLogDir, podName+".log", true)
	}
}

// 处理Http请求
func HttpReq(r *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	return httpClient.Do(r)
}

func RenderSuccess(c *gin.Context, resultJson interface{}) {
	c.JSON(200, resultJson)
}

// 捕获详细panic信息
func panicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) // 4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}

// 直接defer就可以了
func Catch(msg string) {
	if panicErr := recover(); panicErr != nil {
		err := string(panicTrace(32))
		fmt.Println(msg+" catch: ", err)
		log.QyLogger.Error(msg+"#Catch", zap.String("msg", err))
	}
}

// 生成32位UUID
func GenerateUUID() string {
	// 19位纳秒时间戳
	// 加上1位字母+12位数字/英文随机字符
	return fmt.Sprintf("%d%s%s", time.Now().UnixNano(), GetCommonPrivateKeyLetter(1), GetCommonPrivateKey(12))
}

// 获取秘钥方法
func GetCommonPrivateKeyNum(n uint) string {
	var letters = []rune("0123456789")
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 获取秘钥方法
func GetCommonPrivateKeyLetter(n uint) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 获取秘钥方法
func GetCommonPrivateKey(n uint) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 获得零点时间
func GetTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// 获得时间天数差
func GetDays(start, end time.Time) int {
	return int(GetTime(start).Sub(GetTime(end)).Hours() / 24)
}

// 获得时间段相差天数方法,传入的时间格式:"2016-01-02"
func GetTimeArr(start, end string) (int64, error) {
	// 转成时间戳
	startUnix, err := time.Parse("2006-01-02", start)
	if err != nil {
		return 0, err
	}
	endUnix, err := time.Parse("2006-01-02", end)
	if err != nil {
		return 0, err
	}
	if startUnix.After(endUnix) {
		return 0, errors.New("比较的时间不合法")
	}
	return int64(endUnix.Sub(startUnix).Hours()) / 24, nil
}
