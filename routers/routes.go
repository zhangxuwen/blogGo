package routers

import (
	"blogWeb/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})
	return r
}
