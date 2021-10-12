package main

import (
	"github.com/min65535/demo/dfm-test/inter/consume/common"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	NewTaskServer().Run()
}

type TaskMsg struct {
	Addr string `json:"addr"`
	Id   uint   `json:"id"`
}

type TaskServer struct {
	db         mysql.DB
	worker     chan *common.NameAndValue
	running    bool
	stopWorker chan struct{}
	exit       chan os.Signal
}

func NewTaskServer() *TaskServer {
	// todo 处理中间状态
	return &TaskServer{
		db:         common.GetDbConfig(),
		worker:     make(chan *common.NameAndValue, 100),
		stopWorker: make(chan struct{}),
		exit:       make(chan os.Signal, 1),
	}
}

func (t *TaskServer) loadWorker() {
	for {
		fmt.Println("loadWorker Run")
		if !t.running {
			fmt.Println("TaskServer#loadWorker#stop")
			return
		}
		select {
		case data := <-t.worker:
			// 执行任务 todo
			time.Sleep(5 * time.Second)
			fmt.Println("execute worker task:", json.StringifyJson(data))
			addTheRemark(t.db.GetDB(), []*common.NameAndValue{data})

		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (t *TaskServer) Notify() {
	signal.Notify(t.exit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case s := <-t.exit:
			fmt.Println(fmt.Sprintf("server get a signal %s", s.String()))
			log.QyLogger.Info(fmt.Sprintf("server get a signal %s", s.String()))
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				t.close()
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}
}

func (t *TaskServer) close() {
	t.stopWorker <- struct{}{}
	fmt.Println("TaskServer close input")
}

func (t *TaskServer) Start() {
	if t.running {
		return
	}
	fmt.Println("TaskServer Start")
	t.running = true
	go t.run()
}

func (t *TaskServer) Run() {
	if t.running {
		return
	}
	fmt.Println("TaskServer Run")
	t.running = true
	t.run()
}

func (t *TaskServer) workerChan(data []*common.NameAndValue) {
	// todo tasks locks
	closeTheBlock(t.db.GetDB(), data)
	if len(data) > 0 {
		for _, x := range data {
			t.worker <- x
		}
	}
}

func (t *TaskServer) runWorkerServer() {
	if !t.running {
		return
	}
	go t.loadWorker()
	go t.loadWorker()
	go t.loadWorker()
	go t.loadWorker()
	for {
		fmt.Println("runWorkerServer Run")
		select {
		case <-t.stopWorker:
			t.running = false
			fmt.Println("stop runWorkerServer")

		default:
			if !t.running {
				return
			}
			data, err := getDataFromDb(t.db.GetDB())
			if err != nil {
				return
			}
			fmt.Println("-------------runWorkerServer------------")
			if len(data) < 1 {
				time.Sleep(5 * time.Second)
				fmt.Println("-------------runWorkerServer 5 * time.Second------------")
				continue
			}
			t.workerChan(data)
		}
	}
}

func (t *TaskServer) run() {
	if !t.running {
		return
	}
	go t.runWorkerServer()
	t.Notify()
}
