package main

import (
	"errors"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

func main() {
	hn := os.Getenv("HOSTNAME")
	if hn == "" {
		os.Setenv("HOSTNAME", "ali-test1")
	}
	fmt.Println("HOSTNAME:", hn)
	docker := qyenv.GetDockerEnv()
	if docker == "" {
		os.Setenv("docker_env", "1")
	}
	fmt.Println("DockerEnv:", docker)
	InitLogger("ali")
	Ha()
}

func Ha() {
	timer := time.NewTimer(1 * time.Minute) // 新建一个Timer
	//timer := time.NewTimer(2 * time.Second) // 新建一个Timer

	for {
		select {
		case <-timer.C:
			log.QyLogger.Error("RegisterErr", zap.Error(errors.New("fail to register at "+time.Now().Format("2006-01-02 15:04:05"))))
			timer.Reset(1 * time.Minute) // 上一个when执行完毕重新设置
			//timer.Reset(2 * time.Second) // 上一个when执行完毕重新设置
		}
	}
}

// 处理 docker 中 log 到 /var/log/qy 中
func InitLogger(rename ...string) {
	switch qyenv.GetDockerEnv() {
	case "", "0":
	default:
		logNameSuffix := "dfm_test"
		if len(rename) > 1 {
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
