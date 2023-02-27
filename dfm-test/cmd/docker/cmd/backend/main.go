package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/gin-gonic/gin"
	"github.com/min65535/demo/dfm-test/pkg/common/util"
	"github.com/thinkerou/favicon"
)

func getImages(ctx *gin.Context) {
	Host := ctx.Request.Host
	path := ctx.Request.URL.Path
	RawQuery := ctx.Request.URL.RawQuery
	rawQueries := strings.Split(RawQuery, "/")
	rawQuery := rawQueries[len(rawQueries)-1]
	if rawQuery == "" {
		ctx.Data(400, "application/json", json.StringifyJsonToBytes(gin.H{"success": false, "err_msg": "照片查询项目不能为空，举例:http://" + Host + path + "?favicon.ico"}))
		return
	}
	// realUrl := "http://" + Host + "/static/" + rawQuery
	// realUrl := "http://" + "172.30.9.71:3001" + "/static/" + rawQuery
	realUrl := "http://" + "image-back:" + port + "/static/" + rawQuery
	resp, err := proxyReq(realUrl)
	if err != nil {
		ctx.Data(400, "application/json", []byte(systemError()))
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Data(400, "application/json", []byte(systemError()))
		return
	}
	ctx.Data(resp.StatusCode, "image/png", body)
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

var port string

func main() {
	flag.StringVar(&port, "p", "3000", "端口号， 默认为3000")
	flag.Parse()
	env := qyenv.GetUseDocker()
	if env == 2 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode) // 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}

	engine := gin.New()
	root, _ := util.BinaryPathGet()
	engine.Use(favicon.New(root + "/static/favicon.ico"))
	// engine.Use(util.PrintReq())
	engine.Static("/static", root+"/static")
	// engine.Use(util.PrintResp())
	g1 := engine.Group("/images/api/v1")
	{
		g1.GET("/get", getImages)
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
