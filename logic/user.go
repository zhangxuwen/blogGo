package logic

import (
	"blogWeb/dao/mysql"
	"blogWeb/models"
	"blogWeb/pkg/snowflake"
)

// 存放业务逻辑的代码

func SingUp(p *models.ParamSignUp) {

	// 判断用户存不存在
	mysql.QueryUserByID()

	// 生成UID
	snowflake.GetID()

	// 密码加密

	// 保存进数据库
	mysql.InsertUser()

}
