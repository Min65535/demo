package api

import (
	"github.com/min65535/demo/dfm-test/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterGinServer(e *gin.Engine, server Publish, auth middleware.Authorization) {
	_ = newGinHttpServer(e, server, "/api/v1/publisher")

	_ = newGinHttpServer(e, server, "/api/v1/publisher/tool")

}

func admin(gs *ginHttpServer) {
	{
		// admin login
		gs.ginEngine.POST("login", gs.publishLogin())
	}
}
