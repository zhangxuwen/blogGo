package middlewares

import (
	"blogWeb/controller"
	"blogWeb/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 这里 Token 确定为放在 Header 的 Authorization 中，并使用 Bearer 开头
		// Authorization: Bearer xxxxxx.xxx.xxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2003,
			// 	"msg":  "请求头中auth为空",
			// })
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2004,
			// 	"msg":  "请求头中auth格式有误",
			// })
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			// c.JSON(http.StatusOK, gin.H{
			// 	"code": 2005,
			// 	"msg":  "无效的token",
			// })
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的 userID 信息保存到请求的上下文c上
		c.Set(controller.ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
