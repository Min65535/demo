package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/gin-gonic/gin"
	"github.com/min65535/demo/dfm-test/pkg/common/util"
)

func putFiles(ctx *gin.Context) {
	host := ctx.Request.Host
	f, err := ctx.FormFile("f1")
	if err != nil {
		log.Printf("Failed to putFiles file: %s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// 上传文件
	mm := make(map[string]interface{})
	mm["code"] = 200
	mm["msg"] = "ok"
	url, err := fileSave(f)
	if err != nil {
		mm["msg"] = err.Error()
	}
	if url != "" {
		url = fmt.Sprintf("%s%s/static/%s", "http://", host, url)
		mm["url"] = url
	}
	ctx.JSON(http.StatusOK, gin.H(mm))
}

func dirCheck(dir string) error {
	path, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	_, err1 := os.Stat(path)
	if err1 != nil {
		if os.IsNotExist(err1) {
			err2 := os.Mkdir(path, os.ModePerm)
			if err2 != nil {
				return err2
			} else {
				return nil
			}
		} else {
			return err1
		}
	}
	return nil
}

func fileSave(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		log.Printf("Failed to Open file: %s\n", err.Error())
		return "", err
	}
	defer src.Close()
	fileName := file.Filename
	pa := pathFile + "/" + fileName
	fmt.Println("file true path:", pa)
	newFile, err := os.Create(pa)
	if err != nil {
		log.Printf("Failed to create file: %s\n", err.Error())
		return "", err
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, src)
	if err != nil {
		log.Printf("Failed to save data into file: %s\n", err.Error())
		return "", err
	}
	return fileName, nil
}

// 系统异常json结构体
func systemError() string {
	return json.StringifyJson(gin.H{"success": false, "err_msg": "system error"})
}

// 发起请求
func proxyReq(targetUrl string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", targetUrl, nil)
	return util.HttpReq(req)
}

var port, pathFile string

func main() {
	flag.StringVar(&port, "p", "3011", "端口号， 默认为3011")
	flag.Parse()
	env := qyenv.GetUseDocker()
	if env == 2 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode) // 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}

	engine := gin.New()
	root, _ := util.BinaryPathGet()
	fmt.Println("root:", root)
	// engine.Use(favicon.New(root + "/static/favicon.ico"))
	// engine.Use(util.PrintReq())

	switch runtime.GOOS {
	case "windows":
		pathFile = root + "static"
	case "linux", "darwin":
		pathFile = root + "static"
	}
	fmt.Println("pathFile:", pathFile)
	dirCheck(pathFile)
	engine.Static("/static", pathFile)
	// engine.Use(util.PrintResp())
	g1 := engine.Group("/api/v1/file")
	{
		g1.POST("/put_file", putFiles)
	}
	run(port, engine)
}

func run(port string, engine http.Handler) {
	httpSrv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: engine}
	go func() {
		err := httpSrv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	// init signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ginStop(httpSrv)
			return
		default:
			return
		}
	}

}

// gin gracefully stop server
func ginStop(httpSrv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		time.Sleep(time.Second)
		return
	}
}
