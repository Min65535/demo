package server

import (
	"demo/dfm-test/inter/proxy/internal/core/service"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	Core *gin.Engine
	Svc  service.ProxyServer
}

func NewEngine(svc service.ProxyServer) *Engine {
	return &Engine{
		Core: gin.New(),
		Svc:  svc,
	}
}

func InitRouter ()() {

}
