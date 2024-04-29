package routers

import (
	"blogWeb/controller"
	"blogWeb/logger"
	"blogWeb/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1") // 控制版本，版本1

	// 注册业务路由
	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	// r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	// 	// 如果是登录的用户，判断请求头是否有效的JWR
	// 	// isLogin := true
	// 	// c.Request.Header.Get("Authorization")
	// 	// if isLogin {
	// 	// 	c.String(http.StatusOK, "ping")
	// 	// } else {
	// 	// 	// 否则就直接返回请登录
	// 	// 	c.String(http.StatusOK, "请登录")
	// 	// }
	// 	c.String(http.StatusOK, "ping")
	// })

	{
		v1.GET("/community", controller.CommunityHandler)

	}

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "ok")
	// })
	//
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404")
	})

	return r
}
