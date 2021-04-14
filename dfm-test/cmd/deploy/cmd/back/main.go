package main

import (
	"context"
	"demo/dfm-test/pkg/common/util"
	"flag"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
	util.InitLogger("publish")
	env := qyenv.GetUseDocker()
	if env == 2 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode) // 全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	}
	// db :=mysql.MakeDB(db.GetDBConfig())

	engine := gin.New()
	g1 := engine.Group("/api/v1")
	{
		g1.GET("/get", func(c *gin.Context) {})
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
		log.QyLogger.Info("server get a signal ", zap.Any("s", s))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.QyLogger.Info("server exit")
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
		log.QyLogger.Error("server gracefully shutdown fail", zap.Error(err))
		time.Sleep(time.Second)
		return
	}
	log.QyLogger.Info("gin engine stop")
}
