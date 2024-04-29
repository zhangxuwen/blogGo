package controller

import (
	"blogWeb/dao/mysql"
	"blogWeb/logic"
	"blogWeb/models"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// fmt.Printf("%v\n", p)
		// 获取 validator.ValidationErrors 类型的 errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非 validator.ValidationErrors 类型错误直接返回
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			ResponseError(c, CodeIvalidParam)
			return
		}

		// zap.L().Debug("controller user.go at 28")
		// 请求参数有误，直接返回响应
		// zap.L().Error("SignUp with invalid param", zap.Error(err))
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
		// })
		ResponseErrorWithMsg(c, CodeIvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 手动对请求参数进行详细的业务规则校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	// 	zap.L().Error("SignUp with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	zap.L().Debug("controller user.go at 46")

	// 业务处理
	if err := logic.SingUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": "注册失败",
		// })
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	zap.L().Debug("controller user.go at 55")

	// 返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "success",
	// })
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// c.JSON(http.StatusOK, gin.H{
			// 	"msg": err.Error(),
			// })
			ResponseError(c, CodeIvalidParam)
			return
		}

		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// })
		ResponseErrorWithMsg(c, CodeIvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": "用户名或密码错误",
		// })
		ResponseError(c, CodeIvalidPassword)
		return
	}
	// 返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"msg": "登录成功",
	// })
	ResponseSuccess(c, token)
}
