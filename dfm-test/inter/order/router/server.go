package router

import (
	"demo/dfm-test/inter/order/handle"
	"github.com/gin-gonic/gin"
)

type Server struct {
	hld *handle.Handler
	eng *gin.Engine
}

func NewServer(hld *handle.Handler, eng *gin.Engine) *Server {
	return &Server{hld: hld, eng: eng}
}

func (s Server) Run(port string) error {
	s.RegisterRouter()
	return s.eng.Run(port)
}
