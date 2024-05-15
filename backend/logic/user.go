package logic

import (
	"blogWeb/dao/mysql"
	"blogWeb/models"
	"blogWeb/pkg/jwt"
	"blogWeb/pkg/snowflake"
	"errors"

	"go.uber.org/zap"
)

// 存放业务逻辑的代码

func SingUp(p *models.ParamSignUp) (err error) {

	// 判断用户存不存在
	exists, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		zap.L().Debug("logic user.go at 20")
		return err
	}
	if exists {
		zap.L().Debug("logic user.go at 24")
		// 用户已存在
		return errors.New("用户已存在")
	}

	// 生成UID
	userID := snowflake.GetID()
	zap.L().Debug("logic user.go at 31")

	// 构造一个User实例
	u := models.User{
		UserID:   userID,
		UserName: p.Username,
		PassWord: p.Password,
	}

	// 保存进数据库
	err = mysql.InsertUser(&u)

	return
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.Username,
		PassWord: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		zap.L().Error("mysql.Login failed", zap.Error(err))
		return nil, err
	}

	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.Token = token
	return
}
