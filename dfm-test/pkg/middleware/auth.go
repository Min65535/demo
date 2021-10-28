package middleware

import (
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	AuthAdmin() gin.HandlerFunc
}
