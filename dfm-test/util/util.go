package util

import (
	"bytes"
	"errors"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//获取可执行程序当前路径
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
		logNameSuffix := "image-index"
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

//处理Http请求
func HttpReq(r *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	return httpClient.Do(r)
}


func RenderSuccess(c *gin.Context, resultJson interface{}) {
	c.JSON(200, resultJson)
}
