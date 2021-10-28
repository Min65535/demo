package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Publish interface {
	Login(ctx context.Context, req *LoginReq) (*LoginResp,error)
	Auth(ctx context.Context) error
	ImagePush(ctx context.Context, req *ImagePushReq) error
	SvcDeploy(ctx context.Context, req *SvcDeployReq) error
}

type Resp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	ErrCode string      `json:"err_code,omitempty"`
	ErrMsg  string      `json:"err_msg,omitempty"`
}

type ginHttpServer struct {
	ginEngine *gin.Engine
	srv       Publish
	group     string
}


func newGinHttpServer(ginEngine *gin.Engine, srv Publish, group string) *ginHttpServer {
	s := &ginHttpServer{ginEngine: ginEngine, srv: srv, group: group}
	return s
}

func(s *ginHttpServer) withMiddleWare(routers ...*gin.RouterGroup)  {

}

func (s *ginHttpServer) publishLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(LoginReq)
		if err := c.MustBindWith(req, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
			c.JSON(400, &Resp{Success: false, ErrCode: "400"})
			return
		}

		result, err := s.srv.Login(context.Background(), req)
		if err != nil {
			c.JSON(400, &Resp{Success: false, ErrCode: "400", ErrMsg: err.Error()})
			return
		}
		c.JSON(200, &Resp{Success: true, Data: result})
	}
}

