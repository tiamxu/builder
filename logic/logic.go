package logic

import (
	"github.com/gin-gonic/gin"
)

func RegisterHttpRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
